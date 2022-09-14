package impl

import (
	"database/sql"
	"go.uber.org/zap"
	"restful-api-demo/apps/host"
	"restful-api-demo/conf"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Named("host")
	return &HostServiceImpl{
		l:  sugar,
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  *zap.SugaredLogger
	db *sql.DB
}
