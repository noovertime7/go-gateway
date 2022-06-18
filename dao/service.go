package dao

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HttpRule      *HttpRule      `json:"http" description:"http"`
	TcpRule       *TcpRule       `json:"tcp"`
	GRPCRule      *GrpcRule      `json:"grpc" description:"grpc"`
	LoadBalance   *LoadBalance   `json:"loadbalance"`
	AccessControl *AccessControl `json:"accesscontrol" description:"accesscontrol"`
}
