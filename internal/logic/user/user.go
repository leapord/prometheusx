package logic

// 用户业务层

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"
	"github.com/leapord/prometheusx/internal/consts"
	model "github.com/leapord/prometheusx/internal/model/do"
	"github.com/leapord/prometheusx/internal/service"
)

type sUser struct{}

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

// 登陆 并生成Token
func (u *sUser) Login(ctx context.Context, loginName *string, password *string) (token string, userJson string, err error) {
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
	user.Password = nil
	userJson = gjson.MustEncodeString(user)
	return
}

// 注册
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

// 检查用户是否存在
func (s *sUser) CheckUser(ctx context.Context, loginName string) bool {
	if cnt, err := g.Model(model.User{}).Where(model.User{
		LoginName: loginName,
	}).Count(); err == nil && cnt == 1 {
		return true
	} else {
		g.Log().Error(ctx, err)
	}
	return false
}

// 添加
func (s *sUser) Add(ctx context.Context, user model.User) (id int64, err error) {
	id, err = g.Model(model.User{}).InsertAndGetId(user)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 修改
func (s *sUser) Update(ctx context.Context, user model.User) (err error) {
	_, err = g.Model(model.User{}).Where(model.User{Id: user.Id}).UpdateAndGetAffected(user)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 删除
func (s *sUser) Delete(ctx context.Context, id int) (user model.User, err error) {
	gmodel := g.Model(model.User{}).Where(model.User{Id: id})
	cnt, err := gmodel.Count()
	if err != nil || cnt != 1 {
		err = gerror.New("can not find the user")
		g.Log().Info(ctx, "user not find id "+g.NewVar(id).String())
		return
	}
	err = gmodel.Scan(&user)
	if err == nil {
		_, err = gmodel.Delete()
	} else {
		g.Log().Error(ctx, err)
	}
	return
}

// 单个详情
func (s *sUser) Detail(ctx context.Context, id int) (user model.User, err error) {
	err = g.Model(model.User{}).Where(model.User{Id: id}).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// 分页
func (s *sUser) Page(ctx context.Context, pageNo int, pageSize int, user model.User) (total int, users []model.User, err error) {
	gmodel := g.Model(model.User{})

	if !g.NewVar(user.Name).IsEmpty() {
		gmodel.WhereLike("name", "%"+g.NewVar(user.Name).String()+"%")
	}
	if !g.NewVar(user.LoginName).IsEmpty() {
		gmodel.WhereLike("login_name", "%"+g.NewVar(user.LoginName).String()+"%")
	}
	if !g.NewVar(user.Email).IsEmpty() {
		gmodel.WhereLike("email", "%"+g.NewVar(user.Email).String()+"%")
	}
	if !g.NewVar(user.PhoneNumber).IsEmpty() {
		gmodel.WhereLike("phone_number", "%"+g.NewVar(user.PhoneNumber).String()+"%")
	}

	total, err = gmodel.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	err = gmodel.Scan(&users)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}
