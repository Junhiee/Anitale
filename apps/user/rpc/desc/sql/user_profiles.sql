CREATE TABLE
  `user_profiles` (
    `user_id` bigint unsigned NOT NULL COMMENT '与users表的ID关联，用户唯一标识',
    `full_name` varchar(100) DEFAULT NULL COMMENT '用户的全名',
    `bio` text COMMENT '用户个人简介',
    `avatar_url` varchar(255) DEFAULT NULL COMMENT '用户头像的URL',
    `birthday` date DEFAULT NULL COMMENT '用户的生日',
    `gender` varchar(255) DEFAULT NULL COMMENT '用户性别',
    `loc` varchar(255) DEFAULT NULL COMMENT '用户的所在地',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '资料创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '资料更新时间',
    PRIMARY KEY (`user_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '存储用户的个人详细资料'