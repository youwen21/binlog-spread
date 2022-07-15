package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"github.com/gin-gonic/gin"
)

type stateAbnormalHandler struct {
	handler.BaseHandler
}

var (
	StateAbnormalHandler = new(stateAbnormalHandler)
)


func (h stateAbnormalHandler) Query(c *gin.Context) {
	form := &dto.StateAbnormalForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		h.JsonErr(c, err)
		return
	}
	form.OrderBy = append(form.OrderBy, "state_abnormal_id desc")

	ret, err := dal.StateAbnormalDAO.Query(form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}

func (h stateAbnormalHandler) Delete(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")
	if err := dal.StateAbnormalDAO.RealDelete(id); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, "")
}

