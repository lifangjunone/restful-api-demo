package impl

import (
	"context"
	"restful-api-demo/apps/host"
)

// 业务逻辑层(Controller层)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Info("start crate host")
	// 检验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	ins.InjectDefault()
	// 由dao层负责把数据对象入库
	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
