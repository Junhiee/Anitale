package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AnimeModel = (*customAnimeModel)(nil)

type (
	// AnimeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnimeModel.
	AnimeModel interface {
		animeModel
		customAnimeLogicModel
	}

	customAnimeModel struct {
		*defaultAnimeModel
		Condition
	}

	Condition struct {
		Region string
		Tag    string
		Format string
		Sort   string

		// 日期范围
		StartDate *time.Time
		EndDate   *time.Time

		// 季度范围
		StartMonth *time.Time
		EndMonth   *time.Time
	}

	customAnimeLogicModel interface {
		AnimeList(ctx context.Context, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByCondition(ctx context.Context, tx *gorm.DB, page int64, page_size int64, c Condition) ([]Anime, int64, error)
	}
)

// NewAnimeModel returns a model for the database table.
func NewAnimeModel(conn *gorm.DB, c cache.CacheConf) AnimeModel {
	return &customAnimeModel{
		defaultAnimeModel: newAnimeModel(conn, c),
	}
}

func (m *defaultAnimeModel) customCacheKeys(data *Anime) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

// 默认分页, 通过 update_at 排序
func (m *customAnimeModel) AnimeList(ctx context.Context, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return conn.Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})
	return animes, err
}

// 多条件查询
func (m *customAnimeModel) PageByCondition(ctx context.Context, tx *gorm.DB, page int64, page_size int64, c Condition) ([]Anime, int64, error) {
	var totalCount int64
	var anime_list []Anime
	var err error
	err = m.QueryNoCacheCtx(ctx, &anime_list, func(conn *gorm.DB, v interface{}) error {
		db := conn
		if tx != nil {
			db = tx
		}

		q := db.Model(&Anime{})

		// 地区
		if c.Region != "" {
			q = q.Where("region = ?", c.Region)
		}

		// 动画种类
		if c.Format != "" {
			q = q.Where("format = ?", c.Format)
		}

		// 时间范围（发布日期）
		if c.StartDate != nil && c.EndDate != nil && !c.StartDate.IsZero() && !c.EndDate.IsZero() {
			q = q.Where("release_date BETWEEN ? AND ?", *c.StartDate, *c.EndDate)
		}

		// 季度筛选条件（根据月份进行过滤）
		if c.StartMonth != nil && c.EndMonth != nil && !c.StartMonth.IsZero() && !c.EndMonth.IsZero() {
			q = q.Where("MONTH(release_date) BETWEEN ? AND ?", c.StartMonth.Month(), c.EndMonth.Month())
		}

		// 动画标签
		if c.Tag != "" {
			q.Joins("join anime_tags on anime_tags.anime_id = anime.anime_id").
				Joins("join tags on tags.tag_id = anime_tags.tag_id").
				Where("tag = ?", c.Tag)
		}

		// 获取 totalCount
		err := q.Count(&totalCount).Error
		if err != nil {
			return err
		}

		// 排序 默认时间排序
		if c.Sort == "hot" {
			q.Joins("join stats on stats.anime_id = anime.anime_id").Order("hot DESC")
		} else if c.Sort == "rating" {
			q.Order("rating DESC")
		} else {
			q.Order("updated_at")
		}

		// 分页
		q.Limit(int(page_size)).Offset(int((page - 1) * page_size))
		return q.Find(&anime_list).Error
	})

	return anime_list, totalCount, err
}
