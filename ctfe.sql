SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Create database
-- ----------------------------
DROP DATABASE IF EXISTS `ctfe`;
CREATE DATABASE `ctfe` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `ctfe`;

-- ----------------------------
-- Table structure for ctfe_admin
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_admin`;
CREATE TABLE `ctfe_admin`
(
    `admin_id` bigint NOT NULL,
    `user_id`  bigint NOT NULL,
    PRIMARY KEY (`admin_id`) USING BTREE,
    INDEX `admin_user` (`user_id` ASC) USING BTREE,
    CONSTRAINT `admin_user` FOREIGN KEY (`user_id`) REFERENCES `ctfe_user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ctfe_challenge
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_challenge`;
CREATE TABLE `ctfe_challenge`
(
    `challenge_id`      bigint                                                        NOT NULL,
    `challenge_name`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `challenge_type`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `description`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `image_name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `init_score`        bigint                                                        NOT NULL,
    `dynamic_container` tinyint                                                       NOT NULL,
    `dynamic_flag`      tinyint                                                       NOT NULL,
    `competition_id`    bigint                                                        NOT NULL,
    PRIMARY KEY (`challenge_id`) USING BTREE,
    INDEX `challenge_competition` (`competition_id` ASC) USING BTREE,
    CONSTRAINT `challenge_competition` FOREIGN KEY (`competition_id`) REFERENCES `ctfe_competition` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ctfe_challenge
-- ----------------------------

-- ----------------------------
-- Table structure for ctfe_competition
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_competition`;
CREATE TABLE `ctfe_competition`
(
    `competition_id`   bigint                                                        NOT NULL,
    `competition_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `description`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `start_time`       bigint                                                        NOT NULL,
    `left_time`        bigint                                                        NOT NULL,
    PRIMARY KEY (`competition_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ctfe_competition
-- ----------------------------

-- ----------------------------
-- Table structure for ctfe_container
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_container`;
CREATE TABLE `ctfe_container`
(
    `container_id`   bigint                                                        NOT NULL,
    `vessel_tag`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `container_host` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `container_port` int                                                           NOT NULL,
    `start_time`     bigint                                                        NOT NULL,
    `life_time`      bigint                                                        NOT NULL,
    `competition_id` bigint                                                        NOT NULL,
    `challenge_id`   bigint                                                        NOT NULL,
    `group_id`       bigint                                                        NOT NULL,
    `user_id`        bigint                                                        NOT NULL,
    PRIMARY KEY (`container_id`) USING BTREE,
    INDEX `container_competition` (`competition_id` ASC) USING BTREE,
    INDEX `container_challenge` (`challenge_id` ASC) USING BTREE,
    INDEX `container_group` (`group_id` ASC) USING BTREE,
    INDEX `container_user` (`user_id` ASC) USING BTREE,
    CONSTRAINT `container_challenge` FOREIGN KEY (`challenge_id`) REFERENCES `ctfe_challenge` (`challenge_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `container_competition` FOREIGN KEY (`competition_id`) REFERENCES `ctfe_competition` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `container_group` FOREIGN KEY (`group_id`) REFERENCES `ctfe_group` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `container_user` FOREIGN KEY (`user_id`) REFERENCES `ctfe_user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ctfe_container
-- ----------------------------

-- ----------------------------
-- Table structure for ctfe_group
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_group`;
CREATE TABLE `ctfe_group`
(
    `group_id`       bigint                                                        NOT NULL,
    `group_name`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `group_intro`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `competition_id` bigint                                                        NOT NULL,
    PRIMARY KEY (`group_id`) USING BTREE,
    INDEX `group_competition` (`competition_id` ASC) USING BTREE,
    CONSTRAINT `group_competition` FOREIGN KEY (`competition_id`) REFERENCES `ctfe_competition` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ctfe_group
-- ----------------------------

-- ----------------------------
-- Table structure for ctfe_participation
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_participation`;
CREATE TABLE `ctfe_participation`
(
    `participation_id` bigint  NOT NULL,
    `group_id`         bigint  NOT NULL,
    `user_id`          bigint  NOT NULL,
    `competition_id`   bigint  NOT NULL,
    `is_admin`         tinyint NOT NULL,
    PRIMARY KEY (`participation_id`) USING BTREE,
    INDEX `participation_group` (`group_id` ASC) USING BTREE,
    INDEX `participation_user` (`user_id` ASC) USING BTREE,
    INDEX `participation_competition` (`competition_id` ASC) USING BTREE,
    CONSTRAINT `participation_group` FOREIGN KEY (`group_id`) REFERENCES `ctfe_group` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `participation_user` FOREIGN KEY (`user_id`) REFERENCES `ctfe_user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `participation_competition` FOREIGN KEY (`competition_id`) REFERENCES `ctfe_competition` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ctfe_participation
-- ----------------------------

-- ----------------------------
-- Table structure for ctfe_score
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_score`;
CREATE TABLE `ctfe_score`
(
    `score_id`       bigint NOT NULL,
    `score`          bigint NOT NULL,
    `competition_id` bigint NOT NULL,
    `challenge_id`   bigint NOT NULL,
    `group_id`       bigint NOT NULL,
    `user_id`        bigint NOT NULL,
    PRIMARY KEY (`score_id`) USING BTREE,
    INDEX `score_competition` (`competition_id` ASC) USING BTREE,
    INDEX `score_challenge` (`challenge_id` ASC) USING BTREE,
    INDEX `score_group` (`group_id` ASC) USING BTREE,
    INDEX `score_user` (`user_id` ASC) USING BTREE,
    CONSTRAINT `score_challenge` FOREIGN KEY (`challenge_id`) REFERENCES `ctfe_challenge` (`challenge_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `score_competition` FOREIGN KEY (`competition_id`) REFERENCES `ctfe_competition` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `score_group` FOREIGN KEY (`group_id`) REFERENCES `ctfe_group` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `score_user` FOREIGN KEY (`user_id`) REFERENCES `ctfe_user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ctfe_score
-- ----------------------------

-- ----------------------------
-- Table structure for ctfe_solved
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_solved`;
CREATE TABLE `ctfe_solved`
(
    `solved_id`      bigint NOT NULL,
    `solved_time`    bigint NOT NULL,
    `solved_score`   bigint NOT NULL,
    `competition_id` bigint NOT NULL,
    `challenge_id`   bigint NOT NULL,
    `group_id`       bigint NOT NULL,
    `user_id`        bigint NOT NULL,
    PRIMARY KEY (`solved_id`) USING BTREE,
    INDEX `solved_competition` (`competition_id` ASC) USING BTREE,
    INDEX `solved_challenge` (`challenge_id` ASC) USING BTREE,
    INDEX `solved_group` (`group_id` ASC) USING BTREE,
    INDEX `solved_user` (`user_id` ASC) USING BTREE,
    CONSTRAINT `solved_challenge` FOREIGN KEY (`challenge_id`) REFERENCES `ctfe_challenge` (`challenge_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `solved_competition` FOREIGN KEY (`competition_id`) REFERENCES `ctfe_competition` (`competition_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `solved_group` FOREIGN KEY (`group_id`) REFERENCES `ctfe_group` (`group_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `solved_user` FOREIGN KEY (`user_id`) REFERENCES `ctfe_user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ctfe_solved
-- ----------------------------

-- ----------------------------
-- Table structure for ctfe_user
-- ----------------------------
DROP TABLE IF EXISTS `ctfe_user`;
CREATE TABLE `ctfe_user`
(
    `uuid`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `user_id`     bigint                                                        NOT NULL,
    `user_name`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `user_pwd`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `user_sex`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `user_email`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `user_phone`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `user_school` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` bigint                                                        NOT NULL,
    PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Init data
-- ----------------------------

SET @UUID := (SELECT UUID());
SET @USER_ID := (SELECT FLOOR(0 + RAND() * (9223372036854775800 - 0)));
SET @ADMIN_ID := (SELECT FLOOR(0 + RAND() * (9223372036854775800 - 0)));

INSERT INTO `ctfe_user`
VALUES (@UUID, @USER_ID, '方楠', 'A792DE594BDF1906D30241FF3DFAEE5F', '男', '2621737589@qq.com', '18376938582',
        '安徽大学',
        unix_timestamp(NOW(3)));
SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO `ctfe_admin`
VALUES (@ADMIN_ID, @USER_ID);
