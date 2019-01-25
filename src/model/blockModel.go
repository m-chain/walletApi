package model

import (
	"bytes"
	"fmt"
	"math/big"
	"strconv"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
)

type BlockResult struct {
	Status bool      `json:"status" description:"状态"`
	Data   BlockInfo `json:"data" description:"区块信息"`
}

type BlockInfo struct {
	Blockid       string   `json:"blockid" description:"区块编号"`
	Previous_hash string   `json:"previous_hash" description:"上一个区块hash"`
	Data_hash     string   `json:"data_hash" description:"数据hash"`
	Transactions  []TxInfo `json:"transactions" description:"交易数据"`
}

type TxInfo struct {
	Channelname      string            `json:"channelname" description:"链名称"`
	Chaincodename    string            `json:"chaincodename" description:"合约名称"`
	Chaincodeversion string            `json:"chaincodeversion" description:"合约版本"`
	Txhash           string            `json:"txhash" description:"交易hash"`
	Createdt         string            `json:"createdt" description:"交易时间"`
	Trans            []Transaction     `json:"trans" description:"转账信息"`
	Tokens           []Token           `json:"tokens" description:"token信息"`
	TokenMaster      map[string]string `json:"tokenMaster" description:"token的发行者"`
}

type Transaction struct {
	Id          int64  `orm:"auto" json:"id" description:"主健ID"`
	BlockId     string `orm:"size(100);column(blockid)" json:"blockId" description:"区块编号"`
	TxHash      string `orm:"size(100);column(txhash)" json:"txHash" description:"交易hash"`
	TxId        string `orm:"size(64);" json:"txId" description:"转账交易hash"`
	TokenID     string `orm:"size(32);column(token_id)" json:"tokenID" description:"token标识或ID"`
	TokenSymbol string `orm:"size(64)" json:"tokenSymbol" description:"token简称"`
	FromAddress string `orm:"size(100);column(from_address)" json:"fromAddress" description:"付款方"`
	ToAddress   string `orm:"size(100);column(to_address)" json:"toAddress" description:"收款方"`
	Number      string `orm:"size(100);column(number)" json:"number" description:"数量"`
	Fee         string `orm:"size(100);column(fee)" json:"gasUsed" description:"消耗的gas"`
	IsCost      int    `orm:"size(4);column(is_cost)" json:"isCost" description:"是否转账交易 1:是 0:否"`
	TxTime      string `orm:"type(datetime);" json:"time" description:"交易时间"`
	Nonce       string `orm:"size(80);column(nonce)" json:"nonce" description:"避免重复交易"`
	State       int    `orm:"size(11);column(state)" json:"state" description:"交易状态 1:成功　2:失败"`
	Notes       string `orm:"size(200);column(notes)" json:"notes" description:"备注"`
	Msg         string `orm:"size(100);column(msg)" json:"msg" description:"交易失败时的错误信息"`
}

type TransferDetail struct {
	Status bool
	Msg    string
	Data   []Transfer
}

type Transfer struct {
	TxId        string `json:"txId" description:"转账交易ID"`
	TokenID     string `json:"tokenID" description:"token标识或ID"`
	FromAddress string `json:"fromAddress" description:"付款方"`
	ToAddress   string `json:"toAddress" description:"收款方"`
	Number      string `json:"number" description:"数量"`
	GasUsed     string `json:"gasUsed" description:"消耗的gas"`
	Time        string `json:"time" description:"转账时间"`
	Nonce       string `json:"nonce" description:"避免重复交易"`
	Notes       string `json:"notes" description:"备注"`
	State       int    `json:"state" description:"交易状态 1:成功　2:失败"`
	Msg         string `json:"msg" description:"交易失败时的错误信息"`
}

type Token struct {
	TokenID      string `orm:"pk;size(32);column(token_id)" json:"tokenID" description:"token标识或ID"`
	Name         string `orm:"size(100);column(name)" json:"name" description:"token名称"`
	TokenSymbol  string `orm:"size(40);column(token_symbol)" json:"tokenSymbol" description:"token英文简写名称"`
	IconUrl      string `orm:"size(200);column(icon_url)" json:"iconUrl" description:"token图标"`
	IsBaseCoin   bool   `orm:"column(is_base_coin)" json:"isBaseCoin" description:"是否主币"`
	DecimalUnits int    `orm:"size(11);column(decimal_units)" json:"decimalUnits" description:"最大小数点位数"`
	TotalNumber  string `orm:"size(100);column(total_number)" json:"totalNumber" description:"发行总量"`
	RestNumber   string `orm:"-" json:"restNumber" description:"剩余数量"`
	IssuePrice   string `orm:"size(80);column(issue_price)" json:"issuePrice" description:"发行价格"`
	IssueTime    string `orm:"auto_now_add;type(datetime);column(issue_time)" json:"issueTime" description:"发行时间"`
	Status       int    `orm:"size(4);column(status)" json:"status" description:"1、启用　0、禁用"`
	OwnerAddress string `orm:"size(80);" json:"ownerAddress" description:"发行token的地址"`
}

//contract info
type Contract struct {
	Name            string `json:"name" description:"合约名称"`
	ContractAddress string `json:"contractAddress" description:"合约地址"`
	ContractSymbol  string `json:"contractSymbol" description:"合约名称简写"`
	MAddress        string `json:"mAddress" description:"钱包地址"`
	Version         string `json:"version" description:"版本号"`
	CcPath          string `json:"ccPath" description:"合约路径"`
	Remark          string `json:"remark" description:"合约简介"`
	Status          string `json:"status" description:"合约状态 -1、已删除 1、待初始化 2、正在运行 3、余额不足 4、合约已禁用 5、已弃用"`
	CreateTime      string `json:"createTime" description:"合约发布时间"`
	UpdateTime      string `json:"updateTime" description:"合约更新时间"`
}

//股权
type Vesting struct {
	ID                int64  `json:"id" description:"序列号"`
	TokenID           string `json:"tokenID" description:"TOKEN标识"`
	StartTime         int64  `json:"startTime" description:"股权开始时间,以时间戳(秒)为单位"`
	InitReleaseAmount string `json:"initReleaseAmount" description:"初始释放量"`
	Amount            string `json:"amount" description:"总量"`
	Interval          int64  `json:"interval" description:"每次释放间隔时间"`
	Periods           int64  `json:"periods" description:"总期数"`
	Withdrawed        string `json:"withdrawed" description:"此vesting中已经被提取出的币总量"`
	TxID              string `json:"txID" description:"交易ID"`
	VType             int    `json:"vType" description:"类型 0:预售股权 1:实例化合约 2:升级合约 3:合约调用 4:发行token"`
	ContractAddress   string `json:"contractAddress" description:"合约帐户地址"`
	Reason            string `json:"reason" description:"生成此记录的原由"`
	CreateTime        string `json:"createTime" description:"创建时间"`
}

//
type VestingItem struct {
	ID                int64  `json:"id" description:"记录ID"`
	TokenID           string `json:"tokenID" description:"TOKEN标识"`
	Amount            string `json:"amount" description:"总量"`
	InitReleaseAmount string `json:"initReleaseAmount" description:"立即释放数量"`
	Interval          string `json:"interval" description:"每次释放时间间隔"`
	Periods           int64  `json:"periods" description:"总期数"`
	Withdrawed        string `json:"withdrawed" description:"已划入余额的数量"`
	PeriodNum         string `json:"periodNum" description:"每期返还的数量"`
	VType             int    `json:"vType" description:"类型 0:预售股权 1:实例化合约 2:升级合约 3:合约调用 4:发行token"`
	VTypeName         string `json:"vTypeName" description:"类型名称"`
	EffectTime        string `json:"effectTime" description:"生效时间"`
	CreateTime        string `json:"createTime" description:"创建时间"`
	Status            string `json:"status" description:"状态 返还中、待返还"`
}

type TxStatusItem struct {
	TxID  string `json:"txID"`
	State string `json:"state"`
}

func (u *Transaction) TableName() string {

	return "t_transaction"
}

// 多字段唯一键
func (u *Transaction) TableUnique() [][]string {
	return [][]string{
		[]string{"TxId", "IsCost"},
	}
}

func (t *Token) TableName() string {
	return "t_token_info"
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Transaction), new(Token))
}

func AddTransaction(trans Transaction) error {
	o := orm.NewOrm()
	_, err := o.InsertOrUpdate(&trans, "TxId,IsCost")

	return err
}

func UpdateTransferStatus(txID, state string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE t_transaction SET state=? WHERE tx_id=?", state, txID).Exec()

	return err
}

func UpdateMultiTransferStatus(items []TxStatusItem) error {
	o := orm.NewOrm()
	var successBuf, failBuf bytes.Buffer
	for _, item := range items {
		if item.State == "1" {
			if successBuf.Len() > 0 {
				successBuf.WriteString(",")
			}
			successBuf.WriteString("'")
			successBuf.WriteString(item.TxID)
			successBuf.WriteString("'")
		} else if item.State == "2" {
			if failBuf.Len() > 0 {
				failBuf.WriteString(",")
			}
			failBuf.WriteString("'")
			failBuf.WriteString(item.TxID)
			failBuf.WriteString("'")
		}
	}
	var err error
	if successBuf.Len() > 0 {
		_, err = o.Raw("UPDATE t_transaction SET state=1 WHERE tx_id IN(" + successBuf.String() + ")").Exec()
	}
	if failBuf.Len() > 0 {
		_, err = o.Raw("UPDATE t_transaction SET state=2 WHERE tx_id IN(" + failBuf.String() + ")").Exec()
	}

	return err
}

func QueryTransfersNoStatus() []string {
	o := orm.NewOrm()
	var list orm.ParamsList
	num, err := o.Raw("SELECT tx_id FROM t_transaction WHERE state=0 ORDER BY blockid ASC limit 0,100").ValuesFlat(&list)
	var arr []string
	if err == nil && num > 0 {
		for i := range list {
			arr = append(arr, fmt.Sprintf("%v", list[i]))
		}
	}
	return arr
}

func UpdateTokenSymbol() (int64, error) {
	o := orm.NewOrm()
	ret, err := o.Raw("UPDATE t_transaction a,t_token_info b SET a.token_symbol = b.token_symbol WHERE a.token_id=b.token_id AND (ISNULL(a.token_symbol) || LENGTH(trim(a.token_symbol))<1)").Exec()
	if err != nil {
		return 0, err
	}
	num, err := ret.RowsAffected()

	return num, err
}

func AddToken(token *Token) error {
	o := orm.NewOrm()

	rToken := Token{TokenID: token.TokenID}
	err := o.Read(&rToken)
	if rToken.Name == "" { //添加
		_, err = o.Raw("INSERT INTO t_token_info(token_id,name,token_symbol,icon_url,is_base_coin,decimal_units,total_number,issue_price,issue_time,status,owner_address)values(?,?,?,?,?,?,?,?,now(),1,?)", token.TokenID, token.Name, token.TokenSymbol, token.IconUrl, token.IsBaseCoin, token.DecimalUnits, token.TotalNumber, token.IssuePrice, token.OwnerAddress).Exec()
	} else { //修改
		if token.IconUrl != "" {
			o.Update(token, "TotalNumber", "IssuePrice", "OwnerAddress", "IconUrl")
		} else {
			o.Update(token, "TotalNumber", "IssuePrice", "IssueTime", "OwnerAddress")
		}
	}
	return err
}

func UpdateTokenIcon(tokenID, iconUrl string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE t_token_info SET icon_url=? WHERE token_id=?", iconUrl, tokenID).Exec()
	return err
}

func QueryMaxBlockNum() int64 {
	o := orm.NewOrm()
	var list orm.ParamsList
	num, err := o.Raw("SELECT max(blockid+0) as blockNum FROM t_transaction").ValuesFlat(&list)
	if err == nil && num > 0 {
		blockNum, _ := strconv.ParseInt(fmt.Sprintf("%v", list[0]), 10, 64)
		return blockNum
	}
	return 0
}

func QueryTokenInfo(tokenID string) (*Token, error) {
	o := orm.NewOrm()
	token := new(Token)
	err := o.Raw("SELECT * FROM t_token_info WHERE token_id = ?", tokenID).QueryRow(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func QueryAllTokens() []Token {
	o := orm.NewOrm()
	tokens := make([]Token, 0)
	o.Raw("SELECT * FROM t_token_info ORDER BY issue_time ASC").QueryRows(&tokens)
	return tokens
}

func QueryMasterTokenDecimalUnits() (int, error) {
	o := orm.NewOrm()
	var list orm.ParamsList
	num, err := o.Raw("SELECT decimal_units FROM t_token_info WHERE is_base_coin=1").ValuesFlat(&list)
	if err == nil && num > 0 {
		digits, _ := strconv.Atoi(fmt.Sprintf("%v", list[0]))
		return digits, err
	}
	return 0, err
}

func QueryMasterTokenID() string {
	o := orm.NewOrm()
	var list orm.ParamsList
	num, err := o.Raw("SELECT token_id FROM t_token_info WHERE is_base_coin=1").ValuesFlat(&list)
	if err == nil && num > 0 {
		return fmt.Sprintf("%v", list[0])
	}
	return ""
}

func QueryMasterTokenInfo() Token {
	o := orm.NewOrm()
	var token Token
	o.Raw("SELECT * FROM t_token_info WHERE is_base_coin=1").QueryRow(&token)

	return token
}

func QueryTokenDecimalUnits(tokenIds []string) (map[string]interface{}, error) {
	o := orm.NewOrm()
	res := make(orm.Params)
	condition := ""
	for _, v := range tokenIds {
		if len(condition) > 0 {
			condition += " OR "
		}
		condition += "token_id='" + v + "'"
	}
	_, err := o.Raw("SELECT token_id, decimal_units FROM t_token_info WHERE "+condition).RowsToMap(&res, "token_id", "decimal_units")

	return res, err
}

func QueryTokenNames(tokenIds []string) (map[string]interface{}, error) {
	o := orm.NewOrm()
	res := make(orm.Params)
	condition := ""
	for _, v := range tokenIds {
		if len(condition) > 0 {
			condition += " OR "
		}
		condition += "token_id='" + v + "'"
	}
	_, err := o.Raw("SELECT token_id, name FROM t_token_info WHERE "+condition).RowsToMap(&res, "token_id", "name")

	return res, err
}

//查询转账交易详情
func QueryTransferDetails(txID string) (*Transaction, error) {
	o := orm.NewOrm()
	trans := new(Transaction)
	err := o.Raw("SELECT a.* FROM t_transaction a WHERE a.tx_id=? and a.is_cost=1", txID).QueryRow(trans)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func FloatNumber(number, tokenID string) string {
	tokenMap, _ := QueryTokenDecimalUnits([]string{tokenID})
	decimalUnits := tokenMap[tokenID]
	units, _ := strconv.Atoi(fmt.Sprintf("%v", decimalUnits))
	digits := common.GetNumberWithDigits(units)
	bigAmount := new(big.Int)
	bigAmount.SetString(number, 10)

	end := common.BigIntDiv(bigAmount, digits)

	return common.FormatNumber(end)
}

func doWithTransferData(arr []Transaction) {
	var tokenIds []string
	for _, item := range arr {
		tokenIds = append(tokenIds, item.TokenID)
	}
	tokenMap, _ := QueryTokenDecimalUnits(tokenIds)
	masterToken := QueryMasterTokenInfo()
	masterDigits := common.GetNumberWithDigits(masterToken.DecimalUnits)
	for i := range arr {
		decimalUnits := tokenMap[arr[i].TokenID]
		units, _ := strconv.Atoi(fmt.Sprintf("%v", decimalUnits))

		digits := common.GetNumberWithDigits(units)

		restNum := new(big.Int)
		restNum.SetString(arr[i].Number, 10)

		freezeNum := new(big.Int)
		freezeNum.SetString(arr[i].Fee, 10)

		arr[i].Number = common.FormatNumber(common.BigIntDiv(restNum, digits))
		arr[i].Fee = common.FormatNumber(common.BigIntDiv(freezeNum, masterDigits))
	}
}

//查询某地址下某一token的所有转账交易
func QueryTransferListByToken(address, tokenID string, pageSize, pageNo int) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_transaction")

	cond2 := orm.NewCondition()
	cond2 = cond2.Or("FromAddress__exact", address).Or("ToAddress__exact", address)

	cond := orm.NewCondition()
	cond = cond.And("TokenID__exact", tokenID).And("IsCost__exact", 1).AndCond(cond2)
	cond = cond.AndNot("State__in", 2)
	qs = qs.SetCond(cond)

	arr := make([]Transaction, 0)
	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {

		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-TxTime").All(&arr)

		if err == nil && len(arr) > 0 {
			doWithTransferData(arr)

			resultMap["total"] = cnt
			resultMap["data"] = arr
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = arr

	return resultMap
}

//查询两钱包地址之间的转账交易
func QueryTransferListByAddress(address1, address2, tokenID string, pageSize, pageNo int) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_transaction")

	cond2 := orm.NewCondition()
	cond2 = cond2.And("FromAddress__exact", address1).And("ToAddress__exact", address2)

	cond3 := orm.NewCondition()
	cond3 = cond3.And("FromAddress__exact", address2).And("ToAddress__exact", address1)

	cond4 := orm.NewCondition()
	cond4 = cond4.AndCond(cond2).OrCond(cond3)

	cond := orm.NewCondition()
	cond = cond.And("IsCost__exact", 1).AndCond(cond4)
	if tokenID != "" {
		cond = cond.And("TokenID__exact", tokenID)
	}
	cond = cond.AndNot("State__in", 2)
	qs = qs.SetCond(cond)
	arr := make([]Transaction, 0)
	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-TxTime").All(&arr)

		if err == nil && len(arr) > 0 {
			doWithTransferData(arr)
			resultMap["total"] = cnt
			resultMap["data"] = arr
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = arr

	return resultMap
}

//根据地址、年份、起止年月、token查询交易信息
func QueryTransferList(address, year, startYM, endYM, tokenID string, pageSize, pageNo int) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_transaction")

	//地址
	addrCond := orm.NewCondition()
	addrCond = addrCond.And("FromAddress__exact", address).Or("ToAddress__exact", address)

	//时间
	startTime := ""
	endTime := ""
	if startYM != "" || endYM != "" {
		if startYM != "" && endYM != "" {
			_, endDay := common.StartAndEndDayOfMonth(endYM)
			startTime = startYM + "-01 00:00:00"
			endTime = endDay + " 23:59:59"
		} else if startYM != "" {
			_, endDay := common.StartAndEndDayOfMonth(startYM)
			startTime = startYM + "-01 00:00:00"
			endTime = endDay + " 23:59:59"
		} else if endYM != "" {
			_, endDay := common.StartAndEndDayOfMonth(endYM)
			startTime = endYM + "-01 00:00:00"
			endTime = endDay + " 23:59:59"
		}
	} else {
		if year != "" {
			_, endDay := common.StartAndEndDayOfMonth(year + "-12")
			startTime = year + "-01-01 00:00:00"
			endTime = endDay + " 23:59:59"
		}
	}

	cond := orm.NewCondition()
	if startTime != "" && endTime != "" {
		dateCond := orm.NewCondition()
		dateCond = dateCond.And("TxTime__gte", startTime).And("TxTime__lte", endTime)
		cond = cond.And("IsCost__exact", 1).AndCond(addrCond).AndCond(dateCond)
	} else {
		cond = cond.And("IsCost__exact", 1).AndCond(addrCond)
	}
	arr := make([]Transaction, 0)
	resultMap := make(map[string]interface{})
	if tokenID != "" {
		cond = cond.And("TokenID__exact", tokenID)
	}
	cond = cond.AndNot("State__in", 2)
	qs = qs.SetCond(cond)

	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-TxTime").All(&arr)

		if err == nil && len(arr) > 0 {
			doWithTransferData(arr)
			resultMap["total"] = cnt
			resultMap["data"] = arr
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = arr

	return resultMap
}

//根据地址、年份、起止年月、token查询交易统计信息
func QueryTransferStatistics(address, year, startYM, endYM, tokenID string) map[string]interface{} {
	o := orm.NewOrm()
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("SUM(number+0) as amount").
		From("t_transaction").
		Where("1=1").
		And("is_cost=1").
		And("state<>2")

	//时间
	startTime := ""
	endTime := ""
	if startYM != "" || endYM != "" {
		if startYM != "" && endYM != "" {
			_, endDay := common.StartAndEndDayOfMonth(endYM)
			startTime = startYM + "-01 00:00:00"
			endTime = endDay + " 23:59:59"
		} else if startYM != "" {
			_, endDay := common.StartAndEndDayOfMonth(startYM)
			startTime = startYM + "-01 00:00:00"
			endTime = endDay + " 23:59:59"
		} else if endYM != "" {
			_, endDay := common.StartAndEndDayOfMonth(endYM)
			startTime = endYM + "-01 00:00:00"
			endTime = endDay + " 23:59:59"
		}
	} else {
		if year != "" {
			_, endDay := common.StartAndEndDayOfMonth(year + "-12")
			startTime = year + "-01-01 00:00:00"
			endTime = endDay + " 23:59:59"
		}
	}

	if startTime != "" && endTime != "" {
		qb.And("tx_time>=?").And("tx_time<=?")
	}
	resultMap := make(map[string]interface{})
	if tokenID != "" {
		qb.And("token_id=?")
		type AmountResult struct {
			Amount string
		}
		var outAmount AmountResult
		var inAmount AmountResult
		if startTime != "" && endTime != "" {
			o.Raw(qb.String()+" AND from_address='"+address+"'", startTime, endTime, tokenID).QueryRow(&outAmount)
			o.Raw(qb.String()+" AND to_address='"+address+"'", startTime, endTime, tokenID).QueryRow(&inAmount)
		} else {
			o.Raw(qb.String()+" AND from_address='"+address+"'", tokenID).QueryRow(&outAmount)
			o.Raw(qb.String()+" AND to_address='"+address+"'", tokenID).QueryRow(&inAmount)
		}
		resultMap["outAmount"] = FloatNumber(outAmount.Amount, tokenID)
		resultMap["inAmount"] = FloatNumber(inAmount.Amount, tokenID)
		return resultMap
	}
	resultMap["outAmount"] = "0"
	resultMap["inAmount"] = "0"
	return resultMap
}
