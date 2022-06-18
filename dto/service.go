package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/go-gateway/public"
)

type ServiceListInput struct {
	Info     string `form:"info" json:"info" comment:"关键词"   validate:"" example:""`
	PageNo   int    `form:"page_no" json:"page_no" comment:"每页条数"   validate:"" example:"1"`
	PageSize int    `form:"page_size" json:"page_size" comment:"页数"   validate:"" example:"20"`
}

type ServiceListOutput struct {
	Total int64                   `form:"total" json:"total" comment:"总数"   validate:"" example:""`
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表" example:"" validate:""` //列表
}

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" validate:"required"` //id
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id"`                     //id
	ServiceName string `json:"service_name" form:"service_name"` //服务名称
	ServiceDesc string `json:"service_desc" form:"service_desc"` //服务描述
	LoadType    int    `json:"load_type" form:"load_type"`       //类型
	ServiceAddr string `json:"service_addr" form:"service_addr"` //服务地址
	Qps         int64  `json:"qps" form:"qps"`                   //qps
	Qpd         int64  `json:"qpd" form:"qpd"`                   //qpd
	TotalNode   int    `json:"total_node" form:"total_node"`     //节点数
}

func (s *ServiceListInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, s)
}
func (s *ServiceDeleteInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, s)
}
