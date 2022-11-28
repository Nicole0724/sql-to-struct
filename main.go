package main

import (
	"fmt"
	"sqltostruct/cmd"
	"sqltostruct/config"
	"sqltostruct/global"
	"sqltostruct/internal/service"
)

func main() {
	isEnd := userScanLn()
	if isEnd {
		var quit int
		fmt.Println("配置完毕请按任意键退出。。。")
		fmt.Scanln(&quit)
	}
	service.SqlTwoStruct()

	// 如果发生错误则关闭数据库连接
	defer func() {
		_ = global.DBEngine.Close()
	}()

}

func init() {
	config.GetConfToolsInit()
	cmd.InitFileUrl()
	cmd.InitMysql()
}

func userScanLn() bool {
	var isCustom int
	fmt.Println("请输入是否直接执行配置文件 1:是 2:不是：")
	fmt.Scanln(&isCustom)
	if isCustom == 1 {
		return true
	}
	var isAll int
	fmt.Println("是否按数据库输出 1:整库 2:按表：")
	fmt.Scanln(&isAll)
	if isAll == 1 {
		config.Conf.IsAllDataBase = true
		var tableName string
		fmt.Println("请输入表名模糊查询回车确认【直接回车表示转换所有表】：")
		fmt.Scanln(&tableName)
		config.Conf.LikeTableName = tableName
		return true
	} else {
		config.Conf.IsAllDataBase = false
		var tableName string
		fmt.Println("请输入表名回车确认：")
		fmt.Scanln(&tableName)
		config.Conf.TableName = tableName
		return true
	}

}
