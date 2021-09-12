#!/usr/bin/make -f

BRANCH ?= `git rev-parse --abbrev-ref HEAD`
TAG = `git tag | sort -V | tail -1`
COMMIT = `git rev-parse --short HEAD`
DATE = `date`
IMAGE_NAME ?= xxxxx todo:${BRANCH}

.PHONY: all
all: package

.PHONY: test
test: 
	go test -tags=unit -timeout 30s -short -v `go list ./...  | grep -v /vendor/`

.PHONY: init
init: install-lint install-tools

.PHONY: install-lint
install-lint:
	pre-commit install

# only work under go1.16 https://golang.org/doc/go1.16#go-command
.PHONY: install-tools
install-tools:
	cd .. && go get \
    github.com/golang/protobuf/protoc-gen-go@v1.4.3 \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.1 \
	github.com/swaggo/swag/cmd/swag@v1.7.0 \
	gitlab.xiaoduoai.com/base/gomodifytags

.PHONY: package
package: clean build
	docker build -t ${IMAGE_NAME} -f Dockerfile .
	docker push ${IMAGE_NAME}

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'gitlab.xiaoduoai.com/efficiency_engineering/eff-pulsar-proxy/cmd.Version=${TAG}' -X gitlab.xiaoduoai.com/efficiency_engineering/eff-pulsar-proxy/cmd.Commit=${COMMIT} -X 'gitlab.xiaoduoai.com/efficiency_engineering/eff-pulsar-proxy/cmd.Date=${DATE}'" \
	-o ./dist/eff-pulsar-proxy .

.PHONY: release
release: 
	goreleaser

.PHONY: clean	
clean:
	rm -rf ./dist;

.PHONY: swagger
swagger:
	swag init -g doc.go

.PHONY: proto-gen
proto-gen: install-tools
	# this ensure no error hanppens when run protoc generate
	go generate ./...