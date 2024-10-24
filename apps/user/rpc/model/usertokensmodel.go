package model

import (
	"gorm.io/gorm"
)

var _ UserTokensModel = (*customUserTokensModel)(nil)

type (
	// UserTokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTokensModel.
	UserTokensModel interface {
		userTokensModel
		customUserTokensLogicModel
	}

	customUserTokensModel struct {
		*defaultUserTokensModel
	}

	customUserTokensLogicModel interface {
	}
)

// NewUserTokensModel returns a model for the database table.
func NewUserTokensModel(conn *gorm.DB) UserTokensModel {
	return &customUserTokensModel{
		defaultUserTokensModel: newUserTokensModel(conn),
	}
}
