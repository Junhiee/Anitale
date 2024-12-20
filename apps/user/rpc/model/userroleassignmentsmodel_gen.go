// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc"

	"time"

	"gorm.io/gorm"
)

type (
	userRoleAssignmentsModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *UserRoleAssignments) error
		BatchInsert(ctx context.Context, tx *gorm.DB, news []UserRoleAssignments) error
		FindOne(ctx context.Context, assignmentId uint64) (*UserRoleAssignments, error)
		FindOneByUserIdRoleId(ctx context.Context, userId sql.NullInt64, roleId sql.NullInt64) (*UserRoleAssignments, error)
		Update(ctx context.Context, tx *gorm.DB, data *UserRoleAssignments) error
		BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []UserRoleAssignments) error
		BatchDelete(ctx context.Context, tx *gorm.DB, datas []UserRoleAssignments) error

		Delete(ctx context.Context, tx *gorm.DB, assignmentId uint64) error
		// deprecated. recommend add a transaction in service context instead of using this
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultUserRoleAssignmentsModel struct {
		conn  *gorm.DB
		table string
	}

	UserRoleAssignments struct {
		AssignmentId uint64        `gorm:"column:assignment_id;primary_key"` // 角色分配关系唯一标识
		UserId       sql.NullInt64 `gorm:"column:user_id"`                   // 用户唯一标识，与users表的id关联
		RoleId       sql.NullInt64 `gorm:"column:role_id"`                   // 角色唯一标识，与user_roles表的id关联
		AssignedAt   time.Time     `gorm:"column:assigned_at"`               // 角色分配时间
	}
)

func (UserRoleAssignments) TableName() string {
	return "`user_role_assignments`"
}

func newUserRoleAssignmentsModel(conn *gorm.DB) *defaultUserRoleAssignmentsModel {
	return &defaultUserRoleAssignmentsModel{
		conn:  conn,
		table: "`user_role_assignments`",
	}
}

func (m *defaultUserRoleAssignmentsModel) Insert(ctx context.Context, tx *gorm.DB, data *UserRoleAssignments) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(&data).Error
	return err
}
func (m *defaultUserRoleAssignmentsModel) BatchInsert(ctx context.Context, tx *gorm.DB, news []UserRoleAssignments) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Create(&news).Error

	return err
}

func (m *defaultUserRoleAssignmentsModel) FindOne(ctx context.Context, assignmentId uint64) (*UserRoleAssignments, error) {
	var resp UserRoleAssignments
	err := m.conn.WithContext(ctx).Model(&UserRoleAssignments{}).Where("`assignment_id` = ?", assignmentId).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRoleAssignmentsModel) FindOneByUserIdRoleId(ctx context.Context, userId sql.NullInt64, roleId sql.NullInt64) (*UserRoleAssignments, error) {
	var resp UserRoleAssignments
	err := m.conn.WithContext(ctx).Model(&UserRoleAssignments{}).Where("`user_id` = ? and `role_id` = ?", userId, roleId).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRoleAssignmentsModel) Update(ctx context.Context, tx *gorm.DB, data *UserRoleAssignments) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(data).Error
	return err
}
func (m *defaultUserRoleAssignmentsModel) BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []UserRoleAssignments) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(&news).Error

	return err
}

func (m *defaultUserRoleAssignmentsModel) Delete(ctx context.Context, tx *gorm.DB, assignmentId uint64) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Delete(&UserRoleAssignments{}, assignmentId).Error

	return err
}

func (m *defaultUserRoleAssignmentsModel) BatchDelete(ctx context.Context, tx *gorm.DB, datas []UserRoleAssignments) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Delete(&datas).Error

	return err
}

// deprecated. recommend add a transaction in service context instead of using this
func (m *defaultUserRoleAssignmentsModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.conn.WithContext(ctx).Transaction(fn)
}
