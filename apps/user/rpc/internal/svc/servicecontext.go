package svc

import (
	"Anitale/apps/user/rpc/internal/config"
	"Anitale/apps/user/rpc/model"
	"log"

	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config                 config.Config
	UserModel              model.UsersModel
	UserTokensModel        model.UserTokensModel
	UserProfilesModel      model.UserProfilesModel
	UserPreferencesModel   model.UserPreferencesModel
	UserSubscriptionsModel model.UserSubscriptionsModel
	Conn                   *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn, err := mysql.Connect(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:                 c,
		Conn:                   conn,
		UserModel:              model.NewUsersModel(conn),
		UserTokensModel:        model.NewUserTokensModel(conn),
		UserPreferencesModel:   model.NewUserPreferencesModel(conn),
		UserSubscriptionsModel: model.NewUserSubscriptionsModel(conn),
		UserProfilesModel:      model.NewUserProfilesModel(conn),
	}
}
