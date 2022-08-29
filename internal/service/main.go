package service

import (
	"fmt"
	"log"
	"sqltostruct/config"
	"sqltostruct/internal/model"
	"sqltostruct/internal/templates"
	"sqltostruct/libs/words"
	"time"
)

func SqlTwoStruct() {
	var conf = config.Conf
	if conf.IsAllDataBase {
		sqlTwoStructByTable()
	} else if conf.TableName != "" {
		sqlTwoStructByTableName(conf.TableName)
	} else {
		panic("请配置解析方式")
	}

}

func sqlTwoStructByTable() {
	var conf = config.Conf
	tableList, _ := model.GetColumnsByDatabaseName(conf.SqlConfig.DataBaseName)

	for i, column := range tableList {
		fmt.Println(i)
		fmt.Println(column.TableName)
		/*go func() {
			fmt.Println(i)
			fmt.Println(column.TableName)
		}()*/
		sqlTwoStructByTableName(column.TableName)
		time.Sleep(500 * time.Millisecond) //模拟执行的耗时任务
	}
}

func sqlTwoStructByTableName(tableName string) {
	var conf = config.Conf
	// 获取表信息
	tableCol, _ := model.GetColumnsByTableName(conf.SqlConfig.DataBaseName, tableName)
	saveFile(tableCol, tableName)
}

func saveFile(tableCol []*model.TableColumn, tableName string) {
	var conf = config.Conf

	// 使用 templete
	tlp := templates.NewStructTemplate()
	structCol := tlp.AssemblyColumns(tableCol)
	err := tlp.Generate(conf.PackageName, tableName, structCol, conf.SaveFilePwd, fmt.Sprintf("%sModel", words.ToCamelCase(tableName)))
	if err != nil {
		log.Fatalf("template.Generate err: %v", err)
	}

	//println(fmt.Sprintf("%s", structCol))
}
