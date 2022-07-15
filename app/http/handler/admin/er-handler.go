package admin

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/http/handler"
	"binlog_spread/app/service"
	"github.com/gin-gonic/gin"
)

type erHandler struct {
	handler.BaseHandler
}

var (
	ERHandler = new(erHandler)
)

func (h erHandler) Query(c *gin.Context) {
	//DddErSearch DddEr搜索
	form := dto.ErSearch{}
	if err:= c.ShouldBind(&form); err != nil{
		h.JsonErr(c, err)
		return
	}

	ret, err := service.ErSrv.Query(&form)
	if err != nil {
		h.JsonErr(c, err)
		return
	}
	h.JsonOk(c, ret)
}
