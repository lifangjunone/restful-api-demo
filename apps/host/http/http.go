package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps"
	"restful-api-demo/apps/host"
	"restful-api-demo/apps/host/impl"
	"restful-api-demo/conf"
	"restful-api-demo/version"
)

// 通过写一个实体类，把内部的接口通过http协议暴露出去

var handler = &Handler{}

type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	svc := apps.GetService(version.Name)
	if svc == nil {
		panic("dependence host service required")
	}
	h.svc = svc.(*impl.HostServiceImpl)
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}

func (h *Handler) Name() string {
	return version.Name
}

func (h *Handler) InitService() {
	g := gin.Default()
	h.Registry(g)
	g.Run(conf.C().App.HttpAddr())
}

func init() {
	apps.HttpRegistry(handler)
}
