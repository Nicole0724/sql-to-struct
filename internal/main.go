package internal

import (
	"fmt"
	"log"
	"sqltostruct/config"
)

func SqlTwoStruct() {
	var conf = config.Conf

	// 获取表信息
	tableCol, _ := GetColumns(conf.SqlConfig.DataBaseName, conf.TableName)

	// 使用 templete
	tlp := NewStructTemplate()
	structCol := tlp.AssemblyColumns(tableCol)
	err := tlp.Generate(conf.PackageName, conf.TableName, structCol, conf.SaveFilePwd, conf.SaveFileName)
	if err != nil {
		log.Fatalf("template.Generate err: %v", err)
	}

	println(fmt.Sprintf("%v", structCol))
}
