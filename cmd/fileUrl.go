package cmd

import (
	"os"
	"sqltostruct/config"
)

func InitFileUrl() {
	if err := os.MkdirAll(config.Conf.SaveFilePwd, 0755); err != nil {
		return
	}
}
