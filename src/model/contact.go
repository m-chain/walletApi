package model

import (
	"fmt"
	"strings"
	"time"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//联系人
type Contact struct {
	Id            int64     `orm:"auto" json:"id"`
	Name          string    `orm:"size(80)" json:"name"`
	Address       string    `orm:"size(64)" json:"address"`
	ContactAddr   string    `orm:"size(64)" json:"contactAddr"`
	Remark        string    `orm:"size(80)" json:"remark"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime"`
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt"`
}

func (q *Contact) TableName() string {
	return "t_contact"
}

//添加
func AddContact(cont *Contact) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_contact")

	var conts []Contact
	arr := strings.Split(cont.Address, ",")
	cond := orm.NewCondition()
	for _, v := range arr {
		cond = cond.Or("Address__exact", v)
	}
	_, err := qs.SetCond(cond).OrderBy("-CreateTime").All(&conts)
	var addConts []Contact
	if len(conts) == 0 {
		for _, v := range arr {
			var item Contact
			item.Name = cont.Name
			item.ContactAddr = cont.ContactAddr
			item.Address = v
			item.Remark = cont.Remark
			addConts = append(addConts, item)
		}
	} else {
		if cont.Name != "" && cont.ContactAddr != "" {
			for _, v := range arr {
				if !IsExistContact(v, *cont, conts) {
					var item Contact
					item.Name = cont.Name
					item.ContactAddr = cont.ContactAddr
					item.Address = v
					item.Remark = cont.Remark
					addConts = append(addConts, item)
				}
			}
		}
	}

	if len(addConts) > 0 {
		conts = append(conts, addConts...)
	}

	var endConts []Contact
	for _, v := range arr {
		for _, b := range conts {
			if !IsExistContact(v, b, conts) {
				var item Contact
				item.Name = b.Name
				item.ContactAddr = b.ContactAddr
				item.Address = v
				item.Remark = b.Remark
				item.CreateTime = b.CreateTime
				endConts = append(endConts, item)
			}
		}
	}
	if len(addConts) > 0 {
		endConts = append(endConts, addConts...)
	}
	fmt.Printf("%v\n", endConts)
	var num int64
	if len(endConts) > 0 {
		num, err = o.InsertMulti(len(endConts), endConts)
	}
	return num, err
}

func IsExistContact(address string, newCont Contact, oldConts []Contact) bool {
	isExist := false
	for _, v := range oldConts {
		if v.Address == address && v.Name == newCont.Name && v.ContactAddr == newCont.ContactAddr {
			isExist = true
			break
		}
	}
	return isExist
}

//修改
func UpdateContact(cont *Contact) (int64, error) {
	o := orm.NewOrm()

	oldCont := Contact{Id: cont.Id}
	err := o.Read(&oldCont)
	if err != nil {
		return 0, err
	}
	arr := strings.Split(cont.Address, ",")
	var total int64
	o.Begin()
	for _, addr := range arr {
		res, err := o.Raw("UPDATE t_contact SET name=?,contact_addr=?,remark=? WHERE name=? AND contact_addr=? AND address=?", cont.Name, cont.ContactAddr, cont.Remark, oldCont.Name, oldCont.ContactAddr, addr).Exec()
		if err != nil {
			o.Rollback()
			return 0, err
		}
		num, _ := res.RowsAffected()
		total = total + num
	}
	o.Commit()

	return total, err
}

//删除
func DeleteContact(id int64, address string) (int64, error) {
	o := orm.NewOrm()

	cont := Contact{Id: id}
	err := o.Read(&cont)
	if err != nil {
		return 0, err
	}
	addrs := strings.Split(address, ",")
	var total int64
	o.Begin()
	for _, addr := range addrs {
		res, err := o.Raw("DELETE FROM t_contact WHERE name=? AND contact_addr=? AND address=?", cont.Name, cont.ContactAddr, addr).Exec()
		if err != nil {
			o.Rollback()
			return 0, err
		}
		num, _ := res.RowsAffected()
		total = total + num
	}
	o.Commit()
	fmt.Println("mysql row affected nums: ", total)

	return total, err
}

//查询详情
func GetContactInfo(id int64) (*Contact, error) {
	o := orm.NewOrm()
	cont := new(Contact)
	err := o.Raw("SELECT * FROM t_contact WHERE id = ?", id).QueryRow(cont)
	if err != nil {
		return nil, err
	}

	cont.CreateTimeFmt = common.FormatTime(cont.CreateTime)
	return cont, nil
}

//查询所有
func GetContactAll(address string) ([]Contact, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_contact")

	var conts []Contact
	arr := strings.Split(address, ",")
	cond := orm.NewCondition()
	for _, v := range arr {
		cond = cond.Or("Address__exact", v)
	}
	_, err := qs.SetCond(cond).OrderBy("-CreateTime").All(&conts)
	if err != nil {
		return nil, err
	}
	for i := range conts {
		conts[i].CreateTimeFmt = common.FormatTime(conts[i].CreateTime)
	}
	return conts, nil
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Contact))
}
