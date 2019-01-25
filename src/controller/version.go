package controller

import (
	"fmt"
	"strconv"
	"walletApi/src/model"

	"github.com/astaxie/beego"
)

type VersionController struct {
	beego.Controller
}

// @router /Version/InitList/ [get]
func (c *VersionController) InitList() {
	c.TplName = "versionlist.html"
}

// @router /Version/List/ [post]
func (c *VersionController) List() {

	keyWorld := c.GetString("keyWorld")

	pageNo, _ := c.GetInt("current")

	rowCount, _ := c.GetInt("rowCount")

	if pageNo == 0 {
		pageNo = 1
	}
	var version model.AppVersion
	resultMap := version.List(rowCount, pageNo, keyWorld)
	c.Data["json"] = map[string]interface{}{"rows": resultMap["data"], "rowCount": rowCount, "current": pageNo, "total": resultMap["total"]}

	c.ServeJSON()
}

// @router /Version/InitAdd/ [get]
func (c *VersionController) InitAdd() {
	id := c.GetString("id")
	if id != "" {
		var version model.AppVersion
		ver, _ := version.GetVersionInfo(id)
		if ver != nil {
			c.Data["version"] = ver
			c.Data["IsCurrent"] = ver.IsCurrent
			c.Data["UpgradeType"] = ver.UpgradeType
		}
	} else {
		c.Data["IsCurrent"] = 1
		c.Data["UpgradeType"] = 2
	}
	c.TplName = "versionadd.html"
}

// @router /Version/AddVersion [post]
func (c *VersionController) AddVersion() {
	//自动解析绑定到对象中,ParseForm 不支持解析raw data,必须是表单form提交
	version := model.AppVersion{}
	result := new(model.Result)
	var err error
	err = c.ParseForm(&version)
	idStr := c.GetString("id")
	if err == nil {
		if idStr != "" {
			id, _ := strconv.ParseInt(idStr, 10, 64)
			version.Id = id
			err = model.UpdateVersion(&version)
		} else {
			err = model.AddVersion(&version)
		}

		if err != nil {
			result.Status = 1
			result.Msg = err.Error()
		} else {
			result.Status = 0
			result.Data = map[string]int64{"Id": version.Id}
		}
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("操作失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /Version/DeleteVersion/:id [get]
func (c *VersionController) DeleteUser() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.DeleteVersion(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("删除失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}
