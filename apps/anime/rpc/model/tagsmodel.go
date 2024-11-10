package model

import (
	"gorm.io/gorm"
)

var _ TagsModel = (*customTagsModel)(nil)

type (
	// TagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagsModel.
	TagsModel interface {
		tagsModel
		customTagsLogicModel
	}

	customTagsModel struct {
		*defaultTagsModel
	}

	customTagsLogicModel interface {
	}
)

// NewTagsModel returns a model for the database table.
func NewTagsModel(conn *gorm.DB) TagsModel {
	return &customTagsModel{
		defaultTagsModel: newTagsModel(conn),
	}
}
