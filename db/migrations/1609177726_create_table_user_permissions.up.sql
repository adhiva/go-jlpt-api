CREATE TABLE `user_permissions`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `admin_id` int NOT NULL,
  `permission_id` int NOT NULL,
  `created_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(),
  `updated_at` datetime(0) NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`),
  CONSTRAINT `admin_id_foreign` FOREIGN KEY (`admin_id`) REFERENCES `jlpt_n5`.`admins` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
  CONSTRAINT `permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `jlpt_n5`.`permissions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);