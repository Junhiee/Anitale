package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
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
func NewUserProfilesModel(conn *gorm.DB, c cache.CacheConf) UserProfilesModel {
	return &customUserProfilesModel{
		defaultUserProfilesModel: newUserProfilesModel(conn, c),
	}
}

func (m *defaultUserProfilesModel) customCacheKeys(data *UserProfiles) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
