/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：keyworddeletelogic.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月18日 14:59:48
 * # 上次修改时间：2021年07月18日 14:56:06
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

package logic

import (
	"context"

	"aso/rpc/wechat/internal/svc"
	"aso/rpc/wechat/wechat"

	"github.com/tal-tech/go-zero/core/logx"
)

type KeyWordDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKeyWordDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeyWordDeleteLogic {
	return &KeyWordDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// KeyWordDelete 关键词删除
func (l *KeyWordDeleteLogic) KeyWordDelete(in *wechat.KeyWordReplyDeleteReq) (*wechat.KeyWordReplyDeleteResp, error) {
	err := l.svcCtx.KeyReplyModel.Deletes(in.Ids)
	if err != nil {
		return nil, err
	}

	return &wechat.KeyWordReplyDeleteResp{}, nil
}
