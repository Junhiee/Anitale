// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type Anime struct {
	AnimeID     int64    `json:"anime_id"`     // 主键
	Title       string   `json:"title"`        // 标题
	Desc        string   `json:"desc"`         // 简介
	ImgURL      string   `json:"img_url"`      // 图片地址
	Region      string   `json:"region"`       // 国家或地区
	Format      string   `json:"format"`       // 动画种类
	Tags        []string `json:"tags"`         // 动画标签
	Studios     string   `json:"studios"`      // 工作室
	Status      string   `json:"status"`       // 动画状态
	Rating      float64  `json:"rating"`       // 评分
	ReleaseDate string   `json:"release_date"` // 推出日期
	UpdateDate  string   `json:"update_date"`  // 更新日期
	UpdatedAt   string   `json:"updated_at"`   // 更新时间
	CreatedAt   string   `json:"created_at"`   // 创建时间
}

type AnimeListReq struct {
	Page     int64  `form:"page,default=1"`            // 页码，默认为 1
	PageSize int64  `form:"page_size,default=10"`      // 每页显示的动画数量，默认为 10
	Region   string `form:"region,optional"`           // 动画所在的国家或地区，可选
	Tag      string `form:"tag,optional"`              // 动画标签，用于按标签筛选，可选
	Format   string `form:"format,optional"`           // 动画种类，例如 TV、OVA 等，可选
	Year     int64  `form:"year,optional"`             // 动画年份，用于按年份筛选，可选
	Season   int64  `form:"season,optional"`           // 动画季节（如 1 表示春季、2 表示夏季），可选
	Sort     string `form:"sort,default=updated_time"` // 排序方式，支持按热度 (hot) 或更新时间 (updated_time)，默认更新时间
}

type AnimeListResp struct {
	AnimeList  []*Anime `json:"anime_list"`  // 动画数据列表
	Page       int64    `json:"page"`        // 当前页码
	PageSize   int64    `json:"page_size"`   // 每页条目数
	TotalCount int64    `json:"total_count"` // 总记录数
	TotalPages int64    `json:"total_pages"` // 总页数
}

type Character struct {
	CharacterID int64  `json:"character_id"` // 角色的唯一标识
	AnimeID     int64  `json:"anime_id"`     // 所属动画的 ID
	Name        string `json:"name"`         // 角色的名字
	Role        string `json:"role"`         // 角色类型，如 main, supporting, cameo
	Description string `json:"description"`  // 角色的简介
	ImageURL    string `json:"image_url"`    // 角色的图片 URL
}

type Episode struct {
	EpisodeID   int64  `json:"episode_id"`     // 剧集的唯一标识
	AnimeID     int64  `json:"anime_id"`       // 所属动画的 ID
	EpisodeNum  int    `json:"episode_number"` // 剧集编号，如第几集
	Title       string `json:"title"`          // 剧集标题
	ReleaseDate string `json:"release_date"`   // 放送日期
	Duration    int    `json:"duration"`       // 剧集时长，单位为分钟
	Synopsis    string `json:"synopsis"`       // 剧集的内容概要
	VideoURL    string `json:"video_url"`      // 剧集视频的URL
}

type GetCharacterListResp struct {
	Characters []*Character `json:"characters"`  // 角色列表
	TotalCount int64        `json:"total_count"` // 符合条件的角色总数
}

type GetCharacterReq struct {
	AnimeID int64  `path:"anime_id"`      // 动画 ID，用于筛选该动画中的角色
	Role    string `json:"role,optional"` // 角色类型，可选值：main（主角）、supporting（配角）、cameo（客串），默认 supporting
}

type GetEpisodeListResp struct {
	Episodes   []*Episode `json:"episodes"`    // 剧集列表
	TotalCount int64      `json:"total_count"` // 符合条件的剧集总数
}

type GetEpisodeReq struct {
	AnimeID    int64 `path:"anime_id"`                // 动画 ID，用于筛选该动画的剧集
	EpisodeNum int   `json:"episode_number,optional"` // 剧集编号，可选
}

type GetUserProfileReq struct {
	UserId uint64 `path:"user_id"`
}

type GetUserProfileResp struct {
	Profile UserProfile `json:"profile"`
}

type RegisterUserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResp struct {
	User User `json:"user"`
}

type Trending struct {
	AnimeID    int64  `json:"anime_id"`   // 动画的 ID
	Title      string `json:"title"`      // 动画的标题
	Rank       int    `json:"rank"`       // 排名
	Popularity int    `json:"popularity"` // 热度
	ImageURL   string `json:"image_url"`  // 动画封面图片
	Region     string `json:"region"`     // 动画所在地区
	Tag        string `json:"tag"`        // 动画种类
	Format     string `json:"format"`     // 动画格式
}

type TrendingReq struct {
	Since  string `json:"since,default=weekly,optional"` // 获取的时间范围：weekly, monthly, yearly
	Region string `json:"region,optional"`               // 地区筛选，支持不同地区的流行趋势
	Tag    string `json:"tag,optional"`                  // 动画种类：如动作、冒险等
	Format string `json:"format,optional"`               // 动画格式：如 TV、OVA、电影等
}

type TrendingResp struct {
	TrendingList []*Trending `json:"trending_list"` // 动画流行趋势列表
}

type User struct {
	Id         uint64 `json:"id"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	IsActive   bool   `json:"is_active"`
	IsVerified bool   `json:"is_verified"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	AccessToken string `json:"access_token"`
	Expire      int64  `json:"expire"`
}

type UserProfile struct {
	UserId    uint64 `json:"user_id"`
	FullName  string `json:"full_name"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Loc       string `json:"loc"`
}

type UserSubscribeReq struct {
	UserId  uint64 `json:"user_id"`
	AnimeId int64  `json:"anime_id"`
}

type UserSubscribeResp struct {
	Subscriptions UserSubscriptions `json:"subscriptions"`
}

type UserSubscriptions struct {
	SubscriptionId         int64  `json:"subscription_id"`
	UserId                 uint64 `json:"user_id"`
	AnimeId                int64  `json:"anime_id"`
	SubscribedAt           string `json:"subscribed_at"`
	NotificationPreference string `json:"notification_preference"`
	Status                 string `json:"status"`
}
