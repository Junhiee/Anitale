package logic

import (
	"context"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/model"
	"Anitale/apps/user/rpc/pb"
	"Anitale/pkg/errx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrEmailAlreadyRegisterError = errx.NewCustomError(300004, "email has been registered")

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterUserLogic) RegisterUser(in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	// 检查邮箱是否已注册
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "email:%s,err:%v", in.Email, err)

	}

	if user != nil {
		return nil, errors.Wrapf(ErrEmailAlreadyRegisterError, "Register user exists Email:%s,err:%v", in.Email, err)

	}

	// 注册
	l.svcCtx.Conn.Transaction(func(tx *gorm.DB) error {
		// 检查用户名是否已存在
		user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "username:%s,err:%v", in.Username, err)
		}
		if user != nil {
			return errors.Wrapf(errx.NewCustomCode(errx.USERNAME_ALREADY_EXISTS), "username:%s,err:%v", in.Username, err)
		}

		pwd := GenPasswordHash(in.Password)

		err = l.svcCtx.UserModel.Insert(l.ctx, tx, &model.Users{
			Username: in.Username,
			Email:    in.Email,
			PasswordHash: pwd,
			
		})
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "register:%s,err:%v", in.Email, err)
		}

		return nil
	})
	

	return &pb.RegisterUserResponse{}, nil
}

// 密码加密
func GenPasswordHash(passworld string) string {
	
	return passworld
}