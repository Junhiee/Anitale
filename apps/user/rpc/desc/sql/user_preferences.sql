CREATE TABLE
  `user_preferences` (
    `user_id` bigint unsigned NOT NULL COMMENT '用户唯一标识，与users表的id关联',
    `lang` varchar(10) DEFAULT 'en' COMMENT '用户的语言偏好',
    `timezone` varchar(50) DEFAULT 'UTC' COMMENT '用户的时区偏好',
    `receive_newsletter` tinyint(1) DEFAULT '1' COMMENT '用户是否接收邮件通知',
    `receive_sms` tinyint(1) DEFAULT '0' COMMENT '用户是否接收短信通知',
    `receive_push` tinyint(1) DEFAULT '1' COMMENT '用户是否接收推送通知',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '偏好设置的创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '偏好设置的更新时间',
    PRIMARY KEY (`user_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '存储用户的偏好设置，例如语言、时区、通知偏好等'