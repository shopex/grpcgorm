package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"

	"git.shopex.cn/luban/lib"

	"git.shopex.cn/luban/grpcgorm/config"
	"git.shopex.cn/luban/grpcgorm/model"
	"git.shopex.cn/luban/grpcgorm/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var cmdServer = &lib.Command{
	UsageLine: "server",
	Short:     "服务管理",
	Long: `
    start   启动服务
`,
}

func init() {
	cmdServer.Run = server
}

func server(cmd *lib.Command, args []string) int {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(args) < 1 {
		log.Println("缺少参数")
		return 1
	}
	err := model.ConnectDB(config.Config.DataBase)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	switch args[0] {
	case "start":
		lis, err := net.Listen("tcp", "0.0.0.0:"+config.Config.Port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		gs := grpc.NewServer()
		service.RegisterGrpcService(gs)
		reflection.Register(gs)
		log.Println("start server in port:" + config.Config.Port)
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	case "migration":
		d := model.Db().Set("gorm:table_options", "CHARSET=utf8")
		d.AutoMigrate(
			&model.Invoice{},
		)
		fmt.Println(0)
		return 0
	}
	return 0
}
