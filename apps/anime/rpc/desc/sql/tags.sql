CREATE TABLE `tags` (
    `tag_id`bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
    `tag` varchar(255) NOT NULL COMMENT '标签',
    PRIMARY KEY (`tag_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '标签表'