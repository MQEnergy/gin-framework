CREATE TABLE `gin_user_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `role_ids` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '角色ID 例如：1,2,3',
  `created_at` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;