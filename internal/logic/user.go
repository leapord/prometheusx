package logic

// 用户业务层

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"
	"github.com/leapord/prometheusx/internal/consts"
	model "github.com/leapord/prometheusx/internal/model/do"
)

type sUser struct{}

var (
	User = sUser{}
)

func (u *sUser) Login(ctx context.Context, loginName *string, password *string) (token string, err error) {
	user := model.User{}
	errUser := g.Model(model.User{}).Where(model.User{LoginName: loginName, Password: password}).Scan(&user)

	if errUser != nil {
		g.Log().Error(ctx, err)
		err = errUser
		return
	}

	nowTime := time.Now()
	expireTime := jwt.NewNumericDate(nowTime.Add(72 * time.Hour))
	issuer := "leapord"
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: expireTime,
		Subject:   g.NewVar(user).String(),
		IssuedAt:  jwt.NewNumericDate(nowTime),
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(g.NewVar(consts.JWT_SCRET).Bytes())
	return
}

func (u *sUser) Regist(ctx context.Context, user *model.User) (err error) {
	pwd, err := gmd5.Encrypt(user.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	user.Password = pwd
	_, err = g.Model(model.User{}).Insert(user)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}
