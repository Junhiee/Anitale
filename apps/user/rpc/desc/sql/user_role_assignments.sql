CREATE TABLE
  `user_role_assignments` (
    `assignment_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色分配关系唯一标识',
    `user_id` bigint unsigned DEFAULT NULL COMMENT '用户唯一标识，与users表的id关联',
    `role_id` bigint unsigned DEFAULT NULL COMMENT '角色唯一标识，与user_roles表的id关联',
    `assigned_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '角色分配时间',
    PRIMARY KEY (`assignment_id`),
    UNIQUE KEY `user_id` (`user_id`, `role_id`) COMMENT '确保每个用户的每个角色分配唯一'
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '存储用户与其对应的角色关系'