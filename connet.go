package sqltostruct

import (
	"database/sql"
	"fmt"
	"gitee.com/nicole-go-libs/print-colors/color_print"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBEngine *sql.DB
)

func InitMysql(config SqlConnectConfig) {
	//db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/go_frame?charset=utf8&parseTime=True&loc=Local")
	db, err := sql.Open(config.Type, fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Account,
		config.Password,
		config.Host,
		config.Port,
		config.DataBaseName,
	))

	if err != nil {
		fmt.Println(color_print.Red(fmt.Sprintf("初始化数据连接失败，失败原因：\n%s", err)))
		panic(err)
	}

	DBEngine = db
}

func CloseMysql() {
	defer func() {
		_ = DBEngine.Close()
	}()
}
