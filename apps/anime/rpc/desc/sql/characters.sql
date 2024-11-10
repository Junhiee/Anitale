CREATE TABLE
  `characters` (
    `character_id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色的唯一标识',
    `anime_id` bigint DEFAULT NULL COMMENT '所属动画的ID，逻辑上指向 anime 表',
    `name` varchar(255) NOT NULL COMMENT '角色的名字',
    `role` varchar(20) DEFAULT 'supporting' COMMENT '角色类型，如：main, supporting, cameo',
    `description` text COMMENT '角色的简介',
    `image_url` varchar(255) DEFAULT NULL COMMENT '角色的图片URL',
    PRIMARY KEY (`character_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '动画角色信息表'