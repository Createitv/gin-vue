package router

import (
	"fmt"
	"github.com/createitv/gin-vue/controller/user"
	"github.com/createitv/gin-vue/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.POST("/api/user/register", user.Register)
	r.POST("/api/user/login", user.Login)
	r.GET("api/user/info", middleware.AuthMiddleware(), user.Info)
	fmt.Println("server listening on 8080")
	err := r.Run()
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
}
