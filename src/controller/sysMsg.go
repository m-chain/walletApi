package controller

import (
	"fmt"
	"html"
	"strconv"
	"walletApi/src/model"

	"github.com/astaxie/beego"
)

type SysMsgController struct {
	beego.Controller
}

// @router /Message/InitList/ [get]
func (c *SysMsgController) InitList() {
	c.TplName = "msglist.html"
}

// @router /Message/List/ [post]
func (c *SysMsgController) List() {

	keyWord := c.GetString("keyWord")

	pageNo, _ := c.GetInt("current")

	rowCount, _ := c.GetInt("rowCount")

	if pageNo == 0 {
		pageNo = 1
	}
	resultMap := model.QuerySysMsgList(rowCount, pageNo, keyWord)
	c.Data["json"] = map[string]interface{}{"rows": resultMap["data"], "rowCount": rowCount, "current": pageNo, "total": resultMap["total"]}

	c.ServeJSON()
}

// @router /Message/InitAdd/ [get]
func (c *SysMsgController) InitAdd() {
	id := c.GetString("id")
	if id != "" {
		message, _ := model.GetSysMessage(id)
		if message != nil {
			c.Data["message"] = message
		}
	}
	c.TplName = "msgadd.html"
}

// @router /Message/AddMessage [post]
func (c *SysMsgController) AddMessage() {
	//自动解析绑定到对象中,ParseForm 不支持解析raw data,必须是表单form提交
	message := model.SysMessage{}
	result := new(model.Result)
	var err error
	err = c.ParseForm(&message)
	idStr := c.GetString("id")
	if err == nil {
		message.Content = html.EscapeString(message.Content)
		if idStr != "" {
			id, _ := strconv.ParseInt(idStr, 10, 64)
			message.MsgId = id
			err = model.UpdateSysMessage(&message)
		} else {
			err = model.AddSysMessage(&message)
		}

		if err != nil {
			result.Status = 1
			result.Msg = err.Error()
		} else {
			result.Status = 0
			result.Data = map[string]int64{"msgId": message.MsgId}
		}
	} else {
		result.Status = 1
		result.Msg = fmt.Sprintf("操作失败,%s", err.Error())
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /Message/DeleteMessage/:id [get]
func (c *SysMsgController) DeleteMessage() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.DeleteSysMessage(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("删除失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}
