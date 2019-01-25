package controller

import (
	"fmt"
	"walletApi/src/common"
	"walletApi/src/model"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

type LoginController struct {
	beego.Controller
}

//登录页面
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

//登录功能
func (c *LoginController) Post() {
	islogin := 0 //0、表示登录成功 1、表示用户不存在 2、密码错误 3、验证码错误
	captchaID := c.GetString("captchaId")
	captchaValue := c.GetString("captcha")
	if !captcha.VerifyString(captchaID, captchaValue) {
		captcha.Reload(captchaID)
		islogin = 3
		c.Data["json"] = map[string]interface{}{"islogin": islogin}
		c.ServeJSON()
		return
	}

	name := c.GetString("name")
	pwd := c.GetString("pwd")

	user, err := model.GetUserInfoByAccount(name)
	if err != nil || user == nil {
		islogin = 1
		c.Data["json"] = map[string]interface{}{"islogin": islogin}
		c.ServeJSON()
		return
	}
	currPwd := common.GetMD5Str(pwd)
	if currPwd == user.Password { //登录成功
		captcha.Delete(captchaID)
		c.SetSession(common.USER_INFO, user)
	} else { //密码错误
		islogin = 2
	}
	c.Data["json"] = map[string]interface{}{"islogin": islogin}
	c.ServeJSON()
}

func (c *LoginController) InitReg() {
	c.TplName = "register.html"
}

func (c *LoginController) InitPass() {
	c.TplName = "updatePwd.html"
}

func (c *LoginController) Reg() {
	userName := c.GetString("userName")

	//检验用户是否已经注册
	user, err := model.GetUserInfoByAccount(userName)
	if user != nil && user.UserName != "" {
		c.Data["json"] = map[string]interface{}{"status": 0, "msg": "该帐号已经注册!"}
		c.ServeJSON()
		return
	}

	//注册用户
	var userInfo model.User
	userInfo.UserName = userName
	userInfo.Password = common.GetMD5Str(c.GetString("surePwd"))
	userInfo.NickName = c.GetString("surePwd")

	err = model.AddUser(&userInfo)
	if err != nil {
		fmt.Printf("************%#v", userInfo.Id)
		c.Data["json"] = map[string]interface{}{"status": 0, "msg": "添加用户失败!"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"status": 1}
	c.ServeJSON()
}

//更新密码
func (c *LoginController) UpdatePwd() {
	u, ok := c.Ctx.Input.Session(common.USER_INFO).(*model.User)
	if !ok {
		c.Data["json"] = map[string]interface{}{"status": 0, "msg": "登录失效，请重新登录!"}
		c.ServeJSON()
		return
	}

	//检验用户是否已经注册
	user, _ := model.GetUserInfoByAccount(u.UserName)
	if user != nil && user.UserName != "" {
		if user.Password != common.GetMD5Str(c.GetString("oldPwd")) {
			c.Data["json"] = map[string]interface{}{"status": 0, "msg": "旧密码不正确!"}
			c.ServeJSON()
			return
		}
	} else {
		c.Data["json"] = map[string]interface{}{"status": 0, "msg": "该用户不存在!"}
		c.ServeJSON()
		return
	}

	//更新用户密码
	newpwd := c.GetString("pwd")
	newpwd = common.GetMD5Str(newpwd)
	model.UpdateUserPwd(u.UserName, newpwd)

	c.Data["json"] = map[string]interface{}{"status": 1}
	c.ServeJSON()
}

//退出
type LogoutController struct {
	beego.Controller
}

//登录退出功能
func (c *LogoutController) Get() {
	v := c.GetSession(common.USER_INFO)
	if v != nil {
		//删除指定的session
		c.DelSession(common.USER_INFO)
		//销毁全部的session
		c.DestroySession()

		fmt.Println("当前的session:")
		fmt.Println(c.CruSession)
	}
	//c.Data["json"] = map[string]interface{}{"islogin": islogin}
	//c.ServeJSON()
	c.Redirect("/", 302)
}
