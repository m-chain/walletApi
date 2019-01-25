package api

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type (
	BaseController struct {
		beego.Controller
		i18n.Locale
	}
)

func (this *BaseController) Prepare() {
	lang := this.Ctx.Input.Header("language")
	if lang == "en" {
		this.Lang = "en-US"
	} else if lang == "tw" {
		this.Lang = "zh-TW"
	} else {
		this.Lang = "zh-CN"
	}
}
