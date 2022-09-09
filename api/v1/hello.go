package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/version" tags:"Hello" method:"get" summary:"the version of this program"`
}
type HelloRes struct {
	g.Meta `mime:"application/json"`
}
