package impl

import (
	"context"
	"restful-api-demo/apps/host"
	"testing"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestHostCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func init() {
	// 接口的具体实现
	service = NewHostServiceImpl()
}
