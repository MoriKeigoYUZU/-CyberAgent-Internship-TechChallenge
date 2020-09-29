CREATE DATABASE IF NOT EXISTS hlr_db;

CREATE TABLE IF NOT EXISTS `hlr_db`.`user` (
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `password` VARCHAR(64) NOT NULL COMMENT 'パスワード',
  `auth_token` VARCHAR(128) NOT NULL COMMENT '認証トークン',
  `coin` INT UNSIGNED NOT NULL COMMENT '所持コイン',
  PRIMARY KEY (`name`),
ENGINE = InnoDB
COMMENT = 'ユーザ';

CREATE TABLE IF NOT EXISTS `hlr_db`.`user` (
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `password` VARCHAR(64) NOT NULL COMMENT 'パスワード',
  `auth_token` VARCHAR(128) NOT NULL COMMENT '認証トークン',
  `coin` INT UNSIGNED NOT NULL COMMENT '所持コイン',
  PRIMARY KEY (`name`))
ENGINE = InnoDB
COMMENT = 'ユーザ';

CREATE TABLE IF NOT EXISTS `hlr_db`.`rankig` (
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `stage_id` INT NOT NULL COMMENT 'ステージID',
  `score` INT NOT NULL COMMENT 'スコア',
  PRIMARY KEY (`name` , `stage_id`))
ENGINE = InnoDB
COMMENT = 'ランキング';
