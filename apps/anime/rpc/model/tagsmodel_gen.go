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
	cacheAnitaleTagsTagIdPrefix = "cache:anitale:tags:tagId:"
)

type (
	tagsModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Tags) error
		BatchInsert(ctx context.Context, tx *gorm.DB, news []Tags) error
		FindOne(ctx context.Context, tagId int64) (*Tags, error)
		Update(ctx context.Context, tx *gorm.DB, data *Tags) error
		BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []Tags) error
		BatchDelete(ctx context.Context, tx *gorm.DB, datas []Tags) error

		Delete(ctx context.Context, tx *gorm.DB, tagId int64) error
		// deprecated. recommend add a transaction in service context instead of using this
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultTagsModel struct {
		gormc.CachedConn
		table string
	}

	Tags struct {
		TagId int64  `gorm:"column:tag_id;primary_key"` // 主键
		Tag   string `gorm:"column:tag"`                // 标签
	}
)

func (Tags) TableName() string {
	return "`tags`"
}

func newTagsModel(conn *gorm.DB, c cache.CacheConf) *defaultTagsModel {
	return &defaultTagsModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`tags`",
	}
}

func (m *defaultTagsModel) GetCacheKeys(data *Tags) []string {
	if data == nil {
		return []string{}
	}
	anitaleTagsTagIdKey := fmt.Sprintf("%s%v", cacheAnitaleTagsTagIdPrefix, data.TagId)
	cacheKeys := []string{
		anitaleTagsTagIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultTagsModel) Insert(ctx context.Context, tx *gorm.DB, data *Tags) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.GetCacheKeys(data)...)
	return err
}
func (m *defaultTagsModel) BatchInsert(ctx context.Context, tx *gorm.DB, news []Tags) error {

	err := batchx.BatchExecCtx(ctx, m, news, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Create(&news).Error
	})

	return err
}

func (m *defaultTagsModel) FindOne(ctx context.Context, tagId int64) (*Tags, error) {
	anitaleTagsTagIdKey := fmt.Sprintf("%s%v", cacheAnitaleTagsTagIdPrefix, tagId)
	var resp Tags
	err := m.QueryCtx(ctx, &resp, anitaleTagsTagIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Tags{}).Where("`tag_id` = ?", tagId).First(&resp).Error
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

func (m *defaultTagsModel) Update(ctx context.Context, tx *gorm.DB, data *Tags) error {
	old, err := m.FindOne(ctx, data.TagId)
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
func (m *defaultTagsModel) BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []Tags) error {
	clearData := make([]Tags, 0, len(olds)+len(news))
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

func (m *defaultTagsModel) Delete(ctx context.Context, tx *gorm.DB, tagId int64) error {
	data, err := m.FindOne(ctx, tagId)
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
		return db.Delete(&Tags{}, tagId).Error
	}, m.GetCacheKeys(data)...)
	return err
}

func (m *defaultTagsModel) BatchDelete(ctx context.Context, tx *gorm.DB, datas []Tags) error {
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
func (m *defaultTagsModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultTagsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAnitaleTagsTagIdPrefix, primary)
}

func (m *defaultTagsModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Tags{}).Where("`tag_id` = ?", primary).Take(v).Error
}
