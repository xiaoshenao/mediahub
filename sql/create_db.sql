CREATE SCHEMA `mediahub` DEFAULT CHARACTER SET utf8mb4 ;
use mediahub;

CREATE TABLE `mediahub`.`url_map` (
`id` BIGINT(64) NOT NULL AUTO_INCREMENT,
`short_key` VARCHAR(45) NOT NULL DEFAULT '',
`original_url` VARCHAR(512) NOT NULL DEFAULT '',
`times` INT NOT NULL DEFAULT 0,
`create_at` BIGINT(64) NOT NULL DEFAULT 0,
`update_at` BIGINT(64) NOT NULL DEFAULT 0,
PRIMARY KEY (`id`),
INDEX `index_original_url` USING BTREE (`original_url`) VISIBLE,
INDEX `index_short_key` USING BTREE (`short_key`) VISIBLE,
UNIQUE INDEX `unique_original_url` USING BTREE (`original_url`) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;


CREATE TABLE `mediahub`.`url_map_user` (
 `id` BIGINT(64) NOT NULL AUTO_INCREMENT,
 `user_id` BIGINT(64) NOT NULL DEFAULT 0,
 `short_key` VARCHAR(45) NOT NULL DEFAULT '',
 `original_url` VARCHAR(512) NOT NULL DEFAULT '',
 `times` INT NOT NULL DEFAULT 0,
 `create_at` BIGINT(64) NOT NULL DEFAULT 0,
 `update_at` BIGINT(64) NOT NULL DEFAULT 0,
 PRIMARY KEY (`id`),
 INDEX `index_user_id` USING BTREE (`user_id`) VISIBLE,
 INDEX `index_original_url` USING BTREE (`original_url`) VISIBLE,
 INDEX `index_short_key` USING BTREE (`short_key`) VISIBLE,
 UNIQUE INDEX `unique_original_url` USING BTREE (`original_url`) VISIBLE)
 ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;