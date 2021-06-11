/*
 Navicat Premium Data Transfer

 Source Server         : test-cloud
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 13:28:16
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user_download1
-- ----------------------------
DROP TABLE IF EXISTS "cloud_download";
CREATE TABLE IF NOT EXISTS "cloud_download" (
  "id" VARCHAR PRIMARY KEY UNIQUE NOT NULL,-- id
  "user_id" varchar(64) NOT NULL,--用户id,外键  用户表  cloud_user的 id==cloud_download 的 userid
  "name" varchar(128) NOT NULL,--文件名字
  "ptime" date NOT NULL DEFAULT (datetime('now','localtime')),--创建时间
  "cid" varchar(64) NOT NULL,--文件cid
  "size" int(10) NOT NULL,--文件大小
  "path" varchar(128) NOT NULL,--文件路径
  "file_type" int(10) NOT NULL,--文件类型
  "is_downrecord" int(10) NOT NULL--是否下载记录
);

PRAGMA foreign_keys = true;
