package model

import (
	"strings"
	"time"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//签名数据表
type SignData struct {
	Id            int64     `orm:"auto" json:"id" description:"主键ID"`
	QrCode        string    `orm:"size(64)" json:"qrCode" description:"二维码标识"` //二维码
	Address       string    `orm:"size(100)" json:"address" description:"钱包地址"`
	PubKey        string    `orm:"size(100)" json:"pubKey" description:"公钥"`
	SignType      int       `json:"signType" description:"签名类型　1、发布合约 2、实例化合约 3、升级合约 4、发行token 5、设置master、manager 6、manager签名确认 7、添加manager　8、替换manager 9 、删除manager 10、设置manager操作确认的阀值 11、设置发行token所需要的平台币 12、设置发布合约所需要的平台币 13、设置手续费返还规则"`
	TipMsg        string    `orm:"-" json:"tipMsg" description:"提示信息"`
	OriginData    string    `orm:"size(2000)" json:"originData" description:"签名原始数据"`
	SignData      string    `orm:"size(4000)" json:"signData" description:"签名数据"`
	Status        int       `json:"status" description:"状态 : 1、待签名 2、已签名 3、已确认 -1、已失效"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" description:"生成二维码日期"`
	SignTime      time.Time `orm:"type(datetime)" json:"signTime" description:"签名日期"`
	RespResult    string    `orm:"size(2000)" json:"respResult" description:"底层链调用结果"`
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt" description:"日期格式化输出"`
	SignTimeFmt   string    `orm:"-" json:"signTimeFmt" description:"日期格式化输出"`
	ValidTime     int64     `orm:"size(11)" json:"validTime" description:"有效时间戳"`
	LangType      string    `orm:"size(20)" json:"langType" description:"语言类型：zh: 中文 en:英文"`
}

type ContractRequest struct {
	Name           string `json:"name" description:"合约名称"`
	ContractSymbol string `json:"contractSymbol" description:"合约英文间写"`
	Version        string `json:"version" description:"合约版本"`
	Remark         string `json:"remark" description:"备注"`
	CcUrl          string `json:"ccUrl" description:"合约源码路径"`
}

func (q *SignData) TableName() string {
	return "t_sign_data"
}

//添加
func AddSignData(sign *SignData) error {
	o := orm.NewOrm()
	_, err := o.Insert(sign)
	return err
}

//修改
func UpdateSignData(sign *SignData) error {
	o := orm.NewOrm()
	var rSign SignData
	err := o.Raw("SELECT * FROM t_sign_data WHERE qr_code=?", sign.QrCode).QueryRow(&rSign)

	if rSign.Status == 3 {
		_, err = o.Raw("UPDATE t_sign_data SET resp_result=? WHERE qr_code=?", sign.RespResult, sign.QrCode).Exec()
	} else {
		_, err = o.Raw("UPDATE t_sign_data SET address = ?,pub_key=?,sign_data=?,status=?,sign_time=now(),resp_result=? WHERE status<>3 AND qr_code=?", sign.Address, sign.PubKey, sign.SignData, sign.Status, sign.RespResult, sign.QrCode).Exec()
	}

	return err
}

//查询
func GetSingInfo(qrCode string) (*SignData, error) {
	o := orm.NewOrm()
	sign := new(SignData)
	err := o.Raw("SELECT * FROM t_sign_data WHERE qr_code = ?", qrCode).QueryRow(sign)
	if err != nil {
		return nil, err
	}
	timeUnix := time.Now().Unix()
	if timeUnix > sign.ValidTime { //已失效
		ret, err := o.Raw("UPDATE t_sign_data SET status=-1 WHERE status<>3 AND qr_code=?", qrCode).Exec()
		if err == nil {
			num, _ := ret.RowsAffected()
			if num > 0 {
				sign.Status = -1
			}
		}
	}
	sign.CreateTimeFmt = common.FormatTime(sign.CreateTime)
	sign.SignTimeFmt = common.FormatTime(sign.SignTime)
	return sign, nil
}

func GetAllSignInfos(address string, pageSize, pageNo int) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_sign_data")

	cond := orm.NewCondition()
	addrs := strings.Split(address, ",")
	for _, addr := range addrs {
		cond = cond.Or("Address__exact", addr)
	}
	cond = cond.AndNot("Status__exact", -1)

	qs = qs.SetCond(cond)

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		signs := make([]SignData, 0)
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-SignTime").All(&signs)

		if err == nil && len(signs) > 0 {
			resultMap["total"] = cnt
			for i := range signs {
				signs[i].CreateTimeFmt = common.FormatTime(signs[i].CreateTime)
				signs[i].SignTimeFmt = common.FormatTime(signs[i].SignTime)
			}
			resultMap["data"] = signs
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = nil

	return resultMap
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SignData))
}
