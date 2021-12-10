package model

import (
	"github.com/jinzhu/gorm"
	"thsit.com/ginessential/common"
)

// 定一个结构体（model)
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
}

// 入库
func SaveUser(newUser User) {
	common.GetDb().Create(&newUser)
}

// 根据手机号查询
func FindUserByPhone(phone string) User {
	var user User
	common.GetDb().Where("telephone = ?", phone).First(&user)
	return user
}

// 检验手机号是否存在
func IsTelephoneExits(telephone string) bool {
	var user User
	common.GetDb().Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
