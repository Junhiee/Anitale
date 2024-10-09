// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/SpectatorNan/gorm-zero/gormc/batchx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheAnimeAnimeIdPrefix = "cache:anime:animeId:"
)

type (
	animeModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Anime) error
		BatchInsert(ctx context.Context, tx *gorm.DB, news []Anime) error
		FindOne(ctx context.Context, animeId int64) (*Anime, error)
		Update(ctx context.Context, tx *gorm.DB, data *Anime) error
		BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []Anime) error
		BatchDelete(ctx context.Context, tx *gorm.DB, datas []Anime) error

		Delete(ctx context.Context, tx *gorm.DB, animeId int64) error
		// deprecated. recommend add a transaction in service context instead of using this
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultAnimeModel struct {
		gormc.CachedConn
		table string
	}

	Anime struct {
		AnimeId    int64          `gorm:"column:anime_id;primary_key"` // 主键
		Title      sql.NullString `gorm:"column:title"`                // 标题
		Desc       sql.NullString `gorm:"column:desc"`                 // 简介
		Country    sql.NullString `gorm:"column:country"`              // 国家或地区
		AnimeType  sql.NullString `gorm:"column:anime_type"`           // 动画种类
		Tag        sql.NullString `gorm:"column:tag"`                  // 标签
		Studios    sql.NullString `gorm:"column:studios"`              // 工作室
		Status     sql.NullString `gorm:"column:status"`               // 动画状态
		Rating     sql.NullInt64  `gorm:"column:rating"`               // 评分
		RelaseDate sql.NullTime   `gorm:"column:relase_date"`          // 推出日期
		UpdateDate sql.NullTime   `gorm:"column:update_date"`          // 更新日期
		UpdatedAt  time.Time      `gorm:"column:updated_at"`           // 更新时间
		CreatedAt  time.Time      `gorm:"column:created_at"`           // 创建时间
	}
)

func (Anime) TableName() string {
	return "`anime`"
}

func newAnimeModel(conn *gorm.DB, c cache.CacheConf) *defaultAnimeModel {
	return &defaultAnimeModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`anime`",
	}
}

func (m *defaultAnimeModel) GetCacheKeys(data *Anime) []string {
	if data == nil {
		return []string{}
	}
	animeAnimeIdKey := fmt.Sprintf("%s%v", cacheAnimeAnimeIdPrefix, data.AnimeId)
	cacheKeys := []string{
		animeAnimeIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultAnimeModel) Insert(ctx context.Context, tx *gorm.DB, data *Anime) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.GetCacheKeys(data)...)
	return err
}
func (m *defaultAnimeModel) BatchInsert(ctx context.Context, tx *gorm.DB, news []Anime) error {

	err := batchx.BatchExecCtx(ctx, m, news, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Create(&news).Error
	})

	return err
}

func (m *defaultAnimeModel) FindOne(ctx context.Context, animeId int64) (*Anime, error) {
	animeAnimeIdKey := fmt.Sprintf("%s%v", cacheAnimeAnimeIdPrefix, animeId)
	var resp Anime
	err := m.QueryCtx(ctx, &resp, animeAnimeIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Anime{}).Where("`anime_id` = ?", animeId).First(&resp).Error
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

func (m *defaultAnimeModel) Update(ctx context.Context, tx *gorm.DB, data *Anime) error {
	old, err := m.FindOne(ctx, data.AnimeId)
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
func (m *defaultAnimeModel) BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []Anime) error {
	clearData := make([]Anime, 0, len(olds)+len(news))
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

func (m *defaultAnimeModel) Delete(ctx context.Context, tx *gorm.DB, animeId int64) error {
	data, err := m.FindOne(ctx, animeId)
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
		return db.Delete(&Anime{}, animeId).Error
	}, m.GetCacheKeys(data)...)
	return err
}

func (m *defaultAnimeModel) BatchDelete(ctx context.Context, tx *gorm.DB, datas []Anime) error {
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
func (m *defaultAnimeModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultAnimeModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAnimeAnimeIdPrefix, primary)
}

func (m *defaultAnimeModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Anime{}).Where("`anime_id` = ?", primary).Take(v).Error
}
