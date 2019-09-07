package main

import (
	"fmt"
	"github.com/pingcap/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

//Just debugger
func main() {

	p := parser.New()

	// 2. Parse a text SQL into AST([]ast.StmtNode).
	stmtNodes, _, err := p.Parse("select * from tbl where id = 1", "", "")

	// 3. Use AST to do cool things.
	fmt.Println(stmtNodes[0], err)
}