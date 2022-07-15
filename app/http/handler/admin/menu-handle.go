package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"binlog_spread/app/models"
	"binlog_spread/app/service"
	"github.com/gin-gonic/gin"
)

type menuHandler struct {
	handler.BaseHandler
}

var (
	MenuHandler = new(menuHandler)
)

func (h menuHandler) Query(c *gin.Context) {
	form := &dto.MenuForm{}
	if err := c.ShouldBindQuery(&form); err != nil {
		h.JsonErr(c, err)
		return
	}
	form.OrderBy = append(form.OrderBy, "menu_id desc")

	ret, err := service.MenuSrv.Query(form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}

func (h menuHandler) GetInfo(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")

	info, err := service.MenuSrv.Get(id)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, info)
}

func (h menuHandler) Add(c *gin.Context) {
	info := models.Menu{}
	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}
	if err := dal.MenuDAO.Insert(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h menuHandler) Edit(c *gin.Context) {
	info := models.Menu{}
	id, _ := h.ParamInt(c, "id")
	info.MenuId = id

	if err := c.ShouldBind(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	if err := dal.MenuDAO.Update(&info); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, info)
}

func (h menuHandler) Delete(c *gin.Context) {
	id, _ := h.ParamInt(c, "id")
	if err := dal.MenuDAO.Delete(id); err != nil {
		h.JsonErr(c, err)
		return
	}

	h.JsonOk(c, "")
}

func (h menuHandler) GetSelectList(c *gin.Context) {
	tree, _ := service.MenuSrv.GetMenuTree()
	var selectData []dto.MenuSelectItem
	dto.TreeToSelect(&selectData, tree, 0)

	h.JsonOk(c, selectData)
}
