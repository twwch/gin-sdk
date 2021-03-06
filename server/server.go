package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/twwch/gin-sdk/middles"
	myLog "github.com/twwch/gin-sdk/twlog"
	"github.com/twwch/gin-sdk/utils/warper"
	"net/http"
)

type Server interface {
	Name() (name string)
	Run(ctx context.Context) (err error)
	AddMiddles(ms ...middles.Middle)
	GetEngine() (kernel *gin.Engine)
}

type server struct {
	engine  *gin.Engine
	options Options
}

func NewServer(options Options) Server {
	engine := gin.New()
	myLog.NewLog(options.LogConf)
	engine.Use(
		middles.Recovery(),
	)
	s := &server{engine: engine, options: options}
	return s
}

func (s *server) Name() string {
	return s.options.Name
}

func (s *server) Run(ctx context.Context) error {
	log.Infof("http server port is %v", s.options.Address)
	srv := &http.Server{Handler: s.engine, Addr: s.options.Address}
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("httpServer ListenAndServe err:%v", err)
		}
		log.Error("http server is closed")
	}
	return nil
}

func (s *server) AddMiddles(ms ...middles.Middle) {
	s.engine.Use(ms...)
}

func (s *server) GetEngine() *gin.Engine {
	return s.engine
}

// 统一配置导出接口
func Route(routes gin.IRoutes, method string, path string, function interface{}, options ...Options) {
	for _, op := range options {
		if op.WithExport.IsExport {
			routes.Handle(method, fmt.Sprintf("%v/export", path), warper.CreateHandlerFunc(op.WithExport.ExportHandler, true))
		}
	}
	routes.Handle(method, path, warper.CreateHandlerFunc(function, false))
}
