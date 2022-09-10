package middleware

import (
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	jwtUtil "github.com/golang-jwt/jwt/v4"
	"github.com/leapord/prometheusx/internal/consts"
	"github.com/leapord/prometheusx/internal/model/do"
)

func TokenMiddleware(r *ghttp.Request) {
	ctx := gctx.New()

	runtime, err := g.Cfg().Get(ctx, "profile.active")
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	if runtime.String() == "dev" {
		r.Middleware.Next()
		return
	}

	token := r.GetHeader("token")
	if g.IsEmpty(token) {
		result := g.NewVar("请先登录")
		r.Response.Writer.Write(result.Bytes())
		return
	}

	personClaims := jwtUtil.RegisteredClaims{}
	jwtUtil.ParseWithClaims(token, &personClaims, func(t *jwtUtil.Token) (interface{}, error) {
		return consts.JWT_SCRET, nil
	})

	if err := personClaims.Valid(); err != nil {
		r.Response.Writer.Write(g.NewVar("登录信息错误，请重新登录").Bytes())
		g.Log().Errorf(r.Context(), "token error %s", err)
		return
	}

	userString := personClaims.Subject
	u := do.User{}
	json.Unmarshal(g.NewVar(userString).Bytes(), &u)

	r.Header.Add("username", g.NewVar(u.Name).String())
	r.Header.Add("email", g.NewVar(u.Email).String())
	r.Header.Add("loginName", g.NewVar(u.LoginName).String())
	r.Header.Add("phoneNumber", g.NewVar(u.PhoneNumber).String())

	r.Middleware.Next()
}
