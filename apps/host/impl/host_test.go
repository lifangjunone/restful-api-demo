package impl

import (
	"context"
	"fmt"
	"github.com/lifangjunone/restful-api-demo/apps/host"
	"github.com/lifangjunone/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestHostCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Name = "test"
	ins.Id = "ins-061"
	ins.Region = "beijing"
	ins.Type = "11"
	ins.CPU = 1
	ins.Memory = 24
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init() {
	// 测试用例环境变量
	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
	if err != nil {
		panic(err)
	}
	// 接口的具体实现
	service = NewHostServiceImpl()
}
