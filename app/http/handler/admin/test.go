package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/http/handler"
	"github.com/gin-gonic/gin"
)

type testHandler struct {
	handler.BaseHandler
}

var (
	TestHandler = new(testHandler)
)

func (h *testHandler) TestGet(c *gin.Context)  {
	//ret, _ := dal.TestDAO.Query()
	//h.JsonOk(c, ret)

	ret, _ := dal.TestDAO.Get(239)
	h.JsonOk(c, ret)
}