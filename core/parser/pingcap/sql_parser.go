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

func (MySQLWhereRecognizer) GetWhereCondition() (string, error) {
	panic("implement me")
}

type MySQLSelectForUpdateRecognizer struct {
	MySQLWhereRecognizer
}
