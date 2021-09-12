package server

import (
	"github.com/twwch/gin-sdk/twlog"
)

type Options struct {
	Name       string
	Address    string
	LogConf    *twlog.LogConf
	WithExport *ExportOptions
}

type ExportOptions struct {
	IsExport      bool
	ExportHandler interface{}
}
