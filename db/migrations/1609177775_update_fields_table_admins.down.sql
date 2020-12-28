ALTER TABLE `admins` 
MODIFY COLUMN `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT 'Unique Email' AFTER `last_name`,
MODIFY COLUMN `phone_number` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci AFTER `email`,
MODIFY COLUMN `phone_number_country_code` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci AFTER `phone_number`,
MODIFY COLUMN `country` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci AFTER `phone_number_country_code`,
DROP COLUMN `created_at`,
DROP COLUMN `updated_at`;