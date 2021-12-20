package main

import (
	"github.com/createitv/gin-vue/cmd"
	"os"
)

func main() {
	//config.InitConfig()
	//_,err := db.InitDb()
	//if err != nil {
	//	fmt.Printf("datasource init failed: %v", err)
	//}
	//router.Run()

	if err := cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}
