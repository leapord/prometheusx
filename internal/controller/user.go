package controller

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "github.com/leapord/prometheusx/api/v1"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type cUser struct{}

var (
	User = cUser{}
)

// 添加
func (c *cUser) AddUser(ctx context.Context, req *v1.UserAddReq) (res *v1.UserAddRes, err error) {
	pwd, _ := gmd5.Encrypt(req.Password)
	user := model.User{
		Name:        req.Name,
		LoginName:   req.LoginName,
		Password:    pwd,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		CreateTime:  gtime.Now(),
	}
	id, err := service.User().Add(ctx, user)
	user.Id = id
	if err == nil {
		res = &v1.UserAddRes{
			Model: user,
		}
	}
	return
}

// 更新
func (c *cUser) UpdateUser(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	user := model.User{
		Id:          req.Id,
		Name:        req.Name,
		LoginName:   req.LoginName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	err = service.User().Update(ctx, user)
	if err == nil {
		res = &v1.UserUpdateRes{
			Model: user,
		}
	}
	return
}

// 删除
func (c *cUser) RemoveUser(ctx context.Context, req *v1.UserRemoveReq) (res *v1.UserRemoveRes, err error) {
	user, err := service.User().Delete(ctx, req.Id)
	if err == nil {
		res = &v1.UserRemoveRes{
			Model: user,
		}
	}
	return
}

// 单个详情
func (c *cUser) DetailUser(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {
	user, err := service.User().Detail(ctx, req.Id)
	if err == nil {
		res = &v1.UserDetailRes{
			Model: user,
		}
	}
	return
}

// 分页
func (c *cUser) PageUser(ctx context.Context, req *v1.UserPageReq) (res *v1.UserPageRes, err error) {
	user := model.User{
		Name:        req.Name,
		LoginName:   req.LoginName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	total, models, err := service.User().Page(ctx, req.PageNo, req.PageSize, user)
	if err == nil {
		res = &v1.UserPageRes{
			Total:    total,
			Models:   models,
			PageNo:   req.PageNo,
			PageSize: req.PageSize,
		}
	}
	return
}
