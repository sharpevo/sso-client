package common

import (
	"github.com/astaxie/beego"
	"net/http"
	"sso-client/utils/userinfo"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetUserInfo() (userInfo userinfo.UserInfo) {
	userInfo = userinfo.GetUserInfo(c.Ctx.Request)
	beego.Debug("GetUserInfo:", userInfo)
	c.Data["UserInfo"] = userInfo
	return userInfo
}

func (c *BaseController) SetCookie(domain string, value string, exp int64) {
	if value == "" {
		beego.Debug("SetCookie:", "Nothing to set.")
		return
	}
	if exp < 10 {
		exp = 30
	}
	if exp > 7200 { // 5 days
		exp = 7200
	}
	expiration := time.Now().Add(time.Duration(exp) * time.Minute)
	cookie := http.Cookie{Name: "token", Value: value, HttpOnly: true, Domain: domain, Expires: expiration}
	beego.Debug("SetCookie:", cookie)
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)
}
