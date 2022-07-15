package sqlite3_conn

var tableSchema = `
/*
 Navicat Premium Data Transfer

 Source Server         : test2
 Source Server Type    : SQLite
 Source Server Version : 3012001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3012001
 File Encoding         : 65001

 Date: 14/07/2022 22:17:09
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for administrator
-- ----------------------------
DROP TABLE IF EXISTS "administrator";
CREATE TABLE "administrator" (
  "administrator_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "username" VARCHAR(255),
  "password" VARCHAR(255),
  "name" VARCHAR(255),
  "avatar" VARCHAR(255),
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME,
  "is_deleted" TINYINT DEFAULT '1',
  "deleted_at" DATETIME
);

-- ----------------------------
-- Records of administrator
-- ----------------------------
BEGIN;
INSERT INTO "administrator" VALUES (1, 'admin', 'd033e22ae348aeb5660fc2140aec35850c4da997', '主管理帐号', NULL, '2019-05-23 12:09:26', '2021-08-01 07:27:42', 0, '2020-07-20 14:53:30');
INSERT INTO "administrator" VALUES (2, 'owen', '8cb2237d0679ca88db6464eac60da96345513964', 'show me', NULL, '2019-05-23 12:09:26', '2021-08-01 07:27:42', 0, '2020-07-20 14:53:18');
COMMIT;

-- ----------------------------
-- Table structure for api_binlog
-- ----------------------------
DROP TABLE IF EXISTS "api_binlog";
CREATE TABLE "api_binlog" (
	"api_binlog_id" INTEGER NOT NULL  ,
	"db_name" VARCHAR(255) NULL  ,
	"table_name" VARCHAR(255) NULL  ,
	"transaction_tag" VARCHAR(64) NULL  ,
	"event_type" INTEGER NULL DEFAULT '-100' ,
	"columns" TEXT NULL  ,
	"update_columns" TEXT NULL  ,
	"update_value" TEXT NULL  ,
	"ignore_column_value" TEXT NULL  ,
	"comment" TEXT NULL  ,
	"created_at" DATETIME NULL DEFAULT CURRENT_TIMESTAMP ,
	"updated_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ,
	PRIMARY KEY ("api_binlog_id")
);

-- ----------------------------
-- Table structure for api_define
-- ----------------------------
DROP TABLE IF EXISTS "api_define";
CREATE TABLE "api_define" (
  "api_define_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "api_type" VARCHAR(64) DEFAULT 'mysql',
  "api_tag" VARCHAR(255),
  "api_name" VARCHAR(255),
  "stream_ids" VARCHAR(255),
  "api_version" VARCHAR(255),
  "api_link" VARCHAR(255) DEFAULT '',
  "comment" TEXT,
  "is_deleted" INTEGER DEFAULT '0',
  "deleted_at" DATETIME,
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS "menu";
CREATE TABLE "menu" (
  "menu_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "title" VARCHAR(50) NOT NULL DEFAULT '',
  "pid" INTEGER NOT NULL DEFAULT '0',
  "sort" INTEGER NOT NULL DEFAULT '0',
  "hide" TINYINT NOT NULL DEFAULT '0',
  "pathname" VARCHAR(255),
  "iconfont" VARCHAR(255) DEFAULT '',
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "is_deleted" TINYINT DEFAULT '1',
  "deleted_at" DATETIME
);

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO "menu" VALUES (200, '菜单列表', 0, 0, 0, '/admin/menu/list.html', 'fa-bars', '2020-02-16 09:14:38', '2021-08-01 07:27:42', 0, '2020-07-21 13:09:05');
INSERT INTO "menu" VALUES (201, '管理员列表', 0, 0, 0, '/admin/admins/list.html', 'fa-user', '2020-02-16 09:14:38', '2021-08-01 07:27:42', 1, '2021-02-11 07:43:21');
INSERT INTO "menu" VALUES (203, '资源列表', 0, 0, 0, '/admin/resource/list.html', 'fa-tag', '2020-02-16 09:14:38', '2021-08-01 07:27:42', 1, '2021-02-11 07:43:14');
INSERT INTO "menu" VALUES (205, '权限管理', 0, 0, 0, '', 'fa-share', '2020-02-16 09:14:38', '2021-08-01 07:27:42', 1, '2021-02-11 07:42:55');
INSERT INTO "menu" VALUES (206, '权限节点列表', 205, 0, 0, '/admin/permissions/list.html', 'fa-tag', '2020-02-16 09:14:38', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "menu" VALUES (207, '角色列表', 205, 0, 0, '/admin/roles/list.html', 'fa-tag', '2020-02-16 09:14:38', '2021-08-01 07:27:42', 0, '2020-07-20 22:42:33');
INSERT INTO "menu" VALUES (231, '事件管理', 0, 0, 0, '', '', '2020-11-12 18:11:19', '2021-08-01 07:27:42', 0, '2021-03-27 07:45:16');
INSERT INTO "menu" VALUES (232, '事件列表', 231, 0, 0, '/admin/api_define/list.html', '', '2020-11-12 18:11:52', '2021-08-02 00:13:21', 0, '2021-03-27 07:45:12');
INSERT INTO "menu" VALUES (233, 'mysql-binlog', 231, 0, 0, '/admin/api_define/binlog_list.html', '', '2020-11-12 18:26:20', '2021-08-02 00:13:25', 0, '2021-03-27 07:45:10');
INSERT INTO "menu" VALUES (234, '状态管理', 0, 100, 0, '', '', '2021-02-09 10:36:26', '2021-08-01 07:27:42', 0, '2021-03-27 07:45:09');
INSERT INTO "menu" VALUES (235, '状态定义', 234, 0, 0, '/admin/state/list.html', '', '2021-02-09 10:38:17', '2021-08-01 07:27:42', 0, '2021-03-27 07:45:06');
INSERT INTO "menu" VALUES (236, '数据统计', 0, 0, 0, '', '', '2021-02-09 11:22:21', '2021-08-01 07:27:42', 1, '2022-07-13 20:25:35.221282+08:00');
INSERT INTO "menu" VALUES (237, '统计规则', 236, 0, 0, '/admin/statistics/rule_list.html', '', '2021-02-09 11:22:47', '2021-08-01 07:27:42', 0, '2021-03-27 07:44:56');
INSERT INTO "menu" VALUES (238, '异常状态变更', 234, 0, 0, '/admin/state/abnormal_list.html', '', '2021-02-18 20:15:36', '2021-08-01 07:27:42', 0, '2021-03-27 07:44:51');
INSERT INTO "menu" VALUES (239, '每日统计结果', 236, 0, 0, '/admin/statistics/day_list.html', '', '2021-02-18 22:05:32', '2021-08-01 07:27:42', 0, '2021-03-27 07:44:48');
COMMIT;


-- ----------------------------
-- Table structure for state
-- ----------------------------
DROP TABLE IF EXISTS "state";
CREATE TABLE "state" (
  "state_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "state_class_id" INTEGER,
  "state_value" VARCHAR(255),
  "state_value_desc" VARCHAR(255),
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME,
  "is_deleted" INTEGER DEFAULT '0',
  "deleted_at" DATETIME
);

-- ----------------------------
-- Records of state
-- ----------------------------
BEGIN;
INSERT INTO "state" VALUES (1, 1, 1, '待支付', '2021-02-09 11:34:37', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "state" VALUES (2, 1, 2, '部分支付', '2021-02-09 11:35:01', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "state" VALUES (3, 1, 3, '支付完成', '2021-02-09 11:35:12', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "state" VALUES (4, 1, 4, '异步回调失败', '2021-02-09 11:35:40', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "state" VALUES (5, 1, 5, '异步回调成功，待发货', '2021-02-09 11:36:23', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "state" VALUES (6, 1, 6, '已发货', '2021-02-09 11:36:43', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "state" VALUES (7, 1, 7, '已签收', '2021-02-09 11:37:07', '2021-08-01 07:27:42', 0, NULL);
INSERT INTO "state" VALUES (8, 1, 8, '退款中-1', '2021-02-09 11:37:22', '2021-08-01 07:27:42', 0, '2021-02-10 21:00:34');
INSERT INTO "state" VALUES (11, 3, 0, '默认', '2021-02-18 17:10:33', '2021-08-01 07:27:42', 1, '2021-02-19 15:35:01');
INSERT INTO "state" VALUES (12, 0, 1, '活动开始2', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (13, 0, 2, '活动进行中2', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (14, 3, 3, '活动结束', '2021-02-18 17:11:11', '2021-08-01 07:27:42', 1, '2021-02-19 17:28:44');
INSERT INTO "state" VALUES (15, 6, 'aa', 'aaaaaa', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (16, 7, 1, '待抓取', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (17, 7, 2, '抓取中', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (18, 7, 3, '抓取异常', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (19, 7, 4, '待分配', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (20, 7, 5, '待制作', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (21, 7, 6, '待审核', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (22, 7, 7, '待上线', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (23, 7, 8, '上线失败', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (24, 7, 9, '已上线', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (25, 7, 10, '打回CP', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (26, 7, 11, '重新审核', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (27, 7, 12, 'CP处理-取消授权', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (28, 7, 13, '上线中', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (29, 9, 1, '开始', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (30, 9, 2, '过程', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (31, 9, 3, '结束', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (32, 10, 1, 'aaa', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state" VALUES (33, 10, 2, '过程', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for state_abnormal
-- ----------------------------
DROP TABLE IF EXISTS "state_abnormal";
CREATE TABLE "state_abnormal" (
  "state_abnormal_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "db_name" VARCHAR(255) DEFAULT '',
  "table_name" VARCHAR(255) DEFAULT '',
  "field_name" VARCHAR(255) DEFAULT '',
  "event_type" INTEGER DEFAULT '0',
  "state_from" VARCHAR(255),
  "state_to" VARCHAR(255),
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "is_deleted" TINYINT DEFAULT '0',
  "deleted_at" DATETIME
);

-- ----------------------------
-- Records of state_abnormal
-- ----------------------------
BEGIN;
INSERT INTO "state_abnormal" VALUES (1, 'codeper', 'activity', 'status', 0, 2, 22, '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for state_class
-- ----------------------------
DROP TABLE IF EXISTS "state_class";
CREATE TABLE "state_class" (
  "state_class_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "state_name" VARCHAR(255) DEFAULT '',
  "db_name" VARCHAR(255) DEFAULT '',
  "table_name" VARCHAR(255) DEFAULT '',
  "field_name" VARCHAR(255) DEFAULT '',
  "state_describe" VARCHAR(255) DEFAULT '',
  "status" TINYINT DEFAULT '1',
  "is_deleted" TINYINT DEFAULT '0',
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME,
  "deleted_at" DATETIME
);

-- ----------------------------
-- Records of state_class
-- ----------------------------
BEGIN;
INSERT INTO "state_class" VALUES (1, '订单状态', 'test', 'order', 'order_status', '订单流程说明adada', 1, 0, '2021-08-01 07:27:42', '2021-08-01 07:27:42', NULL);
INSERT INTO "state_class" VALUES (2, 'test', 'sss', 'bbb', 'dd', '', 1, 1, '2021-08-01 07:27:42', '2021-08-01 07:27:42', '2021-02-10 23:03:12');
INSERT INTO "state_class" VALUES (7, '审核书籍状态', 'yykeyle', 't_examine_book', 'status', '审核图书状态', 1, 0, '2021-08-01 19:29:33', '2021-08-01 19:29:33', '0000-00-00 00:00:00');
INSERT INTO "state_class" VALUES (9, '测试-a', 'aaa', 'bbb', 'ccc', '审核图书状态-aawwwww', 1, 0, '2021-11-07 20:21:05', '2021-11-07 20:21:05', '0000-00-00 00:00:00');
INSERT INTO "state_class" VALUES (10, '测试-b', 'aaa', 'bbb', 'ccc', '审核图书状态-aawwwwwyy', 1, 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for state_direction
-- ----------------------------
DROP TABLE IF EXISTS "state_direction";
CREATE TABLE "state_direction" (
  "state_direction_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "state_class_id" INTEGER DEFAULT '0',
  "state_from" VARCHAR(255) DEFAULT '0',
  "state_to" VARCHAR(255) DEFAULT '0',
  "label" VARCHAR(255),
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
  "is_deleted" TINYINT DEFAULT '0',
  "deleted_at" DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- ----------------------------
-- Records of state_direction
-- ----------------------------
BEGIN;
INSERT INTO "state_direction" VALUES (1, 1, 1, 2, NULL, '2021-02-10 05:40:00', '2021-08-01 07:27:42', 0, '2021-02-10 05:40:00');
INSERT INTO "state_direction" VALUES (2, 1, 1, 3, NULL, '2021-02-10 05:40:37', '2021-08-01 07:27:42', 0, '2021-02-10 05:40:37');
INSERT INTO "state_direction" VALUES (3, 1, 2, 3, NULL, '2021-02-10 05:40:42', '2021-08-01 07:27:42', 0, '2021-02-10 05:40:42');
INSERT INTO "state_direction" VALUES (5, 1, 3, 4, NULL, '2021-02-10 20:42:33', '2021-08-01 07:27:42', 0, '2021-02-10 20:42:33');
INSERT INTO "state_direction" VALUES (6, 1, 3, 5, NULL, '2021-02-10 20:42:43', '2021-08-01 07:27:42', 0, '2021-02-10 20:42:43');
INSERT INTO "state_direction" VALUES (7, 1, 5, 6, NULL, '2021-02-10 20:42:54', '2021-08-01 07:27:42', 0, '2021-02-10 20:42:54');
INSERT INTO "state_direction" VALUES (8, 1, 6, 7, NULL, '2021-02-10 20:43:22', '2021-08-01 07:27:42', 0, '2021-02-10 20:43:22');
INSERT INTO "state_direction" VALUES (9, 1, 5, 8, NULL, '2021-02-10 20:46:20', '2021-08-01 07:27:42', 0, '2021-02-10 20:46:20');
INSERT INTO "state_direction" VALUES (10, 1, 7, 8, NULL, '2021-02-10 20:47:44', '2021-08-01 07:27:42', 0, '2021-02-10 20:56:19');
INSERT INTO "state_direction" VALUES (11, 3, 0, 1, NULL, '2021-02-18 17:14:05', '2021-08-01 07:27:42', 1, '2021-02-19 15:34:45');
INSERT INTO "state_direction" VALUES (12, 3, 1, 2, NULL, '2021-02-18 17:14:22', '2021-08-01 07:27:42', 0, '2021-02-18 17:14:22');
INSERT INTO "state_direction" VALUES (13, 3, 2, 3, NULL, '2021-02-18 17:16:00', '2021-08-01 07:27:42', 0, '2021-02-18 17:16:00');
INSERT INTO "state_direction" VALUES (14, 3, 1, 3, NULL, '2021-02-18 17:16:22', '2021-08-01 07:27:42', 0, '2021-02-18 17:16:22');
INSERT INTO "state_direction" VALUES (15, 7, 1, 2, NULL, '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state_direction" VALUES (16, 7, 2, 3, NULL, '2021-06-17 14:45:03', '2021-08-01 07:27:42', 0, '2021-06-17 14:45:03');
INSERT INTO "state_direction" VALUES (17, 7, 2, 4, NULL, '2021-06-17 14:45:18', '2021-08-01 07:27:42', 0, '2021-06-17 14:45:18');
INSERT INTO "state_direction" VALUES (18, 7, 4, 5, '分配编辑', '2021-06-17 14:46:46', '2021-08-01 07:27:42', 0, '2021-06-17 14:46:46');
INSERT INTO "state_direction" VALUES (19, 7, 5, 6, '初审', '2021-06-17 14:47:12', '2021-08-01 07:27:42', 0, '2021-06-17 14:47:12');
INSERT INTO "state_direction" VALUES (20, 7, 5, 4, '撤回', '2021-06-17 14:47:51', '2021-08-01 07:27:42', 0, '2021-06-17 14:47:51');
INSERT INTO "state_direction" VALUES (21, 7, 6, 5, '终审退回', '2021-06-17 14:48:36', '2021-08-01 07:27:42', 0, '2021-06-17 14:48:36');
INSERT INTO "state_direction" VALUES (22, 7, 6, 10, '打回CP，不可上线都是状态10', '2021-06-17 14:49:30', '2021-08-01 07:27:42', 0, '2021-06-17 14:49:30');
INSERT INTO "state_direction" VALUES (23, 7, 6, 7, '终审通过', '2021-06-17 14:50:05', '2021-08-01 07:27:42', 0, '2021-06-17 14:50:05');
INSERT INTO "state_direction" VALUES (24, 7, 7, 13, '批量上线', '2021-06-17 14:50:44', '2021-08-01 07:27:42', 0, '2021-06-17 14:50:44');
INSERT INTO "state_direction" VALUES (25, 7, 13, 8, NULL, '2021-06-17 14:50:51', '2021-08-01 07:27:42', 0, '2021-06-17 14:50:51');
INSERT INTO "state_direction" VALUES (26, 7, 13, 9, NULL, '2021-06-17 14:51:13', '2021-08-01 07:27:42', 0, '2021-06-17 14:51:13');
INSERT INTO "state_direction" VALUES (27, 7, 10, 11, '重新审核', '2021-06-17 14:51:34', '2021-08-01 07:27:42', 0, '2021-06-17 14:51:34');
INSERT INTO "state_direction" VALUES (28, 7, 8, 7, '清除上线错误', '2021-06-17 15:03:10', '2021-08-01 07:27:42', 0, '2021-06-17 15:03:10');
INSERT INTO "state_direction" VALUES (29, 7, 11, 1, '编辑处理', '2021-06-17 15:04:13', '2021-08-01 07:27:42', 0, '2021-06-17 15:04:13');
INSERT INTO "state_direction" VALUES (30, 7, 11, 4, '编辑处理', '2021-06-17 15:04:26', '2021-08-01 07:27:42', 0, '2021-06-17 15:04:26');
INSERT INTO "state_direction" VALUES (32, 7, 10, 6, '取消打回', '2021-06-17 15:19:41', '2021-08-01 07:27:42', 0, '2021-06-17 15:19:41');
INSERT INTO "state_direction" VALUES (36, 7, 10, 12, '打回图书，取消授权', '0000-00-00 00:00:00', '2021-08-01 07:27:42', 0, '0000-00-00 00:00:00');
INSERT INTO "state_direction" VALUES (37, 9, 1, 2, 'aaa', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state_direction" VALUES (38, 9, 2, 3, 'vvv', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state_direction" VALUES (39, 9, 1, 3, 'yyy', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state_direction" VALUES (40, 10, 1, 2, 'aaaa', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
INSERT INTO "state_direction" VALUES (41, 10, 2, 1, 'yyyy', '0000-00-00 00:00:00', '0000-00-00 00:00:00', 0, '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Auto increment value for administrator
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 4 WHERE name = 'administrator';

-- ----------------------------
-- Auto increment value for api_define
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 51 WHERE name = 'api_define';

-- ----------------------------
-- Auto increment value for menu
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 239 WHERE name = 'menu';

-- ----------------------------
-- Indexes structure for table menu
-- ----------------------------
CREATE INDEX "main"."pid"
ON "menu" (
  "pid" ASC
);

-- ----------------------------
-- Auto increment value for state
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 33 WHERE name = 'state';

-- ----------------------------
-- Auto increment value for state_abnormal
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 1 WHERE name = 'state_abnormal';

-- ----------------------------
-- Auto increment value for state_class
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 10 WHERE name = 'state_class';

-- ----------------------------
-- Auto increment value for state_direction
-- ----------------------------
UPDATE "main"."sqlite_sequence" SET seq = 41 WHERE name = 'state_direction';

-- ----------------------------
-- Indexes structure for table state_direction
-- ----------------------------
CREATE UNIQUE INDEX "main"."dir_idx"
ON "state_direction" (
  "state_class_id" ASC,
  "state_from" ASC,
  "state_to" ASC
);


-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "sqlite_sequence";
CREATE TABLE sqlite_sequence(name,seq);

-- ----------------------------
-- Records of sqlite_sequence
-- ----------------------------
BEGIN;
INSERT INTO "sqlite_sequence" VALUES ('administrator', 4);
INSERT INTO "sqlite_sequence" VALUES ('api_define', 51);
INSERT INTO "sqlite_sequence" VALUES ('menu', 239);
INSERT INTO "sqlite_sequence" VALUES ('state', 33);
INSERT INTO "sqlite_sequence" VALUES ('state_abnormal', 1);
INSERT INTO "sqlite_sequence" VALUES ('state_class', 10);
INSERT INTO "sqlite_sequence" VALUES ('state_direction', 41);
COMMIT;


PRAGMA foreign_keys = true;


`