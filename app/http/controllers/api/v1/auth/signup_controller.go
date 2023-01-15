// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gin-gorm/app/http/controllers/api/v1"
	"gin-gorm/app/models/user"
	"gin-gorm/app/requests"
	"gin-gorm/pkg/response"
	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
