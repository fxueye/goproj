/*
Navicat MySQL Data Transfer

Source Server         : 121
Source Server Version : 50720
Source Host           : 192.168.1.121:3306
Source Database       : wx_db

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2018-07-15 11:20:10
*/

SET FOREIGN_KEY_CHECKS=0;
drop database IF EXISTS `wx_db`;
create database `wx_db`;
use `wx_db`;
-- ----------------------------
-- Table structure for `stroke`
-- ----------------------------
DROP TABLE IF EXISTS `stroke`;
CREATE TABLE `stroke` (
  `UID` int(11) NOT NULL DEFAULT '0',
  `Send` varchar(200) NOT NULL DEFAULT '',
  `Tel` varchar(200) NOT NULL DEFAULT '',
  `Content` text NOT NULL,
  `Timestamp` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of stroke
-- ----------------------------
