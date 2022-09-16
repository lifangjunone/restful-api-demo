package host

import (
	"context"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

// Service Host app service interface define
type Service interface {
	// CreateHost 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// QueryHost 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// DescribeHost 查询主机详情
	DescribeHost(context.Context, *QueryHostRequest) (*Host, error)
	// UpdateHost 修改主机信息
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// DeleteHost 删除主机
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

type QueryHostRequest struct{}

type UpdateHostRequest struct {
	*Describe
}

// DeleteHostRequest  删除请求参数
type DeleteHostRequest struct {
	Id string
}

// HostSet 主机查询列表
type HostSet struct {
	Items []*Host
	Total int
}
