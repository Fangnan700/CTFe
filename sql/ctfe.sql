/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 172.23.144.1:3306
 Source Schema         : ctfe

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 22/07/2023 10:23:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Create database
-- ----------------------------
DROP DATABASE IF EXISTS `ctfe`;
CREATE DATABASE `ctfe` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `ctfe`;

-- ----------------------------
-- Table structure for challenges
-- ----------------------------
DROP TABLE IF EXISTS `challenges`;
CREATE TABLE `challenges`
(
    `challenge_id`   bigint                                                        NOT NULL,
    `challenge_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `challenge_type` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci   NOT NULL,
    `description`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL,
    `image_name`     varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
    `init_score`     int                                                           NOT NULL,
    PRIMARY KEY (`challenge_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for competitions
-- ----------------------------
DROP TABLE IF EXISTS `competitions`;
CREATE TABLE `competitions`
(
    `competition_id`   bigint                                                        NOT NULL,
    `competition_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `description`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL,
    `start_time`       bigint                                                        NOT NULL,
    `life_circle`      bigint                                                        NOT NULL,
    PRIMARY KEY (`competition_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for containers
-- ----------------------------
DROP TABLE IF EXISTS `containers`;
CREATE TABLE `containers`
(
    `container_id`   bigint                                                        NOT NULL,
    `vessel_id`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `container_ip`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `container_port` int                                                           NOT NULL,
    `start_time`     bigint                                                        NOT NULL,
    `life_circle`    bigint                                                        NOT NULL,
    `challenge_id`   bigint                                                        NOT NULL,
    `group_id`       bigint                                                        NOT NULL,
    PRIMARY KEY (`container_id`) USING BTREE,
    INDEX `containers_challenge_id` (`challenge_id` ASC) USING BTREE,
    INDEX `containers_group_id` (`group_id` ASC) USING BTREE,
    CONSTRAINT `containers_challenge_id` FOREIGN KEY (`challenge_id`) REFERENCES `challenges` (`challenge_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `containers_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for groups
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups`
(
    `group_id`       bigint                                                        NOT NULL,
    `group_name`     varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `intro`          text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL,
    `competition_id` bigint                                                        NOT NULL,
    PRIMARY KEY (`group_id`) USING BTREE,
    INDEX `groups_competition_id` (`competition_id` ASC) USING BTREE,
    CONSTRAINT `groups_competition_id` FOREIGN KEY (`competition_id`) REFERENCES `competitions` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for logs
-- ----------------------------
DROP TABLE IF EXISTS `logs`;
CREATE TABLE `logs`
(
    `log_id`         bigint                                                       NOT NULL,
    `competition_id` bigint                                                       NOT NULL,
    `challenge_id`   bigint                                                       NOT NULL,
    `group_id`       bigint                                                       NOT NULL,
    `user_id`        bigint                                                       NOT NULL,
    `log_time`       bigint                                                       NOT NULL,
    `log_type`       varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `log_content`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci        NOT NULL,
    PRIMARY KEY (`log_id`) USING BTREE,
    INDEX `logs_competition_id` (`competition_id` ASC) USING BTREE,
    INDEX `logs_challenge_id` (`challenge_id` ASC) USING BTREE,
    INDEX `logs_group_id` (`group_id` ASC) USING BTREE,
    INDEX `logs_user_id` (`user_id` ASC) USING BTREE,
    CONSTRAINT `logs_challenge_id` FOREIGN KEY (`challenge_id`) REFERENCES `challenges` (`challenge_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `logs_competition_id` FOREIGN KEY (`competition_id`) REFERENCES `competitions` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `logs_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `logs_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for participation
-- ----------------------------
DROP TABLE IF EXISTS `participation`;
CREATE TABLE `participation`
(
    `group_id` bigint NOT NULL,
    `user_id`  bigint NOT NULL,
    INDEX `participation_group_id` (`group_id` ASC) USING BTREE,
    INDEX `participation_user_id` (`user_id` ASC) USING BTREE,
    CONSTRAINT `participation_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `participation_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for scores
-- ----------------------------
DROP TABLE IF EXISTS `scores`;
CREATE TABLE `scores`
(
    `score_id`       bigint NOT NULL,
    `competition_id` bigint NOT NULL,
    `group_id`       bigint NOT NULL,
    `total_score`    int    NULL DEFAULT NULL,
    PRIMARY KEY (`score_id`) USING BTREE,
    INDEX `scores_competition_id` (`competition_id` ASC) USING BTREE,
    INDEX `scores_group_id` (`group_id` ASC) USING BTREE,
    CONSTRAINT `scores_competition_id` FOREIGN KEY (`competition_id`) REFERENCES `competitions` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `scores_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for solved
-- ----------------------------
DROP TABLE IF EXISTS `solved`;
CREATE TABLE `solved`
(
    `solved_id`      bigint NOT NULL,
    `solved_time`    bigint NOT NULL,
    `competition_id` bigint NOT NULL,
    `challenge_id`   bigint NOT NULL,
    `group_id`       bigint NOT NULL,
    `user_id`        bigint NOT NULL,
    `solved_score`   int    NOT NULL,
    PRIMARY KEY (`solved_id`) USING BTREE,
    INDEX `solved_competition_id` (`competition_id` ASC) USING BTREE,
    INDEX `solved_challenge_id` (`challenge_id` ASC) USING BTREE,
    INDEX `solved_group_id` (`group_id` ASC) USING BTREE,
    INDEX `solved_user_id` (`user_id` ASC) USING BTREE,
    CONSTRAINT `solved_challenge_id` FOREIGN KEY (`challenge_id`) REFERENCES `challenges` (`challenge_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `solved_competition_id` FOREIGN KEY (`competition_id`) REFERENCES `competitions` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `solved_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `solved_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `user_id`     bigint                                                        NOT NULL,
    `user_name`   varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL,
    `user_pwd`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `user_sex`    varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci   NOT NULL,
    `email`       varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `phone`       varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `school`      varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
    `student_num` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
    `create_time` bigint                                                        NOT NULL,
    PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Init data
-- ----------------------------
INSERT INTO users (user_id, user_name, user_pwd, user_sex, email, phone, school, student_num, create_time)
VALUES ('100000', 'init', 'init', 'init', 'init', 'init', 'init', 'init', 0);