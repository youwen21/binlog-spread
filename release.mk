# 项目数据模型管理系统
# 出于操作友好性，mac windows linux 三个makefile 独立分开， make指定文件构建指定平台的程序。
# 默认makefile 就针对下前机器构建就好。  其他开发者构建时，不会多余构建。

# 参考
# 在 Golang 中开发中使用 Makefile
# https://studygolang.com/articles/11131

# Golang的跨平台编译程序
# https://www.cnblogs.com/ghj1976/archive/2013/04/19/3030703.html
#各平台的GOOS和GOARCH参考
#OS                   ARCH                          OS version
#linux                386 / amd64 / arm             >= Linux 2.6
#darwin               386 / amd64                   OS X (Snow Leopard + Lion)
#freebsd              386 / amd64                   >= FreeBSD 7
#windows              386 / amd64                   >= Windows 2000

# Golang 编译Mac、Linux、Windows多平台可执行程序
# https://studygolang.com/articles/28339?fr=sidebar

# Docker命令_各种参数简介（run、v、rm、-w、-u、-e）
# https://blog.csdn.net/sxzlc/article/details/107676425

# Makefile 入门
# https://zhuanlan.zhihu.com/p/149346441

# 使用xgo编译支持CGO

# 定义make变量
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
BINARY_PATH=./bin
CMD_BINARY_NAME=$(BINARY_PATH)/start_up

test:
	$(GOTEST) -v
clean:
	$(GOCLEAN)

mod:
	$(GO) mod tidy

clean-pre:
	@if [ ! -d ./release ]; then mkdir ./release; touch ./release/.gitkeep; fi

clean-dir: clean-pre
	rm -rf ./release/* !(.gitkeep)
	test -e ./release/.env && rm ./release/.env

mv-release-file:
	mv release_*.zip ./storage

publish-common-init:
	mkdir -p ./release/storage/logs && chmod -R 777 ./release/storage
	cp -r ./static ./release
	cp .env.example ./release/.env
	cp ./assistant.sql ./release

publish: publish-linux publish-mac publish-windows mv-release-file

publish-windows: clean-dir build-windows publish-common-init
	cp $(BINARY_PATH)/*.exe ./release
	zip -r release_windows_`date +%Y%m%d`.zip release

publish-mac: clean-dir build-mac publish-common-init
	cp $(CMD_BINARY_NAME) ./release
	zip -r release_mac_`date +%Y%m%d`.zip release

publish-linux: clean-dir build-linux publish-common-init
	cp $(CMD_BINARY_NAME) ./release
	zip -r release_linux_`date +%Y%m%d`.zip release

build-windows:
	xgo --image="youwen21/ali-proxy-xgo" --targets="windows/*" -dest=$(BINARY_PATH) ./cmd/

build-linux:
	export CGO_ENABLED=0 GOOS=linux
	$(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go

build-mac:
	export CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
	$(GOBUILD) -o $(CMD_BINARY_NAME) -v ./cmd/cmd.go
