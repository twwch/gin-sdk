package main

import (
	"context"
	"github.com/twwch/gin-sdk/handler"
	"github.com/twwch/gin-sdk/handler/test"
	"github.com/twwch/gin-sdk/server"
	"github.com/twwch/gin-sdk/twlog"
)

var httphandlers = []handler.Handler{
	&test.TestHanlder{},
}

func main() {
	httpServer := server.NewServer(server.Options{
		Name:    "github.com/twwch/gin-sdk",
		Address: ":8001",
		LogConf: &twlog.LogConf{
			Path:     "D:\\var\\twlog\\gin-sdk",
			ErrorLog: "gin-sdk-error.twlog",
			ApiLog:   "gin-sdk-api.twlog",
			PanicLog: "gin-sdk-panic.twlog",
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
