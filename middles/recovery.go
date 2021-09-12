package middles

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func Recovery() Middle {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Panicf("panic recovered: err = %v, stack = %s", err, debug.Stack())
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
