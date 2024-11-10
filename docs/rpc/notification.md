# 更新推送服务

## 数据库设计

**订阅表**

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

**更新事件表**

```sql
CREATE TABLE anime_updates (
    update_id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '更新事件的唯一标识',
    anime_id BIGINT NOT NULL COMMENT '动画的唯一标识',
    episode_id BIGINT COMMENT '新发布的剧集ID',
    update_type VARCHAR(20) DEFAULT 'new_episode' COMMENT '更新类型，值可为 new_episode, season_update, general_update',
    update_description VARCHAR(255) COMMENT '更新内容描述',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='动画更新事件表';

```

**通知记录表**

```sql
CREATE TABLE notifications (
    notification_id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '通知记录的唯一标识',
    subscription_id BIGINT NOT NULL COMMENT '订阅记录ID，用于追踪是哪个用户的订阅',
    update_id BIGINT NOT NULL COMMENT '更新事件ID，用于标识是哪个更新事件',
    notification_sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '通知发送时间',
    notification_status VARCHAR(10) DEFAULT 'sent' COMMENT '通知状态，值可为 sent 或 failed'
) COMMENT='通知发送记录表';

```

### 业务流程详细步骤

#### 1. 用户订阅动画

用户通过客户端（如移动应用或网站）订阅某个动画，系统会将订阅记录保存到 **`user_subscriptions`** 表中。

- **请求流程**：

  - 客户端发送 **订阅请求**，包括 `user_id` 和 `anime_id`。
  - 服务器处理请求，将 `user_id`、`anime_id`、通知偏好等信息插入到 `user_subscriptions` 表。

- **数据库操作**：

  ```sql
  INSERT INTO user_subscriptions (user_id, anime_id, notification_preference, status)
  VALUES (?, ?, 'app', 'active');
  ```

#### 2. 动画更新事件生成

当动画发布新剧集或有其他内容更新时，更新服务（或内容管理系统）会生成一条更新事件，并将其存储到 **`anime_updates`** 表，同时将该更新事件发送到消息队列中。

- **事件流程**：

  - 更新服务检测到新的剧集或内容发布，将 `anime_id`、`episode_id`、`update_type` 和描述等信息插入到 `anime_updates` 表。
  - 将该更新事件包装成消息，通过消息队列（如 Kafka 或 RabbitMQ）发送到 `anime_update_queue` 队列中。

- **数据库操作**：

  ```sql
  INSERT INTO anime_updates (anime_id, episode_id, update_type, update_description)
  VALUES (?, ?, 'new_episode', '新剧集发布');
  ```

- **消息队列操作**（示例消息内容）：

  ```json
  {
    "update_id": 123,
    "anime_id": 456,
    "episode_id": 789,
    "update_type": "new_episode",
    "description": "新剧集发布"
  }
  ```

#### 3. 通知服务监听消息队列

通知服务作为消费者监听 `anime_update_queue` 消息队列，接收更新事件，并根据 `anime_id` 查询所有订阅该动画的用户，然后逐一推送通知。

- **监听流程**：

  - 通知服务从消息队列中接收更新事件消息，解析出 `anime_id` 等信息。
  - 通过查询 `user_subscriptions` 表，获取订阅该动画的所有 `user_id`。
  - 根据每个用户的 `notification_preference` 设置，选择合适的推送渠道（如 APP、短信、邮件）发送通知。
  - 通知成功后，将通知记录保存到 `notifications` 表，以防重复通知并便于追踪历史。

- **查询用户订阅**：

  ```sql
  SELECT subscription_id, user_id, notification_preference 
  FROM user_subscriptions 
  WHERE anime_id = ? AND status = 'active';
  ```

- **发送通知**（示例）：

  - 对于 `notification_preference` 为 `app` 的用户，可以通过应用内推送服务发送通知。
  - 对于 `notification_preference` 为 `email` 的用户，可以通过邮件服务发送更新信息。

- **通知记录入库**：

  ```sql
  INSERT INTO notifications (subscription_id, update_id, notification_status)
  VALUES (?, ?, 'sent');
  ```

#### 4. 用户收到通知并查看更新

用户收到更新通知后，可以打开应用查看更新的动画信息和新发布的剧集。用户的操作不会影响数据库表和消息队列的存储。



### 消息队列

**API 服务** 是请求驱动的，关注快速响应和同步反馈，适合直接面向用户的业务逻辑。

**消息队列消费者** 是事件驱动的，关注消息的可靠处理和异步处理能力，适合后台任务、延迟容忍的处理逻辑。

一般来说需要把消费者独立设计成一个服务



