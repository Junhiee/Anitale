CREATE TABLE stats (
    `anime_id` bigint COMMENT '主键',
    `view_count` bigint DEFAULT 0 COMMENT '播放数量',
    `like_count` bigint DEFAULT 0 COMMENT '点赞数量',
    `comment_count` bigint DEFAULT 0 COMMENT '评论数量',
    `share_count` bigint DEFAULT 0 COMMENT '转发数量',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`anime_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '数据统计表'