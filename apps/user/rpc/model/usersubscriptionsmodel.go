package model

import (
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
	}
)

// NewUserSubscriptionsModel returns a model for the database table.
func NewUserSubscriptionsModel(conn *gorm.DB) UserSubscriptionsModel {
	return &customUserSubscriptionsModel{
		defaultUserSubscriptionsModel: newUserSubscriptionsModel(conn),
	}
}
