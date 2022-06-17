-- MySQL dump 10.13  Distrib 8.0.28, for macos11 (arm64)
--
-- Host: localhost    Database: shard
-- ------------------------------------------------------
-- Server version	8.0.28

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

--
-- Current Database: `shard`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `shard` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `shard`;

--
-- Table structure for table `shard_admin`
--

DROP TABLE IF EXISTS `shard_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_admin` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `password` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `username` varchar(20) DEFAULT NULL,
  `last_login_time` bigint DEFAULT NULL,
  `role_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_shard_admin_role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_admin`
--

LOCK TABLES `shard_admin` WRITE;
/*!40000 ALTER TABLE `shard_admin` DISABLE KEYS */;
INSERT INTO `shard_admin` VALUES (1,1630981091,1648452703,0,'$2a$10$tMmQf5eholT2v4ikklVuVOq0PCOmObtD2w8chYXAbi4T7KeQOyyFe','6b90add233e508bd03b2bb35ad92e39326257a1d0c796595de2884c6a3c86360.jpeg','admin',1648452703,1),(2,1631276335,1648440382,0,'$2a$10$g/BTPQIe/Fdy9QQN6nyO5OiYGvjZ6gBnU6yec844P2NdRk9KN2sG2','','admin2',1648440382,2);
/*!40000 ALTER TABLE `shard_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_file`
--

DROP TABLE IF EXISTS `shard_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_file` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `name` longtext,
  `hash` longtext,
  `path` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_file`
--

LOCK TABLES `shard_file` WRITE;
/*!40000 ALTER TABLE `shard_file` DISABLE KEYS */;
INSERT INTO `shard_file` VALUES (20,1647004958,1647004958,0,NULL,'87d9bb400c0634691f0e3baaf1e2fd0d','./files'),(21,1648019306,1648019306,0,NULL,'da929f31c7170594e4da6837842bb7ed.py','./files'),(22,1648020635,1648020635,0,NULL,'d72a8d07e77cb7b8fd635f210adbbc85','./files'),(23,1648020721,1648020721,0,NULL,'876a969e6106e584528a701bf66b95ae','./files'),(24,1648085943,1648085943,0,NULL,'046b00af4e7a53cffa6b6d3af1e2a5b5.py','./files'),(25,1648086304,1648086304,0,NULL,'43bcf5a9049f5321127a039ef1ceab87.py','./files'),(26,1648212468,1648212468,0,NULL,'d040176d58101aa7d96dff8b0ae067ed.py','./files'),(27,1648212477,1648212477,0,NULL,'1e965e5aa57b34eb5b39a29ff489546b.py','./files'),(28,1648213801,1648213801,0,NULL,'e8aae68c4a312eb2905c8789c72b3b78.py','./files');
/*!40000 ALTER TABLE `shard_file` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_file_record`
--

DROP TABLE IF EXISTS `shard_file_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_file_record` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `uid` bigint DEFAULT NULL,
  `file_id` bigint DEFAULT NULL,
  `name` longtext,
  `file_type` tinyint unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_file_record`
--

LOCK TABLES `shard_file_record` WRITE;
/*!40000 ALTER TABLE `shard_file_record` DISABLE KEYS */;
INSERT INTO `shard_file_record` VALUES (4,1647004958,1647004963,0,2,20,'run_script.txt',0),(5,1647004975,1648020861,0,1,22,'hello',0),(6,1647005375,1647005375,0,1,20,'run_script.txt',0),(7,1647005376,1647005376,0,1,20,'run_script.txt',0),(8,1647005377,1647005377,0,1,20,'run_script.txt',0),(9,1647005377,1647005377,0,1,20,'run_script.txt',0),(10,1647005378,1647005378,0,1,20,'run_script.txt',0),(11,1648019306,1648086304,0,1,25,'run',1),(12,1648020635,1648020635,0,1,22,'test',0),(13,1648020721,1648020721,0,1,23,'run_script.txt',0),(14,1648212468,1648213801,0,1,28,'test_py',1);
/*!40000 ALTER TABLE `shard_file_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_permission`
--

DROP TABLE IF EXISTS `shard_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_permission` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `name` longtext,
  `parent_id` bigint DEFAULT NULL,
  `key` longtext,
  `icon` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_permission`
--

LOCK TABLES `shard_permission` WRITE;
/*!40000 ALTER TABLE `shard_permission` DISABLE KEYS */;
INSERT INTO `shard_permission` VALUES (5,1631101512,1631110015,0,'管理',0,'','el-icon-s-tools'),(22,1631110062,1631263105,0,'权限管理',5,'1_0_21313131231','el-icon-s-tools'),(25,1631110090,1631273169,0,'路由管理',5,'1_1_1631110079332','el-icon-s-cooperation'),(26,1631171043,1631263096,0,'添加权限',22,'22_1_1631171034438',''),(27,1631189154,1631263090,0,'删除权限',22,'22_2_1631189149693',''),(28,1631260058,1631260058,0,'人员管理',0,'0_0_1631260049220','el-icon-user'),(29,1631260067,1631273182,0,'用户列表',28,'28_1_1631260060784','el-icon-user-solid'),(30,1631260078,1631273188,0,'管理员列表',28,'28_2_1631260069872','el-icon-s-custom'),(31,1631260127,1631273224,0,'角色管理',28,'28_3_1631260121599','el-icon-s-check'),(32,1631263083,1631263083,0,'编辑权限',22,'22_3_1631263050215',''),(33,1631276880,1631276904,0,'新增',30,'30_1_1631276859399',''),(34,1631276901,1631276933,0,'编辑',30,'30_2_1631276890248',''),(35,1631276919,1631276919,0,'删除',30,'30_3_1631276905329',''),(36,1631276958,1631276989,0,'变更角色',30,'30_4_1631276938272',''),(37,1631276984,1631279403,0,'变更权限',30,'30_5_1631276967393',''),(38,1631277288,1631277288,0,'重置密码',30,'30_6_1631277265508',''),(40,1631277638,1631277638,0,'新增',31,'31_1_1631277588024',''),(41,1631277677,1631277677,0,'编辑',31,'31_2_1631277662260',''),(42,1631277686,1631277686,0,'删除',31,'31_3_1631277678919',''),(43,1631277740,1631279397,0,'变更权限',31,'30_5_1631276967393',''),(44,1631589542,1631589542,0,'新增',29,'29_1_1631589517500',''),(45,1631589557,1631589557,0,'编辑',29,'29_2_1631589543207',''),(46,1631589568,1631589568,0,'删除',29,'29_3_1631589558934',''),(47,1631589586,1631589586,0,'重置密码',29,'29_4_1631589569386',''),(48,1631954799,1631954799,0,'调整余额',29,'29_5_1631954785974',''),(49,1632291170,1632291170,0,'修改头像',30,'30_7_1632291149973',''),(50,1646896492,1646919693,0,'工具',0,'0_0_1646895799021','el-icon-s-platform'),(51,1646919685,1646919685,0,'文件管理',50,'50_1_1646919662672','el-icon-s-order'),(52,1648434584,1648434584,0,'新增',51,'51_1_1648434568048',''),(53,1648434614,1648434614,0,'编辑',51,'51_2_1648434600816',''),(54,1648434679,1648434679,0,'删除',51,'51_3_1648434667207',''),(55,1648434688,1648434688,0,'运行',51,'51_4_1648434681188',''),(56,1648439117,1648439117,0,'详情',51,'51_5_1648439101813','');
/*!40000 ALTER TABLE `shard_permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_permission_route`
--

DROP TABLE IF EXISTS `shard_permission_route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_permission_route` (
  `permission_id` bigint NOT NULL,
  `route_id` bigint NOT NULL,
  PRIMARY KEY (`permission_id`,`route_id`),
  KEY `fk_shard_permission_route_route` (`route_id`),
  CONSTRAINT `fk_shard_permission_route_permission` FOREIGN KEY (`permission_id`) REFERENCES `shard_permission` (`id`),
  CONSTRAINT `fk_shard_permission_route_route` FOREIGN KEY (`route_id`) REFERENCES `shard_route` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_permission_route`
--

LOCK TABLES `shard_permission_route` WRITE;
/*!40000 ALTER TABLE `shard_permission_route` DISABLE KEYS */;
INSERT INTO `shard_permission_route` VALUES (22,2),(26,3),(29,6),(30,9),(49,14),(33,15),(34,15),(44,17),(35,31),(45,32),(34,33),(27,34),(46,35),(32,36),(40,38),(41,39),(31,40),(42,41),(37,43),(43,43),(37,44),(43,44),(25,45),(38,48),(36,51),(47,52),(48,53),(52,56),(51,57),(56,60),(55,64),(53,65),(54,83);
/*!40000 ALTER TABLE `shard_permission_route` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_role`
--

DROP TABLE IF EXISTS `shard_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_role` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_role`
--

LOCK TABLES `shard_role` WRITE;
/*!40000 ALTER TABLE `shard_role` DISABLE KEYS */;
INSERT INTO `shard_role` VALUES (1,1631168588,1648439128,0,'超级管理员'),(2,1631168810,1648440377,0,'普通的页面仔'),(3,1631537737,1631537737,0,'admin3'),(4,1631537742,1631537742,0,'admin4'),(5,1631537745,1631537745,0,'admin5'),(6,1631537749,1631537749,0,'admin7'),(7,1631537753,1631537753,0,'admin8'),(8,1631537758,1631537758,0,'admin9');
/*!40000 ALTER TABLE `shard_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_role_permission`
--

DROP TABLE IF EXISTS `shard_role_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_role_permission` (
  `role_id` bigint NOT NULL,
  `permission_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`),
  KEY `fk_shard_role_permission_permission` (`permission_id`),
  CONSTRAINT `fk_shard_role_permission_permission` FOREIGN KEY (`permission_id`) REFERENCES `shard_permission` (`id`),
  CONSTRAINT `fk_shard_role_permission_role` FOREIGN KEY (`role_id`) REFERENCES `shard_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_role_permission`
--

LOCK TABLES `shard_role_permission` WRITE;
/*!40000 ALTER TABLE `shard_role_permission` DISABLE KEYS */;
INSERT INTO `shard_role_permission` VALUES (1,5),(1,22),(1,25),(1,26),(1,27),(1,28),(2,28),(1,29),(2,29),(1,30),(2,30),(1,31),(2,31),(1,32),(1,33),(2,33),(1,34),(2,34),(1,35),(2,35),(1,36),(2,36),(1,37),(2,37),(1,38),(2,38),(1,40),(2,40),(1,41),(2,41),(1,42),(2,42),(1,43),(2,43),(1,44),(2,44),(1,45),(2,45),(1,46),(2,46),(1,47),(2,47),(1,48),(1,49),(1,50),(2,50),(1,51),(2,51),(1,52),(1,53),(1,54),(1,55),(1,56);
/*!40000 ALTER TABLE `shard_role_permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_route`
--

DROP TABLE IF EXISTS `shard_route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_route` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `name` varchar(20) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL,
  `path` varchar(50) DEFAULT NULL,
  `permission_id` bigint DEFAULT NULL,
  `route_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_shard_route_permission` (`permission_id`),
  KEY `fk_shard_route_route` (`route_id`),
  CONSTRAINT `fk_shard_route_permission` FOREIGN KEY (`permission_id`) REFERENCES `shard_permission` (`id`),
  CONSTRAINT `fk_shard_route_route` FOREIGN KEY (`route_id`) REFERENCES `shard_route` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_route`
--

LOCK TABLES `shard_route` WRITE;
/*!40000 ALTER TABLE `shard_route` DISABLE KEYS */;
INSERT INTO `shard_route` VALUES (2,1631017285,1631023195,0,'获取所有权限','GET','/backend/v1/permissions',NULL,NULL),(3,1631017285,1631017285,0,'','POST','/backend/v1/permissions',NULL,NULL),(6,1631017285,1631017285,0,'','GET','/backend/v1/users',NULL,NULL),(9,1631017285,1631017285,0,'','GET','/backend/v1/admins',NULL,NULL),(11,1631017285,1631017285,0,'','POST','/backend/v1/admins/login',NULL,NULL),(14,1631017285,1631017285,0,'','PUT','/backend/v1/admins/avatar',NULL,NULL),(15,1631017285,1631017285,0,'','POST','/backend/v1/admins',NULL,NULL),(17,1631017285,1631017285,0,'','POST','/backend/v1/users',NULL,NULL),(31,1631018134,1631018134,0,'','DELETE','/backend/v1/admins/:id',NULL,NULL),(32,1631018134,1631018134,0,'','PUT','/backend/v1/users/:id',NULL,NULL),(33,1631018134,1631018134,0,'','PUT','/backend/v1/admins/:id',NULL,NULL),(34,1631018134,1631018134,0,'','DELETE','/backend/v1/permissions/:id',NULL,NULL),(35,1631018134,1631018134,0,'','DELETE','/backend/v1/users/:id',NULL,NULL),(36,1631018134,1631018134,0,'','PUT','/backend/v1/permissions/:id',NULL,NULL),(38,1631157768,1631157768,0,NULL,'POST','/backend/v1/roles',NULL,NULL),(39,1631157768,1631157768,0,NULL,'PUT','/backend/v1/roles/:id',NULL,NULL),(40,1631157768,1631157768,0,NULL,'GET','/backend/v1/roles',NULL,NULL),(41,1631157768,1631157768,0,NULL,'DELETE','/backend/v1/roles/:id',NULL,NULL),(43,1631170482,1631170482,0,NULL,'GET','/backend/v1/roles/permissions/:id',NULL,NULL),(44,1631170681,1631170681,0,NULL,'PUT','/backend/v1/roles/permissions/:id',NULL,NULL),(45,1631199982,1631199982,0,NULL,'GET','/backend/v1/routes',NULL,NULL),(48,1631202091,1631202091,0,NULL,'GET','/backend/v1/admins/password/reset/:id',NULL,NULL),(49,1631202123,1631202123,0,NULL,'PUT','/backend/v1/admins/password/change',NULL,NULL),(51,1631242698,1631242698,0,NULL,'PUT','/backend/v1/admins/roles/change/:id',NULL,NULL),(52,1631589013,1631589013,0,NULL,'GET','/backend/v1/users/password/reset/:id',NULL,NULL),(53,1631951620,1631951620,0,NULL,'PATCH','/backend/v1/users/balance/adjust/:id',NULL,NULL),(55,1632279021,1632279021,0,NULL,'GET','/backend/v1/oss/token',NULL,NULL),(56,1646895007,1646895007,0,NULL,'POST','/backend/v1/file',NULL,NULL),(57,1646902642,1646902642,0,NULL,'GET','/backend/v1/file',NULL,NULL),(59,1646903295,1646903295,0,NULL,'GET','/backend/v1/version',NULL,NULL),(60,1646969187,1646969187,0,NULL,'GET','/backend/v1/file/:id',NULL,NULL),(61,1647586623,1647586623,0,NULL,'POST','/backend/v1/task',NULL,NULL),(62,1647591981,1647591981,0,NULL,'GET','/backend/v1/task',NULL,NULL),(63,1647591981,1647591981,0,NULL,'DELETE','/backend/v1/task/:id',NULL,NULL),(64,1648018316,1648018316,0,NULL,'POST','/backend/v1/file/run',NULL,NULL),(65,1648018316,1648018316,0,NULL,'PUT','/backend/v1/file/:id',NULL,NULL),(83,1648434653,1648434653,0,NULL,'DELETE','/backend/v1/file/:id',NULL,NULL);
/*!40000 ALTER TABLE `shard_route` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_task`
--

DROP TABLE IF EXISTS `shard_task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_task` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `name` longtext,
  `start_at` longtext,
  `end_at` longtext,
  `command` longtext,
  `remark` longtext,
  `status` tinyint DEFAULT NULL,
  `begin_at` bigint DEFAULT NULL,
  `spec` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_task`
--

LOCK TABLES `shard_task` WRITE;
/*!40000 ALTER TABLE `shard_task` DISABLE KEYS */;
INSERT INTO `shard_task` VALUES (5,1647589761,1647592166,0,'run_001','* * * * *','* * * * *','','',2,0,NULL),(6,1647592425,1647592425,0,'run_002','* * * * *','* * * * *','','',1,0,NULL),(7,1647592473,1647592473,0,'run_002','* * * * *','* * * * *','','',1,0,NULL),(8,1647593287,1647593287,0,'run_003','@every 1s','@every 1s','','',1,0,NULL),(9,1647593291,1647593291,0,'run_004','@every 1s','@every 1s','','',1,0,NULL),(10,1647593293,1647593293,0,'run_005','@every 1s','@every 1s','','',1,0,NULL);
/*!40000 ALTER TABLE `shard_task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_user`
--

DROP TABLE IF EXISTS `shard_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `mobile` longtext,
  `nick_name` longtext,
  `password` varchar(255) DEFAULT NULL,
  `email` longtext,
  `avatar` varchar(255) DEFAULT NULL,
  `user_name` longtext,
  `username` varchar(20) DEFAULT NULL,
  `last_login_time` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_user`
--

LOCK TABLES `shard_user` WRITE;
/*!40000 ALTER TABLE `shard_user` DISABLE KEYS */;
INSERT INTO `shard_user` VALUES (1,1630727921,1646894793,0,NULL,NULL,'$2a$10$xIprWgvyopi7ecBBHHDqn..S51rAvg7UmVAa5eGxUy56qsNesWkWS',NULL,'','admin','user1',0),(3,1631692728,1631692728,0,NULL,NULL,'$2a$10$egBK/A6TMBiiv7e9V7xAA.k2TOmd5R3JBsul0QcC9QrCUvEy8S1di',NULL,'',NULL,'asdada',0);
/*!40000 ALTER TABLE `shard_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_wallet`
--

DROP TABLE IF EXISTS `shard_wallet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_wallet` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `uid` bigint DEFAULT NULL,
  `balance` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_shard_wallet_uid` (`uid`),
  CONSTRAINT `fk_shard_user_wallet` FOREIGN KEY (`uid`) REFERENCES `shard_user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_wallet`
--

LOCK TABLES `shard_wallet` WRITE;
/*!40000 ALTER TABLE `shard_wallet` DISABLE KEYS */;
INSERT INTO `shard_wallet` VALUES (1,1630727933,1632201698,0,1,501),(3,1631954956,1632201702,0,3,400);
/*!40000 ALTER TABLE `shard_wallet` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shard_wallet_record`
--

DROP TABLE IF EXISTS `shard_wallet_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shard_wallet_record` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_time` bigint DEFAULT NULL,
  `updated_time` bigint DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `wallet_id` bigint DEFAULT NULL,
  `amount` bigint DEFAULT NULL,
  `record_type` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_shard_wallet_record_wallet_id` (`wallet_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shard_wallet_record`
--

LOCK TABLES `shard_wallet_record` WRITE;
/*!40000 ALTER TABLE `shard_wallet_record` DISABLE KEYS */;
INSERT INTO `shard_wallet_record` VALUES (1,1631954922,1631954922,0,0,1,NULL),(2,1631954945,1631954945,0,0,300,NULL),(3,1631954956,1631954956,0,3,100,NULL),(4,1631955131,1631955131,0,0,-300,NULL),(5,1631955258,1631955258,0,0,100,NULL),(6,1631955452,1631955452,0,0,100,NULL),(7,1631955456,1631955456,0,0,100,NULL),(8,1631955479,1631955479,0,0,100,NULL),(9,1632201698,1632201698,0,0,100,NULL),(10,1632201702,1632201702,0,0,300,NULL);
/*!40000 ALTER TABLE `shard_wallet_record` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-30 13:53:14
