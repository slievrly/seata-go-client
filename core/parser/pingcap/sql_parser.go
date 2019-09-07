package pingcap

import (
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"github.com/slievrly/seata-go-client/core/consts"
	"github.com/slievrly/seata-go-client/core/errors"
)

var p = parser.New()

type MySQLRecognizer struct {
	StmtNode   ast.StmtNode
	TableName  string
	TableAlias string
}

func NewMySQLRecognizer(sql string) (*MySQLRecognizer, error) {
	stmtNodes, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		return nil, errors.Error(errors.SQLParserError, err)
	}
	tableVisitor := GetTableVisitor()
	if _, ok := stmtNodes.Accept(&tableVisitor); !ok {
		return nil, errors.Error(errors.SQLParserError)
	}
	return &MySQLRecognizer{
		StmtNode:   stmtNodes,
		TableAlias: tableVisitor.TableAlias,
		TableName:  tableVisitor.TableName,
	}, nil
}

func (*MySQLRecognizer) GetSQLType() (consts.SQLType, error) {
	return consts.SQLTypeNotSupport, errors.Error(errors.NotSupportSqlTypeError)
}

func (mr *MySQLRecognizer) GetTableAlias() (string, error) {
	return mr.TableAlias, nil
}

func (mr *MySQLRecognizer) GetTableName() (string, error) {
	return mr.TableName, nil
}

func (mr *MySQLRecognizer) GetOriginalSQL() (string, error) {
	return mr.StmtNode.Text(), nil
}

type MySQLWhereRecognizer struct {
	MySQLRecognizer
}

type MySQLSelectForUpdateRecognizer struct {
	MySQLWhereRecognizer
}

//type subqueryChecker struct {
//	text string
//}
//
//// Enter implements ast.Visitor interface.
//func (sc *subqueryChecker) Enter(inNode ast.Node) (outNode ast.Node, skipChildren bool) {
//	if expr, ok := inNode.(*ast.SubqueryExpr); ok {
//		fmt.Println(expr.Query)
//		return inNode, false
//	} else if expr, ok := inNode.(*ast.SelectField); ok {
//		fmt.Println(expr.WildCard)
//		return inNode, false
//	} else if expr, ok := inNode.(*ast.TableName); ok {
//		fmt.Println(expr.Name)
//		return inNode, false
//	} else if expr, ok := inNode.(*ast.TableSource); ok {
//		fmt.Println(expr.AsName)
//		return inNode, false
//	} else if expr, ok := inNode.(*ast.FieldList); ok {
//		for _, f := range expr.Fields {
//			if c, ok := f.Expr.(*ast.ColumnNameExpr); ok {
//				fmt.Println(c.Name)
//			}
//		}
//		fmt.Println()
//		return inNode, false
//	} else if expr, ok := inNode.(*ast.SelectStmt); ok {
//		fmt.Println(expr.From.TableRefs.Left.Text())
//		return inNode, false
//	}
//	return inNode, false
//}
//
//// Leave implements ast.Visitor interface.
//func (sc *subqueryChecker) Leave(inNode ast.Node) (node ast.Node, ok bool) {
//	v := reflect.ValueOf(inNode)
//	fmt.Println(v.Elem().String())
//	return inNode, true
//}
//
////Just debugger
//func main() {
//
//	p := parser.New()
//
//	// 2. Parse a text SQL into AST([]ast.StmtNode).
//	stmtNodes, _, err := p.Parse("select name, age from tbl t where id = 1", "", "")
//
//	stmtNodes[0].Accept(&subqueryChecker{})
//	// 3. Use AST to do cool things.
//	fmt.Println(stmtNodes[0].Text(), err)
//}
