package api

import (
	"go.uber.org/zap"

	"study_gva/global"
	"study_gva/model/common/response"
	email_response "study_gva/plugin/email/model/response"
	"study_gva/plugin/email/service"

	"github.com/gin-gonic/gin"
)

type EmailApi struct{}

// 发送测试邮件
func (s *EmailApi) EmailTest(c *gin.Context) {
	err := service.ServiceGroupApp.EmailTest()
	if err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

// 发送邮件
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email email_response.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.SendEmail(email.To, email.Subject, email.Body)
	if err != nil {
		global.GVA_LOG.Error("发送失败！", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}
