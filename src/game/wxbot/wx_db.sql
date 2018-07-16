/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.143
Source Server Version : 50720
Source Host           : 192.168.1.143:3306
Source Database       : wx_db

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2018-07-16 16:10:21
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `stroke`
-- ----------------------------
DROP TABLE IF EXISTS `stroke`;
CREATE TABLE `stroke` (
  `UID` int(11) NOT NULL AUTO_INCREMENT,
  `Send` varchar(200) NOT NULL DEFAULT '',
  `Tel` varchar(200) NOT NULL DEFAULT '',
  `Content` text CHARACTER SET utf8mb4 NOT NULL,
  `Timestamp` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of stroke
-- ----------------------------
