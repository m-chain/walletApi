package model

import (
	"time"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//系统通知
type SysMessage struct {
	MsgId         int64     `orm:"auto" json:"msgId"`
	Title         string    `orm:"size(100)" json:"title" valid:"Required" form:"title"`     //标题
	Content       string    `orm:"size(200)" json:"content" valid:"Required" form:"content"` //消息内容
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`            //创建时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"updateTime"`                //更新时间
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt"`
	UpdateTimeFmt string    `orm:"-" json:"updateTimeFmt"`
}

func (u *SysMessage) TableName() string {
	return "t_sys_msg"
}

//添加系统消息
func AddSysMessage(msg *SysMessage) error {
	o := orm.NewOrm()
	_, err := o.Insert(msg)

	return err
}

//修改系统消息
func UpdateSysMessage(msg *SysMessage) error {
	o := orm.NewOrm()
	_, err := o.Update(msg, "Title", "Content", "UpdateTime")

	return err
}

//删除系统消息
func DeleteSysMessage(id int64) error {
	o := orm.NewOrm()
	msg := SysMessage{MsgId: id}
	err := o.Read(&msg)
	if err != nil {
		return err
	}
	_, err = o.Delete(&msg)
	if err != nil {
		return err
	}
	return nil
}

//根据id查询详细信息
func GetSysMessage(id string) (*SysMessage, error) {
	o := orm.NewOrm()
	msg := new(SysMessage)
	err := o.Raw("SELECT * FROM t_sys_msg WHERE msg_id = ?", id).QueryRow(msg)
	if err != nil {
		return nil, err
	}

	msg.CreateTimeFmt = common.FormatTime(msg.CreateTime)
	msg.UpdateTimeFmt = common.FormatTime(msg.UpdateTime)
	return msg, nil
}

func QuerySysMsgList(pageSize, pageNo int, keyWord string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_sys_msg")

	if keyWord != "" {
		cond := orm.NewCondition()
		cond = cond.And("Title__icontains", keyWord).Or("Content__icontains", keyWord)
		qs = qs.SetCond(cond)
	}

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var messages []SysMessage
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-UpdateTime").All(&messages)

		if err == nil && len(messages) > 0 {
			resultMap["total"] = cnt
			for i := range messages {

				messages[i].CreateTimeFmt = common.FormatTime(messages[i].CreateTime)
				messages[i].UpdateTimeFmt = common.FormatTime(messages[i].UpdateTime)
			}
			resultMap["data"] = messages
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
	orm.RegisterModel(new(SysMessage))
}
