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

type ServiceStatOutput struct {
	Today     []int64 `form:"today" json:"today" comment:"今日流量"   validate:"" example:""`
	Yesterday []int64 `form:"yesterday" json:"yesterday" comment:"昨日流量"   validate:"" example:""`
}

type ServiceADDInput struct {
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required,vaild_service_name" example:""`
	ServiceDesc string `form:"service_desc" json:"service_desc" comment:"接入类型"   validate:"required" example:""`

	RuleType       int    `form:"rule_type" json:"rule_type" comment:"接入路径"   validate:"" example:"max=3,min=0"`
	Rule           string `form:"rule" json:"rule" comment:"接入路径:域名 前缀"   validate:"required,vaild_rule" example:""`
	NeedHttps      int    `form:"need_https" json:"need_https" comment:"是否支持https"   validate:"" example:"max=1,min=0"`
	NeedStripUri   int    `form:"need_strip_uri" json:"need_strip_uri" comment:"是否支持strip uri"   validate:"max=1,min=0" example:""`
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"need_websocket" validate:"max=1,min=0"`
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"url重写功能" validate:"vaild-urlwrite"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header_transfor" validate:"vaild-header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" validate:""`
	WhiteList         string `json:"white_list" form:"white_list" comment:"white_list" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单Ip" validate:""`
	ClientIpFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:"max=3,min=0"`                                //轮询方式
	IpList                 string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"127.0.0.1:80" validate:"required,valid_iplist"`                //ip列表
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"50" validate:"required,valid_weightlist"`             //权重列表
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时, 单位s" example:"" validate:"min=0"`   //建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时, 单位s" example:"" validate:"min=0"` //获取header超时, 单位s
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间, 单位s" example:"" validate:"min=0"`       //链接最大空闲时间, 单位s
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"" validate:"min=0"`                     //最大空闲链接数
}
type ServiceAddTcpInput struct {
	ServiceName    string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc    string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port           int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header头转换" validate:"
"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

type ServiceUpdateTcpInput struct {
	ID                int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

type ServiceUpdateInput struct {
	ID          int64  `json:"id" form:"id" comment:"服务ID" example:"62" validate:"min=1"`
	ServiceName string `form:"service_name" json:"service_name" comment:"服务名"   validate:"required,vaild_service_name" example:""`
	ServiceDesc string `form:"service_desc" json:"service_desc" comment:"服务描述"   validate:"required" example:""`

	RuleType       int    `form:"rule_type" json:"rule_type" comment:"接入类型"   validate:"" example:"max=3,min=0"`
	Rule           string `form:"rule" json:"rule" comment:"接入路径:域名 前缀"   validate:"required,vaild_rule" example:""`
	NeedHttps      int    `form:"need_https" json:"need_https" comment:"是否支持https"   validate:"" example:"max=1,min=0"`
	NeedStripUri   int    `form:"need_strip_uri" json:"need_strip_uri" comment:"是否支持strip uri"   validate:"max=1,min=0" example:""`
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"need_websocket" validate:"max=1,min=0"`
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"url重写功能" validate:"vaild-urlwrite"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"header_transfor" validate:"vaild-header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" validate:""`
	WhiteList         string `json:"white_list" form:"white_list" comment:"white_list" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单Ip" validate:""`
	ClientIpFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式" example:"" validate:"max=3,min=0"`                                //轮询方式
	IpList                 string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"127.0.0.1:80" validate:"required,valid_iplist"`                //ip列表
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"50" validate:"required,valid_weightlist"`             //权重列表
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时, 单位s" example:"" validate:"min=0"`   //建立连接超时, 单位s
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时, 单位s" example:"" validate:"min=0"` //获取header超时, 单位s
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间, 单位s" example:"" validate:"min=0"`       //链接最大空闲时间, 单位s
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数" example:"" validate:"min=0"`                     //最大空闲链接数
}

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" validate:""` //id
}

type ServiceAddGrpcInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

type ServiceUpdateGrpcInput struct {
	ID                int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_iplist"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_iplist"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ipportlist"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weightlist"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_iplist"`
}

func (params *ServiceAddGrpcInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *ServiceUpdateGrpcInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (s *ServiceListInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, s)
}
func (s *ServiceDeleteInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, s)
}
func (s *ServiceADDInput) BindValidParm(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, s)
}
func (s *ServiceAddTcpInput) GetValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, s)
}
func (s *ServiceUpdateTcpInput) GetValidParams(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, s)
}
