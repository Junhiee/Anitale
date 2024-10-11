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
		PageByTag(ctx context.Context, tag string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByCountry(ctx context.Context, country string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByAnimeType(ctx context.Context, anime_type string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByRelaseDate(ctx context.Context, relase_date time.Time, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
		PageByCond(ctx context.Context, cond map[string]string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error)
	}
)

// NewAnimeModel returns a model for the database table.
func NewAnimeModel(conn *gorm.DB, c cache.CacheConf) AnimeModel {
	return &customAnimeModel{
		defaultAnimeModel: newAnimeModel(conn, c),
	}
}

// 默认分页, 通过 update_at 排序
func (m *customAnimeModel) AnimeList(ctx context.Context, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime
	// m.QueryNoCacheCtx(ctx, &animes, func(conn *gorm.DB, v interface{}) error {})
	err := m.ExecCtx(ctx, func(db *gorm.DB) error {
		return db.Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})

	return animes, err
}

// 通过tag分页
func (m *customAnimeModel) PageByTag(ctx context.Context, tag string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	// SELECT tag FROM anime where tag like '%,test_tag,%' or tag like 'test_tag,%' or tag like '%,test_tag' or tag='test_tag';
	err := m.ExecCtx(ctx, func(db *gorm.DB) error {
		return db.Where("tag = ?", tag).Or("tag like ?", tag+",%").Or("tag like ?", "%,"+tag).Or("tag like ?", "%,"+tag+",%").Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})

	return animes, err
}

// 通过 country 分页
func (m *customAnimeModel) PageByCountry(ctx context.Context, country string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	err := m.ExecCtx(ctx, func(db *gorm.DB) error {
		return db.Where("country = ?", country).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})

	return animes, err
}

// TODO 按照季度过滤 需要写一个判断每个季度范围的方法
// 通过 relase_date 分页
func (m *customAnimeModel) PageByRelaseDate(ctx context.Context, relase_date time.Time, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime

	// 大于 relase_date, 小于下个季度 (3个月一个季度)
	err := m.ExecCtx(ctx, func(db *gorm.DB) error {
		return db.Where("relase_date > ? AND relase_date < ?", relase_date, relase_date.AddDate(0, 3, 0)).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})

	return animes, err
}

// 通过 anime_type 分页
func (m *customAnimeModel) PageByAnimeType(ctx context.Context, anime_type string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	var animes []Anime
	err := m.ExecCtx(ctx, func(db *gorm.DB) error {
		return db.Where("anime_type = ?", anime_type).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})
	return animes, err
}



// TODO 多条件查询待优化
func (m *defaultAnimeModel) PageByCond(ctx context.Context, cond map[string]string, page int64, page_size int64, tx *gorm.DB) ([]Anime, error) {
	// cond: 筛选条件
	// country 国家或地区
	// tag 标签
	// season | relase_date 季度
	// anime_type 类型

	var animes []Anime

	err := m.ExecCtx(ctx, func(db *gorm.DB) error {
		return db.Where(cond).Where("country = ?", cond["country"]).Where("tag = ?", cond["tag"]).Or("tag like ?", cond["tag"]+",%").Or("tag like ?", "%,"+cond["tag"]).Or("tag like ?", "%,"+cond["tag"]+",%").Where("anime_type = ?", cond["anime_type"]).Order("updated_at").Limit(int(page_size)).Offset(int((page - 1) * page_size)).Find(&animes).Error
	})

	return animes, err
}

func (m *defaultAnimeModel) customCacheKeys(data *Anime) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
