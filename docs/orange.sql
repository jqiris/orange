/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 127.0.0.1:3306
 Source Schema         : orange

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 23/03/2022 09:55:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auto_code_examples
-- ----------------------------
DROP TABLE IF EXISTS `auto_code_examples`;
CREATE TABLE `auto_code_examples`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `auto_code_example_field` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '仅作示例条目无实际作用',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_auto_code_examples_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auto_code_examples
-- ----------------------------

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/resetPassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/email/emailTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/importExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/loadExcel', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/exportExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/excel/downloadTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/server/getServerList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/mahjongCardSettings/createMahjongCardSettings', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/mahjongCardSettings/deleteMahjongCardSettings', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/mahjongCardSettings/deleteMahjongCardSettingsByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/mahjongCardSettings/updateMahjongCardSettings', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/mahjongCardSettings/findMahjongCardSettings', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/mahjongCardSettings/getMahjongCardSettingsList', 'GET', '', '', '');

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `exa_file_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `file_chunk_number` bigint(20) NULL DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2022-03-10 17:52:54.151', '2022-03-10 17:52:54.151', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '2022-03-10 17:52:54.151', '2022-03-10 17:52:54.151', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `chunk_total` bigint(20) NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_files
-- ----------------------------

-- ----------------------------
-- Table structure for game_notices
-- ----------------------------
DROP TABLE IF EXISTS `game_notices`;
CREATE TABLE `game_notices`  (
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `msg` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `version` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`type`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of game_notices
-- ----------------------------
INSERT INTO `game_notices` VALUES ('notice', '幼麟游戏，卓越品质，值得信赖', '20161128');
INSERT INTO `game_notices` VALUES ('fkgm', '幼麟游戏，卓越品质，值得信赖', '20161128');

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (1, '2022-03-10 17:56:31.459', '2022-03-10 17:56:31.459', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzIyN2NkNjktNjcyNC00YTZjLTg1YjQtOTg3NjA4Njg5OWU5IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0NzUxMDg3NSwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQ2OTA1MDc1fQ.kirwg1I7dxLXQadInbEea3QwkYCVjY6YfmOQ19I_-hk');
INSERT INTO `jwt_blacklists` VALUES (2, '2022-03-10 17:56:31.461', '2022-03-10 17:56:31.461', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzIyN2NkNjktNjcyNC00YTZjLTg1YjQtOTg3NjA4Njg5OWU5IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0NzUxMDg3NSwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQ2OTA1MDc1fQ.kirwg1I7dxLXQadInbEea3QwkYCVjY6YfmOQ19I_-hk');
INSERT INTO `jwt_blacklists` VALUES (3, '2022-03-10 18:37:37.305', '2022-03-10 18:37:37.305', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzIyN2NkNjktNjcyNC00YTZjLTg1YjQtOTg3NjA4Njg5OWU5IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0NzUxMzA1OSwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQ2OTA3MjU5fQ.OjylhThAYAFONBA54M8X4B6KXp1yD0Kg6Cx0LXlENoo');
INSERT INTO `jwt_blacklists` VALUES (4, '2022-03-11 23:57:54.300', '2022-03-11 23:57:54.300', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzIyN2NkNjktNjcyNC00YTZjLTg1YjQtOTg3NjA4Njg5OWU5IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6ODY0MDAsImV4cCI6MTY0NzYxNjgxNCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNjQ3MDExMDE0fQ.UXrjBOWeqwrJxFUlS6fv-0FBl_1EzpAbZYt3pHKhrYk');

-- ----------------------------
-- Table structure for mahjong_action_xlchs
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_action_xlchs`;
CREATE TABLE `mahjong_action_xlchs`  (
  `room_uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `game_index` smallint(6) NOT NULL,
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `base_info` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` int(11) NOT NULL,
  `action_records` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `result` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`room_uuid`, `game_index`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahjong_action_xlchs
-- ----------------------------
INSERT INTO `mahjong_action_xlchs` VALUES ('1647339653281277', 0, 'xlch', '{\"button\":0,\"game_seats\":[[17,3,25,26,0,10,14,4,4,26,14,1,9,8],[19,11,22,1,6,0,17,21,13,17,18,9,23],[5,21,22,15,24,23,16,14,3,10,17,18,24],[20,25,2,16,2,6,2,15,21,25,7,21,14]],\"index\":0,\"mahjongs\":[17,19,5,20,3,11,21,25,25,22,22,2,26,1,15,16,0,6,24,2,10,0,23,6,14,17,16,2,4,21,14,15,4,13,3,21,26,17,10,25,14,18,17,7,1,9,18,21,9,23,24,14,8,4,8,5,11,26,18,24,24,13,19,12,4,3,10,12,7,12,22,20,16,18,10,9,2,16,13,5,19,5,6,26,19,20,7,7,11,1,15,12,25,0,13,6,9,11,23,8,1,20,23,15,0,22,8,3],\"type\":\"xlch\"}', 1647342388, '[0,1,25,3,3,25,3,1,14,0,3,14,0,1,26,1,2,4,1,1,0,2,2,8,2,1,3,3,2,5,3,1,15,0,2,11,0,1,26,1,2,26,1,1,1,2,2,18,2,1,5,3,2,24,3,1,16,0,2,24,0,1,24,2,3,24,2,1,8,3,2,13,3,1,13,0,2,19,0,1,19,1,2,12,1,1,4,0,3,4,0,1,17,1,3,17,1,1,6,2,2,4,2,1,4,3,2,3,3,1,24,0,2,10,0,1,8,1,2,12,1,1,26,2,2,7,2,1,7,3,2,12,3,1,12,1,3,12,1,1,9,2,2,22,2,1,10,3,2,20,3,1,3,0,2,16,0,1,16,1,2,18,1,1,11,2,2,10,2,1,10,3,2,9,3,1,9,0,2,2,0,1,10,1,2,16,1,1,19,2,2,13,2,1,22,3,2,5,3,1,5,0,2,19,0,1,19,1,2,5,1,1,5,2,2,6,2,1,6,3,2,26,3,1,26,0,2,19,0,1,19,1,2,20,1,1,13,2,2,7,2,1,7,3,2,7,3,1,7,0,2,11,0,1,11,1,2,1,1,1,1,2,2,15,2,6,15,3,2,12,3,1,12,2,5,12,3,2,25,3,4,25,3,2,0,3,1,0,0,5,0,1,2,13,1,1,13,2,2,6,2,1,6,3,2,9,3,1,9,0,2,11,0,1,11,1,2,23,1,1,16,2,2,8,2,1,8,3,2,1,3,1,1,0,2,20,0,1,20,3,5,20,0,2,23,0,1,23,1,5,23,2,2,15,2,6,15,3,2,0,3,1,0,0,5,0,1,2,22,1,1,22,2,2,8,2,1,8,3,2,3,3,1,3,0,5,3]', '[0,-4,12,-8]');

-- ----------------------------
-- Table structure for mahjong_action_xzdds
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_action_xzdds`;
CREATE TABLE `mahjong_action_xzdds`  (
  `room_uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `game_index` smallint(6) NOT NULL,
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `base_info` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` int(11) NOT NULL,
  `action_records` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `result` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`room_uuid`, `game_index`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahjong_action_xzdds
-- ----------------------------
INSERT INTO `mahjong_action_xzdds` VALUES ('1646793756416112', 0, '', '{\"button\":0,\"game_seats\":[[13,7,23,25,5,18,8,8,17,21,25,14,19,2],[3,8,9,21,8,13,15,26,16,1,4,16,10],[26,5,20,26,25,7,0,10,7,14,10,20,19],[23,12,7,23,22,3,20,25,24,6,6,3,18]],\"index\":0,\"mahjongs\":[13,3,26,23,7,8,5,12,23,9,20,7,25,21,26,23,5,8,25,22,18,13,7,3,8,15,0,20,8,26,10,25,17,16,7,24,21,1,14,6,25,4,10,6,14,16,20,3,19,10,19,18,2,15,24,0,18,0,26,4,13,12,17,9,2,24,11,20,2,19,12,9,19,21,2,13,17,22,5,11,18,1,14,22,1,16,4,11,14,17,24,3,5,4,21,1,0,9,6,23,10,15,16,12,22,11,15,6],\"type\":\"xzdd\"}', 1646796627, '[0,1,13,1,2,15,1,1,21,2,2,24,2,1,10,3,2,0,3,1,12,0,2,18,0,1,14,1,2,0,1,1,26,2,3,26,2,1,10,3,2,26,3,1,0,0,2,4,0,1,17,1,2,13,1,1,9,2,2,12,2,1,14,3,2,17,3,1,17,0,2,9,0,1,9,1,2,2,1,1,10,2,2,24,2,1,12,3,2,11,3,1,11,0,2,20,0,1,23,3,3,23,3,1,18,0,3,18,0,1,2,1,2,2,1,1,13,2,2,19,2,1,0,3,2,12,3,1,12,0,2,9,0,1,9,1,2,19,1,1,19,2,3,19,2,1,5,3,2,21,3,1,6,0,2,2,0,1,2,1,3,2,1,1,4,2,2,13,2,1,13,3,2,17,3,1,17,0,2,22,0,1,7,2,3,7,2,1,24,3,2,5,3,6,5,0,2,11,0,1,11,1,2,18,1,1,18,2,2,1,2,1,1,0,2,14,0,1,14,1,2,22,1,1,22,2,2,1,2,1,1,0,2,16,0,1,16,1,3,16,1,1,0,2,2,4,2,1,4,0,2,11,0,1,11,1,2,14,1,1,1,2,2,17,2,1,17,0,2,24,0,1,4,1,2,3,1,1,3,2,2,5,2,1,5,0,2,4,0,1,19,1,2,21,1,1,21,2,2,1,2,1,1,0,2,0,0,1,0,1,2,9,1,1,9,2,2,6,2,1,6,0,2,23,0,1,25,1,2,10,1,1,10,2,2,15,2,1,15,1,3,15,1,1,3,0,5,3,1,2,16,1,4,16,1,2,12,1,6,12]', '[-1,3,-8,6]');
INSERT INTO `mahjong_action_xzdds` VALUES ('1646799176796396', 0, '', '{\"button\":0,\"game_seats\":[[11,15,19,20,8,18,7,13,4,9,1,7,25,25],[5,16,15,23,2,3,26,8,21,14,17,26,3],[2,4,17,25,9,22,22,16,19,9,24,24,18],[3,25,1,4,21,7,10,1,13,12,20,9,11]],\"index\":0,\"mahjongs\":[11,5,2,3,15,16,4,25,19,15,17,1,20,23,25,4,8,2,9,21,18,3,22,7,7,26,22,10,13,8,16,1,4,21,19,13,9,14,9,12,1,17,24,20,7,26,24,9,25,3,18,11,25,26,4,3,24,16,0,19,13,10,23,12,11,17,5,19,11,10,14,24,5,13,10,0,14,22,8,15,6,26,7,21,23,6,12,16,14,2,23,6,2,17,5,8,6,18,12,0,18,21,0,20,20,22,15,1],\"type\":\"xzdd\"}', 1646800245, '[0,1,9,2,3,9,2,1,2,3,2,26,3,1,26,0,2,4,0,1,11,1,2,3,1,1,26,2,2,24,2,1,4,0,3,4,0,1,15,1,2,16,1,1,26,2,2,0,2,1,0,3,2,19,3,1,19,0,2,13,0,1,13,1,2,10,1,1,23,2,2,23,2,1,17,3,2,12,3,1,20,0,2,11,0,1,13,1,2,17,1,1,21,2,2,5,2,1,5,3,2,19,3,1,19,0,2,11,0,1,11,1,2,10,1,1,8,2,2,14,2,1,18,3,2,24,3,1,21,0,2,5,0,1,11,1,2,13,1,1,5,2,2,10,2,1,10,1,3,10,1,1,2,2,2,0,2,1,0,3,2,14,3,1,25,0,3,25,0,1,1,3,3,1,3,1,24,2,4,24,2,2,22,2,1,25,3,2,8,3,1,8,0,2,15,0,1,15,1,2,6,1,1,6,2,2,26,2,1,26,3,2,7,3,1,12,0,2,21,0,1,18,1,2,23,1,1,23,2,2,6,2,1,6,3,2,12,3,1,9,0,2,16,0,1,16,1,5,16,2,2,14,2,1,19,3,2,2,3,6,2,0,2,23,0,1,8,2,2,6,2,1,6,0,2,2,0,1,2,2,2,17,2,1,23,0,2,5,0,1,23,2,2,8,2,1,8,0,2,6,0,1,6,2,2,18,2,1,18,0,2,12,0,1,12,2,2,0,2,1,0,0,2,18,0,1,18,2,2,21,2,1,21,0,2,0,0,1,0,2,2,20,2,1,20,0,2,20,0,1,20,2,2,22,2,4,22,2,2,15,2,6,15]', '[-14,1,11,2]');

-- ----------------------------
-- Table structure for mahjong_actions
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_actions`;
CREATE TABLE `mahjong_actions`  (
  `room_uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `game_index` smallint(6) NOT NULL,
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `base_info` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` int(11) NOT NULL,
  `action_records` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `result` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`room_uuid`, `game_index`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahjong_actions
-- ----------------------------
INSERT INTO `mahjong_actions` VALUES ('1646793756416112', 0, '', '{\"button\":0,\"game_seats\":[[13,7,23,25,5,18,8,8,17,21,25,14,19,2],[3,8,9,21,8,13,15,26,16,1,4,16,10],[26,5,20,26,25,7,0,10,7,14,10,20,19],[23,12,7,23,22,3,20,25,24,6,6,3,18]],\"index\":0,\"mahjongs\":[13,3,26,23,7,8,5,12,23,9,20,7,25,21,26,23,5,8,25,22,18,13,7,3,8,15,0,20,8,26,10,25,17,16,7,24,21,1,14,6,25,4,10,6,14,16,20,3,19,10,19,18,2,15,24,0,18,0,26,4,13,12,17,9,2,24,11,20,2,19,12,9,19,21,2,13,17,22,5,11,18,1,14,22,1,16,4,11,14,17,24,3,5,4,21,1,0,9,6,23,10,15,16,12,22,11,15,6],\"type\":\"xzdd\"}', 1646796627, '[0,1,13,1,2,15,1,1,21,2,2,24,2,1,10,3,2,0,3,1,12,0,2,18,0,1,14,1,2,0,1,1,26,2,3,26,2,1,10,3,2,26,3,1,0,0,2,4,0,1,17,1,2,13,1,1,9,2,2,12,2,1,14,3,2,17,3,1,17,0,2,9,0,1,9,1,2,2,1,1,10,2,2,24,2,1,12,3,2,11,3,1,11,0,2,20,0,1,23,3,3,23,3,1,18,0,3,18,0,1,2,1,2,2,1,1,13,2,2,19,2,1,0,3,2,12,3,1,12,0,2,9,0,1,9,1,2,19,1,1,19,2,3,19,2,1,5,3,2,21,3,1,6,0,2,2,0,1,2,1,3,2,1,1,4,2,2,13,2,1,13,3,2,17,3,1,17,0,2,22,0,1,7,2,3,7,2,1,24,3,2,5,3,6,5,0,2,11,0,1,11,1,2,18,1,1,18,2,2,1,2,1,1,0,2,14,0,1,14,1,2,22,1,1,22,2,2,1,2,1,1,0,2,16,0,1,16,1,3,16,1,1,0,2,2,4,2,1,4,0,2,11,0,1,11,1,2,14,1,1,1,2,2,17,2,1,17,0,2,24,0,1,4,1,2,3,1,1,3,2,2,5,2,1,5,0,2,4,0,1,19,1,2,21,1,1,21,2,2,1,2,1,1,0,2,0,0,1,0,1,2,9,1,1,9,2,2,6,2,1,6,0,2,23,0,1,25,1,2,10,1,1,10,2,2,15,2,1,15,1,3,15,1,1,3,0,5,3,1,2,16,1,4,16,1,2,12,1,6,12]', '[-1,3,-8,6]');
INSERT INTO `mahjong_actions` VALUES ('1646799176796396', 0, '', '{\"button\":0,\"game_seats\":[[11,15,19,20,8,18,7,13,4,9,1,7,25,25],[5,16,15,23,2,3,26,8,21,14,17,26,3],[2,4,17,25,9,22,22,16,19,9,24,24,18],[3,25,1,4,21,7,10,1,13,12,20,9,11]],\"index\":0,\"mahjongs\":[11,5,2,3,15,16,4,25,19,15,17,1,20,23,25,4,8,2,9,21,18,3,22,7,7,26,22,10,13,8,16,1,4,21,19,13,9,14,9,12,1,17,24,20,7,26,24,9,25,3,18,11,25,26,4,3,24,16,0,19,13,10,23,12,11,17,5,19,11,10,14,24,5,13,10,0,14,22,8,15,6,26,7,21,23,6,12,16,14,2,23,6,2,17,5,8,6,18,12,0,18,21,0,20,20,22,15,1],\"type\":\"xzdd\"}', 1646800245, '[0,1,9,2,3,9,2,1,2,3,2,26,3,1,26,0,2,4,0,1,11,1,2,3,1,1,26,2,2,24,2,1,4,0,3,4,0,1,15,1,2,16,1,1,26,2,2,0,2,1,0,3,2,19,3,1,19,0,2,13,0,1,13,1,2,10,1,1,23,2,2,23,2,1,17,3,2,12,3,1,20,0,2,11,0,1,13,1,2,17,1,1,21,2,2,5,2,1,5,3,2,19,3,1,19,0,2,11,0,1,11,1,2,10,1,1,8,2,2,14,2,1,18,3,2,24,3,1,21,0,2,5,0,1,11,1,2,13,1,1,5,2,2,10,2,1,10,1,3,10,1,1,2,2,2,0,2,1,0,3,2,14,3,1,25,0,3,25,0,1,1,3,3,1,3,1,24,2,4,24,2,2,22,2,1,25,3,2,8,3,1,8,0,2,15,0,1,15,1,2,6,1,1,6,2,2,26,2,1,26,3,2,7,3,1,12,0,2,21,0,1,18,1,2,23,1,1,23,2,2,6,2,1,6,3,2,12,3,1,9,0,2,16,0,1,16,1,5,16,2,2,14,2,1,19,3,2,2,3,6,2,0,2,23,0,1,8,2,2,6,2,1,6,0,2,2,0,1,2,2,2,17,2,1,23,0,2,5,0,1,23,2,2,8,2,1,8,0,2,6,0,1,6,2,2,18,2,1,18,0,2,12,0,1,12,2,2,0,2,1,0,0,2,18,0,1,18,2,2,21,2,1,21,0,2,0,0,1,0,2,2,20,2,1,20,0,2,20,0,1,20,2,2,22,2,4,22,2,2,15,2,6,15]', '[-14,1,11,2]');
INSERT INTO `mahjong_actions` VALUES ('1647339653281277', 0, 'xlch', '{\"button\":0,\"game_seats\":[[17,3,25,26,0,10,14,4,4,26,14,1,9,8],[19,11,22,1,6,0,17,21,13,17,18,9,23],[5,21,22,15,24,23,16,14,3,10,17,18,24],[20,25,2,16,2,6,2,15,21,25,7,21,14]],\"index\":0,\"mahjongs\":[17,19,5,20,3,11,21,25,25,22,22,2,26,1,15,16,0,6,24,2,10,0,23,6,14,17,16,2,4,21,14,15,4,13,3,21,26,17,10,25,14,18,17,7,1,9,18,21,9,23,24,14,8,4,8,5,11,26,18,24,24,13,19,12,4,3,10,12,7,12,22,20,16,18,10,9,2,16,13,5,19,5,6,26,19,20,7,7,11,1,15,12,25,0,13,6,9,11,23,8,1,20,23,15,0,22,8,3],\"type\":\"xlch\"}', 1647342388, '[0,1,25,3,3,25,3,1,14,0,3,14,0,1,26,1,2,4,1,1,0,2,2,8,2,1,3,3,2,5,3,1,15,0,2,11,0,1,26,1,2,26,1,1,1,2,2,18,2,1,5,3,2,24,3,1,16,0,2,24,0,1,24,2,3,24,2,1,8,3,2,13,3,1,13,0,2,19,0,1,19,1,2,12,1,1,4,0,3,4,0,1,17,1,3,17,1,1,6,2,2,4,2,1,4,3,2,3,3,1,24,0,2,10,0,1,8,1,2,12,1,1,26,2,2,7,2,1,7,3,2,12,3,1,12,1,3,12,1,1,9,2,2,22,2,1,10,3,2,20,3,1,3,0,2,16,0,1,16,1,2,18,1,1,11,2,2,10,2,1,10,3,2,9,3,1,9,0,2,2,0,1,10,1,2,16,1,1,19,2,2,13,2,1,22,3,2,5,3,1,5,0,2,19,0,1,19,1,2,5,1,1,5,2,2,6,2,1,6,3,2,26,3,1,26,0,2,19,0,1,19,1,2,20,1,1,13,2,2,7,2,1,7,3,2,7,3,1,7,0,2,11,0,1,11,1,2,1,1,1,1,2,2,15,2,6,15,3,2,12,3,1,12,2,5,12,3,2,25,3,4,25,3,2,0,3,1,0,0,5,0,1,2,13,1,1,13,2,2,6,2,1,6,3,2,9,3,1,9,0,2,11,0,1,11,1,2,23,1,1,16,2,2,8,2,1,8,3,2,1,3,1,1,0,2,20,0,1,20,3,5,20,0,2,23,0,1,23,1,5,23,2,2,15,2,6,15,3,2,0,3,1,0,0,5,0,1,2,22,1,1,22,2,2,8,2,1,8,3,2,3,3,1,3,0,5,3]', '[0,-4,12,-8]');
INSERT INTO `mahjong_actions` VALUES ('1647962588644426', 0, 'xzdd', '{\"button\":0,\"game_seats\":[[2,3,3,5,6,7,10,15,17,20,21,23,25,26],[2,3,7,9,10,13,13,21,21,23,23,24,24],[2,6,8,8,10,12,16,17,18,18,19,20,23],[0,3,5,6,7,9,11,12,13,15,16,19,25]],\"index\":0,\"mahjongs\":[22,22,4,4,4,0,22,1,11,1,22,17,9,12,1,8,19,17,0,13,21,15,14,20,4,11,26,14,16,16,5,26,14,14,6,25,25,18,9,18,24,7,0,5,11,15,1,26,24,8,20,10,19,2,12],\"type\":\"xzdd\"}', 1647997137, '[0,1,10,1,2,22,1,1,2,2,2,22,2,1,10,3,2,4,3,1,19,0,2,4,0,1,15,1,2,4,1,1,3,2,2,0,2,1,12,3,2,22,3,1,25,0,2,1,0,1,17,1,2,11,1,1,4,2,2,1,2,1,16,3,2,22,3,1,22,0,2,17,0,1,17,1,2,9,1,1,7,2,2,12,2,1,17,3,2,1,3,1,22,0,2,8,0,1,26,1,2,19,1,1,19,2,2,17,2,1,17,3,2,0,3,1,9,0,2,13,0,1,13,1,3,13,1,1,9,2,2,21,2,1,12,3,2,15,3,1,1,0,2,14,0,1,14,1,2,20,1,1,21,2,2,4,2,1,4,3,2,11,3,1,11,0,2,26,0,1,26,1,2,14,1,1,14,2,2,16,2,1,16,3,2,16,3,1,0,0,2,5,0,1,25,1,2,26,1,1,26,2,2,14,2,1,14,3,2,14,3,1,0,0,2,6,0,1,23,1,5,23,2,2,25,2,1,6,3,2,25,3,1,25,0,2,18,0,1,18,2,3,18,2,1,25,3,2,9,3,1,9,0,2,18,0,1,18,2,5,18,3,2,24,3,1,24,0,2,7,0,1,7,3,2,0,3,1,0,0,2,5,0,1,21,3,2,11,3,1,3,0,2,15,0,1,15,3,3,15,3,1,4,0,2,1,0,1,20,3,2,26,3,1,26,0,2,24,0,1,24,3,2,8,3,1,8,0,2,20,0,1,20,3,2,10,3,1,10,0,2,19,0,1,19,3,2,2,3,1,2,0,5,2]', '[1,1,2,-4]');

-- ----------------------------
-- Table structure for mahjong_card_settings
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_card_settings`;
CREATE TABLE `mahjong_card_settings`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `seat_one` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '座位1牌型',
  `seat_two` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '座位2牌型',
  `seat_three` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '座位3牌型',
  `seat_four` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '座位4牌型',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahjong_card_settings
-- ----------------------------
INSERT INTO `mahjong_card_settings` VALUES (1, '2-5-11-15-27-40-48-58-67-71-80-92-96-107', '0-29-33-38-49-51-52-60-64-69-70-72-89', '9-10-20-30-44-46-55-63-68-84-86-91-97', '9-10-13-56-66-75-76-78-82-88-93-100-103', '2022-03-22 20:23:26.981', '2022-03-22 20:23:26.981', '2022-03-22 20:25:37.079');
INSERT INTO `mahjong_card_settings` VALUES (2, '11-14-15-21-25-31-42-63-69-81-87-93-100-104', '9-12-30-37-40-54-55-85-86-94-95-98-99', '10-27-34-35-43-51-64-71-73-75-77-83-92', '0-13-22-26-28-36-45-48-53-62-65-79-101', '2022-03-22 20:26:23.370', '2022-03-22 21:11:11.078', NULL);
INSERT INTO `mahjong_card_settings` VALUES (3, '22-24-44-51-57-58-59-74-77-80-90-91-94-95', '1-7-19-27-31-41-43-50-72-87-93-100-104', '3-5-34-42-49-53-55-68-76-86-102-105-107', '6-13-32-35-37-38-39-45-54-60-64-79-103', '2022-03-22 21:17:21.106', '2022-03-22 21:17:21.106', NULL);
INSERT INTO `mahjong_card_settings` VALUES (4, '7-9-11-15-16-17-27-31-34-38-41-43-71-81', '2-20-24-29-50-52-54-56-59-61-62-78-85', '18-19-33-36-46-51-60-64-69-72-84-93-98', '1-3-6-8-21-63-80-89-90-91-92-100-101', '2022-03-22 21:19:59.330', '2022-03-22 21:19:59.330', NULL);

-- ----------------------------
-- Table structure for mahjong_room_xlchs
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_room_xlchs`;
CREATE TABLE `mahjong_room_xlchs`  (
  `uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `room_id` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '房间号',
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `server_id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器id',
  `landlord` int(11) NOT NULL DEFAULT 0 COMMENT '房主',
  `base_info` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '配置基本信息',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `num_of_turns` int(11) NOT NULL DEFAULT 0 COMMENT '第几轮',
  `next_button` int(11) NOT NULL DEFAULT 0 COMMENT '下一位',
  `user_id0` int(11) NOT NULL DEFAULT 0 COMMENT '用户1',
  `user_icon0` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name0` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score0` int(11) NOT NULL DEFAULT 0,
  `user_id1` int(11) NOT NULL DEFAULT 0,
  `user_icon1` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name1` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score1` int(11) NOT NULL DEFAULT 0,
  `user_id2` int(11) NOT NULL DEFAULT 0,
  `user_icon2` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name2` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score2` int(11) NOT NULL DEFAULT 0,
  `user_id3` int(11) NOT NULL DEFAULT 0,
  `user_icon3` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name3` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score3` int(11) NOT NULL DEFAULT 0,
  `card_id` int(11) NOT NULL DEFAULT 0 COMMENT '牌型id',
  PRIMARY KEY (`uuid`) USING BTREE,
  UNIQUE INDEX `uuid_UNIQUE`(`uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahjong_room_xlchs
-- ----------------------------
INSERT INTO `mahjong_room_xlchs` VALUES ('1647339653281277', '281277', 'xlch', 'mahjong_3001', 1012, '{\"type\":\"xlch\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647339653, 1, 3, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 1011, '', '5Lic6Ziz56iz6LWi', 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0);

-- ----------------------------
-- Table structure for mahjong_room_xzdds
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_room_xzdds`;
CREATE TABLE `mahjong_room_xzdds`  (
  `uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `room_id` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '房间号',
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `server_id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器id',
  `landlord` int(11) NOT NULL DEFAULT 0 COMMENT '房主',
  `base_info` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '配置基本信息',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `num_of_turns` int(11) NOT NULL DEFAULT 0 COMMENT '第几轮',
  `next_button` int(11) NOT NULL DEFAULT 0 COMMENT '下一位',
  `user_id0` int(11) NOT NULL DEFAULT 0 COMMENT '用户1',
  `user_icon0` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name0` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score0` int(11) NOT NULL DEFAULT 0,
  `user_id1` int(11) NOT NULL DEFAULT 0,
  `user_icon1` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name1` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score1` int(11) NOT NULL DEFAULT 0,
  `user_id2` int(11) NOT NULL DEFAULT 0,
  `user_icon2` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name2` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score2` int(11) NOT NULL DEFAULT 0,
  `user_id3` int(11) NOT NULL DEFAULT 0,
  `user_icon3` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name3` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score3` int(11) NOT NULL DEFAULT 0,
  `card_id` int(11) NOT NULL DEFAULT 0 COMMENT '牌型id',
  PRIMARY KEY (`uuid`) USING BTREE,
  UNIQUE INDEX `uuid_UNIQUE`(`uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahjong_room_xzdds
-- ----------------------------
INSERT INTO `mahjong_room_xzdds` VALUES ('1646799009934721', '934721', 'xzdd', 'mahjong_5001', 0, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646799009, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646799176796396', '796396', 'xzdd', 'mahjong_5001', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646799176, 1, 2, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646807113049132', '049132', 'xzdd', 'mahjong_4001', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646807113, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646807161613799', '613799', 'xzdd', 'mahjong_4002', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646807161, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646808441833455', '833455', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646808441, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646808662854645', '854645', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646808662, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646808719257995', '257995', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646808719, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646809027070251', '070251', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646809027, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646809127971069', '971069', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646809127, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646809202034456', '034456', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646809202, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646809225973353', '973353', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646809225, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646809813290861', '290861', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646809813, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646811193087836', '087836', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646811193, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646815841370839', '370839', 'xzdd', 'mahjong_4001', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646815841, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646815915240631', '240631', 'xzdd', 'mahjong_4002', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646815915, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646816822034940', '034940', 'xzdd', 'mahjong_4002', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646816822, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646816926833998', '833998', 'xzdd', 'mahjong_4002', 1008, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1008}', 1646816926, 0, 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646820327336012', '336012', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646820327, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646820484725278', '725278', 'xzdd', 'mahjong_4001', 1006, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1006}', 1646820484, 0, 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646821990042712', '042712', 'xzdd', 'mahjong_4001', 1008, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1008}', 1646821990, 0, 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646827806141663', '141663', 'xzdd', 'mahjong_4001', 1006, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1006}', 1646827806, 0, 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646827869599643', '599643', 'xzdd', 'mahjong_4001', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646827869, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646827972016211', '016211', 'xzdd', 'mahjong_4002', 1008, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1008}', 1646827972, 0, 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646875872546082', '546082', 'xzdd', 'mahjong_4001', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646875872, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646877744492910', '492910', 'xzdd', 'mahjong_4002', 1006, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1006}', 1646877744, 0, 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646877815752155', '752155', 'xzdd', 'mahjong_4001', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646877815, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646878082947657', '947657', 'xzdd', 'mahjong_4001', 1008, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1008}', 1646878082, 0, 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646878211699992', '699992', 'xzdd', 'mahjong_4002', 1009, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1009}', 1646878211, 0, 0, 1009, '', '5a6H5paH56iz6LWi', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646881653824504', '824504', 'xzdd', 'mahjong_4002', 1006, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1006}', 1646881653, 0, 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646881921417437', '417437', 'xzdd', 'mahjong_4001', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646881921, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 1011, '', '5Lic6Ziz56iz6LWi', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646882039381355', '381355', 'xzdd', 'mahjong_4001', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646882039, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646882396902126', '902126', 'xzdd', 'mahjong_4001', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646882396, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646882490638048', '638048', 'xzdd', 'mahjong_4002', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646882490, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646884168500383', '500383', 'xzdd', 'mahjong_4001', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646884168, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646894021254803', '254803', 'xzdd', 'mahjong_4001', 1011, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1011}', 1646894021, 0, 0, 1011, '', '5Lic6Ziz56iz6LWi', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646894032724564', '724564', 'xzdd', 'mahjong_4002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1646894032, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646894098220577', '220577', 'xzdd', 'mahjong_4001', 1007, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1007}', 1646894098, 0, 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646894243670300', '670300', 'xzdd', 'mahjong_4001', 1008, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1008}', 1646894243, 0, 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1646894282758157', '758157', 'xzdd', 'mahjong_4002', 1008, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1008}', 1646894282, 0, 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647017705941461', '941461', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647017705, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647958582161152', '161152', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647958582, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647959064914295', '914295', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647959064, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647959326548819', '548819', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647959326, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647959574134997', '134997', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647959574, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647959602804279', '804279', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647959602, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647959818397525', '397525', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647959818, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647959877553236', '553236', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647959877, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647959958035674', '035674', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647959958, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647960408413728', '413728', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647960408, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647960844219099', '219099', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647960844, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647961759440135', '440135', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647961759, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647961808766169', '766169', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647961808, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647961903037608', '037608', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647961903, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647961951929595', '929595', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647961951, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_room_xzdds` VALUES ('1647962020190693', '190693', 'xzdd', 'mahjong_3002', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647962020, 0, 0, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);

-- ----------------------------
-- Table structure for mahjong_rooms
-- ----------------------------
DROP TABLE IF EXISTS `mahjong_rooms`;
CREATE TABLE `mahjong_rooms`  (
  `uuid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `room_id` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '房间号',
  `game_type` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '游戏类型 xzdd-血战到底',
  `server_id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器id',
  `landlord` int(11) NOT NULL DEFAULT 0 COMMENT '房主',
  `base_info` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '配置基本信息',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `num_of_turns` int(11) NOT NULL DEFAULT 0 COMMENT '第几轮',
  `next_button` int(11) NOT NULL DEFAULT 0 COMMENT '下一位',
  `user_id0` int(11) NOT NULL DEFAULT 0 COMMENT '用户1',
  `user_icon0` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name0` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score0` int(11) NOT NULL DEFAULT 0,
  `user_id1` int(11) NOT NULL DEFAULT 0,
  `user_icon1` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name1` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score1` int(11) NOT NULL DEFAULT 0,
  `user_id2` int(11) NOT NULL DEFAULT 0,
  `user_icon2` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name2` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score2` int(11) NOT NULL DEFAULT 0,
  `user_id3` int(11) NOT NULL DEFAULT 0,
  `user_icon3` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_name3` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_score3` int(11) NOT NULL DEFAULT 0,
  `card_id` int(11) NOT NULL DEFAULT 0 COMMENT '牌型id',
  PRIMARY KEY (`uuid`) USING BTREE,
  UNIQUE INDEX `uuid_UNIQUE`(`uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahjong_rooms
-- ----------------------------
INSERT INTO `mahjong_rooms` VALUES ('1646882008593506', '593506', 'xzdd', 'mahjong_4002', 1006, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1006}', 1646882008, 0, 0, 1006, '', '56uv5pyo6ZuA5Zyj', 0, 0, '', '', 0, 0, '', '', 0, 0, '', '', 0, 0);
INSERT INTO `mahjong_rooms` VALUES ('1647692749041190', '041190', 'xzdd', 'mahjong_3002', 1005, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1647692749, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1011, '', '5Lic6Ziz56iz6LWi', 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0);
INSERT INTO `mahjong_rooms` VALUES ('1647789431884310', '884310', 'xlch', 'mahjong_3001', 1005, '{\"type\":\"xlch\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1005}', 1647789431, 0, 0, 1005, '', '56uv5pyo6ZuA5Zyj', 0, 1011, '', '5Lic6Ziz56iz6LWi', 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 0);
INSERT INTO `mahjong_rooms` VALUES ('1647962588644426', '644426', 'xzdd', 'mahjong_3001', 1012, '{\"type\":\"xzdd\",\"baseScore\":1,\"maxFan\":4,\"maxGames\":2,\"creator\":1012}', 1647962588, 1, 1, 1012, '', '6ZW/5a2Z6Ieq5pG4', 0, 1011, '', '5Lic6Ziz56iz6LWi', 0, 1007, '', '54us5a2k6ZuA5Zyj', 0, 1008, '', '6ZW/5a2Z5LiN6L6T', 0, 2);

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 95 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/base/login', '用户登录(必选)', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES (4, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (6, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (7, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (8, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES (9, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2022-03-10 17:52:52.155', '2022-03-10 18:38:44.327', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (20, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES (24, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (27, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (37, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (38, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (40, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (41, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (42, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (45, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (46, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (47, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` VALUES (48, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` VALUES (49, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` VALUES (50, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (51, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (52, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (53, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (54, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (55, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (56, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (57, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (58, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (59, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (60, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (61, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` VALUES (62, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` VALUES (63, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` VALUES (64, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (65, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (66, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES (67, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` VALUES (68, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` VALUES (69, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (70, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (71, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` VALUES (72, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (73, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (74, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (75, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (76, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST');
INSERT INTO `sys_apis` VALUES (77, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (78, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (79, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (80, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/email/emailSend', '发送邮件示例', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (81, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/excel/importExcel', '导入excel', 'excel', 'POST');
INSERT INTO `sys_apis` VALUES (82, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/excel/loadExcel', '下载excel', 'excel', 'GET');
INSERT INTO `sys_apis` VALUES (83, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/excel/exportExcel', '导出excel', 'excel', 'POST');
INSERT INTO `sys_apis` VALUES (84, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/excel/downloadTemplate', '下载excel模板', 'excel', 'GET');
INSERT INTO `sys_apis` VALUES (85, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (86, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (87, '2022-03-10 17:52:52.155', '2022-03-10 17:52:52.155', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (88, '2022-03-12 00:07:41.287', '2022-03-12 00:07:41.287', NULL, '/server/getServerList', '获取服务器列表', '服务器', 'POST');
INSERT INTO `sys_apis` VALUES (89, '2022-03-21 11:08:41.534', '2022-03-21 11:08:41.534', NULL, '/mahjongCardSettings/createMahjongCardSettings', '新增mahjongCardSettings表', 'mahjongCardSettings', 'POST');
INSERT INTO `sys_apis` VALUES (90, '2022-03-21 11:08:41.541', '2022-03-21 11:08:41.541', NULL, '/mahjongCardSettings/deleteMahjongCardSettings', '删除mahjongCardSettings表', 'mahjongCardSettings', 'DELETE');
INSERT INTO `sys_apis` VALUES (91, '2022-03-21 11:08:41.549', '2022-03-21 11:08:41.549', NULL, '/mahjongCardSettings/deleteMahjongCardSettingsByIds', '批量删除mahjongCardSettings表', 'mahjongCardSettings', 'DELETE');
INSERT INTO `sys_apis` VALUES (92, '2022-03-21 11:08:41.554', '2022-03-21 11:08:41.554', NULL, '/mahjongCardSettings/updateMahjongCardSettings', '更新mahjongCardSettings表', 'mahjongCardSettings', 'PUT');
INSERT INTO `sys_apis` VALUES (93, '2022-03-21 11:08:41.558', '2022-03-21 11:08:41.558', NULL, '/mahjongCardSettings/findMahjongCardSettings', '根据ID获取mahjongCardSettings表', 'mahjongCardSettings', 'GET');
INSERT INTO `sys_apis` VALUES (94, '2022-03-21 11:08:41.566', '2022-03-21 11:08:41.566', NULL, '/mahjongCardSettings/getMahjongCardSettingsList', '获取mahjongCardSettings表列表', 'mahjongCardSettings', 'GET');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2022-03-10 17:52:53.671', '2022-03-21 11:49:23.009', NULL, '888', '普通用户', '0', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2022-03-10 17:52:53.671', '2022-03-12 00:09:53.193', NULL, '8881', '普通用户子角色', '888', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2022-03-10 17:52:53.671', '2022-03-12 00:09:58.812', NULL, '9528', '测试角色', '0', 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns`  (
  `authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, '888');
INSERT INTO `sys_authority_menus` VALUES (1, '8881');
INSERT INTO `sys_authority_menus` VALUES (1, '9528');
INSERT INTO `sys_authority_menus` VALUES (2, '888');
INSERT INTO `sys_authority_menus` VALUES (2, '8881');
INSERT INTO `sys_authority_menus` VALUES (2, '9528');
INSERT INTO `sys_authority_menus` VALUES (3, '888');
INSERT INTO `sys_authority_menus` VALUES (3, '8881');
INSERT INTO `sys_authority_menus` VALUES (3, '9528');
INSERT INTO `sys_authority_menus` VALUES (4, '888');
INSERT INTO `sys_authority_menus` VALUES (4, '9528');
INSERT INTO `sys_authority_menus` VALUES (5, '888');
INSERT INTO `sys_authority_menus` VALUES (5, '9528');
INSERT INTO `sys_authority_menus` VALUES (6, '888');
INSERT INTO `sys_authority_menus` VALUES (6, '9528');
INSERT INTO `sys_authority_menus` VALUES (7, '888');
INSERT INTO `sys_authority_menus` VALUES (7, '9528');
INSERT INTO `sys_authority_menus` VALUES (8, '888');
INSERT INTO `sys_authority_menus` VALUES (8, '8881');
INSERT INTO `sys_authority_menus` VALUES (8, '9528');
INSERT INTO `sys_authority_menus` VALUES (9, '888');
INSERT INTO `sys_authority_menus` VALUES (9, '9528');
INSERT INTO `sys_authority_menus` VALUES (10, '888');
INSERT INTO `sys_authority_menus` VALUES (10, '9528');
INSERT INTO `sys_authority_menus` VALUES (11, '888');
INSERT INTO `sys_authority_menus` VALUES (11, '9528');
INSERT INTO `sys_authority_menus` VALUES (12, '888');
INSERT INTO `sys_authority_menus` VALUES (12, '9528');
INSERT INTO `sys_authority_menus` VALUES (13, '888');
INSERT INTO `sys_authority_menus` VALUES (14, '888');
INSERT INTO `sys_authority_menus` VALUES (14, '9528');
INSERT INTO `sys_authority_menus` VALUES (15, '888');
INSERT INTO `sys_authority_menus` VALUES (15, '9528');
INSERT INTO `sys_authority_menus` VALUES (16, '888');
INSERT INTO `sys_authority_menus` VALUES (16, '9528');
INSERT INTO `sys_authority_menus` VALUES (17, '888');
INSERT INTO `sys_authority_menus` VALUES (17, '9528');
INSERT INTO `sys_authority_menus` VALUES (18, '888');
INSERT INTO `sys_authority_menus` VALUES (19, '888');
INSERT INTO `sys_authority_menus` VALUES (20, '888');
INSERT INTO `sys_authority_menus` VALUES (22, '888');
INSERT INTO `sys_authority_menus` VALUES (23, '888');
INSERT INTO `sys_authority_menus` VALUES (24, '888');
INSERT INTO `sys_authority_menus` VALUES (25, '888');
INSERT INTO `sys_authority_menus` VALUES (26, '888');
INSERT INTO `sys_authority_menus` VALUES (26, '8881');
INSERT INTO `sys_authority_menus` VALUES (26, '9528');
INSERT INTO `sys_authority_menus` VALUES (27, '888');

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `request_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `auto_code_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `injection_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `struct_cn_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `flag` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_histories_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------
INSERT INTO `sys_auto_code_histories` VALUES (1, '2022-03-21 11:08:41.962', '2022-03-21 11:08:41.962', NULL, 'mahjong_card_settings', '{\"structName\":\"MahjongCardSettings\",\"tableName\":\"mahjong_card_settings\",\"packageName\":\"mahjongCardSettings\",\"humpPackageName\":\"mahjong_card_settings\",\"abbreviation\":\"mahjongCardSettings\",\"description\":\"mahjongCardSettings表\",\"autoCreateApiToSql\":true,\"autoMoveFile\":false,\"fields\":[{\"fieldName\":\"SeatOne\",\"fieldDesc\":\"座位1牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatOne\",\"dataTypeLong\":\"45\",\"comment\":\"座位1牌型\",\"columnName\":\"seat_one\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatTwo\",\"fieldDesc\":\"座位2牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatTwo\",\"dataTypeLong\":\"45\",\"comment\":\"座位2牌型\",\"columnName\":\"seat_two\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatThree\",\"fieldDesc\":\"座位3牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatThree\",\"dataTypeLong\":\"45\",\"comment\":\"座位3牌型\",\"columnName\":\"seat_three\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatFour\",\"fieldDesc\":\"座位4牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatFour\",\"dataTypeLong\":\"45\",\"comment\":\"座位4牌型\",\"columnName\":\"seat_four\",\"fieldSearchType\":\"\",\"dictType\":\"\"}]}', '', '', 'MahjongCardSettings', 'mahjongCardSettings表', '89;90;91;92;93;94;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (2, '2022-03-21 11:13:26.810', '2022-03-21 11:13:26.810', NULL, 'mahjong_card_settings', '{\"structName\":\"MahjongCardSettings\",\"tableName\":\"mahjong_card_settings\",\"packageName\":\"mahjongCardSettings\",\"humpPackageName\":\"mahjong_card_settings\",\"abbreviation\":\"mahjongCardSettings\",\"description\":\"mahjongCardSettings表\",\"autoCreateApiToSql\":true,\"autoMoveFile\":false,\"fields\":[{\"fieldName\":\"SeatOne\",\"fieldDesc\":\"座位1牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatOne\",\"dataTypeLong\":\"45\",\"comment\":\"座位1牌型\",\"columnName\":\"seat_one\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatTwo\",\"fieldDesc\":\"座位2牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatTwo\",\"dataTypeLong\":\"45\",\"comment\":\"座位2牌型\",\"columnName\":\"seat_two\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatThree\",\"fieldDesc\":\"座位3牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatThree\",\"dataTypeLong\":\"45\",\"comment\":\"座位3牌型\",\"columnName\":\"seat_three\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatFour\",\"fieldDesc\":\"座位4牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatFour\",\"dataTypeLong\":\"45\",\"comment\":\"座位4牌型\",\"columnName\":\"seat_four\",\"fieldSearchType\":\"\",\"dictType\":\"\"}]}', '', '', 'MahjongCardSettings', 'mahjongCardSettings表', '', 0);
INSERT INTO `sys_auto_code_histories` VALUES (3, '2022-03-21 11:31:26.874', '2022-03-21 11:31:26.874', NULL, 'mahjong_card_settings', '{\"structName\":\"MahjongCardSettings\",\"tableName\":\"mahjong_card_settings\",\"packageName\":\"mahjongCardSettings\",\"humpPackageName\":\"mahjong_card_settings\",\"abbreviation\":\"mahjongCardSettings\",\"description\":\"mahjongCardSettings表\",\"autoCreateApiToSql\":true,\"autoMoveFile\":false,\"fields\":[{\"fieldName\":\"SeatOne\",\"fieldDesc\":\"座位1牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatOne\",\"dataTypeLong\":\"45\",\"comment\":\"座位1牌型\",\"columnName\":\"seat_one\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatTwo\",\"fieldDesc\":\"座位2牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatTwo\",\"dataTypeLong\":\"45\",\"comment\":\"座位2牌型\",\"columnName\":\"seat_two\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatThree\",\"fieldDesc\":\"座位3牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatThree\",\"dataTypeLong\":\"45\",\"comment\":\"座位3牌型\",\"columnName\":\"seat_three\",\"fieldSearchType\":\"\",\"dictType\":\"\"},{\"fieldName\":\"SeatFour\",\"fieldDesc\":\"座位4牌型\",\"fieldType\":\"string\",\"fieldJson\":\"seatFour\",\"dataTypeLong\":\"45\",\"comment\":\"座位4牌型\",\"columnName\":\"seat_four\",\"fieldSearchType\":\"\",\"dictType\":\"\"}]}', '', '', 'MahjongCardSettings', 'mahjongCardSettings表', '', 0);

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_btns_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `menu_level` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '仪表盘', 'odometer', 0);
INSERT INTO `sys_base_menus` VALUES (2, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 7, 0, 0, '关于我们', 'info-filled', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, 0, 0, '超级管理员', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, 1, 0, '菜单管理', 'tickets', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 1, 0, 'api管理', 'platform', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `sys_base_menus` VALUES (8, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, 0, 0, '个人信息', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (9, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 6, 0, 0, '示例文件', 'management', 0);
INSERT INTO `sys_base_menus` VALUES (10, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '9', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, 0, 0, 'excel导入导出', 'takeaway-box', 0);
INSERT INTO `sys_base_menus` VALUES (11, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '9', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, 0, 0, '媒体库（上传下载）', 'upload', 0);
INSERT INTO `sys_base_menus` VALUES (12, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '9', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, 0, 0, '断点续传', 'upload-filled', 0);
INSERT INTO `sys_base_menus` VALUES (13, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '9', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, 0, 0, '客户列表（资源示例）', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (14, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` VALUES (15, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '14', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `sys_base_menus` VALUES (16, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '14', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (17, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '14', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` VALUES (18, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, 0, 0, '字典管理', 'notebook', 0);
INSERT INTO `sys_base_menus` VALUES (19, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 0, 0, '字典详情', 'order', 0);
INSERT INTO `sys_base_menus` VALUES (20, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, 0, 0, '操作历史', 'pie-chart', 0);
INSERT INTO `sys_base_menus` VALUES (21, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '9', 'simpleUploader', 'simpleUploader', 0, 'view/example/simpleUploader/simpleUploader', 6, 0, 0, '断点续传（插件版）', 'upload', 0);
INSERT INTO `sys_base_menus` VALUES (22, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, 0, 0, '官方网站', 'home-filled', 0);
INSERT INTO `sys_base_menus` VALUES (23, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 6, 0, 0, '服务器状态', 'cloudy', 0);
INSERT INTO `sys_base_menus` VALUES (24, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '14', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, 0, 0, '自动化代码管理', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (25, '2022-03-10 17:52:53.163', '2022-03-10 17:52:53.163', NULL, 0, '14', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, 0, 0, '自动化代码（复用）', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (26, '2022-03-11 23:38:56.298', '2022-03-11 23:38:56.298', NULL, 0, '3', 'server', 'server', 0, 'view/superAdmin/server/server.vue', 0, 0, 0, '服务器管理', 'aim', 0);
INSERT INTO `sys_base_menus` VALUES (27, '2022-03-21 11:49:02.100', '2022-03-21 11:49:02.100', NULL, 0, '0', 'mahjongCardSettings', 'mahjongCardSettings', 0, 'view/mahjongCardSettings/mahjongCardSettings.vue', 0, 0, 0, '麻将配牌', 'coffee-cup', 0);

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES ('888', '888');
INSERT INTO `sys_data_authority_id` VALUES ('888', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('888', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '9528');

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2022-03-10 17:52:54.044', '2022-03-10 17:52:54.044', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2022-03-10 17:52:54.044', '2022-03-10 17:52:54.044', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2022-03-10 17:52:54.044', '2022-03-10 17:52:54.044', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2022-03-10 17:52:54.044', '2022-03-10 17:52:54.044', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2022-03-10 17:52:54.044', '2022-03-10 17:52:54.044', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2022-03-10 17:52:54.044', '2022-03-10 17:52:54.044', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) NULL DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, '男', 1, 1, 1, 1);
INSERT INTO `sys_dictionary_details` VALUES (2, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, '女', 2, 1, 2, 1);
INSERT INTO `sys_dictionary_details` VALUES (3, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (6, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (7, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'date', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (11, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (12, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (14, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (15, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (24, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '2022-03-10 17:52:54.112', '2022-03-10 17:52:54.112', NULL, 'tinyint', 0, 1, 0, 6);

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) NULL DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) NULL DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
INSERT INTO `sys_operation_records` VALUES (1, '2022-03-10 17:55:50.699', '2022-03-10 17:55:50.699', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 514800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signingKey\":\"2d9f58fb-903e-4f79-8f7f-5f1b4b4d2d14\",\"expiresTime\":604800,\"bufferTime\":86400,\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"format\":\"console\",\"prefix\":\"[github.com/flipped-aurora/gin-vue-admin/server]\",\"director\":\"log\",\"showLine\":true,\"encodeLevel\":\"LowercaseColorLevelEncoder\",\"stacktraceKey\":\"stacktrace\",\"logInConsole\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"isSSL\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"casbin\":{\"modelPath\":\"./resource/rbac_model.conf\"},\"system\":{\"env\":\"public\",\"addr\":18888,\"dbType\":\"mysql\",\"ossType\":\"local\",\"useMultipoint\":false,\"useRedis\":false,\"iplimitCount\":15000,\"iplimitTime\":3600},\"captcha\":{\"keyLong\":6,\"imgWidth\":240,\"imgHeight\":80},\"autoCode\":{\"transferRestart\":true,\"root\":\"H:\\\\gospace\\\\gin-vue-admin\",\"server\":\"/server\",\"serverApi\":\"/api/v1/autocode\",\"serverInitialize\":\"/initialize\",\"serverModel\":\"/model/autocode\",\"serverRequest\":\"/model/autocode/request/\",\"serverRouter\":\"/router/autocode\",\"serverService\":\"/service/autocode\",\"web\":\"/web/src\",\"webApi\":\"/api\",\"webForm\":\"/view\",\"webTable\":\"/view\"},\"mysql\":{\"path\":\"127.0.0.1\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"dbname\":\"gva\",\"username\":\"root\",\"password\":\"123456\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"error\",\"logZap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false},\"db-list\":[{\"disable\":false,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false}],\"local\":{\"path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"imgPath\":\"\",\"useHttps\":false,\"accessKey\":\"\",\"secretKey\":\"\",\"useCdnDomains\":false},\"aliyunOSS\":{\"endpoint\":\"yourEndpoint\",\"accessKeyId\":\"yourAccessKeyId\",\"accessKeySecret\":\"yourAccessKeySecret\",\"bucketName\":\"yourBucketName\",\"bucketUrl\":\"yourBucketUrl\",\"basePath\":\"yourBasePath\"},\"huaWeiObs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"accessKey\":\"you-access-key\",\"secretKey\":\"you-secret-key\"},\"tencentCOS\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"awsS3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3ForcePathStyle\":false,\"disableSSL\":false,\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (2, '2022-03-10 17:56:06.178', '2022-03-10 17:56:06.178', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 296320300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.18rc1\",\"numGoroutine\":15},\"cpu\":{\"cpus\":[14.285714285714285,0,7.142857142857142,0,0,14.285714285714285,14.285714285714285,0],\"cores\":4},\"ram\":{\"usedMb\":12616,\"totalMb\":32697,\"usedPercent\":38},\"disk\":{\"usedMb\":103398,\"usedGb\":100,\"totalMb\":636928,\"totalGb\":622,\"usedPercent\":16}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (3, '2022-03-10 18:32:08.458', '2022-03-10 18:32:08.458', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 1586600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":18}', '{\"code\":0,\"data\":{\"api\":{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"path\":\"/api/getApiById\",\"description\":\"获取api详细信息\",\"apiGroup\":\"api\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (4, '2022-03-10 18:32:19.514', '2022-03-10 18:32:19.514', NULL, '127.0.0.1', 'POST', '/api/updateApi', 200, 96604300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"path\":\"/api/getApiById\",\"description\":\"获取api详细信息哈哈\",\"apiGroup\":\"api\",\"method\":\"POST\"}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (5, '2022-03-10 18:38:30.512', '2022-03-10 18:38:30.512', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 0, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signingKey\":\"2d9f58fb-903e-4f79-8f7f-5f1b4b4d2d14\",\"expiresTime\":604800,\"bufferTime\":86400,\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"format\":\"console\",\"prefix\":\"[github.com/flipped-aurora/gin-vue-admin/server]\",\"director\":\"log\",\"showLine\":false,\"encodeLevel\":\"LowercaseColorLevelEncoder\",\"stacktraceKey\":\"stacktrace\",\"logInConsole\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"isSSL\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"casbin\":{\"modelPath\":\"./resource/rbac_model.conf\"},\"system\":{\"env\":\"public\",\"addr\":18888,\"dbType\":\"mysql\",\"ossType\":\"local\",\"useMultipoint\":false,\"useRedis\":false,\"iplimitCount\":15000,\"iplimitTime\":3600},\"captcha\":{\"keyLong\":6,\"imgWidth\":240,\"imgHeight\":80},\"autoCode\":{\"transferRestart\":true,\"root\":\"H:\\\\gospace\\\\gin-vue-admin\",\"server\":\"/server\",\"serverApi\":\"/api/v1/autocode\",\"serverInitialize\":\"/initialize\",\"serverModel\":\"/model/autocode\",\"serverRequest\":\"/model/autocode/request/\",\"serverRouter\":\"/router/autocode\",\"serverService\":\"/service/autocode\",\"web\":\"/web/src\",\"webApi\":\"/api\",\"webForm\":\"/view\",\"webTable\":\"/view\"},\"mysql\":{\"path\":\"127.0.0.1\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"dbname\":\"gva\",\"username\":\"root\",\"password\":\"123456\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"error\",\"logZap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false},\"db-list\":[{\"disable\":false,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false}],\"local\":{\"path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"imgPath\":\"\",\"useHttps\":false,\"accessKey\":\"\",\"secretKey\":\"\",\"useCdnDomains\":false},\"aliyunOSS\":{\"endpoint\":\"yourEndpoint\",\"accessKeyId\":\"yourAccessKeyId\",\"accessKeySecret\":\"yourAccessKeySecret\",\"bucketName\":\"yourBucketName\",\"bucketUrl\":\"yourBucketUrl\",\"basePath\":\"yourBasePath\"},\"huaWeiObs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"accessKey\":\"you-access-key\",\"secretKey\":\"you-secret-key\"},\"tencentCOS\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"awsS3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3ForcePathStyle\":false,\"disableSSL\":false,\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (6, '2022-03-10 18:38:40.282', '2022-03-10 18:38:40.282', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 1583700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":18}', '{\"code\":0,\"data\":{\"api\":{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T18:32:19.422+08:00\",\"path\":\"/api/getApiById\",\"description\":\"获取api详细信息哈哈\",\"apiGroup\":\"api\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (7, '2022-03-10 18:38:44.338', '2022-03-10 18:38:44.338', NULL, '127.0.0.1', 'POST', '/api/updateApi', 200, 15304100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T18:32:19.422+08:00\",\"path\":\"/api/getApiById\",\"description\":\"获取api详细信息\",\"apiGroup\":\"api\",\"method\":\"POST\"}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (8, '2022-03-10 18:38:47.977', '2022-03-10 18:38:47.977', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 1570900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":18}', '{\"code\":0,\"data\":{\"api\":{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T18:38:44.327+08:00\",\"path\":\"/api/getApiById\",\"description\":\"获取api详细信息\",\"apiGroup\":\"api\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (9, '2022-03-10 18:38:51.658', '2022-03-10 18:38:51.658', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 1570100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":18}', '{\"code\":0,\"data\":{\"api\":{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T18:38:44.327+08:00\",\"path\":\"/api/getApiById\",\"description\":\"获取api详细信息\",\"apiGroup\":\"api\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (10, '2022-03-10 18:40:45.199', '2022-03-10 18:40:45.199', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 286278200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.18rc1\",\"numGoroutine\":14},\"cpu\":{\"cpus\":[7.6923076923076925,7.6923076923076925,7.6923076923076925,38.46153846153847,0,7.6923076923076925,0,7.6923076923076925],\"cores\":4},\"ram\":{\"usedMb\":13805,\"totalMb\":32697,\"usedPercent\":42},\"disk\":{\"usedMb\":103412,\"usedGb\":100,\"totalMb\":636928,\"totalGb\":622,\"usedPercent\":16}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (11, '2022-03-10 18:40:57.025', '2022-03-10 18:40:57.025', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 259502800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.18rc1\",\"numGoroutine\":14},\"cpu\":{\"cpus\":[14.285714285714285,0,7.142857142857142,14.285714285714285,14.285714285714285,7.142857142857142,21.428571428571427,14.285714285714285],\"cores\":4},\"ram\":{\"usedMb\":13809,\"totalMb\":32697,\"usedPercent\":42},\"disk\":{\"usedMb\":103412,\"usedGb\":100,\"totalMb\":636928,\"totalGb\":622,\"usedPercent\":16}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (12, '2022-03-10 18:41:00.178', '2022-03-10 18:41:00.178', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 258732500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.18rc1\",\"numGoroutine\":14},\"cpu\":{\"cpus\":[13.333333333333334,0,7.142857142857142,0,14.285714285714285,0,7.142857142857142,0],\"cores\":4},\"ram\":{\"usedMb\":13795,\"totalMb\":32697,\"usedPercent\":42},\"disk\":{\"usedMb\":103412,\"usedGb\":100,\"totalMb\":636928,\"totalGb\":622,\"usedPercent\":16}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (13, '2022-03-10 18:43:35.388', '2022-03-10 18:43:35.388', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 260204400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.18rc1\",\"numGoroutine\":13},\"cpu\":{\"cpus\":[28.57142857142857,0,7.6923076923076925,7.6923076923076925,15.384615384615385,0,0,7.6923076923076925],\"cores\":4},\"ram\":{\"usedMb\":13652,\"totalMb\":32697,\"usedPercent\":41},\"disk\":{\"usedMb\":103412,\"usedGb\":100,\"totalMb\":636928,\"totalGb\":622,\"usedPercent\":16}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (14, '2022-03-11 23:38:56.308', '2022-03-11 23:38:56.308', NULL, '127.0.0.1', 'POST', '/menu/addBaseMenu', 200, 14932500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":0,\"path\":\"server\",\"name\":\"server\",\"hidden\":\"\",\"parentId\":\"3\",\"component\":\"view/superAdmin/server/server.vue\",\"meta\":{\"title\":\"服务器管理\",\"icon\":\"aim\",\"defaultMenu\":false,\"keepAlive\":\"\"}}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (15, '2022-03-11 23:41:34.213', '2022-03-11 23:41:34.213', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 31671100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"menus\":[{\"ID\":22,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"https://www.gin-vue-admin.com\",\"name\":\"https://www.gin-vue-admin.com\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":1,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"仪表盘\",\"icon\":\"odometer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"systemTools\",\"name\":\"systemTools\",\"hidden\":false,\"component\":\"view/systemTools/index.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统工具\",\"icon\":\"tools\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":25,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码（复用）\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码（复用）\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"example\",\"name\":\"example\",\"hidden\":false,\"component\":\"view/example/index.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"示例文件\",\"icon\":\"management\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":10,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":21,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"simpleUploader\",\"name\":\"simpleUploader\",\"hidden\":false,\"component\":\"view/example/simpleUploader/simpleUploader\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传（插件版）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":23,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"state\",\"name\":\"state\",\"hidden\":false,\"component\":\"view/system/state.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器状态\",\"icon\":\"cloudy\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":2,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":\"888\"}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (16, '2022-03-11 23:58:11.055', '2022-03-11 23:58:11.055', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 279651100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":4,\"compiler\":\"gc\",\"goVersion\":\"go1.18beta2\",\"numGoroutine\":28},\"cpu\":{\"cpus\":[50,46.15384615384615,30.76923076923077,46.15384615384615],\"cores\":4},\"ram\":{\"usedMb\":16492,\"totalMb\":24534,\"usedPercent\":67},\"disk\":{\"usedMb\":53127,\"usedGb\":51,\"totalMb\":199999,\"totalGb\":195,\"usedPercent\":26}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (17, '2022-03-12 00:06:04.993', '2022-03-12 00:06:04.993', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 1568200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signingKey\":\"2d9f58fb-903e-4f79-8f7f-5f1b4b4d2d14\",\"expiresTime\":604800,\"bufferTime\":86400,\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"format\":\"console\",\"prefix\":\"[github.com/jqiris/orange-admin/server]\",\"director\":\"log\",\"showLine\":false,\"encodeLevel\":\"LowercaseColorLevelEncoder\",\"stacktraceKey\":\"stacktrace\",\"logInConsole\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"isSSL\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"casbin\":{\"modelPath\":\"./resource/rbac_model.conf\"},\"system\":{\"env\":\"public\",\"addr\":18888,\"dbType\":\"mysql\",\"ossType\":\"local\",\"useMultipoint\":false,\"useRedis\":false,\"iplimitCount\":15000,\"iplimitTime\":3600},\"captcha\":{\"keyLong\":6,\"imgWidth\":240,\"imgHeight\":80},\"autoCode\":{\"transferRestart\":true,\"root\":\"D:\\\\webspace\\\\orange-admin\",\"server\":\"/server\",\"serverApi\":\"/api/v1/autocode\",\"serverInitialize\":\"/initialize\",\"serverModel\":\"/model/autocode\",\"serverRequest\":\"/model/autocode/request/\",\"serverRouter\":\"/router/autocode\",\"serverService\":\"/service/autocode\",\"web\":\"/web/src\",\"webApi\":\"/api\",\"webForm\":\"/view\",\"webTable\":\"/view\"},\"mysql\":{\"path\":\"127.0.0.1\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"dbname\":\"orange\",\"username\":\"root\",\"password\":\"123456\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"error\",\"logZap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false},\"db-list\":[{\"disable\":false,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false}],\"local\":{\"path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"imgPath\":\"\",\"useHttps\":false,\"accessKey\":\"\",\"secretKey\":\"\",\"useCdnDomains\":false},\"aliyunOSS\":{\"endpoint\":\"yourEndpoint\",\"accessKeyId\":\"yourAccessKeyId\",\"accessKeySecret\":\"yourAccessKeySecret\",\"bucketName\":\"yourBucketName\",\"bucketUrl\":\"yourBucketUrl\",\"basePath\":\"yourBasePath\"},\"huaWeiObs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"accessKey\":\"you-access-key\",\"secretKey\":\"you-secret-key\"},\"tencentCOS\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/jqiris/orange-admin/server\"},\"awsS3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3ForcePathStyle\":false,\"disableSSL\":false,\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/jqiris/orange-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]},\"discover\":{\"use_type\":\"etcd\",\"dial_timeout\":5,\"endpoints\":[\"127.0.0.1:23791\",\"127.0.0.1:23792\",\"127.0.0.1:23793\"],\"prefix\":\"\"}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (18, '2022-03-12 00:06:21.442', '2022-03-12 00:06:21.442', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 3408500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":13}', '{\"code\":0,\"data\":{\"api\":{\"ID\":13,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"path\":\"/api/createApi\",\"description\":\"创建api\",\"apiGroup\":\"api\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (19, '2022-03-12 00:06:46.998', '2022-03-12 00:06:46.998', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 3248800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":29}', '{\"code\":0,\"data\":{\"api\":{\"ID\":29,\"CreatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:52.155+08:00\",\"path\":\"/menu/getMenu\",\"description\":\"获取菜单树(必选)\",\"apiGroup\":\"菜单\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (20, '2022-03-12 00:07:41.297', '2022-03-12 00:07:41.297', NULL, '127.0.0.1', 'POST', '/api/createApi', 200, 17384200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"path\":\"/server/getServerList\",\"apiGroup\":\"服务器\",\"method\":\"POST\",\"description\":\"获取服务器列表\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (21, '2022-03-12 00:09:05.527', '2022-03-12 00:09:05.527', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 0, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signingKey\":\"2d9f58fb-903e-4f79-8f7f-5f1b4b4d2d14\",\"expiresTime\":604800,\"bufferTime\":86400,\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"format\":\"console\",\"prefix\":\"[github.com/jqiris/orange-admin/server]\",\"director\":\"log\",\"showLine\":false,\"encodeLevel\":\"LowercaseColorLevelEncoder\",\"stacktraceKey\":\"stacktrace\",\"logInConsole\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"isSSL\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"casbin\":{\"modelPath\":\"./resource/rbac_model.conf\"},\"system\":{\"env\":\"public\",\"addr\":18888,\"dbType\":\"mysql\",\"ossType\":\"local\",\"useMultipoint\":false,\"useRedis\":false,\"iplimitCount\":15000,\"iplimitTime\":3600},\"captcha\":{\"keyLong\":6,\"imgWidth\":240,\"imgHeight\":80},\"autoCode\":{\"transferRestart\":true,\"root\":\"D:\\\\webspace\\\\orange-admin\",\"server\":\"/server\",\"serverApi\":\"/api/v1/autocode\",\"serverInitialize\":\"/initialize\",\"serverModel\":\"/model/autocode\",\"serverRequest\":\"/model/autocode/request/\",\"serverRouter\":\"/router/autocode\",\"serverService\":\"/service/autocode\",\"web\":\"/web/src\",\"webApi\":\"/api\",\"webForm\":\"/view\",\"webTable\":\"/view\"},\"mysql\":{\"path\":\"127.0.0.1\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"dbname\":\"orange\",\"username\":\"root\",\"password\":\"123456\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"error\",\"logZap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false},\"db-list\":[{\"disable\":false,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"dbname\":\"\",\"username\":\"\",\"password\":\"\",\"maxIdleConns\":10,\"maxOpenConns\":100,\"logMode\":\"\",\"logZap\":false}],\"local\":{\"path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"imgPath\":\"\",\"useHttps\":false,\"accessKey\":\"\",\"secretKey\":\"\",\"useCdnDomains\":false},\"aliyunOSS\":{\"endpoint\":\"yourEndpoint\",\"accessKeyId\":\"yourAccessKeyId\",\"accessKeySecret\":\"yourAccessKeySecret\",\"bucketName\":\"yourBucketName\",\"bucketUrl\":\"yourBucketUrl\",\"basePath\":\"yourBasePath\"},\"huaWeiObs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"accessKey\":\"you-access-key\",\"secretKey\":\"you-secret-key\"},\"tencentCOS\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/jqiris/orange-admin/server\"},\"awsS3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3ForcePathStyle\":false,\"disableSSL\":false,\"secretID\":\"xxxxxxxx\",\"secretKey\":\"xxxxxxxx\",\"baseURL\":\"https://gin.vue.admin\",\"pathPrefix\":\"github.com/jqiris/orange-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]},\"discover\":{\"use_type\":\"etcd\",\"dial_timeout\":5,\"endpoints\":[\"127.0.0.1:23791\",\"127.0.0.1:23792\",\"127.0.0.1:23793\"],\"prefix\":\"\"}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (22, '2022-03-12 00:09:53.209', '2022-03-12 00:09:53.209', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 30952200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"menus\":[{\"ID\":1,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"仪表盘\",\"icon\":\"odometer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":2,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":\"8881\"}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (23, '2022-03-12 00:09:58.829', '2022-03-12 00:09:58.829', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 28337400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"menus\":[{\"ID\":1,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"仪表盘\",\"icon\":\"odometer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"systemTools\",\"name\":\"systemTools\",\"hidden\":false,\"component\":\"view/systemTools/index.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统工具\",\"icon\":\"tools\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":25,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码（复用）\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"example\",\"name\":\"example\",\"hidden\":false,\"component\":\"view/example/index.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"示例文件\",\"icon\":\"management\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":10,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":21,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"simpleUploader\",\"name\":\"simpleUploader\",\"hidden\":false,\"component\":\"view/example/simpleUploader/simpleUploader\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传（插件版）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":2,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":\"9528\"}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (24, '2022-03-21 11:49:02.112', '2022-03-21 11:49:02.112', NULL, '127.0.0.1', 'POST', '/menu/addBaseMenu', 200, 18869600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":0,\"path\":\"mahjongCardSettings\",\"name\":\"mahjongCardSettings\",\"hidden\":\"\",\"parentId\":\"0\",\"component\":\"view/mahjongCardSettings/mahjongCardSettings.vue\",\"meta\":{\"title\":\"麻将配牌\",\"icon\":\"coffee-cup\",\"defaultMenu\":false,\"keepAlive\":\"\"}}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (25, '2022-03-21 11:49:23.072', '2022-03-21 11:49:23.072', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 161075700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"menus\":[{\"ID\":22,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"https://www.gin-vue-admin.com\",\"name\":\"https://www.gin-vue-admin.com\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":27,\"CreatedAt\":\"2022-03-21T11:49:02.1+08:00\",\"UpdatedAt\":\"2022-03-21T11:49:02.1+08:00\",\"parentId\":\"0\",\"path\":\"mahjongCardSettings\",\"name\":\"mahjongCardSettings\",\"hidden\":false,\"component\":\"view/mahjongCardSettings/mahjongCardSettings.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"麻将配牌\",\"icon\":\"coffee-cup\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":1,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"仪表盘\",\"icon\":\"odometer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"UpdatedAt\":\"2022-03-11T23:38:56.298+08:00\",\"parentId\":\"3\",\"path\":\"server\",\"name\":\"server\",\"hidden\":false,\"component\":\"view/superAdmin/server/server.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"systemTools\",\"name\":\"systemTools\",\"hidden\":false,\"component\":\"view/systemTools/index.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统工具\",\"icon\":\"tools\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":25,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码（复用）\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码（复用）\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"example\",\"name\":\"example\",\"hidden\":false,\"component\":\"view/example/index.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"示例文件\",\"icon\":\"management\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":10,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":21,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"simpleUploader\",\"name\":\"simpleUploader\",\"hidden\":false,\"component\":\"view/example/simpleUploader/simpleUploader\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传（插件版）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":23,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"state\",\"name\":\"state\",\"hidden\":false,\"component\":\"view/system/state.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器状态\",\"icon\":\"cloudy\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":2,\"CreatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"UpdatedAt\":\"2022-03-10T17:52:53.163+08:00\",\"parentId\":\"0\",\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":\"888\"}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (26, '2022-03-21 11:55:06.228', '2022-03-21 11:55:06.228', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 3249700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":88}', '{\"code\":0,\"data\":{\"api\":{\"ID\":88,\"CreatedAt\":\"2022-03-12T00:07:41.287+08:00\",\"UpdatedAt\":\"2022-03-12T00:07:41.287+08:00\",\"path\":\"/server/getServerList\",\"description\":\"获取服务器列表\",\"apiGroup\":\"服务器\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (27, '2022-03-21 11:56:27.873', '2022-03-21 11:56:27.873', NULL, '127.0.0.1', 'POST', '/api/getApiById', 200, 4493800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"id\":88}', '{\"code\":0,\"data\":{\"api\":{\"ID\":88,\"CreatedAt\":\"2022-03-12T00:07:41.287+08:00\",\"UpdatedAt\":\"2022-03-12T00:07:41.287+08:00\",\"path\":\"/server/getServerList\",\"description\":\"获取服务器列表\",\"apiGroup\":\"服务器\",\"method\":\"POST\"}},\"msg\":\"操作成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (28, '2022-03-21 11:57:06.574', '2022-03-21 11:57:06.574', NULL, '127.0.0.1', 'POST', '/api/createApi', 200, 5268900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"path\":\"/mahjongCardSettings/getMahjongCardSettingsList\",\"apiGroup\":\"麻将\",\"method\":\"GET\",\"description\":\"获取麻将配牌列表\"}', '{\"code\":7,\"data\":{},\"msg\":\"创建失败\"}', 1);
INSERT INTO `sys_operation_records` VALUES (29, '2022-03-21 11:58:28.253', '2022-03-21 11:58:28.253', NULL, '127.0.0.1', 'POST', '/casbin/updateCasbin', 200, 228380400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"authorityId\":\"888\",\"casbinInfos\":[{\"path\":\"/base/login\",\"method\":\"POST\"},{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/deleteUser\",\"method\":\"DELETE\"},{\"path\":\"/user/register\",\"method\":\"POST\"},{\"path\":\"/user/getUserList\",\"method\":\"POST\"},{\"path\":\"/user/setUserInfo\",\"method\":\"PUT\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/setUserAuthorities\",\"method\":\"POST\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserAuthority\",\"method\":\"POST\"},{\"path\":\"/user/resetPassword\",\"method\":\"POST\"},{\"path\":\"/api/createApi\",\"method\":\"POST\"},{\"path\":\"/api/deleteApi\",\"method\":\"POST\"},{\"path\":\"/api/updateApi\",\"method\":\"POST\"},{\"path\":\"/api/getApiList\",\"method\":\"POST\"},{\"path\":\"/api/getAllApis\",\"method\":\"POST\"},{\"path\":\"/api/getApiById\",\"method\":\"POST\"},{\"path\":\"/api/deleteApisByIds\",\"method\":\"DELETE\"},{\"path\":\"/authority/copyAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/createAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/deleteAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/updateAuthority\",\"method\":\"PUT\"},{\"path\":\"/authority/getAuthorityList\",\"method\":\"POST\"},{\"path\":\"/authority/setDataAuthority\",\"method\":\"POST\"},{\"path\":\"/casbin/updateCasbin\",\"method\":\"POST\"},{\"path\":\"/casbin/getPolicyPathByAuthorityId\",\"method\":\"POST\"},{\"path\":\"/menu/addBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/menu/deleteBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/updateBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuById\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuList\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuTree\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/menu/addMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/breakpointContinue\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/breakpointContinueFinish\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/removeChunk\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/upload\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/deleteFile\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/getFileList\",\"method\":\"POST\"},{\"path\":\"/system/getServerInfo\",\"method\":\"POST\"},{\"path\":\"/system/getSystemConfig\",\"method\":\"POST\"},{\"path\":\"/system/setSystemConfig\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"PUT\"},{\"path\":\"/customer/customer\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"DELETE\"},{\"path\":\"/customer/customer\",\"method\":\"GET\"},{\"path\":\"/customer/customerList\",\"method\":\"GET\"},{\"path\":\"/autoCode/getDB\",\"method\":\"GET\"},{\"path\":\"/autoCode/getTables\",\"method\":\"GET\"},{\"path\":\"/autoCode/createTemp\",\"method\":\"POST\"},{\"path\":\"/autoCode/preview\",\"method\":\"POST\"},{\"path\":\"/autoCode/getColumn\",\"method\":\"GET\"},{\"path\":\"/autoCode/getMeta\",\"method\":\"POST\"},{\"path\":\"/autoCode/rollback\",\"method\":\"POST\"},{\"path\":\"/autoCode/getSysHistory\",\"method\":\"POST\"},{\"path\":\"/autoCode/delSysHistory\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/updateSysDictionaryDetail\",\"method\":\"PUT\"},{\"path\":\"/sysDictionaryDetail/createSysDictionaryDetail\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/deleteSysDictionaryDetail\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionaryDetail/findSysDictionaryDetail\",\"method\":\"GET\"},{\"path\":\"/sysDictionaryDetail/getSysDictionaryDetailList\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/createSysDictionary\",\"method\":\"POST\"},{\"path\":\"/sysDictionary/deleteSysDictionary\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionary/updateSysDictionary\",\"method\":\"PUT\"},{\"path\":\"/sysDictionary/findSysDictionary\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/getSysDictionaryList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/createSysOperationRecord\",\"method\":\"POST\"},{\"path\":\"/sysOperationRecord/findSysOperationRecord\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/getSysOperationRecordList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecord\",\"method\":\"DELETE\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecordByIds\",\"method\":\"DELETE\"},{\"path\":\"/simpleUploader/upload\",\"method\":\"POST\"},{\"path\":\"/simpleUploader/checkFileMd5\",\"method\":\"GET\"},{\"path\":\"/simpleUploader/mergeFileMd5\",\"method\":\"GET\"},{\"path\":\"/email/emailTest\",\"method\":\"POST\"},{\"path\":\"/excel/importExcel\",\"method\":\"POST\"},{\"path\":\"/excel/loadExcel\",\"method\":\"GET\"},{\"path\":\"/excel/exportExcel\",\"method\":\"POST\"},{\"path\":\"/excel/downloadTemplate\",\"method\":\"GET\"},{\"path\":\"/authorityBtn/setAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/getAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/canRemoveAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/server/getServerList\",\"method\":\"POST\"},{\"path\":\"/mahjongCardSettings/createMahjongCardSettings\",\"method\":\"POST\"},{\"path\":\"/mahjongCardSettings/deleteMahjongCardSettings\",\"method\":\"DELETE\"},{\"path\":\"/mahjongCardSettings/deleteMahjongCardSettingsByIds\",\"method\":\"DELETE\"},{\"path\":\"/mahjongCardSettings/updateMahjongCardSettings\",\"method\":\"PUT\"},{\"path\":\"/mahjongCardSettings/findMahjongCardSettings\",\"method\":\"GET\"},{\"path\":\"/mahjongCardSettings/getMahjongCardSettingsList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (30, '2022-03-22 20:23:27.009', '2022-03-22 20:23:27.009', NULL, '127.0.0.1', 'POST', '/mahjongCardSettings/createMahjongCardSettings', 200, 28295800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"seatOne\":\"2-5-11-15-27-40-48-58-67-71-80-92-96-107\",\"seatTwo\":\"0-29-33-38-49-51-52-60-64-69-70-72-89\",\"seatThree\":\"9-10-20-30-44-46-55-63-68-84-86-91-97\",\"seatFour\":\"9-10-13-56-66-75-76-78-82-88-93-100-103\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (31, '2022-03-22 20:25:37.091', '2022-03-22 20:25:37.091', NULL, '127.0.0.1', 'DELETE', '/mahjongCardSettings/deleteMahjongCardSettings', 200, 13095700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":1}', '{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (32, '2022-03-22 20:26:23.381', '2022-03-22 20:26:23.381', NULL, '127.0.0.1', 'POST', '/mahjongCardSettings/createMahjongCardSettings', 200, 11102900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"seatOne\":\"11-14-15-21-25-31-42-63-69-81-87-93-103-104\",\"seatTwo\":\"9-12-30-37-40-54-55-85-86-94-95-106-107\",\"seatThree\":\"10-27-34-35-43-51-64-71-73-75-77-83-92\",\"seatFour\":\"0-13-22-26-28-36-45-48-53-62-65-79-101\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (33, '2022-03-22 21:06:24.747', '2022-03-22 21:06:24.747', NULL, '127.0.0.1', 'PUT', '/mahjongCardSettings/updateMahjongCardSettings', 200, 23873200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":2,\"CreatedAt\":\"2022-03-22T20:26:23.37+08:00\",\"UpdatedAt\":\"2022-03-22T20:26:23.37+08:00\",\"seatOne\":\"11-14-15-21-25-31-42-63-69-81-87-93-102-103\",\"seatTwo\":\"9-12-30-37-40-54-55-85-86-94-95-106-107\",\"seatThree\":\"10-27-34-35-43-51-64-71-73-75-77-83-92\",\"seatFour\":\"0-13-22-26-28-36-45-48-53-62-65-79-101\"}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (34, '2022-03-22 21:07:01.226', '2022-03-22 21:07:01.226', NULL, '127.0.0.1', 'PUT', '/mahjongCardSettings/updateMahjongCardSettings', 200, 11819800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":2,\"CreatedAt\":\"2022-03-22T20:26:23.37+08:00\",\"UpdatedAt\":\"2022-03-22T21:06:24.723+08:00\",\"seatOne\":\"11-14-15-21-25-31-42-63-69-81-87-93-98-99\",\"seatTwo\":\"9-12-30-37-40-54-55-85-86-94-95-106-107\",\"seatThree\":\"10-27-34-35-43-51-64-71-73-75-77-83-92\",\"seatFour\":\"0-13-22-26-28-36-45-48-53-62-65-79-101\"}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (35, '2022-03-22 21:10:27.809', '2022-03-22 21:10:27.809', NULL, '127.0.0.1', 'PUT', '/mahjongCardSettings/updateMahjongCardSettings', 200, 15700500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":2,\"CreatedAt\":\"2022-03-22T20:26:23.37+08:00\",\"UpdatedAt\":\"2022-03-22T21:07:01.215+08:00\",\"seatOne\":\"11-14-15-21-25-31-42-63-69-81-87-93-100-104\",\"seatTwo\":\"9-12-30-37-40-54-55-85-86-94-95-106-107\",\"seatThree\":\"10-27-34-35-43-51-64-71-73-75-77-83-92\",\"seatFour\":\"0-13-22-26-28-36-45-48-53-62-65-79-101\"}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (36, '2022-03-22 21:10:49.135', '2022-03-22 21:10:49.135', NULL, '127.0.0.1', 'PUT', '/mahjongCardSettings/updateMahjongCardSettings', 200, 14450100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":2,\"CreatedAt\":\"2022-03-22T20:26:23.37+08:00\",\"UpdatedAt\":\"2022-03-22T21:10:27.793+08:00\",\"seatOne\":\"11-14-15-21-25-31-42-63-69-81-87-93-100-104\",\"seatTwo\":\"9-12-30-37-40-54-55-85-86-94-95-99-102\",\"seatThree\":\"10-27-34-35-43-51-64-71-73-75-77-83-92\",\"seatFour\":\"0-13-22-26-28-36-45-48-53-62-65-79-101\"}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (37, '2022-03-22 21:11:11.099', '2022-03-22 21:11:11.099', NULL, '127.0.0.1', 'PUT', '/mahjongCardSettings/updateMahjongCardSettings', 200, 19416400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"ID\":2,\"CreatedAt\":\"2022-03-22T20:26:23.37+08:00\",\"UpdatedAt\":\"2022-03-22T21:10:49.121+08:00\",\"seatOne\":\"11-14-15-21-25-31-42-63-69-81-87-93-100-104\",\"seatTwo\":\"9-12-30-37-40-54-55-85-86-94-95-98-99\",\"seatThree\":\"10-27-34-35-43-51-64-71-73-75-77-83-92\",\"seatFour\":\"0-13-22-26-28-36-45-48-53-62-65-79-101\"}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (38, '2022-03-22 21:17:21.117', '2022-03-22 21:17:21.117', NULL, '127.0.0.1', 'POST', '/mahjongCardSettings/createMahjongCardSettings', 200, 11295300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"seatOne\":\"22-24-44-51-57-58-59-74-77-80-90-91-94-95\",\"seatTwo\":\"1-7-19-27-31-41-43-50-72-87-93-100-104\",\"seatThree\":\"3-5-34-42-49-53-55-68-76-86-102-105-107\",\"seatFour\":\"6-13-32-35-37-38-39-45-54-60-64-79-103\",\"seatThere\":\"\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (39, '2022-03-22 21:19:59.342', '2022-03-22 21:19:59.342', NULL, '127.0.0.1', 'POST', '/mahjongCardSettings/createMahjongCardSettings', 200, 12205400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36', '', '{\"seatOne\":\"7-9-11-15-16-17-27-31-34-38-41-43-71-81\",\"seatTwo\":\"2-20-24-29-50-52-54-56-59-61-62-78-85\",\"seatThree\":\"18-19-33-36-46-51-60-64-69-72-84-93-98\",\"seatFour\":\"1-3-6-8-21-63-80-89-90-91-92-100-101\",\"seatThere\":\"\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, '888');
INSERT INTO `sys_user_authority` VALUES (1, '8881');
INSERT INTO `sys_user_authority` VALUES (1, '9528');
INSERT INTO `sys_user_authority` VALUES (2, '888');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2022-03-10 17:52:52.520', '2022-03-10 17:52:52.520', NULL, 'c227cd69-6724-4a6c-85b4-9876086899e9', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '888', '17611111111', '333333333@qq.com');
INSERT INTO `sys_users` VALUES (2, '2022-03-10 17:52:52.520', '2022-03-10 17:52:52.520', NULL, '229d7ead-68c8-4b8c-8423-4309a50f97ff', 'a303176530', '3ec063004a6f31642261936a379fde3d', 'QMPlusUser', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', '9528', '17611111111', '333333333@qq.com');

-- ----------------------------
-- Table structure for user_members
-- ----------------------------
DROP TABLE IF EXISTS `user_members`;
CREATE TABLE `user_members`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `headimg` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `sex` tinyint(4) NULL DEFAULT 0 COMMENT '性别 0-未知 1-男 2-女',
  `open_id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆open_id',
  `reg_time` datetime NULL DEFAULT NULL COMMENT '注册时间',
  `login_time` datetime NULL DEFAULT NULL COMMENT '登陆时间',
  `lv` smallint(6) NULL DEFAULT 1 COMMENT '用户等级',
  `exp` int(11) NULL DEFAULT 0 COMMENT '用户经验',
  `coins` int(11) NULL DEFAULT 0 COMMENT '用户金币',
  `gems` int(11) NULL DEFAULT 0 COMMENT '用户宝石',
  `room_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '最后进入房间类型 0-麻将 1-扑克',
  `room_id` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '最后进入房间号',
  `history` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作记录',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1013 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_members
-- ----------------------------
INSERT INTO `user_members` VALUES (1005, 'guest_jason1', '56uv5pyo6ZuA5Zyj', 'null', 0, '', '2022-03-09 10:36:10', '2022-03-20 23:45:55', 1, 0, 1000, 21, 0, '884310', '[{\"id\":\"796396\",\"seats\":[{\"name\":\"56uv5pyo6ZuA5Zyj\",\"score\":-14,\"userid\":1005},{\"name\":\"56uv5pyo6ZuA5Zyj\",\"score\":1,\"userid\":1006},{\"name\":\"54us5a2k6ZuA5Zyj\",\"score\":11,\"userid\":1007},{\"name\":\"6ZW/5a2Z5LiN6L6T\",\"score\":2,\"userid\":1008}],\"time\":1646799176,\"uuid\":\"1646799176796396\"}]');
INSERT INTO `user_members` VALUES (1007, 'guest_jason3', '54us5a2k6ZuA5Zyj', 'null', 0, '', '2022-03-09 10:36:40', '2022-03-23 08:47:36', 1, 0, 1000, 19, 0, '644426', '[{\"id\":\"796396\",\"seats\":[{\"name\":\"56uv5pyo6ZuA5Zyj\",\"score\":-14,\"userid\":1005},{\"name\":\"56uv5pyo6ZuA5Zyj\",\"score\":1,\"userid\":1006},{\"name\":\"54us5a2k6ZuA5Zyj\",\"score\":11,\"userid\":1007},{\"name\":\"6ZW/5a2Z5LiN6L6T\",\"score\":2,\"userid\":1008}],\"time\":1646799176,\"uuid\":\"1646799176796396\"},{\"id\":\"281277\",\"seats\":[{\"name\":\"6ZW/5a2Z6Ieq5pG4\",\"score\":0,\"userid\":1012},{\"name\":\"5Lic6Ziz56iz6LWi\",\"score\":0,\"userid\":1011},{\"name\":\"54us5a2k6ZuA5Zyj\",\"score\":0,\"userid\":1007},{\"name\":\"6ZW/5a2Z5LiN6L6T\",\"score\":0,\"userid\":1008}],\"time\":1647339653,\"uuid\":\"1647339653281277\"}]');
INSERT INTO `user_members` VALUES (1008, 'guest_jason4', '6ZW/5a2Z5LiN6L6T', 'null', 0, '', '2022-03-09 10:36:43', '2022-03-23 08:47:33', 1, 0, 1000, 21, 0, '644426', '[{\"id\":\"796396\",\"seats\":[{\"name\":\"56uv5pyo6ZuA5Zyj\",\"score\":-14,\"userid\":1005},{\"name\":\"56uv5pyo6ZuA5Zyj\",\"score\":1,\"userid\":1006},{\"name\":\"54us5a2k6ZuA5Zyj\",\"score\":11,\"userid\":1007},{\"name\":\"6ZW/5a2Z5LiN6L6T\",\"score\":2,\"userid\":1008}],\"time\":1646799176,\"uuid\":\"1646799176796396\"},{\"id\":\"281277\",\"seats\":[{\"name\":\"6ZW/5a2Z6Ieq5pG4\",\"score\":0,\"userid\":1012},{\"name\":\"5Lic6Ziz56iz6LWi\",\"score\":0,\"userid\":1011},{\"name\":\"54us5a2k6ZuA5Zyj\",\"score\":0,\"userid\":1007},{\"name\":\"6ZW/5a2Z5LiN6L6T\",\"score\":0,\"userid\":1008}],\"time\":1647339653,\"uuid\":\"1647339653281277\"}]');
INSERT INTO `user_members` VALUES (1009, 'guest_jason5', '5a6H5paH56iz6LWi', 'null', 0, '', '2022-03-09 14:30:35', '2022-03-19 20:27:50', 1, 0, 1000, 21, 0, '', '');
INSERT INTO `user_members` VALUES (1010, 'guest_jason6', '5Lic5pa56ZuA5Zyj', 'null', 0, '', '2022-03-09 14:32:32', '2022-03-19 20:28:00', 1, 0, 1000, 21, 0, '', '');
INSERT INTO `user_members` VALUES (1011, 'guest_jason2', '5Lic6Ziz56iz6LWi', 'null', 0, '', '2022-03-10 14:31:39', '2022-03-23 08:47:28', 1, 0, 1000, 19, 0, '644426', '[{\"id\":\"281277\",\"seats\":[{\"name\":\"6ZW/5a2Z6Ieq5pG4\",\"score\":0,\"userid\":1012},{\"name\":\"5Lic6Ziz56iz6LWi\",\"score\":0,\"userid\":1011},{\"name\":\"54us5a2k6ZuA5Zyj\",\"score\":0,\"userid\":1007},{\"name\":\"6ZW/5a2Z5LiN6L6T\",\"score\":0,\"userid\":1008}],\"time\":1647339653,\"uuid\":\"1647339653281277\"}]');
INSERT INTO `user_members` VALUES (1012, 'guest_1645804646235', '6ZW/5a2Z6Ieq5pG4', 'null', 0, '', '2022-03-12 00:54:57', '2022-03-23 08:47:26', 1, 0, 1000, 19, 0, '644426', '[{\"id\":\"281277\",\"seats\":[{\"name\":\"6ZW/5a2Z6Ieq5pG4\",\"score\":0,\"userid\":1012},{\"name\":\"5Lic6Ziz56iz6LWi\",\"score\":0,\"userid\":1011},{\"name\":\"54us5a2k6ZuA5Zyj\",\"score\":0,\"userid\":1007},{\"name\":\"6ZW/5a2Z5LiN6L6T\",\"score\":0,\"userid\":1008}],\"time\":1647339653,\"uuid\":\"1647339653281277\"}]');

-- ----------------------------
-- Table structure for user_weixins
-- ----------------------------
DROP TABLE IF EXISTS `user_weixins`;
CREATE TABLE `user_weixins`  (
  `open_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `nickname` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `headimgurl` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户头像',
  `sex` tinyint(4) NOT NULL DEFAULT 0 COMMENT '性别',
  `country` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '国家',
  `province` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省份',
  `city` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '城市',
  `unionid` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联合id',
  PRIMARY KEY (`open_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_weixins
-- ----------------------------

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`sort` AS `sort`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`default_menu` AS `default_menu`,`sys_base_menus`.`close_tab` AS `close_tab`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
