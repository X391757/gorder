package main

import (
	"context"
	"log"

	"github.com/X391757/gorder-v2/common/config"
	"github.com/X391757/gorder-v2/common/genproto/orderpb"
	"github.com/X391757/gorder-v2/common/server"
	"github.com/X391757/gorder-v2/order/ports"
	"github.com/X391757/gorder-v2/order/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal("Failed to load config")
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")
	//serverType := viper.GetString("stock.server-to-run")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server, svc)
	})
	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{app: application}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
