CREATE TABLE
  `user_tokens` (
    `user_id` bigint unsigned NOT NULL COMMENT '用户唯一标识，与users表的id关联',
    `access_token` varchar(255) NOT NULL COMMENT '用户的访问令牌',
    `expires_at` timestamp NULL DEFAULT NULL COMMENT '访问令牌的过期时间',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '令牌创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '令牌更新时间',
    PRIMARY KEY (`user_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '存储用户登录后的访问令牌和刷新令牌信息'