# 安装项目数据模型管理系统

## 安装方式

### 编译安装

### 可执行文件包下载安装

### 下载地址

> https://gitee.com/youwen21/business_data_model-go/releases

## 目录结构重点元素：

```
.
├── assets // html静态资源，无需nginx提供完整服务
│ ├── AdminLTE-3.0.5 // 管理后台框架
│ ├── admin // 管理后台页面
│ ├── dist // 项目js，image等
│ └── plugins // js 外部包
├── storage // app.pid， log日志文件存放位置
├── xxx.exe 可执行文件
└── business_event.sql // 项目初始化的sql文件
```

## 配置环境变量

```
APP_ENV=local 当配置为local时每次启动项目会自动打开浏览器
APP="business_data_model"
APP_PORT=:8000  项目监听的端口
```

### 项目数据库配置

`注意： 被消费mysql-binlog的主库配置不需要写到env文件中。 主库信息由启动命令参数指定`

```
DB_EVENT_HOST=127.0.0.1
DB_EVENT_USERNAME=root
DB_EVENT_PASSWORD=root
DB_EVENT_DATABASE=business_event
DB_EVENT_CHARSET=utf8
DB_EVENT_FILTER=".*\\..*" //指定关心的库包，支持正则，配置方式可参考canal filter配置文式
```

### 项目功能配置

```
# 状态管理
## 状态流程检查是否开启， no关闭状态检查功能 （state_abnormal表）
ENABLE_CHECK_STATE=yes

# 数据统计
## 统计数据变更次数是否开启 no关闭数据统计功能 （statistics_day表）
ENABLE_DATA_STATISTICS=yes
## 统计数据累计次数， 满足立即写入到表statistics_day
DATA_STATISTICS_EVENT_TIMES=200
## 统计数据刷新频率设置 单位秒, 满足立即写入到表statistics_day
DATA_STATISTICS_FLUSH_DURATION=500

# 数据模型
## 启动mysql-binlog 数据流收集  no关闭mysql数据流生成（event_stream表）
ENABLE_MODEL_STREAM=yes
## 累计满多少行数据写一次库, 满足立即写入到表ddd_event_stream
MODEL_STREAM_FLUSH_ROWS=100
##  数据刷新频率设置 单位秒, 满足立即写入到表ddd_event_stream
MODEL_STREAM_FLUSH_DURATION=500
```

## 导入mysql

business_event.sql 文件导入到mysql数据库中

## 启动web服务 和 修改admin密码

```bash
./business_data_model web-start 
```

用户名密码 默认为admin, admin
> http://127.0.0.1:8000/admin/entrance/login.html

`修改登录密码： 把生成的新密码，替换administrator表中password字段既可。`
> echo -n "新密码" | openssl dgst -sha1

## 启动mysql-binlog消费后台进程。

每次启动指定一个mysql master实例。 有多个mysql master需要启动多个后台进程。

```bash
./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password"
```

### mysql-binlog消费服务 - filter过滤数据库表说明

filter是匹配，不是排除，暂不支持exclude参数   
默认过滤条件为: ".*\\..*" , 监听任何库，任何表的binlog事件。 启动mysql-binlog消费者时，指定过滤条件

```bash
#多个过滤表达用，号隔开。
#codeper库的任何表，和test.table1, test.table2都是关系的表。 不符合条件的抛弃，不写入ddd_event_stream表中。
./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password" -filter="codeper\\..*,test\\.table1,test\\.table2"
```
