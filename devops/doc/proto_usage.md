# 安装
go get github.com/gogo/protobuf/protoc-gen-gofast
brew install protobuf

# google protobuf 和gogo protobuf定义上是有区分的
google protobuf 要求 proto文件中有go_package 字段 
gogo ptotobuf不要求go_package字段 ， 会根据package获得包名


# protobuf 说明
> https://zhuanlan.zhihu.com/p/42059170

# gogoprotobuf 速度更多， golang多数使用此proto
> https://github.com/gogo/protobuf

# 很好的grpc demo案例
> https://www.cnblogs.com/sunshenggang/p/12396500.html
# grpc-go
> https://github.com/grpc/grpc-go

# 原生使用proto，不使用grpc的案例
> https://blog.csdn.net/ycc297876771/article/details/79969039

# 第一步，定义proto文件， 没有好的工具把go struct转换成proto file， 所以还是自动定义吧
模板：
```
参考： pb/pbmember/member.proto
```

# 第二步，把proto文件，使用protoc 转换为go文件
protoc --gofast_out=plugins=grpc:. member.proto // 此为 gogo proto, 此方式个人暂没设置好时间
protoc --go_out=plugins=grpc:. member.proto // 此为 google proto

# 第三步， 写grpc server, 实现proto中定义的service, 提供服务

# 第四步， grcp client调用

# 本以为proto可以和传统模式相结合使用。 现在看来是挺不容易。  不能有一个简单的库，中间件之类的东西，把model struct和 proto struct转换。 挺难受。 倒是是可以把proto以json的方式输出。
> https://github.com/grpc-ecosystem/grpc-gateway   grpc提供json restful api 
这个暂时就不实现了。 用到的时候再说吧。


# 版本问题
```
// 这个命令会导致protoc-gen-go 版本    v1.4.3
 go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

// 可以用这个命令现更回来
go get -u github.com/golang/protobuf/protoc-gen-go  v1.25.0

可以再次使用命令： protoc --go_out=plugins=grpc:. member.proto 
```