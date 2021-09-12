package test

import (
	"context"
	"errors"
	"fmt"
	"gin-sdk/constant"
	"gin-sdk/handler"
	"gin-sdk/middles"
	"gin-sdk/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"net/http"
	"net/url"
	"time"
)

type TestHanlder struct {
	handler.Base
}

func (h *TestHanlder) Init(ginRouter *gin.RouterGroup) {
	h.Logger = log.WithField("handler", "test")
	// registry http handler
	if ginRouter != nil {
		appGroup := ginRouter.Group("/test")
		server.Route(appGroup, http.MethodGet, "/", h.Test, server.Options{
			WithExport: &server.ExportOptions{
				IsExport:      true,
				ExportHandler: h.Export,
			},
		})
		//appGroup.GET("/", warper.CreateHandlerFunc(h.Test, false))
	}
}

type Req struct {
	Code int `json:"code" form:"code"`
}

type Resp struct {
	MyData map[string]interface{} `json:"my_data"`
}

func (h *TestHanlder) Test(ctx context.Context, req *Req) (*Resp, error) {
	h.Logger.Info(req)
	return &Resp{
		MyData: map[string]interface{}{"xxx": 1245},
	}, nil
}

func (h *TestHanlder) Export(ctx context.Context, req *Req) error {
	h.Logger.Info(req)
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("续费数据汇总")
	row := sheet.AddRow()
	row.AddCell().Value = ""
	row.AddCell().Value = "总数"
	row.AddCell().Value = "KA"
	row.AddCell().Value = "SMB"
	row.AddCell().Value = "SMALL"
	w, ok := ctx.Value(middles.ResponseKey).(gin.ResponseWriter)
	if !ok {
		return errors.New("no ResponseWriter")
	}
	var filename = url.QueryEscape(fmt.Sprintf("xxxx-%s.xlsx", time.Now().Format(constant.FullDateTimeFormatNotSpace)))
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	err := file.Write(w)
	return err
}
