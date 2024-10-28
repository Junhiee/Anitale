package model

import (
	"context"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"gorm.io/gorm"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		customUsersLogicModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}

	customUsersLogicModel interface {
		FindOneByEmailCtx(ctx context.Context, tx *gorm.DB, email string) (*Users, error)
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn *gorm.DB) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *customUsersModel) FindOneByEmailCtx(ctx context.Context, tx *gorm.DB, email string) (*Users, error) {
	var resp Users
	err := tx.WithContext(ctx).Model(&Users{}).Where("email = ?", email).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}
