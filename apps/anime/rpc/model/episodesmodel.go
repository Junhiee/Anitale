package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ EpisodesModel = (*customEpisodesModel)(nil)

type (
	// EpisodesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEpisodesModel.
	EpisodesModel interface {
		episodesModel
		customEpisodesLogicModel
	}

	customEpisodesModel struct {
		*defaultEpisodesModel
	}

	customEpisodesLogicModel interface {
	}
)

// NewEpisodesModel returns a model for the database table.
func NewEpisodesModel(conn *gorm.DB, c cache.CacheConf) EpisodesModel {
	return &customEpisodesModel{
		defaultEpisodesModel: newEpisodesModel(conn, c),
	}
}

func (m *defaultEpisodesModel) customCacheKeys(data *Episodes) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
