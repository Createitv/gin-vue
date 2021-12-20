package cmd

import (
	"fmt"
	"github.com/createitv/gin-vue/config"
	"github.com/createitv/gin-vue/db"
	"github.com/createitv/gin-vue/router"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		config.InitConfig()
		_, err := db.InitDb()
		if err != nil {
			fmt.Printf("datasource init failed: %v", err)
		}

		router.Run()
		return nil
	}
	return rootCmd.Execute()

}
