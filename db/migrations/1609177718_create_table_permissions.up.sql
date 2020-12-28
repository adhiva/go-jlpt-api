CREATE TABLE `permissions` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`slug` VARCHAR(50) NOT NULL,
	`title` VARCHAR(100),
	`title_en` VARCHAR(100),
	`created_at` DATETIME NOT NULL DEFAULT 'CURRENT_TIMESTAMP()',
	`updated_at` DATETIME NOT NULL DEFAULT 'CURRENT_TIMESTAMP()',
	UNIQUE KEY `unique_slug` (`slug`) USING BTREE,
KEY `index_title` (`title`,`title_en`) USING BTREE,
KEY `index_timestamp` (`created_at`,`updated_at`) USING BTREE,
	PRIMARY KEY (`id`)
);