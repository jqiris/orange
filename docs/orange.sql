/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 80027
Source Host           : 127.0.0.1:3306
Source Database       : orange

Target Server Type    : MYSQL
Target Server Version : 80027
File Encoding         : 65001

Date: 2022-03-08 19:37:08
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for game_notices
-- ----------------------------
DROP TABLE IF EXISTS `game_notices`;
CREATE TABLE `game_notices` (
  `type` varchar(32) COLLATE utf8mb4_general_ci NOT NULL,
  `msg` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL,
  `version` varchar(32) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`type`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mahjong_action_xzdds
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_action_xzdds`;
CREATE TABLE `mahjong_action_xzdds` (
  `room_uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `game_index` smallint NOT NULL,
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `base_info` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` int NOT NULL,
  `action_records` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `result` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`room_uuid`,`game_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mahjong_actions
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_actions`;
CREATE TABLE `mahjong_actions` (
  `room_uuid` char(20) COLLATE utf8mb4_general_ci NOT NULL,
  `game_index` smallint NOT NULL,
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `base_info` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` int NOT NULL,
  `action_records` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `result` char(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`room_uuid`,`game_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mahjong_room_xzdds
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_room_xzdds`;
CREATE TABLE `mahjong_room_xzdds` (
  `uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `room_id` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '房间号',
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `server_id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器id',
  `landlord` int NOT NULL DEFAULT '0' COMMENT '房主',
  `base_info` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '配置基本信息',
  `create_time` int NOT NULL COMMENT '创建时间',
  `num_of_turns` int NOT NULL DEFAULT '0' COMMENT '第几轮',
  `next_button` int NOT NULL DEFAULT '0' COMMENT '下一位',
  `user_id0` int NOT NULL DEFAULT '0' COMMENT '用户1',
  `user_icon0` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name0` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score0` int NOT NULL DEFAULT '0',
  `user_id1` int NOT NULL DEFAULT '0',
  `user_icon1` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name1` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score1` int NOT NULL DEFAULT '0',
  `user_id2` int NOT NULL DEFAULT '0',
  `user_icon2` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name2` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score2` int NOT NULL DEFAULT '0',
  `user_id3` int NOT NULL DEFAULT '0',
  `user_icon3` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name3` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score3` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`uuid`),
  UNIQUE KEY `uuid_UNIQUE` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mahjong_rooms
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_rooms`;
CREATE TABLE `mahjong_rooms` (
  `uuid` char(20) COLLATE utf8mb4_general_ci NOT NULL,
  `room_id` char(8) COLLATE utf8mb4_general_ci NOT NULL COMMENT '房间号',
  `game_type` char(8) COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `server_id` varchar(45) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器id',
  `landlord` int NOT NULL DEFAULT '0' COMMENT '房主',
  `base_info` varchar(256) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '配置基本信息',
  `create_time` int NOT NULL COMMENT '创建时间',
  `num_of_turns` int NOT NULL DEFAULT '0' COMMENT '第几轮',
  `next_button` int NOT NULL DEFAULT '0' COMMENT '下一位',
  `user_id0` int NOT NULL DEFAULT '0' COMMENT '用户1',
  `user_icon0` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name0` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score0` int NOT NULL DEFAULT '0',
  `user_id1` int NOT NULL DEFAULT '0',
  `user_icon1` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name1` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score1` int NOT NULL DEFAULT '0',
  `user_id2` int NOT NULL DEFAULT '0',
  `user_icon2` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name2` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score2` int NOT NULL DEFAULT '0',
  `user_id3` int NOT NULL DEFAULT '0',
  `user_icon3` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name3` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score3` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`uuid`),
  UNIQUE KEY `uuid_UNIQUE` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for user_members
-- ----------------------------
DROP TABLE IF EXISTS `user_members`;
CREATE TABLE `user_members` (
  `user_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `account` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称',
  `headimg` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '头像',
  `sex` tinyint DEFAULT '0' COMMENT '性别 0-未知 1-男 2-女',
  `open_id` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '登陆open_id',
  `reg_time` datetime DEFAULT NULL COMMENT '注册时间',
  `login_time` datetime DEFAULT NULL COMMENT '登陆时间',
  `lv` smallint DEFAULT '1' COMMENT '用户等级',
  `exp` int DEFAULT '0' COMMENT '用户经验',
  `coins` int DEFAULT '0' COMMENT '用户金币',
  `gems` int DEFAULT '0' COMMENT '用户宝石',
  `room_type` tinyint NOT NULL DEFAULT '0' COMMENT '最后进入房间类型 0-麻将 1-扑克',
  `room_id` char(8) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '最后进入房间号',
  `history` varchar(4096) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作记录',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1003 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
