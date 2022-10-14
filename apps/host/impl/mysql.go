package impl

import (
	"database/sql"
	"github.com/lifangjunone/restful-api-demo/apps"
	"github.com/lifangjunone/restful-api-demo/conf"
	"github.com/lifangjunone/restful-api-demo/version"
	"go.uber.org/zap"
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
