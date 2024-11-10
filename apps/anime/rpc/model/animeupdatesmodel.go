package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AnimeUpdatesModel = (*customAnimeUpdatesModel)(nil)

type (
	// AnimeUpdatesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnimeUpdatesModel.
	AnimeUpdatesModel interface {
		animeUpdatesModel
		customAnimeUpdatesLogicModel
	}

	customAnimeUpdatesModel struct {
		*defaultAnimeUpdatesModel
	}

	customAnimeUpdatesLogicModel interface {
	}
)

// NewAnimeUpdatesModel returns a model for the database table.
func NewAnimeUpdatesModel(conn *gorm.DB, c cache.CacheConf) AnimeUpdatesModel {
	return &customAnimeUpdatesModel{
		defaultAnimeUpdatesModel: newAnimeUpdatesModel(conn, c),
	}
}

func (m *defaultAnimeUpdatesModel) customCacheKeys(data *AnimeUpdates) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
