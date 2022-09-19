package impl

import (
	"database/sql"
	"go.uber.org/zap"
	"restful-api-demo/apps"
	"restful-api-demo/conf"
	"restful-api-demo/version"
)

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

func (i *HostServiceImpl) InitService() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Named("host")
	i.l = sugar
	i.db = conf.C().MySQL.GetDB()
}

func (i *HostServiceImpl) Name() string {
	return version.Name
}

type HostServiceImpl struct {
	l    *zap.SugaredLogger
	db   *sql.DB
	name string
}

func init() {
	HostSvc := &HostServiceImpl{}
	apps.Registry(HostSvc)
}
