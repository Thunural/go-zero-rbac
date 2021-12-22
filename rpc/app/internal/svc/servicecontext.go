/*
 * # 此文件版权属于ASO.DESIGN
 * # 文件：servicecontext.go  项目：aso-zero
 * # 作用：
 * # 当前修改时间：2021年07月15日 22:23:58
 * # 上次修改时间：2021年07月15日 21:23:05
 * # 作者：thunur
 * # 此文件不可非法传播、倒卖、共享，否则我们将追究相应的法律责任。
 * # 您如果已获得ASO.DESIGN授权可在原有基础上进行修改使用
 * # 如果您还没获得授权请联系我们 thunur@qq.com
 * # Copyright (c) 2021 aso.design
 */

package svc

import (
	"aso/rpc/app/internal/config"
	"aso/rpc/model/appmodel"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	ItemModel appmodel.AppItemModel
	TypeModel appmodel.AppTypeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config:    c,
		ItemModel: appmodel.NewAppItemModel(sqlConn, c.CacheRedis),
		TypeModel: appmodel.NewAppTypeModel(sqlConn, c.CacheRedis),
	}
}
