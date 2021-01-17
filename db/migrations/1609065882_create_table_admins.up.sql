CREATE TABLE `admins` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`first_name` VARCHAR(25),
	`last_name` VARCHAR(50),
	`email` VARCHAR(50) NOT NULL COMMENT 'Unique Email',
	`password` varchar(255) NOT NULL,
	`token` VARCHAR(255) NULL,
	`phone_number` VARCHAR(15) NOT NULL,
	`phone_number_country_code` VARCHAR(5) NOT NULL,
	`country` VARCHAR(50) NOT NULL,
	`active` SMALLINT(1) NOT NULL DEFAULT 1,
	`created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
	`updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
	UNIQUE KEY `unique_email` (`email`) USING BTREE,
    UNIQUE KEY `unique_phone_number_and_phone_number_country_code` (`phone_number`,`phone_number_country_code`) USING BTREE,
    KEY `search_full_name` (`first_name`,`last_name`) USING BTREE,
    KEY `search_email` (`email`) USING BTREE,
	PRIMARY KEY (`id`)
);