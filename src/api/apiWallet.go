package api

import (
	"math/big"
	"walletApi/src/common"
	"walletApi/src/model"
	"walletApi/src/service"
)

// 钱包相关操作
type WalletController struct {
	BaseController
}

// @Title walletIsExist
// @Description 钱包是否已经存在帐本中
// @Param	address		path 	string	true		"钱包地址"
// @Success 200 {object} model.JsonResult
// @Failure 403 :钱包地址不能为空
// @router /walletIsExist/:address [get]
func (o *WalletController) WalletIsExist() {
	address := o.Ctx.Input.Param(":address")
	var result model.JsonResult
	if address == "" {
		result.Status = false
		result.Msg = o.Tr("address_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.WalletIsExist(address, func(result model.JsonResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title updateWalletPubkey
// @Description 设置公钥
// @Param	rawData		formData 	string	true		"原始数据加上签名后数据的16进制格式"
// @Success 200 {object} model.JsonResult
// @Failure 403 :16进制数据不能为空
// @router /updateWalletPubkey [post]
func (o *WalletController) UpdatePubkey() {
	rawData := o.GetString("rawData")
	var result model.JsonResult
	if rawData == "" {
		result.Status = false
		result.Msg = o.Tr("hex_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.UpdatePubkey(rawData, func(result model.JsonResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryBalance
// @Description 查询余额
// @Param	address		path 	string	true		"钱包地址"
// @Success 200 {object} model.WalletResp
// @Failure 403 :钱包地址不能为空
// @router /queryBalance/:address [get]
func (o *WalletController) QueryBalance() {
	address := o.Ctx.Input.Param(":address")
	var result model.WalletResp
	if address == "" {
		result.Status = false
		result.Msg = o.Tr("address_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.QueryBalance(address, func(result model.WalletResp) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryTransferCount
// @Description 查询余额
// @Param	address		path 	string	true		"钱包地址"
// @Success 200 {object} model.JsonResult
// @Failure 403 :钱包地址不能为空
// @router /queryTransferCount/:address [get]
func (o *WalletController) QueryTransferCount() {
	address := o.Ctx.Input.Param(":address")
	var result model.JsonResult
	if address == "" {
		result.Status = false
		result.Msg = o.Tr("address_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.QueryTransferCount(address, func(result model.JsonResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryChargeGas
// @Description 查询手续费
// @Success 200 {object} model.JsonResult
// @router /queryChargeGas [get]
func (o *WalletController) QueryChargeGas() {
	service.QueryChargeGas(func(result model.JsonResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title sendRawTransaction
// @Description 一对一转账
// @Param	origin		formData 	string	true		"原始字符串数据"
// @Param	signature	formData 	string	true		"对原始签名后的数据"
// @Success 200 {object} model.JsonResult
// @Failure 403 :16进制数据不能为空
// @router /sendRawTransaction [post]
func (o *WalletController) Transfer() {
	originStr := o.GetString("origin")
	signature := o.GetString("signature")
	var result model.JsonResult
	if originStr == "" {
		result.Status = false
		result.Msg = o.Tr("origin_data_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	if signature == "" {
		result.Status = false
		result.Msg = o.Tr("sign_data_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.Transfer(originStr, signature, o.Lang, func(result model.JsonResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title 获取转账交易消息
// @Description 获取转账交易消息
// @Param	address		formData 	string	true	"钱包地址可以有多个，以逗号分隔"
// @Param	page		formData 	int	true		"第几页"
// @Param	pageSize	formData 	int	true		"每页多少条记录""
// @Success 200 {object} model.JpushMsgListResult
// @router /queryJpushMsgList [post]
func (o *WalletController) QueryJpushMsgList() {
	address := o.GetString("address")
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.QueryJpushMsgList(address, pageSize, page)
	var result model.JpushMsgListResult
	if resultMap["data"] == nil {
		result.Status = false
		result.Msg = o.Tr("not_found_msg")
		o.Data["json"] = result
	} else {
		resultMap["status"] = true
		o.Data["json"] = resultMap
	}
	o.ServeJSON()
}

// @Title 获取系统消息
// @Description 获取系统消息
// @Param	page		formData 	int	true		"第几页"
// @Param	pageSize	formData 	int	true		"每页多少条记录""
// @Success 200 {object} model.SysMsgListResult
// @router /querySysMsgList [post]
func (o *WalletController) QuerySysMsgList() {
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.QuerySysMsgList(pageSize, page, "")
	var result model.SysMsgListResult
	if resultMap["data"] == nil {
		result.Status = false
		result.Msg = o.Tr("not_found_msg")
		o.Data["json"] = result
	} else {
		resultMap["status"] = true
		o.Data["json"] = resultMap
	}
	o.ServeJSON()
}

// @Title 获取系统消息详情
// @Description 获取系统消息详情
// @Param	msgId		path 	string	true		"消息记录ID"
// @Success 200 {object} model.SysMsgResult
// @Failure 403 :消息记录ID不能为空
// @router /querySysMsgInfo/:msgId [get]
func (o *WalletController) QuerySysMsgInfo() {
	msgID := o.Ctx.Input.Param(":msgId")
	var result model.SysMsgResult
	if msgID == "" {
		result.Status = false
		result.Msg = o.Tr("id_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	sysMsg, err := model.GetSysMessage(msgID)
	if err != nil {
		result.Status = false
		result.Msg = o.Tr("not_found_msg")
	} else {
		result.Status = true
		result.Data = sysMsg
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title 获取转账交易消息详情
// @Description 获取转账交易消息详情
// @Param	msgId		path 	string	true		"消息记录ID"
// @Success 200 {object} model.JpushMsgResult
// @Failure 403 :消息记录ID不能为空
// @router /queryJpushMsgInfo/:msgId [get]
func (o *WalletController) QueryJpushMsgInfo() {
	msgID := o.Ctx.Input.Param(":msgId")
	var result model.JpushMsgResult
	if msgID == "" {
		result.Status = false
		result.Msg = o.Tr("id_empty")
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	sysMsg, err := model.GetJpushMessage(msgID)
	if err != nil {
		result.Status = false
		result.Msg = o.Tr("not_found_msg")
	} else {
		result.Status = true
		result.Data = sysMsg
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title queryVestingInfo
// @Description 获取股权信息
// @Param	address		formData 	string	true		"钱包地址 13ZtotAWCHeUVY6aUF46B8D992ZUvUEfex"
// @Param	pubKey		formData 	string	true		"钱包地址对应的公钥 0208f27b2189139eba610afaa7a4d48573a0db2d863651b791185eea882df06641"
// @Param	originData	formData 	string	true		"签名原始数据 {&quot;address&quot;:&quot;13ZtotAWCHeUVY6aUF46B8D992ZUvUEfex&quot;,&quot;pubKey&quot;:&quot;0208f27b2189139eba610afaa7a4d48573a0db2d863651b791185eea882df06641&quot;}"
// @Param	signature	formData 	string	true		"签名数据 304402200a63915578b4355926458bcb9ca69d7ecbbfb8813abedc28a629f34bc79f4f2c02201008dff2b72851ab0d749dced0f998a73524b14eb9e392f71cce9f4da5c408a9"
// @Success 200 {object} model.VestingListResult
// @Failure 403 :需要用地址对应的私钥对数据进行签名
// @router /queryVestingInfo [post]
func (o *WalletController) QueryVestingInfo() {
	address := o.GetString("address")
	pubKey := o.GetString("pubKey")
	originData := o.GetString("originData")
	signature := o.GetString("signature")
	service.QueryVestingInfo(address, pubKey, originData, signature, func(result model.VestingListResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryMasterTokenInfo
// @Description 获取主币信息
// @Success 200 {object} model.TokenInfoResult
// @router /queryMasterTokenInfo [post]
func (o *WalletController) QueryMasterTokenInfo() {
	service.QueryMasterTokenInfo(func(result model.TokenInfoResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryTokenInfo
// @Description 获取主币信息
// @Param	tokenId		path 	string	true		"token标识或ID 79866570fea511e8a3cd2f597d577af3"
// @Success 200 {object} model.TokenInfoResult
// @router /queryTokenInfo/:tokenId [get]
func (o *WalletController) QueryTokenInfo() {
	tokenID := o.Ctx.Input.Param(":tokenId")
	service.QueryTokenInfo(tokenID, func(result model.TokenInfoResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryAllTokens
// @Description 获取主币信息
// @Success 200 {object} model.TokenListResult
// @router /queryAllTokens [get]
func (o *WalletController) QueryAllTokens() {
	var result model.TokenListResult
	result.Status = true
	tokens := model.QueryAllTokens()
	for i := range tokens {
		number := common.HexToBigInt(tokens[i].TotalNumber).String()
		tokens[i].TotalNumber = common.FloatNumber(number, tokens[i].DecimalUnits)
	}
	result.Data = tokens
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title queryTransferDetails
// @Description 获取转账交易详情
// @Param	txId		formData 	string	true		"转账交易ID c4c36a80c23011e8a1d1edc8d3951e4a"
// @Success 200 {object} model.TransResult
// @Failure 403 :转账交易ID不能为空
// @router /queryTransferDetails [post]
func (o *WalletController) QueryTransferDetails() {
	txID := o.GetString("txId")
	trans, err := model.QueryTransferDetails(txID)
	var result model.TransResult
	if err != nil || trans == nil {
		service.GetTransferDetail(txID, func(ret *model.Transfer) {
			if ret == nil {
				result.Status = false
				result.Msg = o.Tr("not_found_msg")
				o.Data["json"] = result
				o.ServeJSON()
				return
			}
			var endTrans model.Transaction
			endTrans.Id = 0
			endTrans.BlockId = ""
			endTrans.TxHash = ""
			endTrans.TxId = txID
			endTrans.TokenID = ret.TokenID
			endTrans.FromAddress = ret.FromAddress
			endTrans.ToAddress = ret.ToAddress

			endTrans.IsCost = 1
			endTrans.TxTime = common.FormatDate(ret.Time)
			endTrans.Nonce = ret.Nonce
			endTrans.State = ret.State
			endTrans.Notes = ret.Notes
			endTrans.Msg = ret.Msg

			token, _ := model.QueryTokenInfo(ret.TokenID)
			if token != nil {
				endTrans.TokenSymbol = token.TokenSymbol
				endTrans.Number = common.FloatNumber(ret.Number, token.DecimalUnits)

				units, _ := model.QueryMasterTokenDecimalUnits()
				digits := common.GetNumberWithDigits(units)

				bigGas := new(big.Int)
				bigGas.SetString(ret.GasUsed, 10)

				endTrans.Fee = common.BigIntDiv(bigGas, digits)
			}
			result.Status = true
			result.Data = &endTrans
			o.Data["json"] = result
			o.ServeJSON()
		})
	} else {
		token, _ := model.QueryTokenInfo(trans.TokenID)
		units, _ := model.QueryMasterTokenDecimalUnits()
		digits := common.GetNumberWithDigits(units)

		bigGas := new(big.Int)
		bigGas.SetString(trans.Fee, 10)
		if token != nil {
			trans.Number = common.FloatNumber(trans.Number, token.DecimalUnits)
			trans.Fee = common.BigIntDiv(bigGas, digits)
		}
		result.Status = true
		result.Data = trans
		o.Data["json"] = result
		o.ServeJSON()
	}
}

// @Title queryTransferListByToken
// @Description 获取某一地址下某一token的转账交易记录
// @Param	address		formData 	string	true	"钱包地址 17TJQnX9ytdfMcLMLdHBpm3geBARf2Y4Vc"
// @Param	tokenId		formData 	string	true	"token标识 3e64c9d0c45f11e8b98c29d78a8d37dd"
// @Param	page		formData 	int	true		"第几页"
// @Param	pageSize	formData 	int	true		"每页多少条记录""
// @Success 200 {object} model.TransListResult
// @router /queryTransferListByToken [post]
func (o *WalletController) QueryTransferListByToken() {
	address := o.GetString("address")
	tokenID := o.GetString("tokenId")
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.QueryTransferListByToken(address, tokenID, pageSize, page)

	resultMap["status"] = true
	o.Data["json"] = resultMap

	o.ServeJSON()
}

// @Title queryTransferListByAddress
// @Description 查询两钱包地址之间的转账交易
// @Param	address1    formData 	string	true	"钱包地址1 17TJQnX9ytdfMcLMLdHBpm3geBARf2Y4Vc"
// @Param	address2    formData 	string	true	"钱包地址2 13mK4HGj3HonJniJhsHcK8Fce4RFNXw9Ya"
// @Param	tokenId		formData 	string	false	"token标识 3e64c9d0c45f11e8b98c29d78a8d37dd"
// @Param	page		formData 	int	true		"第几页"
// @Param	pageSize	formData 	int	true		"每页多少条记录""
// @Success 200 {object} model.TransListResult
// @router /queryTransferListByAddress [post]
func (o *WalletController) QueryTransferListByAddress() {
	address1 := o.GetString("address1")
	address2 := o.GetString("address2")
	tokenID := o.GetString("tokenId")
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.QueryTransferListByAddress(address1, address2, tokenID, pageSize, page)
	resultMap["status"] = true
	o.Data["json"] = resultMap

	o.ServeJSON()
}

// @Title queryTransferList
// @Description 综合查询交易信息
// @Param	address     formData 	string	true	"钱包地址1 17TJQnX9ytdfMcLMLdHBpm3geBARf2Y4Vc"
// @Param	year        formData 	string	false	"年份 2018"
// @Param	startYM     formData 	string	false	"开始年月 2018-09"
// @Param	endYM       formData 	string	false	"截止年月 2018-12"
// @Param	tokenId		formData 	string	false	"token标识 3e64c9d0c45f11e8b98c29d78a8d37dd"
// @Param	page		formData 	int	true		"第几页"
// @Param	pageSize	formData 	int	true		"每页多少条记录""
// @Success 200 {object} model.TransListResult
// @router /queryTransferList [post]
func (o *WalletController) QueryTransferList() {
	address := o.GetString("address")
	year := o.GetString("year")
	startYM := o.GetString("startYM")
	endYM := o.GetString("endYM")
	tokenID := o.GetString("tokenId")
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.QueryTransferList(address, year, startYM, endYM, tokenID, pageSize, page)
	resultMap["status"] = true
	o.Data["json"] = resultMap

	o.ServeJSON()
}

// @Title queryTransferStatistics
// @Description 查询交易统计信息
// @Param	address     formData 	string	true	"钱包地址1 17TJQnX9ytdfMcLMLdHBpm3geBARf2Y4Vc"
// @Param	year        formData 	string	false	"年份 2018"
// @Param	startYM     formData 	string	false	"开始年月 2018-09"
// @Param	endYM       formData 	string	false	"截止年月 2018-12"
// @Param	tokenId		formData 	string	false	"token标识 3e64c9d0c45f11e8b98c29d78a8d37dd"
// @Success 200 {object} model.TransStatisticsResult
// @router /queryTransferStatistics [post]
func (o *WalletController) QueryTransferStatistics() {
	address := o.GetString("address")
	year := o.GetString("year")
	startYM := o.GetString("startYM")
	endYM := o.GetString("endYM")
	tokenID := o.GetString("tokenId")
	resultMap := model.QueryTransferStatistics(address, year, startYM, endYM, tokenID)
	ret := make(map[string]interface{}, 0)
	ret["status"] = true
	ret["data"] = resultMap

	o.Data["json"] = ret
	o.ServeJSON()
}

// @Title queryReturnNumOfWait
// @Description 待返还的数量
// @Param	address		formData 	string	true		"钱包地址"
// @Param	tokenId		formData 	string	true		"token标识或ID"
// @Success 200 {object} model.JsonResult
// @Failure 403 :钱包地址或token标识不能为空
// @router /queryReturnNumOfWait [post]
func (o *WalletController) QueryReturnNumOfWait() {
	address := o.GetString("address")
	tokenID := o.GetString("tokenId")

	service.VestingsBalanceCount(address, tokenID, func(result model.JsonResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title queryReturnItemsOfWait
// @Description 待返还的明细记录
// @Param	address		formData 	string	true		"钱包地址"
// @Param	tokenId		formData 	string	true		"token标识或ID"
// @Success 200 {object} model.ReturnListResult
// @Failure 403 :钱包地址或token标识不能为空
// @router /queryReturnItemsOfWait [post]
func (o *WalletController) QueryReturnItemsOfWait() {
	address := o.GetString("address")
	tokenID := o.GetString("tokenId")

	service.VestingBalanceItems(address, tokenID, o.Lang, func(result model.ReturnListResult) {
		o.Data["json"] = result
		o.ServeJSON()
	})
}

// @Title 发行token所需的平台币标准
// @Description 发行token所需的平台币标准
// @Success 200 {object} model.JsonResult
// @router /queryPublishTokenRequireNum [get]
func (o *WalletController) QueryPublishTokenRequireNum() {
	service.QueryTokenRequireNum(func(ret model.JsonResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title 发布合约所需的平台币标准
// @Description 发布合约所需的平台币标准
// @Success 200 {object} model.JsonResult
// @router /queryPublishCCRequireNum [get]
func (o *WalletController) QueryPublishCCRequireNum() {
	service.QueryCCRequireNum(func(ret model.JsonResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title 手续费返还规则
// @Description 手续费返还规则
// @Success 200 {object} model.ReturnGasConfigResponse
// @router /queryReturnGasConfig [get]
func (o *WalletController) QueryReturnGasConfig() {
	service.QueryReturnGasConfig(func(ret model.ReturnGasConfigResponse) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title queryContractList
// @Description 查询某个地址下的合约信息列表
// @Param	address		path 	string	true		"钱包地址或合约地址"
// @Success 200 {object} model.ContractListResult
// @Failure 403 :钱包地址和合约地址必须二选其一
// @router /queryContractList/:address [get]
func (o *WalletController) QueryContractList() {
	address := o.Ctx.Input.Param(":address")
	var result model.ContractListResult
	if address == "" {
		result.Status = false
		result.Msg = o.Tr("address_contractaddr_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.QueryContractList(address, func(ret model.ContractListResult) {
		if len(ret.Data) > 0 {
			for i := range ret.Data {
				ret.Data[i].CreateTime = common.FormatDate(ret.Data[i].CreateTime)
				ret.Data[i].UpdateTime = common.FormatDate(ret.Data[i].UpdateTime)
			}
		}
		o.Data["json"] = ret
		o.ServeJSON()
	})
}

// @Title queryContractInfo
// @Description 查询合约的详情
// @Param	address		path 	string	true		"合约地址"
// @Param	version		path 	string	false		"合约版本号"
// @Success 200 {object} model.ContractResult
// @Failure 403 :合约地址不能为空
// @router /queryContractInfo/:address/:version [get]
func (o *WalletController) QueryContractInfo() {
	address := o.Ctx.Input.Param(":address")
	version := o.Ctx.Input.Param(":version")
	var result model.ContractResult
	if address == "" {
		result.Msg = o.Tr("contractaddr_empty")
		o.Data["json"] = result
		o.ServeJSON()
	}
	service.QueryContractInfo(address, version, func(ret model.ContractResult) {
		o.Data["json"] = ret
		o.ServeJSON()
	})
}
