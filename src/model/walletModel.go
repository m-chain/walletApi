package model

import (
	"strings"
	"time"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
)

type JpushMessage struct {
	MsgId         int64     `orm:"auto" json:"msgId"`
	Target        string    `orm:"size(64)" json:"target"`   //消息接收者
	TxId          string    `orm:"size(80)" json:"txId"`     //转账交易ID
	Alert         string    `orm:"size(80)" json:"alert"`    //提示信息
	Title         string    `orm:"size(100)" json:"title"`   //二级标题
	Content       string    `orm:"size(200)" json:"content"` //消息内容
	MsgType       int       `json:"msgType"`                 //消息类型 1、转账交易信息
	IsRead        int       `json:"isRead"`                  //是否已读 1、已读　2、未读
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt"`
}
type Wallet struct {
	Address     string       `json:"address"`     //钱包地址
	PubKey      string       `json:"pubKey"`      //公钥
	CreateTime  string       `json:"createTime"`  //创建日期
	IsLocked    int          `json:"isLocked"`    //钱包是否被锁定 1、表示锁定　0、表示未锁定
	WalletRests []WalletRest `json:"walletRests"` //钱包余额信息
}

type WalletRest struct {
	TokenID      string `json:"tokenID" description:"token标识"`
	Name         string `json:"name" description:"token名称"`
	TokenSymbol  string `json:"tokenSymbol" description:"token英文简写名称"`
	RestNumber   string `json:"restNumber" description:"余额"`
	FreezeNumber string `json:"freezeNumber" description:"冻结的额度"`
	IsBaseCoin   bool   `json:"isBaseCoin" description:"是否是平台币"`
}

type WalletResp struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   Wallet `json:"data"`
}

func (j *JpushMessage) TableName() string {
	return "t_jpush_msg"
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(JpushMessage))
}

//添加
func AddJpushMsg(msg *JpushMessage) error {
	o := orm.NewOrm()
	_, err := o.Insert(msg)
	return err
}

//更改消息的已读状态
func UpdateJpushMsgReadStatus(msgID int64) error {
	o := orm.NewOrm()
	msg := JpushMessage{MsgId: msgID}
	err := o.Read(&msg)
	if err != nil {
		return err
	}
	if msg.IsRead == 2 {
		msg.IsRead = 1
		_, err = o.Update(&msg)
	}
	return err
}

//根据id查询详细信息
func GetJpushMessage(id string) (*JpushMessage, error) {
	o := orm.NewOrm()
	msg := new(JpushMessage)
	err := o.Raw("SELECT * FROM t_jpush_msg WHERE msg_id = ?", id).QueryRow(msg)
	if err != nil {
		return nil, err
	}

	msg.CreateTimeFmt = common.FormatTime(msg.CreateTime)
	return msg, nil
}

//address可以有多个，以逗号分隔
func QueryJpushMsgList(address string, pageSize, pageNo int) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_jpush_msg")

	cond := orm.NewCondition()
	addrs := strings.Split(address, ",")
	for _, addr := range addrs {
		cond = cond.Or("Target__exact", addr)
	}
	qs = qs.SetCond(cond)

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var messages []JpushMessage
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-CreateTime").All(&messages)

		if err == nil && len(messages) > 0 {
			resultMap["total"] = cnt
			for i := range messages {
				messages[i].CreateTimeFmt = common.FormatTime(messages[i].CreateTime)
			}
			resultMap["data"] = messages
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = nil

	return resultMap
}
