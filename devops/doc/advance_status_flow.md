# 状态管理 - 开发说明

## 模块

状态管理的web-ui模块可以独立使用，不启用ENABLE_DATA_STATISTICS 也可以用来管理状态节点 ，节点关系， 画节点在系图。

### web-ui

- 业务状态管理
    + 业务状态定义
    + 异常状态变更

### binlog-consumer

env配置项ENABLE_DATA_STATISTICS决定是否开启库表统计功能

```
# 数据统计
## 统计数据变更次数是否开启
ENABLE_DATA_STATISTICS=yes
## 统计数据累计次数， 满足立即写入到表statistics_day
DATA_STATISTICS_EVENT_TIMES=2
## 统计数据刷新频率设置 单位秒, 满足立即写入到表statistics_day
DATA_STATISTICS_FLUSH_DURATION=500
```

## 数据库表

- stat_class
- stat
- stat_direction
- stat_abnormal

## 依赖项

- mysql-binlog
- d3-graphviz

## 开发入口

- cmd/cmd.go

## 调用路径

cmd/cmd.go -> StartBinlogClient -> handle_binlog.HandleEvent(ev) -> handleUpdateEventV1(e) -> updateRoutineStatRule(ev)

