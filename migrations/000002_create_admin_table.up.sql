CREATE TABLE IF NOT EXISTS `gin_admin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一id号',
  `account` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `password` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `phone` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `salt` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `real_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `register_time` bigint unsigned NOT NULL COMMENT '注册时间',
  `register_ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '注册ip',
  `login_time` bigint unsigned NOT NULL COMMENT '登录时间',
  `login_ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录ip',
  `role_ids` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色IDs',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1：正常 0：禁用',
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT="后台管理员表";