// Package user 存放用户 Model 相关逻辑
package user

import "gin-gorm/app/models"

type User struct {
	models.BaseModel

	Name     string `json:name,omitempty`
	Email    string `json:"_"`
	Phone    string `json:"_"`
	Password string `json:"_"`

	models.CommonTimestampsField
}
