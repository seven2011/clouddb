/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 14:26:35
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS "cloud_article";
CREATE TABLE IF NOT EXISTS article (
                         id VARCHAR ( 64 ) PRIMARY KEY UNIQUE NOT NULL,--id
                         user_id VARCHAR ( 64 ) NOT NULL,--用户id==peerid
                         accesstory VARCHAR ( 128 ),--附件,cid
                         accessory_type INT ( 10 ) NOT NULL,--附件类型
                         text text NOT NULL,--文本
                         "name" varchar ( 255 ) NOT NULL,--文件名字
                         "ptime" date NOT NULL DEFAULT ( datetime( 'now', 'localtime' ) )--上传时间
);
PRAGMA foreign_keys = true;
