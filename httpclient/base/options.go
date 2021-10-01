package base

import (
	"net/http"
	"time"
)

type Options struct {
	Client           *http.Client
	Host             string
	Timeout          time.Duration
	RetryCount       int           // 重试次数
	RetryWaitTime    time.Duration // 重试间隔等待时间
	RetryMaxWaitTime time.Duration // 重试间隔最大等待时间
	Headers          map[string]string
}

type Option func(*Options)

func SetHost(address string) Option {
	return func(options *Options) {
		options.Host = address
	}
}
func SetHeaders(address string) Option {
	return func(options *Options) {
		options.Host = address
	}
}

func SetTimeout(timeout time.Duration) Option {
	return func(options *Options) {
		options.Timeout = timeout
	}
}

func SetRetryCount(retryCount int) Option  {
	return func(options *Options) {
		options.RetryCount = retryCount
	}
}

func SetRetryWaitTime(retryWaitTime time.Duration) Option  {
	return func(options *Options) {
		options.RetryWaitTime = retryWaitTime
	}
}

func SetRetryMaxWaitTime(retryMaxWaitTime time.Duration) Option  {
	return func(options *Options) {
		options.RetryMaxWaitTime = retryMaxWaitTime
	}
}

func newOptions(opts ...Option) *Options {
	options := &Options{
		Host:             "",
		Timeout:          3 * time.Second,
		RetryCount:       3,
		RetryWaitTime:    time.Duration(100) * time.Millisecond,
		RetryMaxWaitTime: time.Duration(2000) * time.Millisecond,
	}
	for _, opt := range opts {
		opt(options)
	}
	return options
}
