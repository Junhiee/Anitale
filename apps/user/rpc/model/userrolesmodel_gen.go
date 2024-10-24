// Code generated by goctl. DO NOT EDIT!

package model

import (
	"database/sql"
	"context"
	"errors"
	"fmt"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/SpectatorNan/gorm-zero/gormc/batchx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheUserRolesRoleIdPrefix   = "cache:userRoles:roleId:"
	cacheUserRolesRoleNamePrefix = "cache:userRoles:roleName:"
)

type (
	userRolesModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *UserRoles) error
		BatchInsert(ctx context.Context, tx *gorm.DB, news []UserRoles) error
		FindOne(ctx context.Context, roleId uint64) (*UserRoles, error)
		FindOneByRoleName(ctx context.Context, roleName string) (*UserRoles, error)
		Update(ctx context.Context, tx *gorm.DB, data *UserRoles) error
		BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []UserRoles) error
		BatchDelete(ctx context.Context, tx *gorm.DB, datas []UserRoles) error

		Delete(ctx context.Context, tx *gorm.DB, roleId uint64) error
		// deprecated. recommend add a transaction in service context instead of using this
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultUserRolesModel struct {
		gormc.CachedConn
		table string
	}

	UserRoles struct {
		RoleId   uint64         `gorm:"column:role_id;primary_key"` // 角色ID
		RoleName string         `gorm:"column:role_name"`           // 权限的名称，admin|user|guest
		RoleDesc sql.NullString `gorm:"column:role_desc"`           // 角色的描述信息
	}
)

func (UserRoles) TableName() string {
	return "`user_roles`"
}

func newUserRolesModel(conn *gorm.DB, c cache.CacheConf) *defaultUserRolesModel {
	return &defaultUserRolesModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`user_roles`",
	}
}

func (m *defaultUserRolesModel) GetCacheKeys(data *UserRoles) []string {
	if data == nil {
		return []string{}
	}
	userRolesRoleIdKey := fmt.Sprintf("%s%v", cacheUserRolesRoleIdPrefix, data.RoleId)
	userRolesRoleNameKey := fmt.Sprintf("%s%v", cacheUserRolesRoleNamePrefix, data.RoleName)
	cacheKeys := []string{
		userRolesRoleIdKey, userRolesRoleNameKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultUserRolesModel) Insert(ctx context.Context, tx *gorm.DB, data *UserRoles) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.GetCacheKeys(data)...)
	return err
}
func (m *defaultUserRolesModel) BatchInsert(ctx context.Context, tx *gorm.DB, news []UserRoles) error {

	err := batchx.BatchExecCtx(ctx, m, news, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Create(&news).Error
	})

	return err
}

func (m *defaultUserRolesModel) FindOne(ctx context.Context, roleId uint64) (*UserRoles, error) {
	userRolesRoleIdKey := fmt.Sprintf("%s%v", cacheUserRolesRoleIdPrefix, roleId)
	var resp UserRoles
	err := m.QueryCtx(ctx, &resp, userRolesRoleIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserRoles{}).Where("`role_id` = ?", roleId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRolesModel) FindOneByRoleName(ctx context.Context, roleName string) (*UserRoles, error) {
	userRolesRoleNameKey := fmt.Sprintf("%s%v", cacheUserRolesRoleNamePrefix, roleName)
	var resp UserRoles
	err := m.QueryRowIndexCtx(ctx, &resp, userRolesRoleNameKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&UserRoles{}).Where("`role_name` = ?", roleName).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.RoleId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRolesModel) Update(ctx context.Context, tx *gorm.DB, data *UserRoles) error {
	old, err := m.FindOne(ctx, data.RoleId)
	if err != nil && errors.Is(err, ErrNotFound) {
		return err
	}
	clearKeys := append(m.GetCacheKeys(old), m.GetCacheKeys(data)...)
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, clearKeys...)
	return err
}
func (m *defaultUserRolesModel) BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []UserRoles) error {
	clearData := make([]UserRoles, 0, len(olds)+len(news))
	clearData = append(clearData, olds...)
	clearData = append(clearData, news...)
	err := batchx.BatchExecCtx(ctx, m, clearData, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&news).Error
	})

	return err
}

func (m *defaultUserRolesModel) Delete(ctx context.Context, tx *gorm.DB, roleId uint64) error {
	data, err := m.FindOne(ctx, roleId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&UserRoles{}, roleId).Error
	}, m.GetCacheKeys(data)...)
	return err
}

func (m *defaultUserRolesModel) BatchDelete(ctx context.Context, tx *gorm.DB, datas []UserRoles) error {
	err := batchx.BatchExecCtx(ctx, m, datas, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&datas).Error
	})

	return err
}

// deprecated. recommend add a transaction in service context instead of using this
func (m *defaultUserRolesModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultUserRolesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserRolesRoleIdPrefix, primary)
}

func (m *defaultUserRolesModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&UserRoles{}).Where("`role_id` = ?", primary).Take(v).Error
}
