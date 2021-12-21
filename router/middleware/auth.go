package middleware

import (
	"github.com/createitv/gin-vue/db"
	model "github.com/createitv/gin-vue/model/user"
	"github.com/createitv/gin-vue/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取Authorization header
		tokenString := ctx.GetHeader("Authorization")
		//	验证token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort() // 放弃这次请求
			return
		}

		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := utils.ParserToken(tokenString)
		if err != nil || !token.Valid {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort() // 放弃这次请求
			return
		}
		//	验证通过后获取claims中的userId
		userId := claims.UserID
		DB := db.GetDB()
		var user model.User
		DB.First(&user, userId)

		//	 用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort() // 放弃这次请求
			return
		}
		//	用户存在，讲user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
