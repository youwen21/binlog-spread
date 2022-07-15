package handler

import (
	"binlog_spread/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BaseHandler struct {
}

func (h *BaseHandler) ParamInt(c *gin.Context, key string) (int, error) {
	v := c.Param(key)
	return strconv.Atoi(v)
}

func (h *BaseHandler) JsonOk(c *gin.Context, data interface{}) {
	h.JsonRaw(c, app.SUCCESS, app.StatusText(app.SUCCESS), data)
}

func (h *BaseHandler) JsonErr(c *gin.Context, e error) {
	h.JsonRaw(c, app.ERROR, e.Error(), nil)
}

func (h *BaseHandler) JsonFail(c *gin.Context, code int, msg string) {
	h.JsonRaw(c, code, msg, nil)
}

func (h *BaseHandler) JsonRaw(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
}
