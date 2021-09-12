package server

import (
	"gin-sdk/log"
)

type Options struct {
	Name       string
	Address    string
	LogConf    *log.LogConf
	WithExport *ExportOptions
}

type ExportOptions struct {
	IsExport      bool
	ExportHandler interface{}
}
