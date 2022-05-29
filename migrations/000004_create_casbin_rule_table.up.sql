CREATE TABLE IF NOT EXISTS `casbin_rule` (
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `v0`    varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `v1`    varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `v2`    varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `v3`    varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `v4`    varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `v5`    varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;