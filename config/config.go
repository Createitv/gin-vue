package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	configFile := workDir + "\\config"
	fmt.Println(workDir + "\\config")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFile)
	viper.AddConfigPath(".")
	//读取配置文件内容
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("host", viper.GetString("datasource.host"))
	fmt.Println("port", viper.GetString("datasource.port"))
}
