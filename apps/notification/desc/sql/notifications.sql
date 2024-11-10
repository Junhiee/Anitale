CREATE TABLE
  `notifications` (
    `notification_id` bigint NOT NULL AUTO_INCREMENT COMMENT '通知记录的唯一标识',
    `subscription_id` bigint NOT NULL COMMENT '订阅记录ID，用于追踪是哪个用户的订阅',
    `update_id` bigint NOT NULL COMMENT '更新事件ID，用于标识是哪个更新事件',
    `notification_sent_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '通知发送时间',
    `notification_status` varchar(10) DEFAULT 'sent' COMMENT '通知状态，值可为 sent 或 failed',
    PRIMARY KEY (`notification_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '通知发送记录表'