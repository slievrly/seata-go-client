/**
2 * @Author: Nico
3 * @Date: 2019/9/7 20:02
4 */
package pingcap

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	driver "github.com/pingcap/tidb/types/parser_driver"
	"reflect"
	"testing"
)

func TestNewMySQLRecognizer(t *testing.T) {
	sql := "select name, age from tbl t where id = 1"

	mr, err := NewMySQLRecognizer(sql)
	assert.Equal(t, err, nil)
	assert.Equal(t, mr.StmtNode.Text(), sql)
	assert.Equal(t, mr.TableName, "tbl")
	assert.Equal(t, mr.TableAlias, "t")
}

func TestPingCapParser(t *testing.T) {
	p := parser.New()

	// 2. Parse a text SQL into AST([]ast.StmtNode).
	stmtNodes, _, err := p.Parse("select name, age from tbl t where id = 1", "", "")

	stmtNodes[0].Accept(&subqueryChecker{})
	// 3. Use AST to do cool things.
	fmt.Println(stmtNodes[0].Text(), err)
}

type subqueryChecker struct {
	text string
}

// Enter implements ast.Visitor interface.
func (sc *subqueryChecker) Enter(inNode ast.Node) (outNode ast.Node, skipChildren bool) {
	if expr, ok := inNode.(*ast.SubqueryExpr); ok {
		fmt.Println(expr.Query)
		return inNode, false
	} else if expr, ok := inNode.(*ast.SelectField); ok {
		fmt.Println(expr.WildCard)
		return inNode, false
	} else if expr, ok := inNode.(*ast.TableName); ok {
		fmt.Println(expr.Name)
		return inNode, false
	} else if expr, ok := inNode.(*ast.TableSource); ok {
		fmt.Println(expr.AsName)
		return inNode, false
	} else if expr, ok := inNode.(*ast.FieldList); ok {
		for _, f := range expr.Fields {
			if c, ok := f.Expr.(*ast.ColumnNameExpr); ok {
				fmt.Println(c.Name)
			}
		}
		fmt.Println()
		return inNode, false
	} else if expr, ok := inNode.(*ast.SelectStmt); ok {
		fmt.Println(expr.From.TableRefs.Left.Text())
		return inNode, false
	} else if expr, ok := inNode.(*ast.BinaryOperationExpr); ok {
		if c, ok := expr.L.(*ast.ColumnNameExpr); ok {
			fmt.Println("L", c.Name)
		}
		PrintType(expr.R)
		if c, ok := expr.R.(*ast.ColumnNameExpr); ok {
			fmt.Println("R", c.Name)
		}
		if c, ok := expr.R.(*driver.ValueExpr); ok {
			fmt.Println("R", c.GetValue())
		}

		return inNode, false
	}
	return inNode, false
}

// Leave implements ast.Visitor interface.
func (sc *subqueryChecker) Leave(inNode ast.Node) (node ast.Node, ok bool) {
	PrintType(inNode)
	return inNode, true
}

func PrintType(obj interface{}) {
	v := reflect.ValueOf(obj)
	fmt.Println(v.Elem().String())
}
