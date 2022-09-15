package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/host"
)

// 通过写一个实体类，把内部的接口通过http协议暴露出去

func NewHostHandler(svc host.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

type Handler struct {
	svc host.Service
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
