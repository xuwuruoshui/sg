package config

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"log"
	"os"
	"path/filepath"
)

const configFile = "/conf/conf.ini"

var File *goconfig.ConfigFile

func init() {
	// 获取当前工作目录的路径
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath := filepath.Join(currentDir, configFile)

	if !fileExist(configPath) {
		panic("文件不存在")
	}

	// 可以传参任意路径读取
	if len(os.Args) > 1 {
		dir := os.Args[1]
		if dir != "" {
			configPath = dir + "/" + configFile
		}
	}

	// 文件系统的读取
	File, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		log.Fatal("读取配置文件出错:", err)
	}
}

// 判断文件是否存在
func fileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func A() {
	fmt.Println()
}
