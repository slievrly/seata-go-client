package parser

import "github.com/slievrly/seata-go-client/core/consts"

//The interface Sql recognizer.
type SQLRecognizer interface {
	//Type of the SQL. INSERT/UPDATE/DELETE ...
	GetSQLType() (consts.SQLType, error)
	//TableRecords source related in the SQL, including alias if any.
	//SELECT id, name FROM user u WHERE ...
	//Alias should be 'u' for this SQL.
	GetTableAlias() (string, error)
	//TableRecords name related in the SQL.
	//SELECT id, name FROM user u WHERE ...
	//TableRecords name should be 'user' for this SQL, without alias 'u'.
	GetTableName() (string, error)
	//Return the original SQL input by the upper application.
	GetOriginalSQL() (string, error)
}

//The interface Where recognizer.
type WhereRecognizer interface {
	GetWhereCondition(columnName string) (string, error)
}

//The interface Sql update recognizer.
type SQLUpdateRecognizer interface {
	//Gets update columns.
	GetUpdateColumns() ([]string, error)
	//Gets update values.
	GetUpdateValues() ([]interface{}, error)
}

//The interface Sql select recognizer.
type SQLSelectRecognizer interface {
	WhereRecognizer
}

//The interface Sql insert recognizer.
type SQLInsertRecognizer interface {
	//Gets insert columns.
	GetInsertColumns() ([]string, error)
	//Gets insert rows.
	GetInsertRows() ([][]interface{}, error)
}

//The interface Sql delete recognizer.
type SQLDeleteRecognizer interface {
	WhereRecognizer
}