# 清理mysql数据

```sql
DELETE 
FROM
	ddd_event_stream 
WHERE
	created_at < DATE_SUB( NOW( ), INTERVAL 7 DAY ) 
	AND ddd_event_stream_id NOT IN ( SELECT GROUP_CONCAT( `stream_ids` SEPARATOR ',' ) FROM ddd_event )
```