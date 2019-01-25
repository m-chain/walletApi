package controller

import (
	"fmt"
	"strconv"
	"walletApi/src/common"
	"walletApi/src/model"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @router /User/InitList/ [get]
func (c *UserController) InitList() {
	c.TplName = "userlist.html"
}

// @router /User/List/ [post]
func (c *UserController) List() {
	keyWorld := c.GetString("keyWorld")

	pageNo, _ := c.GetInt("current")

	rowCount, _ := c.GetInt("rowCount")

	if pageNo == 0 {
		pageNo = 1
	}

	resultMap := model.SearchUsers(rowCount, pageNo, keyWorld)
	c.Data["json"] = map[string]interface{}{"rows": resultMap["data"], "rowCount": rowCount, "current": pageNo, "total": resultMap["total"]}

	c.ServeJSON()
}

// @router /User/InitAdd/ [get]
func (c *UserController) InitAdd() {
	c.TplName = "useradd.html"
}

// @router /User/DeleteUser/:id [get]
func (c *UserController) DeleteUser() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.DeleteUser(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("删除失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /User/AddUser [post]
func (c *UserController) AddUser() {
	//自动解析绑定到对象中,ParseForm 不支持解析raw data,必须是表单form提交
	user := model.User{}
	result := new(model.Result)
	err := c.ParseForm(&user)
	if err == nil {
		user.Password = common.GetMD5Str(user.Password)
		err = model.AddUser(&user)
		if err != nil {
			result.Status = 1
			result.Msg = err.Error()
		} else {
			result.Status = 0
			result.Data = map[string]int64{"Id": user.Id}
		}
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("添加用户失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /User/Qrcode [get]
func (c *UserController) Qrcode() {
	c.TplName = "qrcode.html"
}

// @router /User/CCQrcode [get]
func (c *UserController) CCQrcode() {
	c.TplName = "ccqrcode.html"
}

// @router /User/InitCCQrcode [get]
func (c *UserController) InitCCQrcode() {
	c.TplName = "initccqrcode.html"
}

func (c *UserController) Get() {

	c.TplName = "list.html"
}
func (c *UserController) Post() {
	pageno, err := c.GetInt("pageno")
	if err != nil {
		fmt.Println(err)
	}
	search := c.GetString("search")
	userList := model.SearchDataList(3, pageno, search)
	listnum := model.GetRecordNum(search)
	c.Data["json"] = map[string]interface{}{"Count": listnum, "PageSize": 3, "Page": pageno, "DataList": userList}
	c.ServeJSON()
}

type YonghuController struct {
	beego.Controller
}

func (c *YonghuController) Post() {
	pageno, err := c.GetInt("pageno")
	if err != nil {
		fmt.Println(err)
	}
	search := c.GetString("search")
	userList := model.SearchDataList(3, pageno, search)
	listnum := model.GetRecordNum(search)
	c.Data["json"] = map[string]interface{}{"Count": listnum, "PageSize": 3, "Page": pageno, "DataList": userList}
	c.ServeJSON()
}
