package consts

//The enum Sql type
type SQLType int

const (
	//No Support
	SQLTypeNotSupport = -1
	//Select sql type.
	SQLTypeSelect = 0
	//Insert sql type.
	SQLTypeInsert = 1
	//Update sql type.
	SQLTypeUpdate = 2
	//Delete sql type.
	SQLTypeDelete = 3
	//Select for update sql type.
	SQLTypeSelectForUpdate = 4
	//Replace sql type.
	SQLTypeReplace = 5
	//Truncate sql type.
	SQLTypeTruncate = 6
	//Create sql type.
	SQLTypeCreate = 7
	//Drop sql type.
	SQLTypeDrop = 8
	//Load sql type.
	SQLTypeLoad = 9
	//Merge sql type.
	SQLTypeMerge = 10
	//Show sql type.
	SQLTypeShow = 11
	//Alter sql type.
	SQLTypeAlter = 12
	//Rename sql type.
	SQLTypeRename = 13
	//Dump sql type.
	SQLTypeDump = 14
	//Debug sql type.
	SQLTypeDebug = 15
	//Explain sql type.
	SQLTypeExplan = 16
	//Stored procedure
	SQLTypeProcedure = 17
	//Desc sql type.
	SQLTypeDesc = 18
	//Select last insert id
	SQLTypeSelectLastInsertId = 19
	//Select without table sql type.
	SQLTypeSelectWithoutTableSqlType = 20
	//Create sequence sql type.
	SQLTypeCreateSequence = 21
	//Show sequences sql type.
	SQLTypeShowSequences = 22
	//Get sequence sql type.
	SQLTypeGetSequence = 23
	//Alter sequence sql type.
	SQLTypeAlterSequence = 24
	//Drop sequence sql type.
	SQLTypeSequence = 25
	//Tddl show sql type.
	SQLTypeTddlShow = 26
	//Set sql type.
	SQLTypeSet = 27
	//Reload sql type.
	SQLTypeReload = 28
	//Select union sql type.
	SQLTypeSelectUnion = 29
	//Create table sql type.
	SQLTypeCreateTable = 30
	//Drop table sql type.
	SQLTypeDropTable = 31
	//Alter table sql type.
	SQLTypeAlterTable = 32
	//Save point sql type.
	SQLTypeSavePoint = 33
	//Select from update sql type.
	SQLTypeSelectFromUpdate = 34
	//multi delete/update
	SQLTypeMultiDelete = 35
	//Multi update sql type.
	SQLTypeMultiUpdate = 36
	//Create index sql type.
	SQLTypeCreateIndex = 37
	//Drop index sql type.
	SQLTypeDropIndex = 38
	//Kill sql type.
	SQLTypeKill = 39
	//Release dblock sql type.
	SQLTypeReleaseDblock = 40
	//Lock tables sql type.
	SQLTypeLockTables = 41
	//Unlock tables sql type.
	SQLTypeUnlockTables = 42
	//Check table sql type.
	SQLTypeCheckTable = 43
	//Select found rows.
	SQLTypeSelectFoundRows = 44
	//Insert ingore sql type.
	SQLTypeInsertIngore = 101
	//Insert on duplicate update sql type.
	SQLTypeInsertOnDuplicateUpdate = 102
)
