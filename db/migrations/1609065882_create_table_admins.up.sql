CREATE TABLE `admins` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`first_name` VARCHAR(25),
	`last_name` VARCHAR(50),
	`email` VARCHAR(50) COMMENT 'Unique Email',
	`phone_number` VARCHAR(15),
	`phone_number_country_code` VARCHAR(5),
	`country` VARCHAR(50),
	UNIQUE KEY `unique_email` (`email`) USING BTREE,
    UNIQUE KEY `unique_phone_number_and_phone_number_country_code` (`phone_number`,`phone_number_country_code`) USING BTREE,
    KEY `search_full_name` (`first_name`,`last_name`) USING BTREE,
    KEY `search_email` (`email`) USING BTREE,
	PRIMARY KEY (`id`)
);