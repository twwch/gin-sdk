package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"runtime"
	"time"
)

// https://github.com/go-resty/resty/
type HttpClient interface {
	GetKernel() (kernel *resty.Client)

	NewRequest(ctx context.Context) (request *resty.Request)

	PostJSON(ctx context.Context, path string, req, resp interface{}) (err error)
	Put(ctx context.Context, path string, req, resp interface{}) (err error)
	Patch(ctx context.Context, path string, req, resp interface{}) (err error)
	Delete(ctx context.Context, path string, req, resp interface{}) (err error)
	Get(ctx context.Context, path string, req map[string]string, resp interface{}) (err error)
}

type client struct {
	kernel  *resty.Client
	options *Options
}

func NewClient(opts ...Option) (HttpClient, error) {
	return NewClientWithOptions(newOptions(opts...))
}

func NewClientWithOptions(options *Options) (HttpClient, error) {
	var c *client
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
	c = &client{
		kernel:  resty.NewWithClient(&http.Client{Transport: transport}),
		options: options,
	}
	c.kernel.SetHostURL(c.options.Host)
	c.kernel.SetTimeout(c.options.Timeout)
	c.kernel.SetRetryCount(c.options.RetryCount)
	c.kernel.SetRetryWaitTime(c.options.RetryWaitTime)
	c.kernel.SetRetryMaxWaitTime(c.options.RetryMaxWaitTime)
	if c.options.Headers != nil {
		c.kernel.SetHeaders(c.options.Headers)
	} else {
		c.kernel.SetHeader("Accept", "application/json")
	}
	return c, nil
}

func (c *client) NewRequest(ctx context.Context) *resty.Request {
	rr := c.kernel.NewRequest().SetContext(ctx)
	return rr
}

func (c *client) PostJSON(ctx context.Context, path string, req, resp interface{}) error {
	r := c.kernel.NewRequest().SetContext(ctx)
	if resp != nil {
		r.SetResult(resp)
	}
	if req != nil {
		r.SetBody(req)
	}
	rsp, err := r.Post(path)
	if err != nil {
		return fmt.Errorf("post:{%s} param:%+v err: %v", r.URL, req, err)
	}
	httpStatusCode := rsp.StatusCode()
	if httpStatusCode < http.StatusOK || httpStatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("post:{%s} param:%+v failed: %v", rsp.Request.URL, req, httpStatusCode)
	}
	return nil
}

func (c *client) Put(ctx context.Context, path string, req, resp interface{}) error {
	r := c.kernel.NewRequest().SetContext(ctx)
	if resp != nil {
		r.SetResult(resp)
	}
	if req != nil {
		r.SetBody(req)
	}
	rsp, err := r.Put(path)
	if err != nil {
		return fmt.Errorf("put:{%s} param:%+v err: %v", r.URL, req, err)
	}
	httpStatusCode := rsp.StatusCode()
	if httpStatusCode < http.StatusOK || httpStatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("put:{%s} param:%+v failed: %v", rsp.Request.URL, req, httpStatusCode)
	}
	return nil
}

func (c *client) Patch(ctx context.Context, path string, req, resp interface{}) error {
	r := c.kernel.NewRequest().SetContext(ctx)
	if resp != nil {
		r.SetResult(resp)
	}
	if req != nil {
		r.SetBody(req)
	}
	rsp, err := r.Patch(path)
	if err != nil {
		return fmt.Errorf("patch:{%s} param:%+v err: %v", r.URL, req, err)
	}
	httpStatusCode := rsp.StatusCode()
	if httpStatusCode < http.StatusOK || httpStatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("patch:{%s} param:%+v failed: %v", rsp.Request.URL, req, httpStatusCode)
	}
	return nil
}

func (c *client) Delete(ctx context.Context, path string, req, resp interface{}) error {
	r := c.kernel.NewRequest().SetContext(ctx)
	if resp != nil {
		r.SetResult(resp)
	}
	if req != nil {
		r.SetBody(req)
	}
	rsp, err := r.Delete(path)
	if err != nil {
		return fmt.Errorf("delete:{%s} param:%+v err: %v", r.URL, req, err)
	}
	httpStatusCode := rsp.StatusCode()
	if httpStatusCode < http.StatusOK || httpStatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("delete:{%s} param:%+v failed: %v", rsp.Request.URL, req, httpStatusCode)
	}
	return nil
}

func (c *client) Get(ctx context.Context, path string, req map[string]string, resp interface{}) error {
	r := c.kernel.NewRequest().SetContext(ctx)
	if resp != nil {
		r.SetResult(resp)
	}
	if req != nil {
		r.SetQueryParams(req)
	}
	rsp, err := r.Get(path)
	if err != nil {
		return fmt.Errorf("get:{%s} err: %v", r.URL, err)
	}
	if !rsp.IsSuccess() {
		return errors.New(fmt.Sprintf("check getResp.IsSuccess() failed, code is %v", rsp.StatusCode()))
	}
	return nil
}

func (c *client) GetKernel() *resty.Client {
	return c.kernel
}
