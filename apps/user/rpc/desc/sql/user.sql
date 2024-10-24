CREATE TABLE
  `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户唯一标识',
    `username` varchar(50) NOT NULL COMMENT '用户名，必须唯一',
    `email` varchar(100) NOT NULL COMMENT '用户邮箱，必须唯一',
    `password_hash` varchar(255) NOT NULL COMMENT '加密后的用户密码',
    `is_active` tinyint(1) DEFAULT '1' COMMENT '用户账户是否激活',
    `is_verified` tinyint(1) DEFAULT '0' COMMENT '用户邮箱是否经过验证',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用户账户创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '用户账户最近更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `email` (`email`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '存储用户的基本信息，包括用户名、邮箱、密码等'