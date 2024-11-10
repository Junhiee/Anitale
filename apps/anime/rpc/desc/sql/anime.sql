CREATE TABLE `anime` (
    `anime_id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title` varchar(255) DEFAULT NULL COMMENT '标题',
    `desc` varchar(255) DEFAULT NULL COMMENT '简介',
    `region` varchar(255) DEFAULT NULL COMMENT '国家或地区',
    `format` varchar(50) DEFAULT NULL COMMENT '动画种类',
    `img_url` varchar(255) DEFAULT NULL COMMENT '图片地址',
    `studios` varchar(255) DEFAULT NULL COMMENT '工作室',
    `status` varchar(255) DEFAULT NULL COMMENT '动画状态',
    `rating` float DEFAULT NULL COMMENT '评分',
    `release_date` timestamp NULL DEFAULT NULL COMMENT '推出日期',
    `update_date` timestamp NULL DEFAULT NULL COMMENT '更新日期',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`anime_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '动画表'