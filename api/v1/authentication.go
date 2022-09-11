package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta    `path:"/login" tags:"Login" method:"post" summary:"user login api"`
	LoginName string `p:"loginName" v:"required" dc:"user's login name"`
	Password  string `p:"password" v:"password3#密码6位字符必须包含大小写字母、数字和特殊字符" dc:"password"`
}
type LoginRes struct {
	g.Meta `mime:"text/html" example:"success"`
	Token  string `p:"token" dc:"login token"`
}

type RegisterReq struct {
	g.Meta      `path:"/regist" tags:"Regist" method:"post"  summary:"regist account"`
	LoginName   string `p:"loginName" v:"required" dc:"user's login name"`
	Name        string `p:"name" v:"required" dc:"user's name"`
	Password    string `p:"password" v:"password3#密码6位字符必须包含大小写字母、数字和特殊字符" dc:"login password, there is going to be encrypt md5"`
	Repassword  string `p:"repassword" v:"same:password#密码必须相同" dc:"password check"`
	Email       string `p:"email" v:"required" dc:"user's email address , all alert message will send to this address"`
	PhoneNumber string `p:"phoneNumber" v:"required" dc:"telephone number"`
}

type RegisterRes struct {
	g.Meta `mime:"text/html" example:"success"`
}
