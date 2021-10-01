package base

import (
	"context"
	"fmt"
	"log"
	"testing"
)

var clientDemo HttpClient
var err error

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Init() {
	clientDemo, err = NewClient(
		SetHost("http://127.0.0.1:9002"),
		SetRetryCount(5),
	)
	if err != nil {
		panic(err)
	}
}

func TestClient_Get(t *testing.T) {
	Init()
	ctx := context.Background()
	resp := &Resp{}

	err := clientDemo.Get(ctx, "/v1/test/get", map[string]string{"code": "123"}, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Code, resp.Message, resp.Data)
}

func TestClient_Post(t *testing.T) {
	Init()
	ctx := context.Background()
	resp := &Resp{}
	err := clientDemo.PostJSON(ctx, "/v1/test/post", map[string]interface{}{"code": 12346}, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Code, resp.Message, resp.Data)
}

func TestClient_Put(t *testing.T) {
	Init()
	ctx := context.Background()
	resp := &Resp{}

	err := clientDemo.Put(ctx, "/v1/test/put", map[string]interface{}{"code": 123}, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Code, resp.Message, resp.Data)
}

func TestClient_Patch(t *testing.T) {
	Init()
	ctx := context.Background()
	resp := &Resp{}

	err := clientDemo.Patch(ctx, "/v1/test/patch", map[string]interface{}{"code": 123}, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Code, resp.Message, resp.Data)
}
