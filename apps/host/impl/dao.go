package impl

import (
	"context"
	"restful-api-demo/apps/host"
)

// 把Host保存数据库
func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	// 检验数据合法性
	if err := ins.Validate(); err != nil {
		return err
	}
	return nil
}
