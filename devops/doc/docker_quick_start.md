# 项目数据模型管理系统 - docker 快速开始

## 下载镜像

docker pull youwen21/project_data_model:latest

## 创建好数据库
`为存储数据创建的库， mysql-binlog consumer拉取的binlog是针对实际业务中的mysql实例`

SQL文件：
> project_data_model/business_event.sql

## 启动 web
```bash
docker run --rm --name pdm_web -p 8000:8000 \
-e DB_EVENT_HOST="rncs.com" \
-e DB_EVENT_PORT="3306" \
-e DB_EVENT_USERNAME="myn" \
-e DB_EVENT_PASSWORD="ow@" \
-e DB_EVENT_DATABASE="business_event" \
-e DB_EVENT_CHARSET="utf8" \
youwen21/project-data-model /go/src/start_up web-start
```

> 浏览器访问 http://127.0.0.1:8000 

## 启动 mysql-binlog consumer
`注意，consumer消费的mysql主库您的业务库， 和上面创建的库没有关系。`

```bash
docker run --rm --name pdm_binlog_consumer \
-e DB_EVENT_HOST="127.0.0.1" \
-e DB_EVENT_PORT="3306" \
-e DB_EVENT_USERNAME="mun" \
-e DB_EVENT_PASSWORD="ow@" \
-e DB_EVENT_DATABASE="business_event" \
-e DB_EVENT_CHARSET="utf8" \
youwen21/project-data-model /go/src/start_up binlog-start -host='riyuncs.com' -username='mun' -password='ow@@'

```
