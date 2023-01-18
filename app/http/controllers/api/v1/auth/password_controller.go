package auth

import (
	v1 "gin-gorm/app/http/controllers/api/v1"
	"gin-gorm/app/models/user"
	"gin-gorm/app/requests"
	"gin-gorm/pkg/response"
	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	// 2. 更新密码
	userModel := user.GetByPhone(request.Phone)

	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
