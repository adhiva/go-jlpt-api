CREATE TABLE `user_permission`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `admin_id` int NULL,
  `permission_id` int NULL,
  `created_at` datetime(0) NULL DEFAULT NOW(),
  `updated_at` datetime(0) NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`),
  UNIQUE INDEX `user_permission_unique`(`admin_id`, `permission_id`) USING BTREE COMMENT 'User cannot duplicate permission',
  CONSTRAINT `admin_id_foreign` FOREIGN KEY (`admin_id`) REFERENCES `admins` (`id`) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT `permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE SET NULL ON UPDATE SET NULL
);