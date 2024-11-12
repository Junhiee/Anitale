package model

import (
	"context"
	"gorm.io/gorm"
)

var _ UserSubscriptionsModel = (*customUserSubscriptionsModel)(nil)

type (
	// UserSubscriptionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSubscriptionsModel.
	UserSubscriptionsModel interface {
		userSubscriptionsModel
		customUserSubscriptionsLogicModel
	}

	customUserSubscriptionsModel struct {
		*defaultUserSubscriptionsModel
	}

	customUserSubscriptionsLogicModel interface {
		IsSubscribed(ctx context.Context, userId uint64, animeId int64) (bool, error)
		DeleteByAnimeIdAndUserId(ctx context.Context, userId uint64, animeId int64) error
	}
)

// NewUserSubscriptionsModel returns a model for the database table.
func NewUserSubscriptionsModel(conn *gorm.DB) UserSubscriptionsModel {
	return &customUserSubscriptionsModel{
		defaultUserSubscriptionsModel: newUserSubscriptionsModel(conn),
	}
}

// IsSubscribed 判断用户是否已经订阅
func (m *customUserSubscriptionsModel) IsSubscribed(ctx context.Context, userId uint64, animeId int64) (bool, error) {
	var count int64
	err := m.conn.WithContext(ctx).Model(&UserSubscriptions{}).
		Where("user_id = ? AND anime_id = ? AND status = 'active'", userId, animeId).
		Count(&count).Error

	// 如果 count == 0，表示没有找到符合条件的记录，返回 false。
	// 如果 count > 0，表示用户已订阅该动画，返回 true
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *customUserSubscriptionsModel) DeleteByAnimeIdAndUserId(ctx context.Context, userId uint64, animeId int64) error {
	err := m.conn.WithContext(ctx).
		Where("user_id = ? AND anime_id = ?", userId, animeId).
		Delete(&UserSubscriptions{}).Error

	if err != nil {
		return err
	}

	return nil
}
