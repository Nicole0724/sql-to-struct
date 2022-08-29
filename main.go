package main

import (
	"sqltostruct/cmd"
	"sqltostruct/config"
	"sqltostruct/global"
	"sqltostruct/internal"
)

func init() {
	config.GetConfToolsInit()
	cmd.InitFileUrl()
	cmd.InitMysql()
}

func main() {
	internal.SqlTwoStruct()

	// 如果发生错误则关闭数据库连接
	defer func() {
		_ = global.DBEngine.Close()
	}()
}
