package sqltostruct

import (
	"fmt"
	"log"
)

type SqlConnectConfig struct {
	Type         string
	Host         string
	Port         string
	Account      string
	Password     string
	DataBaseName string
}

type SqlToStructParams struct {
	SqlConfig    SqlConnectConfig
	PackageName  string
	TableName    string
	SaveFileName string
	SaveFilePwd  string
}

func (s *SqlToStructParams) SqlTwoStruct() {
	InitMysql(s.SqlConfig)
	// 获取表信息
	tableCol, _ := GetColumns(s.SqlConfig.DataBaseName, s.TableName)

	// 使用 templete
	tlp := NewStructTemplate()
	structCol := tlp.AssemblyColumns(tableCol)
	err := tlp.Generate(s.PackageName, s.TableName, structCol, s.SaveFilePwd, s.SaveFileName)
	if err != nil {
		log.Fatalf("template.Generate err: %v", err)
	}

	println(fmt.Sprintf("%v", structCol))
}
