package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/go-gateway/dto"
	"github.com/noovertime7/go-gateway/middleware"
	"github.com/noovertime7/go-gateway/public"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	admininfo := &AdminController{}
	group.GET("/admin_info", admininfo.AdminInfo)

}

func (a *AdminController) AdminInfo(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sessinfo := sess.Get(public.AdminSessionInfoKey)
	adminsessioninfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessinfo)), adminsessioninfo); err != nil {
		middleware.ResponseError(ctx, 3000, err)
		return
	}
	//1、读取sessionkey对应的json，转换为结构体
	//2、取出数据然后封装输出
	out := &dto.AdminInfoOutput{
		ID:           adminsessioninfo.ID,
		Name:         adminsessioninfo.UserName,
		LoginTime:    adminsessioninfo.LoginTime,
		Avatar:       "",
		Introduction: "我是介绍",
		Rules:        []string{"admin"},
	}
	middleware.ResponseSuccess(ctx, out)
}
