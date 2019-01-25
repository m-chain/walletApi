package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"walletApi/src/common"
	"walletApi/src/model"

	"github.com/beego/i18n"

	"github.com/astaxie/beego"
)

type Result struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func parseResult(ret Result) model.JsonResult {
	var jsonResult model.JsonResult
	if ret.Status == "true" {
		jsonResult.Status = true
		jsonResult.Data = ret.Data
	} else {
		jsonResult.Status = false
		jsonResult.Msg = ret.Msg
	}
	return jsonResult
}

//获取区块信息
func SyncBlockInfo(blockNum int64) error {
	fmt.Printf("blockNum=%d\n", blockNum)
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/blocks/%d", blockNum)

	bytes, err := common.SyncHttpGet(chainurl, "en")
	if err != nil {
		return err
	}

	var blockResult model.BlockResult
	err = json.Unmarshal(bytes, &blockResult)
	if err != nil {
		return err
	}
	if blockResult.Status == false {
		return fmt.Errorf("未找到对应的区块信息")
	}
	var arr []model.Transaction
	for _, v := range blockResult.Data.Transactions {
		txhash := v.Txhash
		for _, t := range v.Trans {
			t.TxHash = txhash
			t.BlockId = blockResult.Data.Blockid
			if t.Fee != "" && t.Fee != "0" {
				t.IsCost = 1
			}

			//获取交易状态
			t.State = queryTransferStatus(t.TxId)
			arr = append(arr, t)
		}

		for _, token := range v.Tokens {
			token.OwnerAddress = v.TokenMaster[token.TokenID]
			model.AddToken(&token)
		}
	}
	//往数据库中插入交易信息
	for _, v := range arr {
		model.AddTransaction(v)
	}

	if len(arr) > 0 {
		//设置token名称
		model.UpdateTokenSymbol()
	}
	return nil
}

//更新转账交易状态
func SyncTransferStatus(callback func(error)) error {
	txIDs := model.QueryTransfersNoStatus()
	if len(txIDs) > 0 {
		data := map[string]interface{}{
			"data": txIDs,
		}
		chainurl := beego.AppConfig.String("chainurl") + "/queryMultiTransferStatus"
		bytes, _ := common.SyncHttpPost(chainurl, "en", data)

		if len(bytes) > 0 {

			type Result struct {
				Status bool
				Msg    string
				Data   []model.TxStatusItem
			}

			var ret Result
			err := json.Unmarshal(bytes, &ret)
			if err != nil {
				callback(err)
				return err
			}
			if ret.Status == false {
				callback(fmt.Errorf(ret.Msg))
				return err
			}
			if len(ret.Data) > 0 {
				model.UpdateMultiTransferStatus(ret.Data)
			}
		}
	}

	return nil
}

//钱包是否已经存在帐本中
func WalletIsExist(address string, callback func(model.JsonResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/walletIsExist/%s", address)
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.JsonResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//设置公钥
func UpdatePubkey(rawData string, callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/updateWalletPubkey"
	data := map[string]interface{}{
		"rawData": rawData,
	}
	common.HttpPost(chainurl, "en", data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.JsonResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//获取签名情况
func QueryConfirmInfo(tokenID, confirmData string, callback func(model.ConfirmResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/queryConfirmInfo"
	data := map[string]interface{}{
		"tokenID": tokenID,
		"data":    confirmData,
	}
	common.HttpPost(chainurl, "en", data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ConfirmResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//查询某一币种的待返还数量
func VestingsBalanceCount(address, tokenID string, callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/vestingsBalance"
	data := map[string]interface{}{
		"address": address,
		"tokenID": tokenID,
	}
	common.HttpPost(chainurl, "en", data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonRet model.JsonResult
				err := json.Unmarshal(data, &jsonRet)
				if err != nil {
					jsonRet.Status = false
					jsonRet.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonRet.Status == true && jsonRet.Data != "0" {
					jsonRet.Data = model.FloatNumber(jsonRet.Data, tokenID)
				}
				callback(jsonRet)
			}
		}
	})
}

//查询某一币种的待返还数量
func VestingBalanceItems(address, tokenID, lang string, callback func(model.ReturnListResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/vestingsBalanceDetail"
	data := map[string]interface{}{
		"address": address,
		"tokenID": tokenID,
	}
	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var ret model.VestingListResult
				var jsonRet model.ReturnListResult
				err := json.Unmarshal(data, &ret)
				if err != nil {
					jsonRet.Status = false
					jsonRet.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				var arr []model.VestingItem
				if ret.Data != nil && len(ret.Data) > 0 {
					token, _ := model.QueryTokenInfo(tokenID)
					for _, v := range ret.Data {
						var item model.VestingItem
						item.ID = v.ID
						item.TokenID = v.TokenID
						item.Periods = v.Periods
						item.VType = v.VType
						if v.VType == 0 {
							item.VTypeName = i18n.Tr(lang, "tx_presale_type")
						} else if v.VType == 1 {
							item.VTypeName = i18n.Tr(lang, "tx_init_cc_type")
						} else if v.VType == 2 {
							item.VTypeName = i18n.Tr(lang, "tx_upgrade_cc_type")
						} else if v.VType == 3 {
							item.VTypeName = i18n.Tr(lang, "tx_invoke_cc_type")
						} else if v.VType == 4 {
							item.VTypeName = i18n.Tr(lang, "tx_isuue_token_type")
						}
						cmpRet := common.TimestampCmp(fmt.Sprintf("%v", v.StartTime))
						if cmpRet == 1 {
							item.Status = i18n.Tr(lang, "will_return")
						} else {
							item.Status = i18n.Tr(lang, "returning")
						}

						amount := common.HexToBigInt(v.Amount)
						item.Amount = common.FloatNumber(amount.String(), token.DecimalUnits)

						initReleaseAmount := common.HexToBigInt(v.InitReleaseAmount)
						item.InitReleaseAmount = common.FloatNumber(initReleaseAmount.String(), token.DecimalUnits)

						withdrawed := common.HexToBigInt(v.Withdrawed)
						item.Withdrawed = common.FloatNumber(withdrawed.String(), token.DecimalUnits)

						ip := big.NewInt(0)
						iAmount := ip.Sub(amount, initReleaseAmount).Div(ip, big.NewInt(v.Periods))
						item.PeriodNum = common.FloatNumber(iAmount.String(), token.DecimalUnits)

						item.EffectTime = common.TimestampToStr(v.StartTime)
						item.CreateTime = v.CreateTime

						item.Interval = common.ConverTimestamp(v.Interval, lang)

						arr = append(arr, item)
					}
				}
				jsonRet.Status = true
				jsonRet.Data = arr

				callback(jsonRet)
			}
		}
	})
}

//查询余额
func QueryBalance(address string, callback func(model.WalletResp)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/queryBalance/%s", address)
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.WalletResp
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				var tokenIds []string
				for _, rest := range jsonResult.Data.WalletRests {
					tokenIds = append(tokenIds, rest.TokenID)
				}

				if len(tokenIds) > 0 {
					tokenMap, _ := model.QueryTokenDecimalUnits(tokenIds)
					nameMap, _ := model.QueryTokenNames(tokenIds)
					for i := range jsonResult.Data.WalletRests {
						decimalUnits := tokenMap[jsonResult.Data.WalletRests[i].TokenID]
						name := nameMap[jsonResult.Data.WalletRests[i].TokenID]
						jsonResult.Data.WalletRests[i].Name = fmt.Sprintf("%v", name)
						units, _ := strconv.Atoi(fmt.Sprintf("%v", decimalUnits))

						digits := common.GetNumberWithDigits(units)
						restNum := new(big.Int)
						restNum.SetString(jsonResult.Data.WalletRests[i].RestNumber, 10)

						freezeNum := new(big.Int)
						freezeNum.SetString(jsonResult.Data.WalletRests[i].FreezeNumber, 10)

						jsonResult.Data.WalletRests[i].RestNumber = common.FormatNumber(common.BigIntDiv(restNum, digits))
						jsonResult.Data.WalletRests[i].FreezeNumber = common.FormatNumber(common.BigIntDiv(freezeNum, digits))
					}
				}

				if len(jsonResult.Data.WalletRests) == 0 {
					jsonResult.Data.WalletRests = make([]model.WalletRest, 0)
				}

				callback(jsonResult)
			}
		}
	})
}

//获取某地址下的转账次数
func QueryTransferCount(address string, callback func(model.JsonResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/queryTransferCount/%s", address)
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				type TempResult struct {
					Status bool   `json:"status"`
					Msg    string `json:"msg"`
					Count  int64  `json:"count"`
				}
				var tempResult TempResult
				err := json.Unmarshal(data, &tempResult)

				var jsonResult model.JsonResult
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				jsonResult.Status = tempResult.Status
				jsonResult.Msg = tempResult.Msg
				jsonResult.Data = fmt.Sprintf("%d", tempResult.Count)
				callback(jsonResult)
			}
		}
	})
}

//获取手续费
func QueryChargeGas(callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/queryChargeGas"
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				type TempResult struct {
					Status       bool   `json:"status"`
					Msg          string `json:"msg"`
					Gas          string `json:"gas"`
					DecimalUnits string `json:"decimalUnits"`
				}
				var tempResult TempResult
				err := json.Unmarshal(data, &tempResult)

				var jsonResult model.JsonResult
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				jsonResult.Status = tempResult.Status
				jsonResult.Msg = tempResult.Msg

				units, _ := strconv.Atoi(tempResult.DecimalUnits)
				digits := common.GetNumberWithDigits(units)

				bigGas := new(big.Int)
				bigGas.SetString(tempResult.Gas, 10)
				jsonResult.Data = common.BigIntDiv(bigGas, digits)

				callback(jsonResult)
			}
		}
	})
}

//一对一转账
func Transfer(originStr, signature, lang string, callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/sendRawTransaction"
	data := map[string]interface{}{
		"origin":    originStr,
		"signature": signature,
	}
	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				type TransferResp struct {
					Status bool   `json:"status"`
					Msg    string `json:"msg"`
					TxID   string `json:"txId"`
				}
				var ret TransferResp
				var jsonRet model.JsonResult
				err := json.Unmarshal(data, &ret)
				if err != nil {
					jsonRet.Status = false
					jsonRet.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				} else {
					if ret.Status == true { //转账成功
						jsonRet.Status = true
						jsonRet.Data = ret.TxID

						var originObj map[string]interface{}
						json.Unmarshal([]byte(originStr), &originObj)

						tokenID := fmt.Sprintf("%v", originObj["tokenID"])
						from := fmt.Sprintf("%v", originObj["fromAddress"])
						to := fmt.Sprintf("%v", originObj["toAddress"])
						number := fmt.Sprintf("%v", originObj["number"])
						token, _ := model.QueryTokenInfo(tokenID)
						var msgObj model.JpushMessage
						msgObj.Alert = i18n.Tr(lang, "transfer_tips")
						msgObj.MsgType = 1
						msgObj.IsRead = 2
						msgObj.TxId = ret.TxID
						msgObj.Target = to
						msgObj.Title = i18n.Tr(lang, "transfer_success")

						msgObj.Content = i18n.Tr(lang, "transfer_notice", from, number, token.TokenSymbol)

						err = model.AddJpushMsg(&msgObj)
						if err == nil {
							msgBytes, err := makeMessage(msgObj)
							if err == nil {
								jpushurl := beego.AppConfig.String("jpushurl")
								common.SendMessage(jpushurl, msgBytes)
							}
						}
					} else {
						jsonRet.Status = ret.Status
						jsonRet.Msg = ret.Msg
					}
				}
				callback(jsonRet)
			}
		}
	})
}

func makeMessage(msgObj model.JpushMessage) ([]byte, error) {
	notice := make(map[string]interface{})
	tags := make(map[string]([]string))
	tags["tag"] = []string{msgObj.Target}
	notice["platform"] = []string{"android", "ios"}
	notice["audience"] = tags

	notification := make(map[string]interface{})
	android := make(map[string]interface{})
	ios := make(map[string]interface{})
	extras := make(map[string]interface{})

	extras["msgid"] = msgObj.MsgId
	extras["txId"] = msgObj.TxId

	android["alert"] = msgObj.Alert
	android["title"] = msgObj.Title
	android["builder_id"] = 1 //通知栏样式
	android["extras"] = extras

	ios["alert"] = msgObj.Alert
	ios["sound"] = "default"
	ios["badge"] = "+1"
	ios["extras"] = extras

	notification["android"] = android
	notification["ios"] = ios
	notice["notification"] = notification

	message := make(map[string]interface{}) //消息内容体。是被推送到客户端的内容。与 notification 一起二者必须有其一，可以二者并存
	//msgExtras := make(map[string]interface{})
	//msgExtras["key"] = "value"

	message["msg_content"] = msgObj.Content
	message["content_type"] = "text"
	message["title"] = msgObj.Title
	//message["extras"] = msgExtras

	notice["message"] = message

	options := make(map[string]interface{})
	options["time_to_live"] = 86400    //推送当前用户不在线时，为该用户保留多长时间的离线消息，保留１天
	options["apns_production"] = false //True 表示推送生产环境，False 表示要推送开发环境；如果不指定则为推送生产环境
	//options["apns_collapse_id"] = "jiguang_test_201706011100" //ios可以覆盖原有的消息

	notice["options"] = options
	return json.Marshal(notice)
}

//查询某地址下的股权信息
func QueryVestingInfo(address, pubKey, originData, signature string, callback func(model.VestingListResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/queryVestingInfo"
	//data := url.Values{"address": {address}, "pubKey": {pubKey}, "originData": {originData}, "signature": {signature}}
	data := map[string]interface{}{
		"address":   address,
		"pubKey":    pubKey,
		"origin":    originData,
		"signature": signature,
	}
	common.HttpPost(chainurl, "en", data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.VestingListResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("json解析出错,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//查询主币信息
func QueryMasterTokenInfo(callback func(model.TokenInfoResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/queryMasterTokenInfo"
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.TokenInfoResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonResult.Status == true {
					totalNumber := common.HexToBigInt(jsonResult.Data.TotalNumber).String()
					restNumber := common.HexToBigInt(jsonResult.Data.RestNumber).String()
					decimalUnits := jsonResult.Data.DecimalUnits
					jsonResult.Data.TotalNumber = common.FloatNumber(totalNumber, decimalUnits)
					jsonResult.Data.RestNumber = common.FloatNumber(restNumber, decimalUnits)
				}
				callback(jsonResult)
			}
		}
	})
}

//查询token信息
func QueryTokenInfo(tokenID string, callback func(model.TokenInfoResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/queryTokenInfo/" + tokenID
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.TokenInfoResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				} else {
					token, _ := model.QueryTokenInfo(tokenID)
					if token != nil {
						jsonResult.Data.IconUrl = token.IconUrl
					}
					jsonResult.Data.TotalNumber = common.HexToBigInt(jsonResult.Data.TotalNumber).String()
					jsonResult.Data.RestNumber = common.HexToBigInt(jsonResult.Data.RestNumber).String()
					jsonResult.Data.TotalNumber = common.FloatNumber(jsonResult.Data.TotalNumber, jsonResult.Data.DecimalUnits)
					jsonResult.Data.RestNumber = common.FloatNumber(jsonResult.Data.RestNumber, jsonResult.Data.DecimalUnits)
				}
				callback(jsonResult)
			}
		}
	})
}

//根据签名的结果调用底层链接口
func InvokeChainCC(signData *model.SignData, lang string, callback func(model.CCInvokeResult)) {
	rSign, err := model.GetSingInfo(signData.QrCode)

	var jsonResult model.CCInvokeResult
	if err != nil || rSign == nil {
		jsonResult.Status = false
		jsonResult.Msg = i18n.Tr(lang, "not_found_qrcode")
		callback(jsonResult)
		return
	}
	if rSign.Status == 3 { //1、待签名 2、已签名 3、已确认 -1、已过期
		jsonResult.Status = false
		jsonResult.Msg = i18n.Tr(lang, "repeat_operation")
		callback(jsonResult)
		return
	}
	signData.OriginData = rSign.OriginData
	signData.SignType = rSign.SignType

	if rSign.LangType != "" {
		lang = rSign.LangType
	}

	if rSign.SignType == 1 { //发布合约
		deployeeCC(signData, lang, callback)
	} else if rSign.SignType == 2 { //初始化合约
		initCC(signData, lang, callback)
	} else if rSign.SignType == 3 { //升级合约
		deployeeCC(signData, lang, callback)
	} else if rSign.SignType == 4 { //发行token
		issueToken(signData, lang, callback)
	} else if rSign.SignType == 5 { //设置master、manager
		setMasterAndManagers(signData, lang, callback)
	} else if rSign.SignType == 6 { //manager签名确认
		managerExec("udo_confirm", lang, signData, callback)
	} else if rSign.SignType == 7 { //master
		masterExecAction("udo_addManager", lang, signData, callback)
	} else if rSign.SignType == 8 { //master
		masterExecAction("udo_replaceManager", lang, signData, callback)
	} else if rSign.SignType == 9 { //master
		masterExecAction("udo_removeManager", lang, signData, callback)
	} else if rSign.SignType == 10 { //设置manager操作所需的确认数阀值
		masterExecAction("setMajorityThreshold", lang, signData, callback)
	} else if rSign.SignType == 11 { //设置发行token所需要的平台币 manager
		regularOfMasterToken("udo_publishTokenRequireNum", lang, signData, callback)
	} else if rSign.SignType == 12 { //设置发布合约所需要的平台币 manager
		regularOfMasterToken("udo_publishCCRequireNum", lang, signData, callback)
	} else if rSign.SignType == 13 { //设置手续费返还规则 manager
		regularOfMasterToken("udo_returnGasConfig", lang, signData, callback)
	} else if rSign.SignType == 14 { //删除合约
		deleteCC(signData, lang, callback)
	} else if rSign.SignType == 15 { //设置master操作所需的确认数阀值
		masterExecAction("setMasterThreshold", lang, signData, callback)
	} else if rSign.SignType == 16 { //设置token图标
		updateTokenIcon(signData, lang, callback)
	}
}

//master执行，manager签名确认
func masterExecAction(method, lang string, signData *model.SignData, callback func(model.CCInvokeResult)) {
	type Origin struct {
		SignType int
		TokenID  string
	}
	var origin Origin
	originBytes, _ := hex.DecodeString(signData.OriginData)
	json.Unmarshal(originBytes, &origin)
	address := signData.Address
	tokenID := origin.TokenID
	chainurl := beego.AppConfig.String("chainurl") + "/isMaterOrManagerOfAddr/" + address + "/" + tokenID
	common.HttpGet(chainurl, lang, func(result interface{}, err error) {
		if err == nil && result != nil {

			if data, ok := result.([]byte); ok {
				var ret model.JsonResult
				err := json.Unmarshal(data, &ret)
				var jsonResult model.CCInvokeResult
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s", err.Error())
					callback(jsonResult)
					return
				}
				if ret.Status == false {
					jsonResult.Status = false
					jsonResult.Msg = ret.Msg
					callback(jsonResult)
					return
				}

				if ret.Data == "3" { //非master或manager
					jsonResult.Status = false
					jsonResult.Msg = i18n.Tr(lang, "no_auth_operation")
					callback(jsonResult)
					return
				} else if ret.Data == "1" { //master execute
					managerExec(method, lang, signData, callback)
				} else if ret.Data == "2" { //manager confirm
					managerExec("udo_confirm", lang, signData, callback)
				}
			}
		}
	})
}

func isMajorityConfirmed(signData *model.SignData) model.CCInvokeResult {
	chainurl := beego.AppConfig.String("chainurl") + "/isMajorityConfirmed"
	var result model.CCInvokeResult
	result.Status = false
	params := map[string]interface{}{
		"origin": signData.OriginData,
	}
	bytes, err := common.SyncHttpPost(chainurl, "en", params)
	if err != nil {
		result.Msg = err.Error()
		return result
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		result.Msg = fmt.Sprintf("JSON parsing error,%s！", string(bytes))
		return result
	}
	return result
}

//manager签名确认或执行
func regularOfMasterToken(method, lang string, signData *model.SignData, callback func(model.CCInvokeResult)) {
	result := isMajorityConfirmed(signData)
	if result.Status == true { //达到确认的阀值
		managerExec(method, lang, signData, callback)
	} else { //manager签名确认
		if signData.SignType == 11 || signData.SignType == 12 || signData.SignType == 13 {
			managerExec("udo_confirm", lang, signData, func(result model.CCInvokeResult) {
				if result.Status == true {
					ret := isMajorityConfirmed(signData)
					if ret.Status == true {
						managerExec(method, lang, signData, callback)
					} else {
						callback(result)
					}
				} else {
					callback(result)
				}
			})
		} else {
			managerExec("udo_confirm", lang, signData, callback)
		}
	}
}

//manager签名确认
func managerExec(method, lang string, signData *model.SignData, callback func(model.CCInvokeResult)) {

	chainurl := beego.AppConfig.String("chainurl") + "/" + method
	var jsonResult model.CCInvokeResult
	//data := url.Values{"address": {signData.Address}, "pubKey": {signData.PubKey}, "origin": {signData.OriginData}, "signature": {signData.SignData}}
	data := map[string]interface{}{
		"address":   signData.Address,
		"pubKey":    signData.PubKey,
		"origin":    signData.OriginData,
		"signature": signData.SignData,
	}

	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {

				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonResult.Status == false {
					signData.Status = 2 //已签名
					signData.RespResult = string(data)
					model.UpdateSignData(signData)
				} else {
					if method == "udo_confirm" {
						jsonResult.Msg = i18n.Tr(lang, "tip_sign_success")
					} else {
						if signData.SignType == 6 || signData.SignType == 11 || signData.SignType == 12 || signData.SignType == 13 {
							jsonResult.Msg = i18n.Tr(lang, common.SignTips[signData.SignType].SuccessMsg)
						} else if signData.SignType == 7 || signData.SignType == 8 || signData.SignType == 9 || signData.SignType == 10 || signData.SignType == 15 {
							jsonResult.Msg = i18n.Tr(lang, common.SignTips[signData.SignType].SuccessMsg)
						}
					}
					signData.Status = 3 //已确认且最终得以执行
					jsonBytes, _ := json.Marshal(jsonResult)
					signData.RespResult = string(jsonBytes)
					model.UpdateSignData(signData)
				}
				callback(jsonResult)
			}
		}
	})
}

//设置master、manager
func setMasterAndManagers(signData *model.SignData, lang string, callback func(model.CCInvokeResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_provideAuthority"

	//data := url.Values{"address": {signData.Address}, "pubKey": {signData.PubKey}, "origin": {signData.OriginData}, "signature": {signData.SignData}}
	data := map[string]interface{}{
		"address":   signData.Address,
		"pubKey":    signData.PubKey,
		"origin":    signData.OriginData,
		"signature": signData.SignData,
	}
	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.CCInvokeResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonResult.Status == false {
					signData.Status = 2 //已签名
					signData.RespResult = string(data)
					model.UpdateSignData(signData)
				} else {
					jsonResult.Msg = i18n.Tr(lang, common.SignTips[signData.SignType].SuccessMsg)

					signData.Status = 3 //已确认且最终得以执行
					jsonBytes, _ := json.Marshal(jsonResult)
					signData.RespResult = string(jsonBytes)
					model.UpdateSignData(signData)
				}
				callback(jsonResult)
			}
		}
	})
}

//发行token
func issueToken(signData *model.SignData, lang string, callback func(model.CCInvokeResult)) {
	originBytes, _ := hex.DecodeString(signData.OriginData)
	var token model.Token
	err := json.Unmarshal(originBytes, &token)
	var jsonResult model.CCInvokeResult
	if err != nil {
		jsonResult.Status = false
		jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", err.Error())
		callback(jsonResult)
		return
	}
	token.IsBaseCoin = false
	token.Status = 1

	chainurl := beego.AppConfig.String("chainurl") + "/udo_issueToken"

	sendData := make(map[string]interface{}, 0)
	sendData["address"] = signData.Address
	sendData["pubKey"] = signData.PubKey
	sendData["origin"] = signData.OriginData
	sendData["signature"] = signData.SignData

	jsonBytes, _ := json.Marshal(&sendData)
	hexStr := hex.EncodeToString(jsonBytes)

	data := map[string]interface{}{
		"rawData": hexStr,
	}
	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonResult.Status == false {
					signData.Status = 2 //已签名
					signData.RespResult = string(data)
					model.UpdateSignData(signData)
				} else {
					jsonResult.Msg = i18n.Tr(lang, common.SignTips[signData.SignType].SuccessMsg)
					signData.Status = 3 //已确认且最终得以执行
					jsonBytes, _ := json.Marshal(jsonResult)
					signData.RespResult = string(jsonBytes)
					model.UpdateSignData(signData)

					//添加token
					token.TokenID = jsonResult.TokenId
					token.OwnerAddress = signData.Address
					model.AddToken(&token)
				}
				callback(jsonResult)
			}
		}
	})
}

//发布合约
func deployeeCC(signData *model.SignData, lang string, callback func(model.CCInvokeResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/cc_deploye_url"
	if signData.SignType == 3 {
		chainurl = beego.AppConfig.String("chainurl") + "/cc_upgrade_url"
	}
	sendData := make(map[string]interface{}, 0)
	sendData["address"] = signData.Address
	sendData["pubKey"] = signData.PubKey
	sendData["origin"] = signData.OriginData
	sendData["signature"] = signData.SignData

	var jsonResult model.CCInvokeResult
	/*
		originBytes, _ := hex.DecodeString(signData.OriginData)
		var contract model.ContractRequest
		json.Unmarshal(originBytes, &contract)

		ccBuff, err := common.SyncDownloadFile(contract.CcUrl)

		if err != nil || len(ccBuff) == 0 {
			jsonResult.Status = false
			jsonResult.Msg = fmt.Sprintf("获取合约源代码出错,%s！", err.Error())
			callback(jsonResult)
			return
		}
		//base64编码
		encodeString := base64.StdEncoding.EncodeToString(ccBuff)

		sendData["ccData"] = encodeString
	*/

	jsonBytes, _ := json.Marshal(&sendData)
	hexStr := hex.EncodeToString(jsonBytes)
	data := map[string]interface{}{
		"rawData": hexStr,
	}
	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonResult.Status == false {
					signData.Status = 2 //已签名
					signData.RespResult = string(data)
					model.UpdateSignData(signData)
				} else {
					jsonResult.Msg = i18n.Tr(lang, common.SignTips[signData.SignType].SuccessMsg)
					signData.Status = 3 //已确认且最终得以执行
					jsonBytes, _ := json.Marshal(jsonResult)
					signData.RespResult = string(jsonBytes)
					model.UpdateSignData(signData)
				}
				callback(jsonResult)
			}
		} else {
			jsonResult.Status = false
			jsonResult.Msg = fmt.Sprintf("Failure of platform chain publishing contract,%v！", err)
			callback(jsonResult)
		}
	})
}

//初始化合约
func initCC(signData *model.SignData, lang string, callback func(model.CCInvokeResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_cc_init"

	sendData := make(map[string]interface{}, 0)
	sendData["pubKey"] = signData.PubKey
	sendData["origin"] = signData.OriginData
	sendData["signature"] = signData.SignData

	jsonBytes, _ := json.Marshal(&sendData)
	hexStr := hex.EncodeToString(jsonBytes)

	data := map[string]interface{}{
		"rawData": hexStr,
	}
	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.CCInvokeResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonResult.Status == false {
					signData.Status = 2 //已签名
					signData.RespResult = string(data)
					model.UpdateSignData(signData)
				} else {
					jsonResult.Msg = i18n.Tr(lang, common.SignTips[signData.SignType].SuccessMsg)
					signData.Status = 3 //已确认且最终得以执行

					jsonBytes, _ := json.Marshal(jsonResult)
					signData.RespResult = string(jsonBytes)

					model.UpdateSignData(signData)
				}
				callback(jsonResult)
			}
		}
	})
}

//删除合约
func deleteCC(signData *model.SignData, lang string, callback func(model.CCInvokeResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_cc_delete"

	sendData := make(map[string]interface{}, 0)
	sendData["pubKey"] = signData.PubKey
	sendData["origin"] = signData.OriginData
	sendData["signature"] = signData.SignData

	jsonBytes, _ := json.Marshal(&sendData)
	hexStr := hex.EncodeToString(jsonBytes)
	data := map[string]interface{}{
		"rawData": hexStr,
	}
	common.HttpPost(chainurl, lang, data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.CCInvokeResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				if jsonResult.Status == false {
					signData.Status = 2 //已签名
					signData.RespResult = string(data)
					model.UpdateSignData(signData)
				} else {
					jsonResult.Msg = i18n.Tr(lang, common.SignTips[signData.SignType].SuccessMsg)
					signData.Status = 3 //已确认且最终得以执行

					jsonBytes, _ := json.Marshal(jsonResult)
					signData.RespResult = string(jsonBytes)

					model.UpdateSignData(signData)
				}
				callback(jsonResult)
			}
		}
	})
}

//获取转账交易详情
func GetTransferDetail(txID string, callback func(*model.Transfer)) {
	chainurl := beego.AppConfig.String("chainurl") + "/queryTransferDetails"
	data := map[string]interface{}{
		"txId": txID,
	}
	common.HttpPost(chainurl, "en", data, func(result interface{}, err error) {

		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var ret model.TransferDetail
				err := json.Unmarshal(data, &ret)
				if err != nil {
					callback(nil)
				} else {
					transfer := new(model.Transfer)
					for _, v := range ret.Data {
						if v.GasUsed != "" && v.GasUsed != "0" {
							transfer = &v
							break
						}
					}
					callback(transfer)
				}
			}
		}
	})
}

//设置token图标
func updateTokenIcon(signData *model.SignData, lang string, callback func(model.CCInvokeResult)) {

	var jsonResult model.CCInvokeResult
	if common.GetAddress(signData.PubKey) != signData.Address {
		jsonResult.Msg = i18n.Tr(lang, "illegal_use_address")
		callback(jsonResult)
		return
	}
	ok, _ := common.Verify(signData.PubKey, signData.OriginData, signData.SignData)
	if !ok {
		jsonResult.Msg = i18n.Tr(lang, "sign_data_failure")
		callback(jsonResult)
		return
	}
	type TokenIcon struct {
		TokenID string `json:"tokenID"`
		IconUrl string `json:"iconUrl"`
	}

	hexBytes, _ := hex.DecodeString(signData.OriginData)
	var tokenIcon TokenIcon
	json.Unmarshal(hexBytes, &tokenIcon)

	if tokenIcon.TokenID == "" {
		jsonResult.Msg = i18n.Tr(lang, "token_id_empty")
		callback(jsonResult)
		return
	}
	if tokenIcon.IconUrl == "" {
		jsonResult.Msg = i18n.Tr(lang, "token_icon_url_empty")
		callback(jsonResult)
		return
	}
	chainurl := beego.AppConfig.String("chainurl") + "/queryMasterAdressOfToken/" + tokenIcon.TokenID
	retBytes, _ := common.SyncHttpGet(chainurl, lang)
	if len(retBytes) == 0 {
		jsonResult.Msg = i18n.Tr(lang, "invalid_issue_token_address")
		callback(jsonResult)
		return
	}
	var retObj model.JsonResult
	json.Unmarshal(retBytes, &retObj)
	if retObj.Status == false {
		jsonResult.Msg = retObj.Msg
		callback(jsonResult)
		return
	}
	if retObj.Data != signData.Address {
		jsonResult.Msg = i18n.Tr(lang, "invalid_token_sign_address")
		callback(jsonResult)
		return
	}
	err := model.UpdateTokenIcon(tokenIcon.TokenID, tokenIcon.IconUrl)
	if err != nil {
		jsonResult.Msg = i18n.Tr(lang, "set_token_icon_fail")
		callback(jsonResult)
		return
	}
	jsonResult.Status = true
	jsonResult.TokenId = tokenIcon.TokenID
	jsonResult.Msg = i18n.Tr(lang, "set_token_con_success")

	jsonBytes, _ := json.Marshal(jsonResult)
	signData.RespResult = string(jsonBytes)
	signData.Status = 3 //已确认且最终得以执行

	model.UpdateSignData(signData)

	callback(jsonResult)
}

//获取发行token所需要的平台币标准
func QueryTokenRequireNum(callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_queryPublishTokenRequireNum"
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.JsonResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//获取发布合约所需要的平台币标准
func QueryCCRequireNum(callback func(model.JsonResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_queryPublishCCRequireNum"
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.JsonResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//获取手续费返还规则
func QueryReturnGasConfig(callback func(model.ReturnGasConfigResponse)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_queryReturnGasConfig"
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ReturnGasConfigResponse
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//根据地址查询合约列表
func QueryContractList(address string, callback func(model.ContractListResult)) {
	chainurl := beego.AppConfig.String("chainurl") + "/udo_cc_query"
	data := map[string]interface{}{
		"address": address,
	}
	common.HttpPost(chainurl, "en", data, func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ContractListResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//根据合约地址获取合约信息
func QueryContractInfo(address, version string, callback func(model.ContractResult)) {
	chainurl := fmt.Sprintf(beego.AppConfig.String("chainurl")+"/udo_cc_info/%s/%s", address, version)
	common.HttpGet(chainurl, "en", func(result interface{}, err error) {
		if err == nil && result != nil {
			if data, ok := result.([]byte); ok {
				var jsonResult model.ContractResult
				err := json.Unmarshal(data, &jsonResult)
				if err != nil {
					jsonResult.Status = false
					jsonResult.Msg = fmt.Sprintf("JSON parsing error,%s！", string(data))
				}
				callback(jsonResult)
			}
		}
	})
}

//查询交易的状态
func queryTransferStatus(txID string) int {
	chainurl := beego.AppConfig.String("chainurl") + "/queryTransferStatus/" + txID

	bytes, err := common.SyncHttpGet(chainurl, "en")
	if err != nil {
		return 2
	}
	var result model.JsonResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return 2
	}
	if result.Status == true {
		if result.Data != "" {
			state, _ := strconv.Atoi(result.Data)
			return state
		}
	}
	return 2
}
