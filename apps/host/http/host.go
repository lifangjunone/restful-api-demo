package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/host"
)

// 用于暴露host service 接口

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	h.svc.CreateHost(c.Request.Context(), ins)
}
