// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/SpectatorNan/gorm-zero/gormc/batchx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheAnimeTagsAnimeTagsIdPrefix  = "cache:animeTags:animeTagsId:"
	cacheAnimeTagsAnimeIdTagIdPrefix = "cache:animeTags:animeId:tagId:"
)

type (
	animeTagsModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *AnimeTags) error
		BatchInsert(ctx context.Context, tx *gorm.DB, news []AnimeTags) error
		FindOne(ctx context.Context, animeTagsId int64) (*AnimeTags, error)
		FindOneByAnimeIdTagId(ctx context.Context, animeId int64, tagId int64) (*AnimeTags, error)
		Update(ctx context.Context, tx *gorm.DB, data *AnimeTags) error
		BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []AnimeTags) error
		BatchDelete(ctx context.Context, tx *gorm.DB, datas []AnimeTags) error

		Delete(ctx context.Context, tx *gorm.DB, animeTagsId int64) error
		// deprecated. recommend add a transaction in service context instead of using this
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultAnimeTagsModel struct {
		gormc.CachedConn
		table string
	}

	AnimeTags struct {
		AnimeTagsId int64 `gorm:"column:anime_tags_id;primary_key"` // 主键
		AnimeId     int64 `gorm:"column:anime_id"`                  // 动画ID
		TagId       int64 `gorm:"column:tag_id"`                    // 标签ID
	}
)

func (AnimeTags) TableName() string {
	return "`anime_tags`"
}

func newAnimeTagsModel(conn *gorm.DB, c cache.CacheConf) *defaultAnimeTagsModel {
	return &defaultAnimeTagsModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`anime_tags`",
	}
}

func (m *defaultAnimeTagsModel) GetCacheKeys(data *AnimeTags) []string {
	if data == nil {
		return []string{}
	}
	animeTagsAnimeIdTagIdKey := fmt.Sprintf("%s%v:%v", cacheAnimeTagsAnimeIdTagIdPrefix, data.AnimeId, data.TagId)
	animeTagsAnimeTagsIdKey := fmt.Sprintf("%s%v", cacheAnimeTagsAnimeTagsIdPrefix, data.AnimeTagsId)
	cacheKeys := []string{
		animeTagsAnimeIdTagIdKey, animeTagsAnimeTagsIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultAnimeTagsModel) Insert(ctx context.Context, tx *gorm.DB, data *AnimeTags) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.GetCacheKeys(data)...)
	return err
}
func (m *defaultAnimeTagsModel) BatchInsert(ctx context.Context, tx *gorm.DB, news []AnimeTags) error {

	err := batchx.BatchExecCtx(ctx, m, news, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Create(&news).Error
	})

	return err
}

func (m *defaultAnimeTagsModel) FindOne(ctx context.Context, animeTagsId int64) (*AnimeTags, error) {
	animeTagsAnimeTagsIdKey := fmt.Sprintf("%s%v", cacheAnimeTagsAnimeTagsIdPrefix, animeTagsId)
	var resp AnimeTags
	err := m.QueryCtx(ctx, &resp, animeTagsAnimeTagsIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&AnimeTags{}).Where("`anime_tags_id` = ?", animeTagsId).First(&resp).Error
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

func (m *defaultAnimeTagsModel) FindOneByAnimeIdTagId(ctx context.Context, animeId int64, tagId int64) (*AnimeTags, error) {
	animeTagsAnimeIdTagIdKey := fmt.Sprintf("%s%v:%v", cacheAnimeTagsAnimeIdTagIdPrefix, animeId, tagId)
	var resp AnimeTags
	err := m.QueryRowIndexCtx(ctx, &resp, animeTagsAnimeIdTagIdKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&AnimeTags{}).Where("`anime_id` = ? and `tag_id` = ?", animeId, tagId).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.AnimeTagsId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAnimeTagsModel) Update(ctx context.Context, tx *gorm.DB, data *AnimeTags) error {
	old, err := m.FindOne(ctx, data.AnimeTagsId)
	if err != nil && errors.Is(err, ErrNotFound) {
		return err
	}
	clearKeys := append(m.GetCacheKeys(old), m.GetCacheKeys(data)...)
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, clearKeys...)
	return err
}
func (m *defaultAnimeTagsModel) BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []AnimeTags) error {
	clearData := make([]AnimeTags, 0, len(olds)+len(news))
	clearData = append(clearData, olds...)
	clearData = append(clearData, news...)
	err := batchx.BatchExecCtx(ctx, m, clearData, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&news).Error
	})

	return err
}

func (m *defaultAnimeTagsModel) Delete(ctx context.Context, tx *gorm.DB, animeTagsId int64) error {
	data, err := m.FindOne(ctx, animeTagsId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&AnimeTags{}, animeTagsId).Error
	}, m.GetCacheKeys(data)...)
	return err
}

func (m *defaultAnimeTagsModel) BatchDelete(ctx context.Context, tx *gorm.DB, datas []AnimeTags) error {
	err := batchx.BatchExecCtx(ctx, m, datas, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&datas).Error
	})

	return err
}

// deprecated. recommend add a transaction in service context instead of using this
func (m *defaultAnimeTagsModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultAnimeTagsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAnimeTagsAnimeTagsIdPrefix, primary)
}

func (m *defaultAnimeTagsModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&AnimeTags{}).Where("`anime_tags_id` = ?", primary).Take(v).Error
}
