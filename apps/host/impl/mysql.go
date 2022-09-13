package impl

import (
	"go.uber.org/zap"
	"restful-api-demo/apps/host"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return &HostServiceImpl{
		l: sugar,
	}
}

type HostServiceImpl struct {
	l *zap.SugaredLogger
}
