/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 14:42:02
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE IF NOT EXISTS cloud_user (
                      id VARCHAR ( 64 ) PRIMARY KEY UNIQUE NOT NULL,--id
                      peer_id VARCHAR ( 64 ) NOT NULL,--用户id==peerid
                      name VARCHAR ( 128 ) NOT NULL,--文件名字
                      phone VARCHAR ( 64 ) NOT NULL UNIQUE,--手机号
                      sex INT ( 10 ) NOT NULL DEFAULT ( 1 ),--性别
                      ptime DATE NOT NULL DEFAULT ( datetime( 'now', 'localtime' ) ),--创建时间
                      utime DATE NOT NULL DEFAULT ( datetime( 'now', 'localtime' ) ),--更新时间
                      nickname VARCHAR ( 128 )--别名
);
PRAGMA foreign_keys = true;
