package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ StatsModel = (*customStatsModel)(nil)

type (
	// StatsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStatsModel.
	StatsModel interface {
		statsModel
		customStatsLogicModel
	}

	customStatsModel struct {
		*defaultStatsModel
	}

	customStatsLogicModel interface {
		GetStasByAinmeIds(ctx context.Context, animeIds []int64) ([]Stats, error)
	}
)

// NewStatsModel returns a model for the database table.
func NewStatsModel(conn *gorm.DB, c cache.CacheConf) StatsModel {
	return &customStatsModel{
		defaultStatsModel: newStatsModel(conn, c),
	}
}

func (m *defaultStatsModel) customCacheKeys(data *Stats) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *customStatsModel) GetStasByAinmeIds(ctx context.Context, animeIds []int64) ([]Stats, error) {
	var stats []Stats
	// err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
	// 	err := conn.Where("anime_id IN ?", animeIds).Find(&stats).Error
	// 	return err
	// })

	err := m.QueryCtx(ctx, &stats, "cache:stats:rank", func(conn *gorm.DB, v interface{}) error {

		return conn.Where("anime_id IN ?", animeIds).Find(&stats).Error
	})

	return stats, err
}

type AnimeHotScore struct {
	AnimeId int64
	Score   float64
}

func (m *customStatsModel) SetAnimeidToStats(ctx context.Context, anime_id int64, score float64) {

}
