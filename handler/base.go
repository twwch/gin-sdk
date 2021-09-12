package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler interface {
	Init(ginRouter *gin.RouterGroup)
}

type Base struct {
	Logger *log.Entry
}
