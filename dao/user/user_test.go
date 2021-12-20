package dao

import (
	"fmt"
	"github.com/createitv/gin-vue/db"
	"testing"
)

func TestGetUserByPhone(t *testing.T) {
	DB := db.GetDB()
	if IsTelephoneExist(DB, "18727792922") {
		fmt.Println("true")
	}
}
