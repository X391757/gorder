package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTPServer(servicename string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub("order").GetString("http-addr")
	RunHTTPServerOnAddr(addr, wrapper)
}

func RunHTTPServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api").GET("ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}
