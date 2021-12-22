/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：routes.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月18日 01:05:05
 * # 上次修改时间：2021年07月16日 15:36:36
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	wechatinterface "aso/wechat-api/internal/handler/wechat/interface"
	"aso/wechat-api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/wechat/index",
				Handler: wechatinterface.WechatInterfaceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/wechat/index",
				Handler: wechatinterface.WechatHandler(serverCtx),
			},
		},
	)
}
