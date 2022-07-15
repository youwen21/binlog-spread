package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"binlog_spread/app/service"
	"github.com/gin-gonic/gin"
)

type binlogHandler struct {
	handler.BaseHandler
}

var (
	BinlogHandler = new(binlogHandler)
)

func (h binlogHandler) Query(c *gin.Context) {
	form := dto.ApiBinlogForm{}
	form.OrderBy = append(form.OrderBy, "api_binlog_id desc")

	ret, err := service.ApiBinlogSrv.Query(&form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}

func (h binlogHandler) Delete(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")
	if err := dal.ApiBinlogDAO.RealDelete(id); err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, "")
}

func (h binlogHandler) Diff(c *gin.Context) {
	searchParams := dto.BinlogDiffSearch{}
	if err := c.ShouldBind(&searchParams); err != nil {
		h.JsonErr(c, err)
		return
	}

	result, err := service.ApiBinlogSrv.Diff(c, searchParams)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, result)
}
