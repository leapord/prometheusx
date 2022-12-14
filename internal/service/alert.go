// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	model "github.com/leapord/prometheusx/internal/model/do"
	vo "github.com/leapord/prometheusx/internal/model/vo"
)

type IAlert interface {
	AddAlert(ctx context.Context, alerts []model.Alert) (err error)
	Page(ctx context.Context, pageNo int, pageSize int, alert model.Alert) (models []vo.Alert, total int, err error)
}

var localAlert IAlert

func Alert() IAlert {
	if localAlert == nil {
		panic("implement not found for interface IAlert, forgot register?")
	}
	return localAlert
}

func RegisterAlert(i IAlert) {
	localAlert = i
}
