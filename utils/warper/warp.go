package warper

import (
	"context"
	"gin-sdk/middles"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
)

type HandlerFunc func(c *gin.Context) error

type APIException struct {
	HttpStatus int         `json:"-"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type RespStruct struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func (e *APIException) Error() string {
	return e.Message
}

func newAPIException(httpStatus int, code int, msg string) *APIException {
	return &APIException{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    msg,
	}
}

const (
	ParameterError = 1 // 参数错误
	ServerError    = 2 // 系统错误
	AuthError      = 4 // 认证错误
)

func NewParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, ParameterError, message)
}

func NewAuthError(message string) *APIException {
	return newAPIException(http.StatusUnauthorized, AuthError, message)
}

func NewServerError(message string) *APIException {
	return newAPIException(http.StatusOK, ServerError, message)
}

func CreateHandlerFunc(method interface{}, isExport bool) gin.HandlerFunc {
	return createHandlerFuncWithLogger(method, isExport)
}

func createHandlerFuncWithLogger(method interface{}, isExport bool) gin.HandlerFunc {
	mV := reflect.ValueOf(method)
	if !mV.IsValid() {
		panic(errors.Errorf("method(%s) not found", method))
	}
	mT := mV.Type()
	if mT.NumIn() != 2 {
		panic(errors.Errorf("method(%s) must has 2 ins", method))
	}
	reqT := mT.In(1).Elem()
	if mT.NumOut() != 2 && !isExport {
		panic(errors.Errorf("method(%s) must has 2 out", method))
	}
	return func(c *gin.Context) {
		var (
			ctx = c.Request.Context()
			req = reflect.New(reqT)
			err error
		)

		if c.Request.Method == http.MethodGet && c.Request.Body == http.NoBody {
			err = c.BindQuery(req.Interface())
		} else {
			err = c.ShouldBind(req.Interface())
		}

		if err != nil {
			var apiException *APIException
			apiException = NewParameterError(err.Error())
			c.JSON(apiException.HttpStatus, apiException)
			return
		}

		ctx = context.WithValue(ctx, middles.RequestKey, c.Request)
		ctx = context.WithValue(ctx, middles.ResponseKey, c.Writer)
		ctx = context.WithValue(ctx, middles.GinContextKey, c)

		results := mV.Call([]reflect.Value{reflect.ValueOf(ctx), req})
		var response, e interface{}
		if !isExport {
			response, e = results[0].Interface(), results[1].Interface()
		}else{
			e = results[0].Interface()
		}
		if e != nil {
			var apiException *APIException
			if h, ok := e.(*APIException); ok {
				apiException = h
			} else if e, ok := e.(error); ok {
				apiException = NewServerError(e.Error())
			} else {
				apiException = NewServerError("unknown error")
			}
			apiException.Data = response
			c.JSON(apiException.HttpStatus, apiException)

			return
		}
		if isExport {
			return
		}
		ret := getRetData(response)
		c.PureJSON(http.StatusOK, ret)
	}
}

func getRetData(value interface{}) interface{} {
	ret := &RespStruct{Code: 0, Data: value}
	return ret
}
