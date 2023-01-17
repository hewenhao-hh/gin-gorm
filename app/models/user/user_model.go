// Package user 存放用户 Model 相关逻辑
package user

import (
	"gin-gorm/app/models"
	"gin-gorm/pkg/database"
)

type User struct {
	models.BaseModel

	Name     string `json:name,omitempty`
	Email    string `json:"_"`
	Phone    string `json:"_"`
	Password string `json:"_"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
