package admin

import (
	"binlog_spread/app/http/handler"
	"binlog_spread/app/service"
	"github.com/gin-gonic/gin"
)

type sidebarHandler struct {
	handler.BaseHandler
}

var (
	SidebarHandler = new(sidebarHandler)
)

func (h sidebarHandler) GetSideBarTree(c *gin.Context) {
	ret, _ := service.MenuSrv.GetMenuTree()
	h.JsonOk(c, ret)
}
