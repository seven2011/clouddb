/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 14:34:35
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS "chat_record";
CREATE TABLE IF NOT EXISTS chat_record (
                             id VARCHAR ( 64 ) PRIMARY KEY UNIQUE NOT NULL,--id
                             name varchar ( 64 ) NOT NULL,--聊天对方名称
                             topic VARCHAR ( 64 ) UNIQUE NOT NULL,--主题
                             img VARCHAR ( 128 ) NOT NULL,--头像
                             create_by VARCHAR ( 64 ) NOT NULL,--创建者
                             ptime DATE NOT NULL DEFAULT ( datetime( 'now', 'localtime' ) ),--创建时间
                             last_msg TEXT NOT NULL --最后的消息
);
PRAGMA foreign_keys = true;
