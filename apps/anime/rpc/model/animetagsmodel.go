package model

import (
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"

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
		FindTagIdByAnimeId(ctx context.Context, animeId int64) ([]AnimeTags, error)
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

// 通过 anime_id 查找 tag_id
func (m *customAnimeTagsModel) FindTagIdByAnimeId(ctx context.Context, animeId int64) ([]AnimeTags, error) {
	var resp []AnimeTags
	var animeTags = AnimeTags{}
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Model(&animeTags).Where("anime_id = ?", animeId).Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
