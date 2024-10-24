package model

import (
	"gorm.io/gorm"
)

var _ UserProfilesModel = (*customUserProfilesModel)(nil)

type (
	// UserProfilesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserProfilesModel.
	UserProfilesModel interface {
		userProfilesModel
		customUserProfilesLogicModel
	}

	customUserProfilesModel struct {
		*defaultUserProfilesModel
	}

	customUserProfilesLogicModel interface {
	}
)

// NewUserProfilesModel returns a model for the database table.
func NewUserProfilesModel(conn *gorm.DB) UserProfilesModel {
	return &customUserProfilesModel{
		defaultUserProfilesModel: newUserProfilesModel(conn),
	}
}
