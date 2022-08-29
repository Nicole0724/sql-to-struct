package model

import (
	"fmt"
	"sqltostruct/global"
)

type TableColumn struct {
	TableName     string
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}
type TableNameColumn struct {
	TableName string
}

// DBTypeToStructType 数据库字段类型 转 go 数据字段类型
var DBTypeToStructType = map[string]string{
	"int":       "int32",
	"tinyint":   "int",
	"smallint":  "int",
	"mediumint": "int64",
	"bigint":    "int64",
	"bit":       "int",
	"bool":      "bool",
	"enum":      "string",
	"set":       "string",
	"varchar":   "string",
	"text":      "string",
	"datetime":  "string",
}

func GetColumnsByTableName(databaseName string, tableName string) ([]*TableColumn, error) {
	jq := fmt.Sprintf("SELECT TABLE_NAME,COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE,COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA= '%s' and TABLE_NAME = '%s'", databaseName, tableName)
	rows, err := global.DBEngine.Query(jq)

	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	defer rows.Close()

	var res []*TableColumn

	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.TableName, &column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		res = append(res, &column)
	}

	return res, nil
}

func GetColumnsByDatabaseName(databaseName string) ([]*TableNameColumn, error) {
	jq := fmt.Sprintf("SELECT TABLE_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA= '%s' group by TABLE_NAME", databaseName)
	rows, err := global.DBEngine.Query(jq)

	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, err
	}
	defer rows.Close()

	var res []*TableNameColumn

	for rows.Next() {
		var column TableNameColumn
		err := rows.Scan(&column.TableName)
		if err != nil {
			return nil, err
		}
		res = append(res, &column)
	}

	return res, nil
}
