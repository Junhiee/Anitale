CREATE TABLE
  `episodes` (
    `episode_id` bigint NOT NULL AUTO_INCREMENT COMMENT '剧集的唯一标识',
    `anime_id` bigint DEFAULT NULL COMMENT '所属动画的ID，逻辑上指向 anime 表',
    `episode_number` int DEFAULT NULL COMMENT '剧集的编号，如第几集',
    `title` varchar(255) DEFAULT NULL COMMENT '剧集标题',
    `release_date` timestamp NULL DEFAULT NULL COMMENT '放送日期',
    `duration` int DEFAULT NULL COMMENT '剧集时长，单位为分钟',
    `synopsis` text COMMENT '剧集的内容概要',
    `video_url` varchar(255) DEFAULT NULL COMMENT '剧集视频的URL',
    PRIMARY KEY (`episode_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '动画剧集信息表'