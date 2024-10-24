package model

import (
	"context"

	"github.com/SpectatorNan/gorm-zero/gormc"
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
		FindOneNoCacheCtx(ctx context.Context, animeId int64) (*Stats, error)
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

func (m *customStatsModel) FindOneNoCacheCtx(ctx context.Context, animeId int64) (*Stats, error) {
	var resp Stats
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Stats{}).Where("`anime_id` = ?", animeId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
