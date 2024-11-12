CREATE TABLE `user_subscriptions` (
    `subscription_id` bigint NOT NULL AUTO_INCREMENT COMMENT '订阅记录的唯一标识',
    `user_id` bigint unsigned NOT NULL COMMENT '用户的唯一标识',
    `anime_id` bigint NOT NULL COMMENT '动画的唯一标识',
    `subscribed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '订阅时间',
    `notification_preference` varchar(20) DEFAULT 'web' COMMENT '通知方式，值可为 email, sms, app, none',
    `status` varchar(10) DEFAULT 'active' COMMENT '订阅状态，值可为 active 或 inactive',
    PRIMARY KEY (`subscription_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户订阅动画表'