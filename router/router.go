package router

import (
	"fmt"
	"github.com/createitv/gin-vue/controller/user"
	"github.com/createitv/gin-vue/router/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	userRoute := r.Group("/api/user")
	{
		userRoute.POST("/register", user.Register)
		userRoute.POST("/login", user.Login)
		userRoute.GET("/info", middleware.AuthMiddleware(), user.Info)

	}
	fmt.Println("server listening on 8080")
	err := r.Run()
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
}
