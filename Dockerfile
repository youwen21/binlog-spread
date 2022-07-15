# base go

FROM golang:1.14 as build

# 镜像维护者
MAINTAINER owen <youwen21@yeah.net>

# 修改国内源
#RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
#RUN sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

# 环境变量
ENV GOROOT="/usr/local/go"
ENV GO111MODULE="on"

# 添加目录与文件
RUN mkdir -p /go/src/
ENV GOPROXY=https://goproxy.cn,direct
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY "./" "/go/src/"

WORKDIR /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -o start_up cmd/cmd.go



# 添加所有需要编译的应用代码， 这个地方用column， 二阶段编译用add
#ADD . .

# 环境变量
# ENV

# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
# （必看）使用 scratch Docker 镜像部署 Go 应用
# https://baijiahao.baidu.com/s?id=1617163590078677512&wfr=spider&for=pc
# （必看）如何为你的Go应用创建轻量级Docker镜像？
# http://dockone.io/article/8196
#
# 使用docker scratch 构建 golang 之坑
# https://www.cnblogs.com/feiquan/p/13413008.html


# 生产阶段， scratch是一个最小镜像,可以理解为null， 用来构建原生二进制执行的最小镜像包， 但不支持cgo
# alpine优缺点对比 有/bin/sh 但没有/bin/bash. 目前了解到相比scratch多包含了busybox
#
#FROM scratch AS prod
FROM alpine AS prod
#FROM centos AS prod
RUN mkdir -p /go/src/storage && chmod -R 777 /go/src/storage
WORKDIR /go/src/

# 从buil阶段拷贝二进制文件
COPY --from=build "/go/src/start_up" /go/src/start_up
COPY --from=build "/go/src/.env.example" /go/src/.env
COPY ./assets /go/src/assets
#COPY "./,env.example" "/go/src/.env"

# 环境变量
# ENV

EXPOSE 8000
CMD ["/go/src/start_up"]
