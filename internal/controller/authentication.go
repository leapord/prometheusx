package controller

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

var (
	Authentication = cAuthentication{}
)

type cAuthentication struct{}

func (a *cAuthentication) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	pwd := req.Password
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	token, user, err := service.User().Login(ctx, &req.LoginName, &pwd)

	res = &v1.LoginRes{Token: token, UserInfo: user}
	return
}

func (a *cAuthentication) RegisterUser(ctx context.Context, req *v1.RegisterReq) (res v1.RegisterRes, err error) {
	if req.Password != req.Repassword {
		err = errors.New("password must be the same")
		return
	}
	user := &model.User{
		Name:        req.Name,
		LoginName:   req.LoginName,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		CreateTime:  gtime.Now(),
	}
	err = service.User().Regist(ctx, user)

	return
}
