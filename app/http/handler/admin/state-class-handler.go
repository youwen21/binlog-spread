package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"binlog_spread/app/models"
	"github.com/gin-gonic/gin"
)

type stateClassHandler struct {
	handler.BaseHandler
}

var (
	StateClassHandler = new(stateClassHandler)
)


func (h stateClassHandler) Query(c *gin.Context) {
	form := &dto.StateClassForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		h.JsonErr(c, err)
		return
	}
	form.OrderBy = append(form.OrderBy, "state_class_id desc")

	ret, err := dal.StatClassDAO.Query(form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}

func (h stateClassHandler) GetInfo(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")

	info, err := dal.StatClassDAO.Get(id)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, info)
}

func (h stateClassHandler) Add(c *gin.Context) {
	info := models.StateClass{}
	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}
	if err := dal.StatClassDAO.Insert(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h stateClassHandler) Edit(c *gin.Context) {
	info := models.StateClass{}
	id, _ := h.ParamInt(c, "id")
	info.StateClassId = id

	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	if err := dal.StatClassDAO.Update(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h stateClassHandler) Delete(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")
	if err := dal.StatClassDAO.RealDelete(id); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, "")
}

