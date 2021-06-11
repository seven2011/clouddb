/*
 Navicat Premium Data Transfer

 Source Server         : test-cloud
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 13:29:42
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for cloud_file
-- ----------------------------
DROP TABLE IF EXISTS "userfile";
CREATE TABLE IF NOT EXISTS "cloud_file" (
                             "id"  VARCHAR PRIMARY KEY UNIQUE NOT NULL,-- id
                             "user_id" varchar ( 64 ) NOT NULL,--用户表的外键id
                             "name" varchar ( 255 ) NOT NULL,--文件名字
                             "parent_id" int ( 64 ) NOT NULL,--父id
                             "ptime" date NOT NULL DEFAULT ( datetime( 'now', 'localtime' ) ),--上传时间
                             "cid" varchar ( 255 ) NOT NULL,--文件cid
                             "size" int ( 10 ) NOT NULL,--文件大小
                             "status" int ( 10 ) NOT NULL,--状态(0正常1已删除
                             "file_type" int ( 10 ) NOT NULL,--文件类型
                             "folder" int ( 10 ) NOT NULL --是否是文件夹
);

PRAGMA foreign_keys = true;
