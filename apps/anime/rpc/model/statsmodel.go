package model

import (
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
