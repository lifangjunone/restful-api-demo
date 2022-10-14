package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lifangjunone/restful-api-demo/apps/host"
	"github.com/lifangjunone/restful-api-demo/response"
	"net/http"
)

// 用于暴露host service 接口

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		c.JSON(http.StatusOK, response.Failed(err))
		return
	}
	data, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ParamMissing(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(data))
}
