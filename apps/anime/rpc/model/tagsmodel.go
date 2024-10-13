package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TagsModel = (*customTagsModel)(nil)

type (
	// TagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagsModel.
	TagsModel interface {
		tagsModel
		customTagsLogicModel
	}

	customTagsModel struct {
		*defaultTagsModel
	}

	customTagsLogicModel interface {
		FindOneByTag(ctx context.Context, tag string) (int64, error)
	}
)

// NewTagsModel returns a model for the database table.
func NewTagsModel(conn *gorm.DB, c cache.CacheConf) TagsModel {
	return &customTagsModel{
		defaultTagsModel: newTagsModel(conn, c),
	}
}

func (m *defaultTagsModel) customCacheKeys(data *Tags) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

// 通过 tag 查找 tag_id
func (m *customTagsModel) FindOneByTag(ctx context.Context, tag string) (int64, error) {
	var tags Tags
	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return conn.Select("tag_id").Where("tag = ?", tag).Find(&tags).Error
	})

	return tags.TagId, err
}
