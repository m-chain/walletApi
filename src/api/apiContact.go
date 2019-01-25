package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"walletApi/src/model"
)

// 联系人相关操作
type ApiContactController struct {
	BaseController
}

// @Title Create
// @Description 添加联系人
// @Param	body		body 	model.Contact	true		"Contact 对象"
// @Success 200 {object} model.JsonResult
// @Failure 403 address可以有多个，以逗号隔开;在添加钱包调用时，只需要传递所有钱包地址即可，其它的联系人信息可不用传递
// @router /addContact [post]
func (o *ApiContactController) Post() {
	var result model.JsonResult
	var cont model.Contact
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &cont)
	if err != nil {
		result.Status = false
		result.Msg = o.Tr("json_parse_error")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if cont.Address == "" {
		result.Status = false
		result.Msg = o.Tr("address_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	arr := strings.Split(cont.Address, ",")
	if len(arr) == 1 {
		if cont.Name == "" || cont.ContactAddr == "" {
			result.Status = false
			result.Msg = o.Tr("contact_name_empty")
			o.Data["json"] = result
			o.ServeJSON()
			return
		}
	}

	num, err := model.AddContact(&cont)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
	} else {
		result.Status = true
		result.Data = fmt.Sprintf("%d", num)
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Update
// @Description 修改联系人
// @Param	body		body 	model.Contact	true		"Contact 对象"
// @Success 200 {object} model.JsonResult
// @Failure 403 address可以有多个，以逗号隔开
// @router /updateContact [post]
func (o *ApiContactController) Update() {
	var result model.JsonResult
	var cont model.Contact
	//使用requestBody需要在app.conf里面设置copyrequestbody = true
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &cont)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if cont.Name == "" || cont.Address == "" || cont.ContactAddr == "" {
		result.Status = false
		result.Msg = o.Tr("no_valid_args")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if cont.Id == 0 {
		result.Status = false
		result.Msg = o.Tr("id_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	total, err := model.UpdateContact(&cont)
	if err != nil {
		result.Status = false
		result.Msg = err.Error()
	} else {
		result.Status = true
		result.Data = fmt.Sprintf("%d", total)
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Delete
// @Description 删除联系人
// @Param	id		formData 	int	    true		"联系人记录ID"
// @Param	address	formData 	string	true		"钱包地址，多个以逗号分隔"
// @Success 200 {object} model.JsonResult
// @Failure 403 :记录ID、address不能为空
// @router /deleteContact [post]
func (o *ApiContactController) Delete() {
	id := o.GetString("id", "0")
	address := o.GetString("address", "")
	var result model.JsonResult

	if id == "" || id == "0" {
		result.Status = false
		result.Msg = o.Tr("id_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}

	if address == "" {
		result.Status = false
		result.Msg = o.Tr("address_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	_id, _ := strconv.ParseInt(id, 10, 64)
	total, err := model.DeleteContact(_id, address)
	if err != nil {
		result.Status = false
		result.Msg = o.Tr("delete_failure")
	} else {
		result.Status = true
		result.Data = fmt.Sprintf("%d", total)
		result.Msg = o.Tr("delete_success")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Query
// @Description 查询联系人详情
// @Param	id		path 	int	true		"联系人记录ID"
// @Success 200 {object} model.ContactResult
// @Failure 403 :记录ID不能为空
// @router /getContactInfo/:id [get]
func (o *ApiContactController) Query() {
	id := o.Ctx.Input.Param(":id")
	var result model.ContactResult
	if id != "" && id != "0" {
		_id, _ := strconv.ParseInt(id, 10, 64)
		cont, err := model.GetContactInfo(_id)
		if err != nil {
			result.Status = false
			result.Msg = o.Tr("not_found_msg")
		} else {
			result.Status = true
			result.Data = *cont
		}
	} else {
		result.Status = false
		result.Msg = o.Tr("id_empty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title QueryAll
// @Description 查询我的所有联系人
// @Param	address		path 	string	true		"钱包地址,只需传递任意一个地址"
// @Success 200 {object} model.ContactListResult
// @Failure 403 :钱包地址不能为空,只需传递任意一个地址
// @router /getContactAll/:address [get]
func (o *ApiContactController) QueryAll() {
	address := o.Ctx.Input.Param(":address")
	var result model.ContactListResult
	if address != "" {
		arr, err := model.GetContactAll(address)
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
