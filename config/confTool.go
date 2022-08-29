package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var dirRoot = "./"

// GetConfToolsInit 利用init 函数会在主函数执行之前运行的特点，在main函数执行之前读取文件 /**
func GetConfToolsInit() {
	yamlFile := fileReading("config")
	err := yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(fmt.Sprintf("公共配置：%s\n", Conf))
}

func fileReading(fileName string) []byte {
	evnUrl := fmt.Sprintf("%s/%s.yaml", dirRoot, fileName)
	yamlFile, err := ioutil.ReadFile(evnUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	return yamlFile
}
