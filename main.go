package main

import (
	"context"
	"gin-sdk/handler"
	"gin-sdk/handler/test"
	"gin-sdk/log"
	"gin-sdk/server"
)

var httphandlers = []handler.Handler{
	&test.TestHanlder{},
}

func main() {
	httpServer := server.NewServer(server.Options{
		Name:    "gin-sdk",
		Address: ":8001",
		LogConf: &log.LogConf{
			Path:     "D:\\var\\log\\gin-sdk",
			ErrorLog: "gin-sdk-error.log",
			ApiLog:   "gin-sdk-api.log",
			PanicLog: "gin-sdk-panic.log",
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
