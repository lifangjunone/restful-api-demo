package impl

import (
	"log"
	"restful-api-demo/apps/host"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		l: log.Logger{},
	}
}

type HostServiceImpl struct {
	l log.Logger
}
