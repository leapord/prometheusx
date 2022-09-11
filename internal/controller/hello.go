package controller

import (
	"context"

	"github.com/gogf/gf/v2/container/gmap"
	v1 "github.com/leapord/prometheusx/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	hashMap := gmap.New(true)
	hashMap.Set("version", "v1.0")
	hashMap.Set("author", "leapord")
	res = &v1.HelloRes{
		Model: hashMap,
	}
	return
}
