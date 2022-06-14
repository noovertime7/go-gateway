package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/go-gateway/public"
)

type AdminLoginInput struct {
	UserName string `form:"username" json:"username" comment:"用户名"  validate:"required,is_valid_username" example:"admin"`
	Password string `form:"password" json:"password" comment:"密码"   validate:"required" example:"123456"`
}

type AdminLoginOut struct {
	Token string `form:"token" json:"token" comment:"token"  example:"token"`
}

// BindValidParm 绑定并校验参数

func (a *AdminLoginInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
