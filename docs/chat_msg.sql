/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 14:51:39
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for chat_msg
-- ----------------------------
DROP TABLE IF EXISTS "chat_msg";
CREATE TABLE IF NOT EXISTS chat_msg (
                          id VARCHAR ( 64 ) PRIMARY KEY UNIQUE NOT NULL,--id
                          content_type INT ( 10 ) NOT NULL,--文件内容类型
                          content TEXT NOT NULL,--文件内容
                          from_id VARCHAR ( 64 ) NOT NULL,--发送者id
                          to_id VARCHAR ( 64 ) NOT NULL,--接收者id
                          ptime DATE NOT NULL DEFAULT ( datetime( 'now', 'localtime' ) ),--创建时间
                          revocation INTEGER ( 10 ) NOT NULL DEFAULT ( 1 ),--撤回
                          is_read INT ( 10 ) NOT NULL DEFAULT ( 0 )--是否已读
);
PRAGMA foreign_keys = true;
