/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 80027
Source Host           : 127.0.0.1:3306
Source Database       : orange

Target Server Type    : MYSQL
Target Server Version : 80027
File Encoding         : 65001

Date: 2022-02-21 19:02:21
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user_members
-- ----------------------------
DROP TABLE IF EXISTS `user_members`;
CREATE TABLE `user_members` (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `nickname` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '头像',
  `gender` tinyint DEFAULT '0' COMMENT '性别 0-未知 1-男 2-女',
  `open_id` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '登陆open_id',
  `reg_time` datetime DEFAULT NULL COMMENT '注册时间',
  `login_time` datetime DEFAULT NULL COMMENT '登陆时间',
  `offline_state` tinyint DEFAULT '0' COMMENT '离线状态 0-在线 1-离线',
  `offline_time` datetime DEFAULT NULL COMMENT '离线时间',
  `is_new` tinyint(1) DEFAULT '0' COMMENT '是否是新用户 0-否 1-是',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_members
-- ----------------------------
