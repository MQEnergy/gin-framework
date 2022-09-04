CREATE TABLE `casbin_rule`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '策略类型',
    `v0`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色ID',
    `v1`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api路径',
    `v2`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api访问方法',
    `v3`    varchar(100) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `v4`    varchar(100) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `v5`    varchar(100) COLLATE utf8mb4_unicode_ci                       DEFAULT NULL,
    `v6`    varchar(25) COLLATE utf8mb4_unicode_ci                        DEFAULT NULL,
    `v7`    varchar(25) COLLATE utf8mb4_unicode_ci                        DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;