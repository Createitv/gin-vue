package user

import (
	dao "github.com/createitv/gin-vue/dao/user"
	"github.com/createitv/gin-vue/db"
	model "github.com/createitv/gin-vue/model/user"
	"github.com/createitv/gin-vue/pkg/response"
	"github.com/createitv/gin-vue/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	// 获取参数
	DB := db.GetDB()
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(phone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 如果名字没有给一个随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	log.Printf(name, phone, password)
	// 判断手机号是否存在
	if dao.IsTelephoneExist(DB, phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "mgs": "用户已经存在"})
		return
	}
	// 创建用户
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")

		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "mgs": "加密错误"})
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	response.Success(ctx, nil, "注册成功")

}

func Login(ctx *gin.Context) {
	DB := db.GetDB()
	// 获取参数
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	//	数据验证
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "mgs": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//	判断手机号是否存在
	var user model.User
	DB.Where("phone = ? ", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//	判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发放token
	toke, err := utils.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "token生成失败"})
		log.Printf("token generation error: %v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": toke}, "登录成功")

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200, "data": gin.H{"user": dao.ToUserDao(user.(model.User))}, "msg": "获取成功",
	})
}
