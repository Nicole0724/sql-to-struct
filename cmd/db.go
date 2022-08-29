package cmd

import (
	"database/sql"
	"fmt"
	"gitee.com/nicole-go-libs/print-colors/color_print"
	_ "github.com/go-sql-driver/mysql"
	"sqltostruct/config"
	"sqltostruct/global"
)

func InitMysql() {
	db, err := sql.Open(config.Conf.SqlConfig.Type, fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Conf.SqlConfig.Account,
		config.Conf.SqlConfig.Password,
		config.Conf.SqlConfig.Host,
		config.Conf.SqlConfig.Port,
		config.Conf.SqlConfig.DataBaseName,
	))

	if err != nil {
		fmt.Println(color_print.Red(fmt.Sprintf("初始化数据连接失败，失败原因：\n%s", err)))
		panic(err)
	}

	//fmt.Println(color_print.Green("数据库连接成功!"))
	global.DBEngine = db
}

func CloseMysql() {
	defer func() {
		_ = global.DBEngine.Close()
	}()
}
