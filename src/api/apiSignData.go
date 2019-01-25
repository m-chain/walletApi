package api

import (
	"encoding/hex"
	"encoding/json"
	"walletApi/src/common"
	"walletApi/src/model"
	"walletApi/src/service"
)

// 签名相关操作
type ApiSignController struct {
	BaseController
}

// @Title Update
// @Description 设置签名数据
// @Param	qrCode	    formData 	string	true		"二维码"
// @Param	address		formData 	string	true		"钱包地址"
// @Param	pubKey		formData 	string	true		"钱包地址公钥"
// @Param	signData	formData 	string	true		"已签名的数据"
// @Success 200 {object} model.CCInvokeResult
// @Failure 403 只需要传递地址、公钥、签名数据、二维码标识即可
// @router /updateSignData [post]
func (o *ApiSignController) Update() {

	var qrCode = o.GetString("qrCode")
	var address = o.GetString("address")
	var pubKey = o.GetString("pubKey")
	var signData = o.GetString("signData")

	var result model.CCInvokeResult
	//使用requestBody需要在app.conf里面设置copyrequestbody = true
	//err := json.Unmarshal(o.Ctx.Input.RequestBody, &sign)
	if qrCode == "" {
		result.Status = false
		result.Msg = o.Tr("qrcode_empty")
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
	if pubKey == "" {
		result.Status = false
		result.Msg = o.Tr("public_key_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	if signData == "" {
		result.Status = false
		result.Msg = o.Tr("sign_data_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	var sign model.SignData
	sign.QrCode = qrCode
	sign.Address = address
	sign.PubKey = pubKey
	sign.SignData = signData

	service.InvokeChainCC(&sign, o.Lang, func(ret model.CCInvokeResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title querySignInfo
// @Description 获取签名信息
// @Param	qrCode		path 	string	true		"二维码标识 9f9acb8f749b4f32b2e9dcab30601560"
// @Success 200 {object} model.SignInfoResult
// @Failure 403 :二维码标识不能为空
// @router /querySignInfo/:qrCode [get]
func (o *ApiSignController) QuerySignInfo() {
	qrCode := o.Ctx.Input.Param(":qrCode")
	var result model.SignInfoResult
	if qrCode != "" {
		sign, err := model.GetSingInfo(qrCode)
		if err != nil || sign == nil {
			result.Status = false
			result.Msg = o.Tr("not_found_sign")
		} else {
			if sign.Status == -1 {
				result.Status = false
				result.Msg = o.Tr("no_valid_qrcode")
			} else {
				result.Status = true
				//1、发布合约 2、实例化合约 3、升级合约 4、发行token 5、设置master、manager 6、manager签名确认 7、添加manager　8、替换manager 9 、删除manager 10、设置manager操作确认的阀值 11、设置发行token所需要的平台币 12、设置发布合约所需要的平台币 13、设置手续费返还规则 14、删除合约 15、设置master操作所需的确认数阀值
				sign.TipMsg = o.Tr(common.SignTips[sign.SignType].TipMsg)
				result.Data = sign
			}
		}
	} else {
		result.Status = false
		result.Msg = o.Tr("qrcode_empty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Query
// @Description 查询签名是否成功
// @Param	qrCode		path 	string	true		"二维码标识"
// @Success 200 {object} model.CCInvokeResult
// @Failure 403 :二维码标识不能为空
// @router /getSignStatus/:qrCode [get]
func (o *ApiSignController) GetSignStatus() {
	qrCode := o.Ctx.Input.Param(":qrCode")
	var result model.CCInvokeResult
	if qrCode != "" {
		sign, err := model.GetSingInfo(qrCode)
		if err != nil || sign == nil {
			result.Status = false
			result.Msg = o.Tr("not_found_sign")
		} else {
			if sign.Status == -1 {
				result.Status = false
				result.Msg = o.Tr("no_valid_qrcode")
			} else {
				if sign.Status == 3 {
					json.Unmarshal([]byte(sign.RespResult), &result)
					result.Msg = o.Tr("confirm_success")
				} else {
					result.Status = false
					result.Msg = o.Tr("rescan_qrcode")
				}
			}

		}
	} else {
		result.Status = false
		result.Msg = o.Tr("qrcode_empty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title queryAllSignInfos
// @Description 获取某地址下所有的签名记录
// @Param	address		formData 	string	true	"钱包地址"
// @Param	page		formData 	int	true		"第几页"
// @Param	pageSize	formData 	int	true		"每页多少条记录默认20"
// @Success 200 {object} model.SignListResult
// @router /queryAllSignInfos [post]
func (o *ApiSignController) QueryAllSignInfos() {
	address := o.GetString("address")
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.GetAllSignInfos(address, pageSize, page)
	resultMap["status"] = true
	o.Data["json"] = resultMap
	o.ServeJSON()
}

// @Title queryConfirmInfo
// @Description 获取签名情况
// @Param	qrCode		path 	string	true  "二维码标识 platsignfbc9e0bd97f348d1840843f3c9d7aaf1"
// @Success 200 {object} model.ConfirmResult
// @Failure 403 :二维码标识不能为空
// @router /queryConfirmInfo/:qrCode [get]
func (o *ApiSignController) QueryConfirmInfo() {
	qrCode := o.Ctx.Input.Param(":qrCode")
	var result model.ConfirmResult
	if qrCode == "" {
		result.Msg = o.Tr("qrcode_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	sign, err := model.GetSingInfo(qrCode)
	if err != nil || sign == nil {
		result.Msg = o.Tr("not_found_sign")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	type Origin struct {
		SignType int
		TokenID  string
	}
	var origin Origin
	originBytes, _ := hex.DecodeString(sign.OriginData)
	json.Unmarshal(originBytes, &origin)

	service.QueryConfirmInfo(origin.TokenID, sign.OriginData, func(ret model.ConfirmResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}
