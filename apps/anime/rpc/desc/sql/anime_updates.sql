CREATE TABLE
  `anime_updates` (
    `update_id` bigint NOT NULL AUTO_INCREMENT COMMENT '更新事件的唯一标识',
    `anime_id` bigint NOT NULL COMMENT '动画的唯一标识',
    `episode_id` bigint DEFAULT NULL COMMENT '新发布的剧集ID',
    `update_type` varchar(20) DEFAULT 'new_episode' COMMENT '更新类型，值可为 new_episode, season_update, general_update',
    `update_description` varchar(255) DEFAULT NULL COMMENT '更新内容描述',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`update_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '动画更新事件表'