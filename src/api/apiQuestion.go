package api

import (
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"walletApi/src/model"
)

// 意见反馈相关操作
type ApiQuestionController struct {
	BaseController
}

// @Title Create
// @Description 意见反馈
// @Param	body		body 	model.Question	true		"Question 对象"
// @Success 200 {object} model.JsonResult
// @Failure 403 body is empty
// @router /addQuestion [post]
func (o *ApiQuestionController) Post() {
	var result model.JsonResult
	var quest model.Question
	//使用requestBody需要在app.conf里面设置copyrequestbody = true
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &quest)
	if err != nil {
		result.Status = false
		result.Msg = o.Tr("json_parse_error")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if quest.Title == "" || quest.Address == "" || quest.Question == "" {
		result.Status = false
		result.Msg = o.Tr("no_valid_args")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	quest.Question = html.EscapeString(quest.Question)

	err = model.AddQuestion(&quest)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
	} else {
		result.Status = true
		result.Data = fmt.Sprintf("%d", quest.Id)
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Update
// @Description 修改意见反馈
// @Param	body		body 	model.Question	true		"Question 对象"
// @Success 200 {object} model.JsonResult
// @Failure 403 body is empty
// @router /updateQuestion [post]
func (o *ApiQuestionController) Update() {
	var result model.JsonResult
	var quest model.Question
	//使用requestBody需要在app.conf里面设置copyrequestbody = true
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &quest)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if quest.Title == "" || quest.Address == "" || quest.Question == "" {
		result.Status = false
		result.Msg = o.Tr("no_valid_args")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if quest.Id == 0 {
		result.Status = false
		result.Msg = o.Tr("id_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	err = model.UpdateQuestion(&quest)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
	} else {
		result.Status = true
		result.Data = fmt.Sprintf("%d", quest.Id)
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Delete
// @Description 删除反馈意见
// @Param	id		path 	int	true		"意见记录ID"
// @Success 200 {object} model.JsonResult
// @Failure 403 :记录ID不能为空
// @router /deleteQuestion/:id [get]
func (o *ApiQuestionController) Delete() {
	id := o.Ctx.Input.Param(":id")
	var result model.JsonResult
	if id != "" && id != "0" {
		_id, _ := strconv.ParseInt(id, 10, 64)
		err := model.DeleteQuestion(_id)
		if err != nil {
			result.Status = false
			result.Msg = o.Tr("delete_failure")
		} else {
			result.Status = true
			result.Msg = o.Tr("delete_success")
		}
	} else {
		result.Status = false
		result.Msg = o.Tr("id_empty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Query
// @Description 查询反馈意见详情
// @Param	id		path 	int	true		"意见记录ID"
// @Success 200 {object} model.QuestionResult
// @Failure 403 :记录ID不能为空
// @router /getQuestionInfo/:id [get]
func (o *ApiQuestionController) Query() {
	id := o.Ctx.Input.Param(":id")
	var result model.QuestionResult
	if id != "" && id != "0" {
		_id, _ := strconv.ParseInt(id, 10, 64)
		quest, err := model.GetQuestionInfo(_id)
		if err != nil {
			result.Status = false
			result.Msg = o.Tr("not_found_msg")
		} else {
			result.Status = true
			result.Data = *quest
		}
	} else {
		result.Status = false
		result.Msg = o.Tr("id_empty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title QueryAll
// @Description 查询我的反馈意见
// @Param	address		formData 	string	true		"钱包地址"
// @Success 200 {object} model.QuestionListResult
// @Failure 403 :钱包地址不能为空,如果有多个可以逗号分隔
// @router /getQuestionAll [post]
func (o *ApiQuestionController) QueryAll() {
	address := o.GetString("address")
	var result model.QuestionListResult
	if address != "" {
		arr, err := model.GetQuestionAll(address)
		if err != nil {
			result.Status = false
			result.Msg = o.Tr("not_found_msg")
		} else {
			result.Status = true
			result.Data = arr
		}
	} else {
		result.Status = false
		result.Msg = o.Tr("address_empty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}
