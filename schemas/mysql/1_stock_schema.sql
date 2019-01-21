CREATE DATABASE IF NOT EXISTS `stock`
  COLLATE 'utf8_general_ci'
  DEFAULT CHARSET 'utf8';

USE `stock`;

CREATE TABLE `entities` (
  `account_id` bigint(20) unsigned NOT NULL COMMENT 'Идентификатор',
  `entity_id` binary(16) NOT NULL COMMENT 'Идентификатор сущности',
  `entity_type` varchar(255) NOT NULL COMMENT 'Тип сущности',
  `schema` json NOT NULL,
  `entity` blob NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Дата создания',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'состояние удален',
  PRIMARY KEY (`account_id`,`entity_id`),
  UNIQUE KEY `entity_id` (`entity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `edges` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` bigint(20) unsigned NOT NULL COMMENT 'Идентификатор',
  `subject_id` binary(16) NOT NULL COMMENT 'Идентификатор субъекта',
  `predicate` varchar(64) NOT NULL COMMENT 'предикат',
  `object_id` binary(16) NOT NULL COMMENT 'Идентификатор объекта',
  `priority` bigint(20) unsigned NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `account_id_2` (`account_id`,`object_id`,`predicate`,`deleted`),
  KEY `priority_index` (`account_id`,`subject_id`,`predicate`,`deleted`,`priority`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
