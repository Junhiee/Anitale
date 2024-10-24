package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserRoleAssignmentsModel = (*customUserRoleAssignmentsModel)(nil)

type (
	// UserRoleAssignmentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleAssignmentsModel.
	UserRoleAssignmentsModel interface {
		userRoleAssignmentsModel
		customUserRoleAssignmentsLogicModel
	}

	customUserRoleAssignmentsModel struct {
		*defaultUserRoleAssignmentsModel
	}

	customUserRoleAssignmentsLogicModel interface {
	}
)

// NewUserRoleAssignmentsModel returns a model for the database table.
func NewUserRoleAssignmentsModel(conn *gorm.DB, c cache.CacheConf) UserRoleAssignmentsModel {
	return &customUserRoleAssignmentsModel{
		defaultUserRoleAssignmentsModel: newUserRoleAssignmentsModel(conn, c),
	}
}

func (m *defaultUserRoleAssignmentsModel) customCacheKeys(data *UserRoleAssignments) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
