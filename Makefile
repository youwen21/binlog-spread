# 项目数据模型管理系统
# makefile参考
# 在 Golang 中开发中使用 Makefile
# https://studygolang.com/articles/11131

# 定义make变量
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
BINARY_PATH=./bin
CMD_BINARY_NAME=$(BINARY_PATH)/start_up

# make 不指定动作时，默认执行第一个动作
default:build

test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)

mod:
	$(GO) mod tidy

build-local:
	export CGO_ENABLED=0
	$(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
	shasum -a 256 $(CMD_BINARY_NAME)

build: mod clean test build-local
	echo "build done"

include release.mk


