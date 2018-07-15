/*
Navicat MySQL Data Transfer

Source Server         : 121
Source Server Version : 50720
Source Host           : 192.168.1.121:3306
Source Database       : wx_db

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2018-07-15 15:15:16
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
  `Content` text NOT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;
