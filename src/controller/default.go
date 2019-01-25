package controller

import (
	"walletApi/src/common"
	"walletApi/src/model"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//需要先获取一下session，以免c.CruSession未空
	if c.CruSession == nil {
		c.GetSession(common.USER_INFO)
	}
	u, ok := c.Ctx.Input.Session(common.USER_INFO).(*model.User)
	if !ok {
		d := struct {
			CaptchaId string
		}{
			captcha.NewLen(4),
		}
		c.Data["CaptchaId"] = d.CaptchaId
		c.TplName = "login.html"
	} else {
		c.Data["userName"] = u.UserName
		c.Data["nickName"] = u.NickName
		if u.UserName == "admin" {
			c.Data["IsAdmin"] = true
		}
		model.UpdateLastLoginTime(u.UserName)

		c.TplName = "index.html"
	}
}
