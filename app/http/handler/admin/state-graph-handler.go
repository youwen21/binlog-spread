package admin

import (
	"binlog_spread/app/http/handler"
	"binlog_spread/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type stateGraphHandler struct {
	handler.BaseHandler
}

var (
	StateGraphHandler = new(stateGraphHandler)
)

func (h stateGraphHandler) GetInfo(c *gin.Context) {
	stateClassId, _ := h.ParamInt(c, "id")

	ret, err := service.StateGraphSrv.Graph(stateClassId)
	if err != nil {
		h.JsonErr(c, err)
	}

	svg, err := service.StateGraphSrv.Svg(ret)
	if err != nil {
		h.JsonErr(c, err)
	}

	c.String(http.StatusOK, svg)
}
