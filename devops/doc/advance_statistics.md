# 数据变更统计 - 开发说明

数据变更统计目前是以天为单位， 统计每天规则包含的库表增、改、删影响行数。  
如 sql: update db.t set status=1 where id in (1,2,3) ， 将会统计 db库t表更新行数3  

## 模块

### web-ui

- 数据变更统计
    + 统计规则
    + 每日统计列表

### binlog-consumer

env配置项ENABLE_CHECK_STATE决定是否开启状态校验功能

```
# 状态管理
## 状态流程检查是否开启
ENABLE_CHECK_STATE=yes
```

## 数据库表

- statistics_rule
- statistics_day

## 依赖项

- mysql-binlog

## 开发文件入口

- cmd/cmd.go

## 调用路径

cmd/cmd.go -> StartBinlogClient -> handle_binlog.HandleEvent(ev) -> handleUpdateEventV1(e) -> updateRoutineStatistics(
ev)

