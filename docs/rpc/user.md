# 用户服务

## 数据库设计

**用户订阅表**

```sql
CREATE TABLE user_subscriptions (
    subscription_id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '订阅记录的唯一标识',
    user_id BIGINT NOT NULL COMMENT '用户的唯一标识',
    anime_id BIGINT NOT NULL COMMENT '动画的唯一标识',
    subscribed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '订阅时间',
    notification_preference VARCHAR(20) DEFAULT 'app' COMMENT '通知方式，值可为 email, sms, app, none',
    status VARCHAR(10) DEFAULT 'active' COMMENT '订阅状态，值可为 active 或 inactive'
) COMMENT='用户订阅动画表';

```

