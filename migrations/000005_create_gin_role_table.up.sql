-- 创建角色表
CREATE TABLE IF NOT EXISTS `gin_role` (
    `id`         int                                                          NOT NULL AUTO_INCREMENT,
    `name`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
    `desc`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色描述',
    `status`     tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1正常(默认) 0停用',
    `role_type`  tinyint(1) NOT NULL DEFAULT '1' COMMENT '角色类型 1：web角色 2：app角色',
    `created_at` int                                                          NOT NULL,
    `updated_at` int                                                          NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 创建超级管理员角色
INSERT INTO `gin_role`(`id`, `name`, `desc`, `status`, `role_type`, `created_at`, `updated_at`) VALUES (1, '超级管理员', '超级管理员', 1, 1, 1645932052, 1645932052);