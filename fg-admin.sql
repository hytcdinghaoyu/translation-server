SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `username` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role_id` int unsigned DEFAULT NULL,
  `email` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- ----------------------------
--  Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `raw_content` longtext COLLATE utf8mb4_general_ci DEFAULT NULL,
  `translate_content` longtext COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
