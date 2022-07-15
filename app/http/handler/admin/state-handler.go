package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"binlog_spread/app/models"
	"binlog_spread/app/service"
	"github.com/gin-gonic/gin"
)

type stateHandler struct {
	handler.BaseHandler
}

var (
	StateHandler = new(stateHandler)
)


func (h stateHandler) Query(c *gin.Context) {
	form := &dto.StateForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		h.JsonErr(c, err)
		return
	}
	form.OrderBy = append(form.OrderBy, "state_id desc")

	ret, err := service.StateSrv.Query(form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}

func (h stateHandler) GetInfo(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")

	info, err := dal.StateDAO.Get(id)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, info)
}

func (h stateHandler) Add(c *gin.Context) {
	info := models.State{}
	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}
	if err := dal.StateDAO.Insert(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h stateHandler) Edit(c *gin.Context) {
	info := models.State{}
	id, _ := h.ParamInt(c, "id")
	info.StateId = id

	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	if err := dal.StateDAO.Update(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h stateHandler) Delete(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")
	if err := dal.StateDAO.Delete(id); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, "")
}

