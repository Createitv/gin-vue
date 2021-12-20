package dao

import (
	"github.com/createitv/gin-vue/db"
	model "github.com/createitv/gin-vue/model/user"
	"gorm.io/gorm"
)

func GetUserByPhone(phone string) (*model.User, error) {
	user := new(model.User)
	if err := db.DB.Debug().Where("phone=?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func IsTelephoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ? ", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

type UserDao struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserDao(user model.User) UserDao {
	return UserDao{
		Name:  user.Name,
		Phone: user.Phone,
	}

}
