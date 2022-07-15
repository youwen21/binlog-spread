# 数据变更统计 - 开发说明

## 模块

### web-ui

- 业务事件数据模型
    + 事件列表
    + binlog数据流

### binlog-consumer

env配置项ENABLE_MODEL_STREAM决定 把mysql-binlog转换为ddd_event_stream数据

```
# 数据模型
## 启动mysql-binlog 数据流收集
ENABLE_MODEL_STREAM=yes
## 累计满多少行数据写一次库, 满足立即写入到表ddd_event_stream
MODEL_STREAM_FLUSH_ROWS=100
##  数据刷新频率设置 单位秒, 满足立即写入到表ddd_event_stream
MODEL_STREAM_FLUSH_DURATION=500
```

## 数据库表

- ddd_event
- ddd_event_stream

## 依赖项

- mysql-binlog
- jsdifflib

## 开发入口

- cmd/cmd.go

## 调用路径

cmd/cmd.go -> StartBinlogClient -> handle_binlog.HandleEvent(ev) -> handleUpdateEventV1(e) -> updateRoutineModelStream(
ev)




