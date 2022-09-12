package v1

import "github.com/gogf/gf/v2/frame/g"

type NodeTargetReq struct {
	g.Meta `path:"/node/target" tags:"target" method:"post" dc:"prometheus http_sd target"`
}

type NodeTargetRes struct {
	g.Meta `mime:"application/json"`
}
