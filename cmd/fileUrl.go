package cmd

import (
	"fmt"
	"gitee.com/nicole-go-libs/print-colors/color_print"
	"os"
	"sqltostruct/config"
)

func InitFileUrl() {
	if err := os.MkdirAll(config.Conf.SaveFilePwd, 0755); err != nil {
		fmt.Println(color_print.Red(fmt.Sprintf("初始化文件路径创建失败，失败原因：\n%s", err)))
		panic(err)
	}
}
