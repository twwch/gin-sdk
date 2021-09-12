package middles

import "github.com/gin-gonic/gin"

type requestKey struct{}
type responseKey struct{}
type ginContextKey struct{}
type Middle = gin.HandlerFunc

var (
	RequestKey    = requestKey{}
	ResponseKey   = responseKey{}
	GinContextKey = ginContextKey{}
)
