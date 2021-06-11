/*
 Navicat Premium Data Transfer

 Source Server         : test-cloud
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 13:17:22
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user_like
-- ----------------------------
DROP TABLE IF EXISTS "user_like";
CREATE TABLE IF NOT EXISTS "article_like" (
                             "id" VARCHAR PRIMARY KEY UNIQUE NOT NULL,-- id
                             "user_id" varchar ( 64 ) DEFAULT NULL,--用户id 用户表的外键
                             "file_id" int ( 64 ) DEFAULT NULL--文件id
);

PRAGMA foreign_keys = true;
