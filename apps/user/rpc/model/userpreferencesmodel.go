package model

import (
	"gorm.io/gorm"
)

var _ UserPreferencesModel = (*customUserPreferencesModel)(nil)

type (
	// UserPreferencesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPreferencesModel.
	UserPreferencesModel interface {
		userPreferencesModel
		customUserPreferencesLogicModel
	}

	customUserPreferencesModel struct {
		*defaultUserPreferencesModel
	}

	customUserPreferencesLogicModel interface {
	}
)

// NewUserPreferencesModel returns a model for the database table.
func NewUserPreferencesModel(conn *gorm.DB) UserPreferencesModel {
	return &customUserPreferencesModel{
		defaultUserPreferencesModel: newUserPreferencesModel(conn),
	}
}
