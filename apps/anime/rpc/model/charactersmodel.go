package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CharactersModel = (*customCharactersModel)(nil)

type (
	// CharactersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCharactersModel.
	CharactersModel interface {
		charactersModel
		customCharactersLogicModel
	}

	customCharactersModel struct {
		*defaultCharactersModel
	}

	customCharactersLogicModel interface {
	}
)

// NewCharactersModel returns a model for the database table.
func NewCharactersModel(conn *gorm.DB, c cache.CacheConf) CharactersModel {
	return &customCharactersModel{
		defaultCharactersModel: newCharactersModel(conn, c),
	}
}

func (m *defaultCharactersModel) customCacheKeys(data *Characters) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
