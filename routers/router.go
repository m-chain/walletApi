// @APIVersion 1.0.0
// @Title wallet API
// @Description 钱包相关api接口
// @Contact 278985177@qq.com
// @TermsOfServiceUrl http://www.m-chain.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"walletApi/src/api"
	"walletApi/src/controller"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	//默认的
	beego.Router("/", &controller.MainController{})

	//验证码
	beego.Handler("/captcha/*.png", captcha.Server(106, 36))

	//JS分页
	beego.Router("/Home/PageNextData", &controller.YonghuController{})

	//Bootstrap运用
	beego.Router("/Home/Index", &controller.PageController{})
	//Easyui使用
	beego.Router("/Home/EasyUI", &controller.EasyUIController{})
	beego.Router("/Home/EasyUIData", &controller.EasyUIDataController{})
	//Api接口部分
	beego.Router("/api/Html", &api.ApiController{})
	beego.Router("/api/GetJson", &api.ApiJsonController{})
	beego.Router("/api/GetXml", &api.ApiXMLController{})
	beego.Router("/api/GetJsonp", &api.ApiJsonpController{})
	beego.Router("/api/GetDictionary", &api.ApiDictionaryController{})
	beego.Router("/api/GetParams", &api.ApiParamsController{})
	//session部分
	beego.Router("/Home/Login", &controller.LoginController{})
	beego.Router("/Home/InitReg", &controller.LoginController{}, "get:InitReg")
	beego.Router("/Home/InitPass", &controller.LoginController{}, "get:InitPass")
	beego.Router("/Home/Reg", &controller.LoginController{}, "post:Reg")
	beego.Router("/Home/UpdatePwd", &controller.LoginController{}, "post:UpdatePwd")

	//用户管理
	beego.Include(&controller.UserController{})

	//版本管理
	beego.Include(&controller.VersionController{})

	//系统通知管理
	beego.Include(&controller.SysMsgController{})

	//意见反馈
	beego.Include(&controller.QuestionController{})

	beego.Router("/Home/Logout", &controller.LogoutController{})
	//布局页面部分
	beego.Router("/Home/Layout", &controller.LayoutController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/version",
			beego.NSInclude(
				&api.ApiVersionController{},
			),
		),
		beego.NSNamespace("/question",
			beego.NSInclude(
				&api.ApiQuestionController{},
			),
		),
		beego.NSNamespace("/contact",
			beego.NSInclude(
				&api.ApiContactController{},
			),
		),
		beego.NSNamespace("/sign",
			beego.NSInclude(
				&api.ApiSignController{},
			),
		),
		beego.NSNamespace("/wallet",
			beego.NSInclude(
				&api.WalletController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
