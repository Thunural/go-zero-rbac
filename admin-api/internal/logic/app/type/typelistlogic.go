/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：typelistlogic.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月16日 00:55:19
 * # 上次修改时间：2021年07月15日 22:45:13
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

package logic

import (
	"aso/rpc/app/appclient"
	"context"

	"aso/admin-api/internal/svc"
	"aso/admin-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) TypeListLogic {
	return TypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeListLogic) TypeList() (*types.Resp, error) {
	resp, err := l.svcCtx.App.AppListType(l.ctx, &appclient.AppListTypeReq{})
	if err != nil {
		return nil, err
	}
	var list []*types.TypeListData
	for _, item := range resp.List {
		list = append(list, &types.TypeListData{
			Id:         item.Id,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
			Name:       item.Name,
		})
	}
	return &types.Resp{
		Code: 200,
		Data: list,
		Msg:  "success",
	}, nil
}
