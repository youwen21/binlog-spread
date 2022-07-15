#!/usr/bin/env bash

# https://qq52o.me/2688.html
# 解决依赖
go mod tidy
# 重新编译
go build
# 重启服务
kill -1 `cat storage/app.pid`

# 百度云 go get 需要proxy
# https://blog.csdn.net/xuelang532777032/article/details/108584203