package main

import (
	"context"
	"log"

	"github.com/X391757/gorder-v2/common/config"
	"github.com/X391757/gorder-v2/common/genproto/stockpb"
	"github.com/X391757/gorder-v2/common/server"
	"github.com/X391757/gorder-v2/stock/ports"
	"github.com/X391757/gorder-v2/stock/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal("Failed to load config")
	}
}

func main() {
	serviceName := viper.GetString("stock.Service-name")
	serverType := viper.GetString("stock.server-to-run")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)

	switch serverType {
	case "grpc":
		// 注册grpc服务
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		//暂时不用
	default:
		panic("unexpected server type")
	}

}
