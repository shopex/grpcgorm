package service

import (
	lb "git.shopex.cn/luban/go-sdk"
	"google.golang.org/grpc"
)

func RegisterGrpcService(gs *grpc.Server) {
	lb.RegisterDealerServer(gs, &Dealer{})
}
