/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : localhost:3306
 Source Schema         : shard

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 14/09/2021 16:07:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for shard_admin
-- ----------------------------
DROP TABLE IF EXISTS `shard_admin`;
CREATE TABLE `shard_admin` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
                               `created_time` bigint(20) DEFAULT NULL,
                               `updated_time` bigint(20) DEFAULT NULL,
                               `is_deleted` tinyint(1) DEFAULT '0',
                               `password` varchar(255) DEFAULT NULL,
                               `avatar` varchar(255) DEFAULT NULL,
                               `username` varchar(20) DEFAULT NULL,
                               `last_login_time` bigint(20) DEFAULT NULL,
                               `role_id` bigint(20) DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               KEY `idx_shard_admin_role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_admin
-- ----------------------------
BEGIN;
INSERT INTO `shard_admin` VALUES (1, 1630981091, 1631603273, 0,  '$2a$10$rmY3VZuXJl2/Wz7okf.ttuS1fCqv4shS2JIz.afwQVkU420gWotxi', '', 'admin', 1631603273, 1);
INSERT INTO `shard_admin` VALUES (2, 1631276335, 1631589859, 0,  '$2a$10$g/BTPQIe/Fdy9QQN6nyO5OiYGvjZ6gBnU6yec844P2NdRk9KN2sG2', '', 'admin2', 1631589859, 2);
COMMIT;


-- ----------------------------
-- Table structure for shard_permission
-- ----------------------------
DROP TABLE IF EXISTS `shard_permission`;
CREATE TABLE `shard_permission` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                    `created_time` bigint(20) DEFAULT NULL,
                                    `updated_time` bigint(20) DEFAULT NULL,
                                    `is_deleted` tinyint(1) DEFAULT '0',
                                    `name` longtext,
                                    `parent_id` bigint(20) DEFAULT NULL,
                                    `key` longtext,
                                    `icon` longtext,
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_permission
-- ----------------------------
BEGIN;
INSERT INTO `shard_permission` VALUES (5, 1631101512, 1631110015, 0, '管理', 0, '', 'el-icon-s-tools');
INSERT INTO `shard_permission` VALUES (22, 1631110062, 1631263105, 0, '权限管理', 5, '1_0_21313131231', 'el-icon-s-tools');
INSERT INTO `shard_permission` VALUES (25, 1631110090, 1631273169, 0, '路由管理', 5, '1_1_1631110079332', 'el-icon-s-cooperation');
INSERT INTO `shard_permission` VALUES (26, 1631171043, 1631263096, 0, '添加权限', 22, '22_1_1631171034438', '');
INSERT INTO `shard_permission` VALUES (27, 1631189154, 1631263090, 0, '删除权限', 22, '22_2_1631189149693', '');
INSERT INTO `shard_permission` VALUES (28, 1631260058, 1631260058, 0, '人员管理', 0, '0_0_1631260049220', 'el-icon-user');
INSERT INTO `shard_permission` VALUES (29, 1631260067, 1631273182, 0, '用户列表', 28, '28_1_1631260060784', 'el-icon-user-solid');
INSERT INTO `shard_permission` VALUES (30, 1631260078, 1631273188, 0, '管理员列表', 28, '28_2_1631260069872', 'el-icon-s-custom');
INSERT INTO `shard_permission` VALUES (31, 1631260127, 1631273224, 0, '角色管理', 28, '28_3_1631260121599', 'el-icon-s-check');
INSERT INTO `shard_permission` VALUES (32, 1631263083, 1631263083, 0, '编辑权限', 22, '22_3_1631263050215', '');
INSERT INTO `shard_permission` VALUES (33, 1631276880, 1631276904, 0, '新增', 30, '30_1_1631276859399', '');
INSERT INTO `shard_permission` VALUES (34, 1631276901, 1631276933, 0, '编辑', 30, '30_2_1631276890248', '');
INSERT INTO `shard_permission` VALUES (35, 1631276919, 1631276919, 0, '删除', 30, '30_3_1631276905329', '');
INSERT INTO `shard_permission` VALUES (36, 1631276958, 1631276989, 0, '变更角色', 30, '30_4_1631276938272', '');
INSERT INTO `shard_permission` VALUES (37, 1631276984, 1631279403, 0, '变更权限', 30, '30_5_1631276967393', '');
INSERT INTO `shard_permission` VALUES (38, 1631277288, 1631277288, 0, '重置密码', 30, '30_6_1631277265508', '');
INSERT INTO `shard_permission` VALUES (40, 1631277638, 1631277638, 0, '新增', 31, '31_1_1631277588024', '');
INSERT INTO `shard_permission` VALUES (41, 1631277677, 1631277677, 0, '编辑', 31, '31_2_1631277662260', '');
INSERT INTO `shard_permission` VALUES (42, 1631277686, 1631277686, 0, '删除', 31, '31_3_1631277678919', '');
INSERT INTO `shard_permission` VALUES (43, 1631277740, 1631279397, 0, '变更权限', 31, '30_5_1631276967393', '');
INSERT INTO `shard_permission` VALUES (44, 1631589542, 1631589542, 0, '新增', 29, '29_1_1631589517500', '');
INSERT INTO `shard_permission` VALUES (45, 1631589557, 1631589557, 0, '编辑', 29, '29_2_1631589543207', '');
INSERT INTO `shard_permission` VALUES (46, 1631589568, 1631589568, 0, '删除', 29, '29_3_1631589558934', '');
INSERT INTO `shard_permission` VALUES (47, 1631589586, 1631589586, 0, '重置密码', 29, '29_4_1631589569386', '');
COMMIT;

-- ----------------------------
-- Table structure for shard_permission_route
-- ----------------------------
DROP TABLE IF EXISTS `shard_permission_route`;
CREATE TABLE `shard_permission_route` (
                                          `permission_id` bigint(20) NOT NULL,
                                          `route_id` bigint(20) NOT NULL,
                                          PRIMARY KEY (`permission_id`,`route_id`),
                                          KEY `fk_shard_permission_route_route` (`route_id`),
                                          CONSTRAINT `fk_shard_permission_route_permission` FOREIGN KEY (`permission_id`) REFERENCES `shard_permission` (`id`),
                                          CONSTRAINT `fk_shard_permission_route_route` FOREIGN KEY (`route_id`) REFERENCES `shard_route` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_permission_route
-- ----------------------------
BEGIN;
INSERT INTO `shard_permission_route` VALUES (22, 2);
INSERT INTO `shard_permission_route` VALUES (26, 3);
INSERT INTO `shard_permission_route` VALUES (29, 6);
INSERT INTO `shard_permission_route` VALUES (30, 9);
INSERT INTO `shard_permission_route` VALUES (33, 15);
INSERT INTO `shard_permission_route` VALUES (34, 15);
INSERT INTO `shard_permission_route` VALUES (44, 17);
INSERT INTO `shard_permission_route` VALUES (35, 31);
INSERT INTO `shard_permission_route` VALUES (45, 32);
INSERT INTO `shard_permission_route` VALUES (34, 33);
INSERT INTO `shard_permission_route` VALUES (27, 34);
INSERT INTO `shard_permission_route` VALUES (46, 35);
INSERT INTO `shard_permission_route` VALUES (32, 36);
INSERT INTO `shard_permission_route` VALUES (40, 38);
INSERT INTO `shard_permission_route` VALUES (41, 39);
INSERT INTO `shard_permission_route` VALUES (31, 40);
INSERT INTO `shard_permission_route` VALUES (42, 41);
INSERT INTO `shard_permission_route` VALUES (37, 43);
INSERT INTO `shard_permission_route` VALUES (43, 43);
INSERT INTO `shard_permission_route` VALUES (37, 44);
INSERT INTO `shard_permission_route` VALUES (43, 44);
INSERT INTO `shard_permission_route` VALUES (25, 45);
INSERT INTO `shard_permission_route` VALUES (38, 48);
INSERT INTO `shard_permission_route` VALUES (36, 51);
INSERT INTO `shard_permission_route` VALUES (47, 52);
COMMIT;

-- ----------------------------
-- Table structure for shard_role
-- ----------------------------
DROP TABLE IF EXISTS `shard_role`;
CREATE TABLE `shard_role` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `created_time` bigint(20) DEFAULT NULL,
                              `updated_time` bigint(20) DEFAULT NULL,
                              `is_deleted` tinyint(1) DEFAULT '0',
                              `name` longtext,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_role
-- ----------------------------
BEGIN;
INSERT INTO `shard_role` VALUES (1, 1631168588, 1631589630, 0, '超级管理员');
INSERT INTO `shard_role` VALUES (2, 1631168810, 1631589848, 0, '普通的页面仔');
INSERT INTO `shard_role` VALUES (3, 1631537737, 1631537737, 0, 'admin3');
INSERT INTO `shard_role` VALUES (4, 1631537742, 1631537742, 0, 'admin4');
INSERT INTO `shard_role` VALUES (5, 1631537745, 1631537745, 0, 'admin5');
INSERT INTO `shard_role` VALUES (6, 1631537749, 1631537749, 0, 'admin7');
INSERT INTO `shard_role` VALUES (7, 1631537753, 1631537753, 0, 'admin8');
INSERT INTO `shard_role` VALUES (8, 1631537758, 1631537758, 0, 'admin9');
INSERT INTO `shard_role` VALUES (9, 1631537764, 1631537764, 0, 'admin10');
INSERT INTO `shard_role` VALUES (10, 1631538572, 1631538572, 0, 'admin11');
INSERT INTO `shard_role` VALUES (11, 1631538576, 1631538576, 0, 'admin12');
COMMIT;

-- ----------------------------
-- Table structure for shard_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `shard_role_permission`;
CREATE TABLE `shard_role_permission` (
                                         `role_id` bigint(20) NOT NULL,
                                         `permission_id` bigint(20) NOT NULL,
                                         PRIMARY KEY (`role_id`,`permission_id`),
                                         KEY `fk_shard_role_permission_permission` (`permission_id`),
                                         CONSTRAINT `fk_shard_role_permission_permission` FOREIGN KEY (`permission_id`) REFERENCES `shard_permission` (`id`),
                                         CONSTRAINT `fk_shard_role_permission_role` FOREIGN KEY (`role_id`) REFERENCES `shard_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_role_permission
-- ----------------------------
BEGIN;
INSERT INTO `shard_role_permission` VALUES (1, 5);
INSERT INTO `shard_role_permission` VALUES (1, 22);
INSERT INTO `shard_role_permission` VALUES (1, 25);
INSERT INTO `shard_role_permission` VALUES (1, 26);
INSERT INTO `shard_role_permission` VALUES (1, 27);
INSERT INTO `shard_role_permission` VALUES (1, 28);
INSERT INTO `shard_role_permission` VALUES (2, 28);
INSERT INTO `shard_role_permission` VALUES (1, 29);
INSERT INTO `shard_role_permission` VALUES (2, 29);
INSERT INTO `shard_role_permission` VALUES (1, 30);
INSERT INTO `shard_role_permission` VALUES (2, 30);
INSERT INTO `shard_role_permission` VALUES (1, 31);
INSERT INTO `shard_role_permission` VALUES (2, 31);
INSERT INTO `shard_role_permission` VALUES (1, 32);
INSERT INTO `shard_role_permission` VALUES (1, 33);
INSERT INTO `shard_role_permission` VALUES (2, 33);
INSERT INTO `shard_role_permission` VALUES (1, 34);
INSERT INTO `shard_role_permission` VALUES (2, 34);
INSERT INTO `shard_role_permission` VALUES (1, 35);
INSERT INTO `shard_role_permission` VALUES (2, 35);
INSERT INTO `shard_role_permission` VALUES (1, 36);
INSERT INTO `shard_role_permission` VALUES (2, 36);
INSERT INTO `shard_role_permission` VALUES (1, 37);
INSERT INTO `shard_role_permission` VALUES (2, 37);
INSERT INTO `shard_role_permission` VALUES (1, 38);
INSERT INTO `shard_role_permission` VALUES (2, 38);
INSERT INTO `shard_role_permission` VALUES (1, 40);
INSERT INTO `shard_role_permission` VALUES (2, 40);
INSERT INTO `shard_role_permission` VALUES (1, 41);
INSERT INTO `shard_role_permission` VALUES (2, 41);
INSERT INTO `shard_role_permission` VALUES (1, 42);
INSERT INTO `shard_role_permission` VALUES (2, 42);
INSERT INTO `shard_role_permission` VALUES (1, 43);
INSERT INTO `shard_role_permission` VALUES (2, 43);
INSERT INTO `shard_role_permission` VALUES (1, 44);
INSERT INTO `shard_role_permission` VALUES (2, 44);
INSERT INTO `shard_role_permission` VALUES (1, 45);
INSERT INTO `shard_role_permission` VALUES (2, 45);
INSERT INTO `shard_role_permission` VALUES (1, 46);
INSERT INTO `shard_role_permission` VALUES (2, 46);
INSERT INTO `shard_role_permission` VALUES (1, 47);
INSERT INTO `shard_role_permission` VALUES (2, 47);
COMMIT;

-- ----------------------------
-- Table structure for shard_route
-- ----------------------------
DROP TABLE IF EXISTS `shard_route`;
CREATE TABLE `shard_route` (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
                               `created_time` bigint(20) DEFAULT NULL,
                               `updated_time` bigint(20) DEFAULT NULL,
                               `is_deleted` tinyint(1) DEFAULT '0',
                               `name` varchar(20) DEFAULT NULL,
                               `method` varchar(20) DEFAULT NULL,
                               `path` varchar(50) DEFAULT NULL,
                               `permission_id` bigint(20) DEFAULT NULL,
                               `route_id` bigint(20) DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               KEY `fk_shard_route_permission` (`permission_id`),
                               KEY `fk_shard_route_route` (`route_id`),
                               CONSTRAINT `fk_shard_route_permission` FOREIGN KEY (`permission_id`) REFERENCES `shard_permission` (`id`),
                               CONSTRAINT `fk_shard_route_route` FOREIGN KEY (`route_id`) REFERENCES `shard_route` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_route
-- ----------------------------
BEGIN;
INSERT INTO `shard_route` VALUES (2, 1631017285, 1631023195, 0, '获取所有权限', 'GET', '/backend/v1/permissions', NULL, NULL);
INSERT INTO `shard_route` VALUES (3, 1631017285, 1631017285, 0, '', 'POST', '/backend/v1/permissions', NULL, NULL);
INSERT INTO `shard_route` VALUES (6, 1631017285, 1631017285, 0, '', 'GET', '/backend/v1/users', NULL, NULL);
INSERT INTO `shard_route` VALUES (9, 1631017285, 1631017285, 0, '', 'GET', '/backend/v1/admins', NULL, NULL);
INSERT INTO `shard_route` VALUES (11, 1631017285, 1631017285, 0, '', 'POST', '/backend/v1/admins/login', NULL, NULL);
INSERT INTO `shard_route` VALUES (14, 1631017285, 1631017285, 0, '', 'PUT', '/backend/v1/admins/avatar', NULL, NULL);
INSERT INTO `shard_route` VALUES (15, 1631017285, 1631017285, 0, '', 'POST', '/backend/v1/admins', NULL, NULL);
INSERT INTO `shard_route` VALUES (17, 1631017285, 1631017285, 0, '', 'POST', '/backend/v1/users', NULL, NULL);
INSERT INTO `shard_route` VALUES (31, 1631018134, 1631018134, 0, '', 'DELETE', '/backend/v1/admins/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (32, 1631018134, 1631018134, 0, '', 'PUT', '/backend/v1/users/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (33, 1631018134, 1631018134, 0, '', 'PUT', '/backend/v1/admins/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (34, 1631018134, 1631018134, 0, '', 'DELETE', '/backend/v1/permissions/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (35, 1631018134, 1631018134, 0, '', 'DELETE', '/backend/v1/users/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (36, 1631018134, 1631018134, 0, '', 'PUT', '/backend/v1/permissions/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (38, 1631157768, 1631157768, 0, NULL, 'POST', '/backend/v1/roles', NULL, NULL);
INSERT INTO `shard_route` VALUES (39, 1631157768, 1631157768, 0, NULL, 'PUT', '/backend/v1/roles/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (40, 1631157768, 1631157768, 0, NULL, 'GET', '/backend/v1/roles', NULL, NULL);
INSERT INTO `shard_route` VALUES (41, 1631157768, 1631157768, 0, NULL, 'DELETE', '/backend/v1/roles/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (43, 1631170482, 1631170482, 0, NULL, 'GET', '/backend/v1/roles/permissions/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (44, 1631170681, 1631170681, 0, NULL, 'PUT', '/backend/v1/roles/permissions/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (45, 1631199982, 1631199982, 0, NULL, 'GET', '/backend/v1/routes', NULL, NULL);
INSERT INTO `shard_route` VALUES (48, 1631202091, 1631202091, 0, NULL, 'GET', '/backend/v1/admins/password/reset/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (49, 1631202123, 1631202123, 0, NULL, 'PUT', '/backend/v1/admins/password/change', NULL, NULL);
INSERT INTO `shard_route` VALUES (51, 1631242698, 1631242698, 0, NULL, 'PUT', '/backend/v1/admins/roles/change/:id', NULL, NULL);
INSERT INTO `shard_route` VALUES (52, 1631589013, 1631589013, 0, NULL, 'GET', '/backend/v1/users/password/reset/:id', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for shard_user
-- ----------------------------
DROP TABLE IF EXISTS `shard_user`;
CREATE TABLE `shard_user` (
                              `id` bigint(20) NOT NULL AUTO_INCREMENT,
                              `created_time` bigint(20) DEFAULT NULL,
                              `updated_time` bigint(20) DEFAULT NULL,
                              `is_deleted` tinyint(1) DEFAULT '0',
                              `mobile` longtext,
                              `nick_name` longtext,
                              `password` varchar(255) DEFAULT NULL,
                              `email` longtext,
                              `avatar` varchar(255) DEFAULT NULL,
                              `user_name` longtext,
                              `username` varchar(20) DEFAULT NULL,
                              `last_login_time` bigint(20) DEFAULT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_user
-- ----------------------------
BEGIN;
INSERT INTO `shard_user` VALUES (1, 1630727921, 1630985167, 0, NULL, NULL, '$2a$10$oBW4P.v6dMzauKiXBnKmPugh/kUOsE76u7/UtWp1TIhqEbC/KsqYq', NULL, '', 'admin', 'abcdeabcdeabcdeabcde', 0);
INSERT INTO `shard_user` VALUES (2, 1630984741, 1631088195, 0, NULL, NULL, '$2a$10$wsEDuEohmfUCk2WuxQcdle5Xv8FixK7Bqt3ToPKnI5UxNIFo/6fWO', NULL, '', NULL, 'user1', 1631088195);
COMMIT;

-- ----------------------------
-- Table structure for shard_wallet
-- ----------------------------
DROP TABLE IF EXISTS `shard_wallet`;
CREATE TABLE `shard_wallet` (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                `created_time` bigint(20) DEFAULT NULL,
                                `updated_time` bigint(20) DEFAULT NULL,
                                `is_deleted` tinyint(1) DEFAULT '0',
                                `uid` bigint(20) DEFAULT NULL,
                                `balance` bigint(20) DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                KEY `idx_shard_wallet_uid` (`uid`),
                                CONSTRAINT `fk_shard_user_wallet` FOREIGN KEY (`uid`) REFERENCES `shard_user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of shard_wallet
-- ----------------------------

BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
