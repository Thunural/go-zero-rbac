/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：roleaddlogic.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月11日 15:48:15
 * # 上次修改时间：2021年07月11日 13:18:16
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

package logic

import (
	"aso/admin-api/internal/svc"
	"aso/admin-api/internal/types"
	"aso/rpc/sys/sysclient"
	"aso/utils/common/errorx"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleAddLogic {
	return RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req types.RoleAddReq) (*types.Resp, error) {
	_, err := l.svcCtx.Sys.RoleAdd(l.ctx, &sysclient.RoleAddRequest{
		Name:             req.Name,
		Label:            req.Label,
		Remark:           req.Remark,
		Relevance:        req.Relevance,
		MenuIdList:       req.MenuIdList,
		DepartmentIdList: req.DepartmentIdList,
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.Resp{
		Code: 200,
		Data: nil,
		Msg:  "success",
	}, nil
}
