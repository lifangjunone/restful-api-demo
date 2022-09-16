package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps"
	"restful-api-demo/apps/host"
	"restful-api-demo/apps/host/impl"
)

// 通过写一个实体类，把内部的接口通过http协议暴露出去

func NewHostHandler() *Handler {
	return &Handler{}
}

type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	if apps.ServicesCenter[impl.HostSvc.Name()] == nil {
		panic("dependence host service required")
	}
	h.svc = impl.HostSvc
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
