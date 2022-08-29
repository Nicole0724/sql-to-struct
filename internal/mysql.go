package internal

import (
	"fmt"
	"sqltostruct/global"
)

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
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

func GetColumns(databaseName string, tableName string) ([]*TableColumn, error) {
	jq := fmt.Sprintf("SELECT COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE,COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA= '%s' and TABLE_NAME = '%s'", databaseName, tableName)
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
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		res = append(res, &column)
	}

	return res, nil
}
