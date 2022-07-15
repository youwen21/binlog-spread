package admin

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/http/handler"
	lib2 "binlog_spread/lib"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type authHandler struct {
	handler.BaseHandler
}

var (
	AuthHandler = new(authHandler)
)

type login struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	KeepLogin string `json:"keep_login" form:"keep_login"`
	JumpURL   string `json:"jump_url" form:"jump_url"`
}

//DefaultJumpURL 登录成功默认跳转页
var DefaultJumpURL string = "/member/userinfo_detail.html"

//GetLoginParams 获取登录入参
func GetLoginParams(c *gin.Context) (*login, error) {
	loginParams := &login{}

	err := c.ShouldBind(&loginParams)
	if nil != err {
		return nil, err
	}

	return loginParams, nil
}

func (h authHandler) Index(c *gin.Context) {
	h.JsonOk(c, "welcome")
}

func (h authHandler) Login(c *gin.Context) {
	loginParams, err := GetLoginParams(c)
	if nil != err {
		h.JsonErr(c, err)
		return
	}

	adminInfo, _ := dal.AdminDAO.GetByUsername(loginParams.Username)
	if nil == adminInfo {
		h.JsonFail(c, 800, "用户名或密码错误")
		return
	}

	// 密码校验
	genPwd := EncryptWord(loginParams.Password, nil)
	if adminInfo.Password != genPwd {
		h.JsonFail(c, 800, "用户名或密码错误2")
		return
	}

	tokenString, err := lib2.GenToken(adminInfo.AdministratorId)
	if err != nil {
		h.JsonErr(c, err)
	}

	// 设置cookie
	c.SetCookie("AdminAuthorization", tokenString, 86400, "/", c.Request.Host, false, true)

	// 如果是ajax，不需要跳转，直接输出结果
	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		h.JsonOk(c, gin.H{"token": tokenString, "info": gin.H{"admin_id": adminInfo.AdministratorId}, "jump_url": loginParams.JumpURL})
		return
	}
	// 登录完成跳转到jump_url
	c.Redirect(302, loginParams.JumpURL)
	return
}

func (h authHandler) Logout(c *gin.Context) {
	h.JsonOk(c, "welcome")
}

func EncryptWord(passwd string, salt interface{}) (hashResult string) {

	h := sha1.New()

	h.Write([]byte(passwd))
	if nil != salt {
		h.Write([]byte(salt.(string)))
	}

	// 16进制输出的结果才和php是一样的。  php默认按16进制进行输出。
	//@see https://segmentfault.com/q/1010000007510284
	hashResult = fmt.Sprintf("%x", h.Sum(nil))
	return
}
