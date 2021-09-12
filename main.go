package main

import (
	"context"
	"github.com/twwch/gin-sdk/handler"
	"github.com/twwch/gin-sdk/handler/test"
	"github.com/twwch/gin-sdk/log"
	"github.com/twwch/gin-sdk/server"
)

var httphandlers = []handler.Handler{
	&test.TestHanlder{},
}

func main() {
	httpServer := server.NewServer(server.Options{
		Name:    "github.com/twwch/gin-sdk",
		Address: ":8001",
		LogConf: &log.LogConf{
			Path:     "D:\\var\\log\\github.com/twwch/gin-sdk",
			ErrorLog: "github.com/twwch/gin-sdk-error.log",
			ApiLog:   "github.com/twwch/gin-sdk-api.log",
			PanicLog: "github.com/twwch/gin-sdk-panic.log",
		},
	})
	router := httpServer.GetEngine()
	v1 := router.Group("/v1")
	{
		for _, hd := range httphandlers {
			hd.Init(v1)
		}
	}
	_ = httpServer.Run(context.Background())
}
