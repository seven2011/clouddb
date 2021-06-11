/*
 Navicat Premium Data Transfer

 Source Server         : test-cloud
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 13:37:38
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for usershare
-- ----------------------------
DROP TABLE IF EXISTS "cloud_share";
CREATE TABLE IF NOT EXISTS "cloud_share" (
                              "id" VARCHAR PRIMARY KEY UNIQUE NOT NULL,-- id
                              "user_id" varchar ( 64 ) NOT NULL,-- 用户表外键id
                              "name" varchar ( 128 ) NOT NULL,-- 文件名字
                              "ptime" date NOT NULL DEFAULT ( datetime( 'now', 'localtime' ) ),--上传时间
                              "cid" varchar ( 64 ) NOT NULL,-- 文件cid
                              "code" text NOT NULL,--分享码
                              "status" int ( 10 ) NOT NULL,-- '0无密码|1七天|2十四天|3一个月|4永久'
                              "share_link" text NOT NULL,-- 分享链接
                              "share_info" text NOT NULL --分享信息

);

PRAGMA foreign_keys = true;
