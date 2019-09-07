/**
2 * @Author: Nico
3 * @Date: 2019/9/7 19:03
4 */
package pingcap

import (
	"github.com/pingcap/parser/ast"
)

type TableVisitor struct {
	TableName  string
	TableAlias string
}

func (tv *TableVisitor) Enter(inNode ast.Node) (outNode ast.Node, skipChildren bool) {
	if expr, ok := inNode.(*ast.TableName); ok {
		tv.TableName = expr.Name.String()
	} else if expr, ok := inNode.(*ast.TableSource); ok {
		tv.TableAlias = expr.AsName.String()
	}
	return inNode, false
}

func (tv *TableVisitor) Leave(inNode ast.Node) (node ast.Node, ok bool) {
	return inNode, true
}

func GetTableVisitor() TableVisitor {
	return TableVisitor{}
}
