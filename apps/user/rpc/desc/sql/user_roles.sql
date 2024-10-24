CREATE TABLE
  `user_roles` (
    `role_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `role_name` varchar(50) NOT NULL DEFAULT 'guest' COMMENT '权限的名称，admin|user|guest',
    `role_desc` varchar(255) DEFAULT NULL COMMENT '角色的描述信息',
    PRIMARY KEY (`role_id`),
    UNIQUE KEY `role_name` (`role_name`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '定义系统中不同用户的角色'