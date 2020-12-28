ALTER TABLE `admins` 
MODIFY COLUMN `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'Unique Email' AFTER `last_name`,
MODIFY COLUMN `phone_number` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL AFTER `email`,
MODIFY COLUMN `phone_number_country_code` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL AFTER `phone_number`,
MODIFY COLUMN `country` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL AFTER `phone_number_country_code`,
ADD COLUMN `created_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) AFTER `country`,
ADD COLUMN `updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) AFTER `created_at`;