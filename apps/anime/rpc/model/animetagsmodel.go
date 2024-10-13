package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AnimeTagsModel = (*customAnimeTagsModel)(nil)

type (
	// AnimeTagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnimeTagsModel.
	AnimeTagsModel interface {
		animeTagsModel
		customAnimeTagsLogicModel
	}

	customAnimeTagsModel struct {
		*defaultAnimeTagsModel
	}

	customAnimeTagsLogicModel interface {
		PageByTags(ctx context.Context, tag_id int64, tx *gorm.DB) ([]AnimeTags, error)
	}
)

// NewAnimeTagsModel returns a model for the database table.
func NewAnimeTagsModel(conn *gorm.DB, c cache.CacheConf) AnimeTagsModel {
	return &customAnimeTagsModel{
		defaultAnimeTagsModel: newAnimeTagsModel(conn, c),
	}
}

func (m *defaultAnimeTagsModel) customCacheKeys(data *AnimeTags) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

// 通过 tag_id 查找 anime_id
func (m *customAnimeTagsModel) PageByTags(ctx context.Context, tag_id int64, tx *gorm.DB) ([]AnimeTags, error) {
	var anime_tags []AnimeTags
	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return conn.Select("anime_id").Where("tag_id = ?", tag_id).Find(&anime_tags).Error
	})

	return anime_tags, err
}
