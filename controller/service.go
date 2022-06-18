package controller

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/go-gateway/dao"
	"github.com/noovertime7/go-gateway/dto"
	"github.com/noovertime7/go-gateway/middleware"
	"github.com/noovertime7/go-gateway/public"
)

type ServiceController struct{}

func ServiceRegister(group *gin.RouterGroup) {
	Serviceinfo := &ServiceController{}
	group.POST("/service_list", Serviceinfo.ServiceList)
	group.POST("/service_delete", Serviceinfo.ServiceDelete)

}

func (s *ServiceController) ServiceList(ctx *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	// 从数据库中分页读取信息
	serviceinfo := &dao.ServiceInfo{}
	list, total, err := serviceinfo.PageList(ctx, tx, params)
	if err != nil {
		middleware.ResponseError(ctx, 30003, err)
		return
	}
	// 格式化输出信息
	outList := []dto.ServiceListItemOutput{}
	for _, listIterm := range list {
		//1、http后缀接入 clusterIP + clusterPort + path
		//2、http域名接入  domain
		//3、tcp、grpc接入  clusterIP + servicePort
		serviceAddr := "unknow"
		clusterIP := lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort := lib.GetStringConf("base.cluster.cluster_port")
		clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")
		//拿到detail
		serviceDetail, err := listIterm.ServiceDetail(ctx, tx, &listIterm)
		if err != nil {
			middleware.ResponseError(ctx, 30004, err)
			return
		}
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP && serviceDetail.HttpRule.RuleType == public.HTTPRuleTypePrefixURL {
			//需要https
			if serviceDetail.HttpRule.NeedHttps == 1 {
				serviceAddr = fmt.Sprintf("%s:%s", clusterIP, clusterSSLPort+serviceDetail.HttpRule.Rule)
			}
			serviceAddr = fmt.Sprintf("%s:%s", clusterIP, clusterPort+serviceDetail.HttpRule.Rule)
		}
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP && serviceDetail.HttpRule.RuleType == public.HTTPRuleTypeDomain {
			serviceAddr = serviceDetail.HttpRule.Rule
		}
		if serviceDetail.Info.LoadType == public.LoadTypeTCP {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.TcpRule.Port)
		}
		if serviceDetail.Info.LoadType == public.LoadTypeGRPC {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.GRPCRule.Port)
		}
		iplist := serviceDetail.LoadBalance.GetIPListByModle()
		outItem := dto.ServiceListItemOutput{
			ID:          listIterm.ID,
			ServiceName: listIterm.ServiceName,
			ServiceDesc: listIterm.ServiceDesc,
			ServiceAddr: serviceAddr,
			Qpd:         0,
			Qps:         0,
			TotalNode:   len(iplist),
		}
		outList = append(outList, outItem)
	}
	out := &dto.ServiceListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(ctx, out)
}

func (s *ServiceController) ServiceDelete(ctx *gin.Context) {
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParm(ctx); err != nil {
		middleware.ResponseError(ctx, 30001, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 30002, err)
		return
	}
	// 读取基本信息
	serviceinfo := &dao.ServiceInfo{ID: params.ID}
	serviceinfo, err = serviceinfo.Find(ctx, tx, serviceinfo)
	if err != nil {
		middleware.ResponseError(ctx, 30003, err)
		return
	}
	serviceinfo.IsDelete = 1
	if err := serviceinfo.Save(ctx, tx); err != nil {
		middleware.ResponseError(ctx, 30004, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}
