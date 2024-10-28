package logic

import (
	"context"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/model"
	"Anitale/apps/user/rpc/pb"
	"Anitale/pkg/errx"
	"Anitale/pkg/util"

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
	var resp = &pb.RegisterUserResponse{}
	// 邮箱格式校验
	if !util.CheckEmail(in.Email) {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.INVALID_EMAIL_FORMAT_ERROR), "email:%s", in.Email)
	}

	// 用户名校验
	if !util.CheckUsername(in.Username) {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.INVALID_USERNAME_FORMAT_ERROR), "username:%s", in.Username)
	}

	// 检查密码是否过于简单
	if !util.CheckPassword(in.Password) {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.INVALID_PASSWORD_FORMAT_ERROR), "password:%s", in.Password)
	}

	// 检查邮箱是否已注册
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "FindOneByEmail:%s,err:%v", in.Email, err)
	}
	if user != nil {
		return nil, errors.Wrapf(ErrEmailAlreadyRegisterError, "Register email exists:%s,err:%v", in.Email, err)

	}

	// 检查用户名是否已存在
	user, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "FindOneByUsername:%s,err:%v", in.Username, err)
	}
	if user != nil {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.USERNAME_ALREADY_EXISTS), "Register usename exists:%s,err:%v", in.Username, err)
	}

	// 注册
	err = l.svcCtx.Conn.Transaction(func(tx *gorm.DB) error {
		// 密码加密
		pwd := util.GenPasswordHash(in.Password)

		// 注册登陆表 user
		err = l.svcCtx.UserModel.Insert(l.ctx, tx, &model.Users{
			Username:     in.Username,
			Email:        in.Email,
			PasswordHash: pwd,
		})
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "register:%s,err:%v", in.Email, err)
		}

		// 查询到插入的数据后， 插入个人信息表 user_profiles
		u, err := l.svcCtx.UserModel.FindOneByEmailCtx(l.ctx, tx, in.Email)
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "FindOneByEmail:%s,err:%v", in.Email, err)
		}
		// 插入个人信息表 user_profiles
		err = l.svcCtx.UserProfiles.Insert(l.ctx, tx, &model.UserProfiles{
			UserId: u.Id,
		})
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "UserProfiles:%s,err:%v", in.Email, err)
		}

		resp.Message = "register success"
		resp.Success = true
		resp.User = &pb.User{
			Id:       u.Id,
			Username: in.Username,
			Email:    in.Email,
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomCode(errx.DB_ERROR), "Transaction:%s,err:%v", in.Email, err)
	}

	return &pb.RegisterUserResponse{Success: true, Message: "register success", User: &pb.User{
		Username: in.Username,
		Email:    in.Email,
	}}, nil
}
