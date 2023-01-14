package auth

import (
	v1 "gin-gorm/app/http/controllers/api/v1"
	"gin-gorm/pkg/captcha"
	"gin-gorm/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生产验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 对等级的日志
	logger.LogIf(err)
	// 返回给用户
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
