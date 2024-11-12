package logic

import (
	"context"
	"database/sql"
	"time"

	"Anitale/pkg/e"
	"Anitale/pkg/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"Anitale/apps/user/rpc/internal/svc"
	"Anitale/apps/user/rpc/model"
	"Anitale/apps/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginUser 用户登录
func (l *LoginUserLogic) LoginUser(in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// 邮箱格式校验
	if !util.CheckEmail(in.Email) {
		return nil, errors.Wrapf(e.NewCustomCode(e.INVALID_EMAIL_FORMAT_ERROR), "email:%s", in.Email)
	}

	// 检查密码是否过于简单
	if !util.CheckPassword(in.Password) {
		return nil, errors.Wrapf(e.NewCustomCode(e.INVALID_PASSWORD_FORMAT_ERROR), "password:%s", in.Password)
	}

	// 验证邮箱是否已经注册
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if user == nil && err != nil {
		return nil, errors.Wrap(e.NewCustomCode(e.EMAIL_NOT_REGISTER_ERROR), "邮箱不存在，请先注册")
	}

	// 验证密码是否正确
	ok := util.ComparePassword(user.PasswordHash, in.Password)
	if !ok {
		return nil, errors.Wrap(e.NewCustomCode(e.PASSWORLD_ERROR), "密码输入错误")
	}

	// 生成token
	token, exp, err := l.generateJwtToken(user.Username, int64(user.Id))
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Conn.Transaction(func(tx *gorm.DB) error {
		// 用户再次登陆，应该先删除之前的token
		err = l.svcCtx.UserTokensModel.Delete(l.ctx, tx, user.Id)
		if err != nil {
			return errors.Wrapf(e.NewCustomError(e.DB_ERROR, e.GetMessage(e.DB_ERROR)), "delete user_tokens error: %v", err)
		}

		// 插入user_tokens
		err = l.svcCtx.UserTokensModel.Insert(l.ctx, tx, &model.UserTokens{
			UserId:      user.Id,
			AccessToken: token,
			ExpiresAt: sql.NullTime{
				Time:  time.Unix(exp, 0),
				Valid: true,
			},
		})
		if err != nil {
			return errors.Wrapf(e.NewCustomError(e.DB_ERROR, e.GetMessage(e.DB_ERROR)), "insert user_tokens error: %v", err)
		}

		return nil
	})

	return &pb.LoginUserResponse{
		Success:     true,
		Message:     "success",
		AccessToken: token,
		Expire:      exp,
	}, nil
}

// generateJwtToken 生成JwtToken
func (l *LoginUserLogic) generateJwtToken(sub string, userId int64) (string, int64, error) {
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret

	iat := time.Now().Unix()
	exp := iat + int64(l.svcCtx.Config.JwtAuth.AccessExpire)
	claims := jwt.MapClaims{
		"sub": sub,
		"jti": userId,
		"exp": exp,
		"iat": iat,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(accessSecret))
	return token, exp, err
}
