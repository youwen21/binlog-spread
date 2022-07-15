package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"binlog_spread/app/models"
	"binlog_spread/app/service"
	"github.com/gin-gonic/gin"
)

type stateDirectionHandler struct {
	handler.BaseHandler
}

var (
	StateDirectionHandler = new(stateDirectionHandler)
)

func (h stateDirectionHandler) Query(c *gin.Context) {
	form := &dto.StateDirectionForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		h.JsonErr(c, err)
		return
	}
	form.OrderBy = append(form.OrderBy, "state_direction_id desc")

	ret, err := service.StateDirectionSrv.Query(form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}

func (h stateDirectionHandler) GetInfo(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")

	info, err := dal.StateDirectionDAO.Get(id)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, info)
}

func (h stateDirectionHandler) Add(c *gin.Context) {
	info := models.StateDirection{}
	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}
	if err := dal.StateDirectionDAO.Insert(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h stateDirectionHandler) Edit(c *gin.Context) {
	info := models.StateDirection{}
	id, _ := h.ParamInt(c, "id")
	info.StateDirectionId = id

	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	if err := dal.StateDirectionDAO.Update(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h stateDirectionHandler) Delete(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")
	if err := dal.StateDirectionDAO.RealDelete(id); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, "")
}
