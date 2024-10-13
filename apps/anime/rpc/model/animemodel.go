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
	}

	customAnimeLogicModel interface {
		AnimeList(ctx context.Context, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		// PageByTag(ctx context.Context, tag string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByRegion(ctx context.Context, region string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByAnimeType(ctx context.Context, anime_type string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageBySeasonDate(ctx context.Context, start_date time.Time, end_date time.Time, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByCond(ctx context.Context, cond map[string]interface{}, page int64, page_size int64, anime_ids []int64, tx *gorm.DB) ([]Anime, error)
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

// 通过 region 分页
func (m *customAnimeModel) PageByRegion(ctx context.Context, region string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return conn.Where("region = ?", region).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})
	return animes, err
}

// 通过 season 分页
func (m *customAnimeModel) PageBySeasonDate(ctx context.Context, start_date time.Time, end_date time.Time, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	/*
		SELECT id, name, release_date FROM anime
		WHERE release_date BETWEEN ? AND ?
		ORDER BY release_date
		LIMIT ? OFFSET ?`
	*/
	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return conn.Where("release_date BETWEEN ? AND ?", start_date, end_date).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})

	return animes, err
}

// 通过 anime_type 分页
func (m *customAnimeModel) PageByAnimeType(ctx context.Context, anime_type string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return conn.Where("anime_type = ?", anime_type).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})
	return animes, err
}

type Conditon struct {
	Region    string
	AnimeType string
}

// 多条件查询待优化
func (m *defaultAnimeModel) PageByCond(ctx context.Context, cond map[string]interface{}, page int64, page_size int64, anime_ids []int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime
	var err error

	conditon := Conditon{
		Region:    cond["region"].(string),
		AnimeType: cond["anime_type"].(string),
	}

	if anime_ids == nil && cond["start_date"] == nil && cond["end_date"] == nil {
		err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
			return conn.Where(&conditon).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
		})
	}

	if anime_ids == nil {
		err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
			return conn.Where(&conditon).Where("release_date BETWEEN ? AND ?", cond["start_date"], cond["end_date"]).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
		})
	}

	if cond["start_date"] == nil && cond["end_date"] == nil {
		err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
			return conn.Where("anime_id IN ?", anime_ids).Where(&conditon).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
		})
	}

	if anime_ids != nil && cond["start_date"] != nil && cond["end_date"] != nil {
		err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
			return conn.Where("anime_id IN ?", anime_ids).Where(&conditon).Where("release_date BETWEEN ? AND ?", cond["start_date"], cond["end_date"]).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
		})
	}

	return animes, err
}
