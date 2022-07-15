package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"binlog_spread/app/models"
	"github.com/gin-gonic/gin"
)

type apiDefineHandler struct {
	handler.BaseHandler
}

var (
	ApiDefineHandler = new(apiDefineHandler)
)


func (h apiDefineHandler) Query(c *gin.Context) {
	form := &dto.PageParam{}
	if err := c.ShouldBind(&form); err != nil {
		h.JsonErr(c, err)
		return
	}
	form.OrderBy = append(form.OrderBy, "api_define_id desc")

	ret, err := dal.ApiDefineDAO.Query(form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}

func (h apiDefineHandler) GetInfo(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")

	info, err := dal.ApiDefineDAO.Get(id)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, info)
}

func (h apiDefineHandler) Add(c *gin.Context) {
	info := models.ApiDefine{}
	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}
	if err := dal.ApiDefineDAO.Insert(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h apiDefineHandler) Edit(c *gin.Context) {
	info := models.ApiDefine{}
	id, _ := h.ParamInt(c, "id")
	info.ApiDefineId = id

	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	if err := dal.ApiDefineDAO.Update(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h apiDefineHandler) Delete(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")
	if err := dal.ApiDefineDAO.RealDelete(id); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, "")
}

