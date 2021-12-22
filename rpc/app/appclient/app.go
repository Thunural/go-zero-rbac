/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：app.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月16日 00:54:39
 * # 上次修改时间：2021年07月15日 23:04:39
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

// Code generated by goctl. DO NOT EDIT!
// Source: app.proto

//go:generate mockgen -destination ./app_mock.go -package appclient -source $GOFILE

package appclient

import (
	"context"

	"aso/rpc/app/app"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	AppListTypeData   = app.AppListTypeData
	AppAddTypeReq     = app.AppAddTypeReq
	AppAddTypeResp    = app.AppAddTypeResp
	AppAddItemResp    = app.AppAddItemResp
	AppListTypeResp   = app.AppListTypeResp
	AppUpdateTypeReq  = app.AppUpdateTypeReq
	AppDeleteTypeReq  = app.AppDeleteTypeReq
	AppPageItemData   = app.AppPageItemData
	AppPageItemResp   = app.AppPageItemResp
	AppAddItemReq     = app.AppAddItemReq
	AppListTypeReq    = app.AppListTypeReq
	AppUpdateTypeResp = app.AppUpdateTypeResp
	AppDeleteTypeResp = app.AppDeleteTypeResp
	AppPageItemReq    = app.AppPageItemReq

	App interface {
		//  相册类型
		AppListType(ctx context.Context, in *AppListTypeReq) (*AppListTypeResp, error)
		//  相册类型添加
		AppAddType(ctx context.Context, in *AppAddTypeReq) (*AppAddTypeResp, error)
		//  相册类型更新
		AppUpdateType(ctx context.Context, in *AppUpdateTypeReq) (*AppUpdateTypeResp, error)
		//  相册类型删除
		AppDeleteType(ctx context.Context, in *AppDeleteTypeReq) (*AppDeleteTypeResp, error)
		//  相册列表
		AppPageItem(ctx context.Context, in *AppPageItemReq) (*AppPageItemResp, error)
		//  相册列表插入
		AppAddItem(ctx context.Context, in *AppAddItemReq) (*AppAddItemResp, error)
	}

	defaultApp struct {
		cli zrpc.Client
	}
)

func NewApp(cli zrpc.Client) App {
	return &defaultApp{
		cli: cli,
	}
}

//  相册类型
func (m *defaultApp) AppListType(ctx context.Context, in *AppListTypeReq) (*AppListTypeResp, error) {
	client := app.NewAppClient(m.cli.Conn())
	return client.AppListType(ctx, in)
}

//  相册类型添加
func (m *defaultApp) AppAddType(ctx context.Context, in *AppAddTypeReq) (*AppAddTypeResp, error) {
	client := app.NewAppClient(m.cli.Conn())
	return client.AppAddType(ctx, in)
}

//  相册类型更新
func (m *defaultApp) AppUpdateType(ctx context.Context, in *AppUpdateTypeReq) (*AppUpdateTypeResp, error) {
	client := app.NewAppClient(m.cli.Conn())
	return client.AppUpdateType(ctx, in)
}

//  相册类型删除
func (m *defaultApp) AppDeleteType(ctx context.Context, in *AppDeleteTypeReq) (*AppDeleteTypeResp, error) {
	client := app.NewAppClient(m.cli.Conn())
	return client.AppDeleteType(ctx, in)
}

//  相册列表
func (m *defaultApp) AppPageItem(ctx context.Context, in *AppPageItemReq) (*AppPageItemResp, error) {
	client := app.NewAppClient(m.cli.Conn())
	return client.AppPageItem(ctx, in)
}

//  相册列表插入
func (m *defaultApp) AppAddItem(ctx context.Context, in *AppAddItemReq) (*AppAddItemResp, error) {
	client := app.NewAppClient(m.cli.Conn())
	return client.AppAddItem(ctx, in)
}
