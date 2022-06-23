package dao

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HttpRule      *HttpRule      `json:"http_rule" description:"http"`
	TcpRule       *TcpRule       `json:"tcp"`
	GRPCRule      *GrpcRule      `json:"grpc" description:"grpc"`
	LoadBalance   *LoadBalance   `json:"load_balance"`
	AccessControl *AccessControl `json:"access_control" description:"accesscontrol"`
}
