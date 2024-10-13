CREATE TABLE `anime_tags` (
    `anime_tags_id` bigint NOT NUll AUTO_INCREMENT COMMENT '主键',
    `anime_id` bigint NOT NULL COMMENT '动画ID',
    `tag_id` bigint NOT NULL COMMENT '标签ID',
    PRIMARY KEY (`anime_tags_id`),
    UNIQUE KEY (`anime_id`, `tag_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '动画标签关系表'