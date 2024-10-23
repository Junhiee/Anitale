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
		PageBySesson(ctx context.Context, startMonth int, endMonth int, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
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

// 多条件查询
func (m *customAnimeModel) PageByCond(ctx context.Context, cond map[string]interface{}, page int64, page_size int64, anime_ids []int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime
	var err error

	query := tx.Model(&Anime{})

	if region, ok := cond["region"].(string); ok && region != "" {
		query = query.Where("region = ?", region)
	}

	if animeType, ok := cond["anime_type"].(string); ok && animeType != "" {
		query = query.Where("anime_type = ?", animeType)
	}

	// 日期范围条件
	if startDate, startOk := cond["start_date"].(time.Time); startOk {
		if endDate, endOk := cond["end_date"].(time.Time); endOk {
			query = query.Where("release_date BETWEEN ? AND ?", startDate, endDate)
		}
	}

	// 处理季度筛选条件（根据月份进行过滤）
	if startMonth, startOk := cond["start_month"].(time.Time); startOk {
		if endMonth, endOk := cond["end_month"].(time.Time); endOk {
			query = query.Where("MONTH(release_date) BETWEEN ? AND ?", startMonth.Month(), endMonth.Month())
		}
	}

	// Anime ID 列表
	if len(anime_ids) > 0 {
		query = query.Where("anime_id IN ?", anime_ids)
	}

	// 排序和分页
	query = query.Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size))

	// 执行查询
	err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return query.Find(&animes).Error
	})

	return animes, err
}

// 通过 season 分页，返回该季度范围内的anime动画(所有年份)
func (m *customAnimeModel) PageBySesson(ctx context.Context, startMonth int, endMonth int, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		// 使用 MONTH(release_date) 进行月份过滤，忽略年份
		return conn.Where("MONTH(release_date) BETWEEN ? AND ?", startMonth, endMonth).
			Order("updated_at").
			Limit(int(page_size)).
			Offset(int((page - 1) * page_size)).
			Find(&animes).Error
	})

	return animes, err
}

// conditon := Conditon{
// 	Region:    cond["region"].(string),
// 	AnimeType: cond["anime_type"].(string),
// }

// if anime_ids == nil && cond["start_date"] == nil && cond["end_date"] == nil {
// 	err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
// 		return conn.Where(&conditon).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
// 	})
// }

// if anime_ids == nil {
// 	err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
// 		return conn.Where(&conditon).Where("release_date BETWEEN ? AND ?", cond["start_date"], cond["end_date"]).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
// 	})
// }

// if cond["start_date"] == nil && cond["end_date"] == nil {
// 	err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
// 		return conn.Where("anime_id IN ?", anime_ids).Where(&conditon).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
// 	})
// }

// if anime_ids != nil && cond["start_date"] != nil && cond["end_date"] != nil {
// 	err = m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
// 		return conn.Where("anime_id IN ?", anime_ids).Where(&conditon).Where("release_date BETWEEN ? AND ?", cond["start_date"], cond["end_date"]).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
// 	})
// }
