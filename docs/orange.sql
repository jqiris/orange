/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 80027
Source Host           : 127.0.0.1:3306
Source Database       : orange

Target Server Type    : MYSQL
Target Server Version : 80027
File Encoding         : 65001

Date: 2022-03-09 12:19:31
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
-- Records of game_notices
-- ----------------------------
INSERT INTO `game_notices` VALUES ('notice', '幼麟游戏，卓越品质，值得信赖', '20161128');
INSERT INTO `game_notices` VALUES ('fkgm', '幼麟游戏，卓越品质，值得信赖', '20161128');

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
-- Records of mahjong_action_xzdds
-- ----------------------------
INSERT INTO `mahjong_action_xzdds` VALUES ('1646793756416112', '0', '', '{\"button\":0,\"game_seats\":[[13,7,23,25,5,18,8,8,17,21,25,14,19,2],[3,8,9,21,8,13,15,26,16,1,4,16,10],[26,5,20,26,25,7,0,10,7,14,10,20,19],[23,12,7,23,22,3,20,25,24,6,6,3,18]],\"index\":0,\"mahjongs\":[13,3,26,23,7,8,5,12,23,9,20,7,25,21,26,23,5,8,25,22,18,13,7,3,8,15,0,20,8,26,10,25,17,16,7,24,21,1,14,6,25,4,10,6,14,16,20,3,19,10,19,18,2,15,24,0,18,0,26,4,13,12,17,9,2,24,11,20,2,19,12,9,19,21,2,13,17,22,5,11,18,1,14,22,1,16,4,11,14,17,24,3,5,4,21,1,0,9,6,23,10,15,16,12,22,11,15,6],\"type\":\"xzdd\"}', '1646796627', '[0,1,13,1,2,15,1,1,21,2,2,24,2,1,10,3,2,0,3,1,12,0,2,18,0,1,14,1,2,0,1,1,26,2,3,26,2,1,10,3,2,26,3,1,0,0,2,4,0,1,17,1,2,13,1,1,9,2,2,12,2,1,14,3,2,17,3,1,17,0,2,9,0,1,9,1,2,2,1,1,10,2,2,24,2,1,12,3,2,11,3,1,11,0,2,20,0,1,23,3,3,23,3,1,18,0,3,18,0,1,2,1,2,2,1,1,13,2,2,19,2,1,0,3,2,12,3,1,12,0,2,9,0,1,9,1,2,19,1,1,19,2,3,19,2,1,5,3,2,21,3,1,6,0,2,2,0,1,2,1,3,2,1,1,4,2,2,13,2,1,13,3,2,17,3,1,17,0,2,22,0,1,7,2,3,7,2,1,24,3,2,5,3,6,5,0,2,11,0,1,11,1,2,18,1,1,18,2,2,1,2,1,1,0,2,14,0,1,14,1,2,22,1,1,22,2,2,1,2,1,1,0,2,16,0,1,16,1,3,16,1,1,0,2,2,4,2,1,4,0,2,11,0,1,11,1,2,14,1,1,1,2,2,17,2,1,17,0,2,24,0,1,4,1,2,3,1,1,3,2,2,5,2,1,5,0,2,4,0,1,19,1,2,21,1,1,21,2,2,1,2,1,1,0,2,0,0,1,0,1,2,9,1,1,9,2,2,6,2,1,6,0,2,23,0,1,25,1,2,10,1,1,10,2,2,15,2,1,15,1,3,15,1,1,3,0,5,3,1,2,16,1,4,16,1,2,12,1,6,12]', '[-1,3,-8,6]');

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
-- Records of mahjong_actions
-- ----------------------------
INSERT INTO `mahjong_actions` VALUES ('1646793756416112', '0', '', '{\"button\":0,\"game_seats\":[[13,7,23,25,5,18,8,8,17,21,25,14,19,2],[3,8,9,21,8,13,15,26,16,1,4,16,10],[26,5,20,26,25,7,0,10,7,14,10,20,19],[23,12,7,23,22,3,20,25,24,6,6,3,18]],\"index\":0,\"mahjongs\":[13,3,26,23,7,8,5,12,23,9,20,7,25,21,26,23,5,8,25,22,18,13,7,3,8,15,0,20,8,26,10,25,17,16,7,24,21,1,14,6,25,4,10,6,14,16,20,3,19,10,19,18,2,15,24,0,18,0,26,4,13,12,17,9,2,24,11,20,2,19,12,9,19,21,2,13,17,22,5,11,18,1,14,22,1,16,4,11,14,17,24,3,5,4,21,1,0,9,6,23,10,15,16,12,22,11,15,6],\"type\":\"xzdd\"}', '1646796627', '[0,1,13,1,2,15,1,1,21,2,2,24,2,1,10,3,2,0,3,1,12,0,2,18,0,1,14,1,2,0,1,1,26,2,3,26,2,1,10,3,2,26,3,1,0,0,2,4,0,1,17,1,2,13,1,1,9,2,2,12,2,1,14,3,2,17,3,1,17,0,2,9,0,1,9,1,2,2,1,1,10,2,2,24,2,1,12,3,2,11,3,1,11,0,2,20,0,1,23,3,3,23,3,1,18,0,3,18,0,1,2,1,2,2,1,1,13,2,2,19,2,1,0,3,2,12,3,1,12,0,2,9,0,1,9,1,2,19,1,1,19,2,3,19,2,1,5,3,2,21,3,1,6,0,2,2,0,1,2,1,3,2,1,1,4,2,2,13,2,1,13,3,2,17,3,1,17,0,2,22,0,1,7,2,3,7,2,1,24,3,2,5,3,6,5,0,2,11,0,1,11,1,2,18,1,1,18,2,2,1,2,1,1,0,2,14,0,1,14,1,2,22,1,1,22,2,2,1,2,1,1,0,2,16,0,1,16,1,3,16,1,1,0,2,2,4,2,1,4,0,2,11,0,1,11,1,2,14,1,1,1,2,2,17,2,1,17,0,2,24,0,1,4,1,2,3,1,1,3,2,2,5,2,1,5,0,2,4,0,1,19,1,2,21,1,1,21,2,2,1,2,1,1,0,2,0,0,1,0,1,2,9,1,1,9,2,2,6,2,1,6,0,2,23,0,1,25,1,2,10,1,1,10,2,2,15,2,1,15,1,3,15,1,1,3,0,5,3,1,2,16,1,4,16,1,2,12,1,6,12]', '[-1,3,-8,6]');

-- ----------------------------
-- Table structure for mahjong_room_xzdds
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_room_xzdds`;
CREATE TABLE `mahjong_room_xzdds` (
  `uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `room_id` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '房间号',
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
-- Records of mahjong_room_xzdds
-- ----------------------------
INSERT INTO `mahjong_room_xzdds` VALUES ('1646799009934721', '934721', 'xzdd', 'mahjong_5001', '0', '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', '1646799009', '0', '0', '1005', '', '56uv5pyo6ZuA5Zyj', '0', '0', '', '', '0', '0', '', '', '0', '0', '', '', '0');

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
-- Records of mahjong_rooms
-- ----------------------------
INSERT INTO `mahjong_rooms` VALUES ('1646799176796396', '796396', 'xzdd', 'mahjong_5001', '1005', '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', '1646799176', '0', '0', '1005', '', '56uv5pyo6ZuA5Zyj', '0', '1006', '', '56uv5pyo6ZuA5Zyj', '0', '1007', '', '54us5a2k6ZuA5Zyj', '0', '1008', '', '6ZW/5a2Z5LiN6L6T', '0');

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
) ENGINE=InnoDB AUTO_INCREMENT=1009 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_members
-- ----------------------------
INSERT INTO `user_members` VALUES ('1005', 'guest_jason1', '56uv5pyo6ZuA5Zyj', 'null', '0', '', '2022-03-09 10:36:10', '2022-03-09 10:36:10', '1', '0', '1000', '21', '0', '796396', '');
INSERT INTO `user_members` VALUES ('1006', 'guest_jason2', '56uv5pyo6ZuA5Zyj', 'null', '0', '', '2022-03-09 10:36:35', '2022-03-09 10:36:35', '1', '0', '1000', '19', '0', '796396', '');
INSERT INTO `user_members` VALUES ('1007', 'guest_jason3', '54us5a2k6ZuA5Zyj', 'null', '0', '', '2022-03-09 10:36:40', '2022-03-09 10:36:40', '1', '0', '1000', '21', '0', '796396', '');
INSERT INTO `user_members` VALUES ('1008', 'guest_jason4', '6ZW/5a2Z5LiN6L6T', 'null', '0', '', '2022-03-09 10:36:43', '2022-03-09 10:36:43', '1', '0', '1000', '21', '0', '796396', '');
