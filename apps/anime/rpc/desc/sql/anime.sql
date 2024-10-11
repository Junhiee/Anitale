CREATE TABLE `anime` (
    `anime_id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title` varchar(255) DEFAULT NULL COMMENT '标题',
    `desc` varchar(255) DEFAULT NULL COMMENT '简介',
    `region` varchar(255) DEFAULT NULL COMMENT '国家或地区',
    `anime_type` varchar(50) DEFAULT NULL COMMENT '动画种类',
    `img_url` varchar(255) DEFAULT NULL COMMENT '图片地址',
    `studios` varchar(255) DEFAULT NULL COMMENT '工作室',
    `status` varchar(255) DEFAULT NULL COMMENT '动画状态',
    `rating` float DEFAULT NULL COMMENT '评分',
    `relase_date` timestamp NULL DEFAULT NULL COMMENT '推出日期',
    `update_date` timestamp NULL DEFAULT NULL COMMENT '更新日期',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`anime_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '动画表'

CREATE TABLE `tags` (
    `tag_id`bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
    `tag` varchar(255) NOT NULL COMMENT '标签',
    PRIMARY KEY (`tag_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '标签表'

CREATE TABLE `anime_tags` (
    `anime_id` bigint NOT NULL COMMENT '动画ID',
    `tag_id` bigint NOT NULL COMMENT '标签ID',
    PRIMARY KEY (`anime_id`, `tag_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '动画标签关系表'