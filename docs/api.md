# 业务

动画服务 API

增：向数据库新增一条 anime 数据

删：从数据库删除一条 anime 数据

改：从数据库修改一条 anime 数据

查：分页查找、搜索、筛选

- 分页查找
- 搜索：搜索番剧名字
  - [meilisearch](https://www.meilisearch.com/)
  - [meilisearch-go](https://github.com/meilisearch/meilisearch-go?tab=readme-ov-file#-documentation)

- 筛选
  - tag 标签
  - anime_type 种类
  - region 国家
  - season 年份 - 季度
  - update_date | update_time 按更新日期排序
  - 按热度排序 后面做

返回的基本格式

```go
type Base {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
    // Data []*Data `json:"data"`
}
```
