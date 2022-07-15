# 项目实现说明

## 三个独立功能

项目共包含三个独立功能， 三个功能都可以独立使用。 三个功能都依赖于mysql-binlog解析。

### 数据变更统计 - 开发说明

[数据变更统计-开发说明文档](advance_statistics.md)

### 状态流定义和图形化展示，状态变更校验 - 开发说明

[状态管理-开发说明文档](advance_status_flow.md)

### 事件数据模型 - 开发说明

[事件数据模型-开发说明文档](advance_data_model.md)

## 本地启动binlog-consumer

```bash
go run cmd/cmd.go binlog-start -host='host' -username='username' -password='password'
```

## debug断点调试命令行

### 配置命令行启动参数

配置golang    
run -> Edit Configurations 打开配置窗口

Go Build ->(Configuration项) -> Program arguments

```
binlog-start -host='host' -username='username' -password='password'
```

## 使用go-callvis绘制调用关系

```bash
go-callvis -focus=owen2020/cmd/command/handle_binlog -group pkg,type -limit owen2020/cmd/command  ./cmd
```

![handle_binlog_package_graph](/devops/docops/doc/images/package_graph.jpg)

参考
> https://www.lagou.com/lgeduarticle/96545.html  
> https://github.com/ofabry/go-callvis/tree/master/examples

## docker和make说明
### docker镜像构建
项目根目中的Dockerfile是为构建项目镜像准备的。还需要改动读取mysql，redis等配置的方式。

### Makefile说明
三个Makefile的做用：   
Makefile - 构建本机环境的可执行文件  
linux.Makefile - 构建linux系统可执行文件, 非项目维护者不需要关心  
window.Makefile - 构建window系统可执行文件, 非项目维护者不需要关心  