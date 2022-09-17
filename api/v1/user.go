package v1

import "github.com/gogf/gf/v2/frame/g"

type UserAddReq struct {
	g.Meta      `path:"/user/add" tags:"User" method:"put" summary:"add user"`
	LoginName   string `p:"loginName" v:"required" dc:"user's login name"`
	Name        string `p:"name" v:"required" dc:"user's name"`
	Password    string `p:"password" v:"password3#密码6位字符必须包含大小写字母、数字和特殊字符" dc:"login password, there is going to be encrypt md5"`
	Email       string `p:"email" v:"required|email" dc:"user's email address , all alert message will send to this address"`
	PhoneNumber string `p:"phoneNumber" v:"required|phone" dc:"telephone number"`
}
type UserAddRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

type UserUpdateReq struct {
	g.Meta      `path:"/user/update" tags:"User" method:"post" summary:"update user"`
	Id          int    `p:"id"  v:"required|integer" dc:"User record id"`
	LoginName   string `p:"loginName" v:"required" dc:"user's login name"`
	Name        string `p:"name" v:"required" dc:"user's name"`
	Email       string `p:"email" v:"required|email" dc:"user's email address , all alert message will send to this address"`
	PhoneNumber string `p:"phoneNumber" v:"required|phone" dc:"telephone number"`
}
type UserUpdateRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

type UserRemoveReq struct {
	g.Meta `path:"/user/delete/{id}" tags:"User" method:"delete" summary:"delete user"`
	Id     int `p:"id" v:"required|integer" dc:"user record id"`
}
type UserRemoveRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

type UserDetailReq struct {
	g.Meta `path:"/user/detail/{id}" tags:"User" method:"get" summary:"get single user detail"`
	Id     int `p:"id" v:"required|integer" dc:"user record id"`
}
type UserDetailRes struct {
	g.Meta `mime:"application/json"`
	Model  interface{} `json:"model"`
}

type UserPageReq struct {
	g.Meta      `path:"/user/page" tags:"User" method:"post" summary:"find user by page"`
	PageNo      int    `p:"page" v:"min:1" d:"1" dc:"page number"`
	PageSize    int    `p:"pageSize" v:"max:50" d:"10" dc:"page size"`
	LoginName   string `p:"loginName" dc:"user's login name"`
	Name        string `p:"name" dc:"user's name"`
	Email       string `p:"email" dc:"user's email address , all alert message will send to this address"`
	PhoneNumber string `p:"phoneNumber"  dc:"telephone number"`
}

type UserPageRes struct {
	g.Meta   `mime:"application/json"`
	Models   interface{} `json:"rows"`
	Total    int         `json:"total" dc:"number of this condition"`
	PageNo   int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
