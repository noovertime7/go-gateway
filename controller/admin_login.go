package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/go-gateway/dao"
	"github.com/noovertime7/go-gateway/dto"
	"github.com/noovertime7/go-gateway/middleware"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)

}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /admin_login/login [post]
func (a *AdminLoginController) AdminLogin(ctx *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(ctx, tx, params)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	out := &dto.AdminLoginOut{Token: admin.UserName}
	middleware.ResponseSuccess(ctx, out)
}
