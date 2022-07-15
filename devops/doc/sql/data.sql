-- MySQL dump 10.13  Distrib 8.0.22, for osx10.14 (x86_64)
--
-- Host: rm-2zemxuvee9kii2b55so.mysql.rds.aliyuncs.com    Database: business_event
-- ------------------------------------------------------
-- Server version	8.0.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ '2d2a7af0-ecdb-11e9-b1b5-00163e164aa3:1-894246';

--
-- Dumping data for table `administrator`
--

LOCK TABLES `administrator` WRITE;
/*!40000 ALTER TABLE `administrator` DISABLE KEYS */;
INSERT INTO `administrator` VALUES (1,'admin','d033e22ae348aeb5660fc2140aec35850c4da997','主管理帐号',NULL,'2019-05-23 12:09:26','2020-07-21 22:47:03',0,'2020-07-20 14:53:30'),(3,'owen','8cb2237d0679ca88db6464eac60da96345513964','show me',NULL,'2019-05-23 12:09:26','2020-07-22 10:21:48',0,'2020-07-20 14:53:18');
/*!40000 ALTER TABLE `administrator` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (200,'菜单列表',0,0,0,'/admin/menu/list.html','fa-bars','2020-02-16 09:14:38','2021-02-09 10:35:08',0,'2020-07-21 13:09:05'),(201,'管理员列表',0,0,0,'/admin/admins/list.html','fa-user','2020-02-16 09:14:38','2021-02-11 07:43:21',1,'2021-02-11 07:43:21'),(203,'资源列表',0,0,0,'/admin/resource/list.html','fa-tag','2020-02-16 09:14:38','2021-02-11 07:43:14',1,'2021-02-11 07:43:14'),(205,'权限管理',0,0,0,'','fa-share','2020-02-16 09:14:38','2021-02-11 07:42:55',1,'2021-02-11 07:42:55'),(206,'权限节点列表',205,0,0,'/admin/permissions/list.html','fa-tag','2020-02-16 09:14:38','2020-07-22 13:48:33',0,NULL),(207,'角色列表',205,0,0,'/admin/roles/list.html','fa-tag','2020-02-16 09:14:38','2020-07-22 13:48:36',0,'2020-07-20 22:42:33'),(231,'事件管理',0,0,0,'','','2020-11-12 18:11:19','2021-06-09 10:23:26',0,'2021-03-27 07:45:16'),(232,'事件列表',231,0,0,'/admin/event/list.html','','2020-11-12 18:11:52','2021-04-13 11:32:36',0,'2021-03-27 07:45:12'),(233,'mysql-binlog',231,0,0,'/admin/event/stream_list.html','','2020-11-12 18:26:20','2021-06-09 10:24:19',0,'2021-03-27 07:45:10'),(234,'状态管理',0,100,0,'','','2021-02-09 10:36:26','2021-06-09 10:23:39',0,'2021-03-27 07:45:09'),(235,'状态定义',234,0,0,'/admin/state/list.html','','2021-02-09 10:38:17','2021-06-09 10:24:42',0,'2021-03-27 07:45:06'),(236,'数据统计',0,0,0,'','','2021-02-09 11:22:21','2021-06-09 10:23:58',0,'2021-03-27 07:45:04'),(237,'统计规则',236,0,0,'/admin/statistics/rule_list.html','','2021-02-09 11:22:47','2021-04-13 11:32:41',0,'2021-03-27 07:44:56'),(238,'异常状态变更',234,0,0,'/admin/state/abnormal_list.html','','2021-02-18 20:15:36','2021-04-13 11:32:42',0,'2021-03-27 07:44:51'),(239,'每日统计结果',236,0,0,'/admin/statistics/day_list.html','','2021-02-18 22:05:32','2021-06-09 10:25:14',0,'2021-03-27 07:44:48');
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-06-12 22:55:57
