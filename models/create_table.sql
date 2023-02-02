
CREATE DATABASE IF NOT EXISTS `bluebell`;

CREATE TABLE `user`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT `0`,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `idx_username` (`username`) USING BTREE,
     UNIQUE KEY `idx_user_id`  (`user_id`)  USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `community_id` int(10) unsigned NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_general_ci,
    `introduct_name` varchar(256) COLLATE utf8mb4_general_ci,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`) ,
    UNIQUE KEY `idx_community_name`  (`community_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community` VALUES('1','1','GO','golang','2020-11-12 08:10:10','2020-12-11 09:09:09');
INSERT INTO `community` VALUES('2','2','python','python','2021-11-12 08:10:10','2019-12-11 09:09:09');
INSERT INTO `community` VALUES('3','3','CS:GO','A通','2021-12-22 13:10:10','2020-12-23 19:19:19');
INSERT INTO `community` VALUES('4','4','LOL','欢迎来到英雄联盟','2022-10-12 09:10:10','2020-2-11 18:09:09');

DROP TABLE IF EXISTS `post`;
CREATE TABLE IF NOT EXISTS `post`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL COMMENT '帖子id',
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_id` bigint(20) NOT NULL COMMENT  '作者的用户id',
    `community_id` bigint(20) NOT NULL COMMENT '所属社区',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;