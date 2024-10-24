package model

import (
	"gorm.io/gorm"
)

var _ UserRolesModel = (*customUserRolesModel)(nil)

type (
	// UserRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRolesModel.
	UserRolesModel interface {
		userRolesModel
		customUserRolesLogicModel
	}

	customUserRolesModel struct {
		*defaultUserRolesModel
	}

	customUserRolesLogicModel interface {
	}
)

// NewUserRolesModel returns a model for the database table.
func NewUserRolesModel(conn *gorm.DB) UserRolesModel {
	return &customUserRolesModel{
		defaultUserRolesModel: newUserRolesModel(conn),
	}
}
