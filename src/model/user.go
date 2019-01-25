package model

import (
	"errors"
	"fmt"
	"time"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//用户
type User struct {
	Id               int64     `orm:"auto" from:"id"`
	UserName         string    `orm:"size(100);column(UserName)" valid:"Required" form:"userName"`
	NickName         string    `orm:"size(100);column(nickName)" valid:"Required" form:"nickName"`
	Password         string    `orm:"size(100);column(password)" valid:"Required" form:"pwd"`
	CreateTime       time.Time `orm:"auto_now_add;type(datetime);column(createTime)"`
	LastLoginTime    time.Time `orm:"auto_now;type(datetime);column(lastLoginTime)"` //'最近登录时间'
	CreateTimeFmt    string    `orm:"-"`                                             //创建时间
	LastLoginTimeFmt string    `orm:"-"`                                             //'最近登录时间'
}

func (u *User) TableName() string {
	return "t_user"
}

//更新验证码信息
func AddUser(user *User) error {
	u, err := GetUserInfoByAccount(user.UserName)
	if err != nil {
		return err
	}

	if u != nil && u.UserName != "" {
		return fmt.Errorf("用户已经存在，不能重复添加!")
	}
	o := orm.NewOrm()
	_, err = o.Insert(user)
	return err
}

//更新用户密码
func UpdateUserPwd(userName, newpwd string) bool {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE t_user SET password = ? WHERE userName = ?", newpwd, userName).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			return true
		}
	}
	return false
}

//更新用户最后一次登录时间
func UpdateLastLoginTime(userName string) bool {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE t_user SET lastLoginTime = now() WHERE userName = ?", userName).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			return true
		}
	}
	return false
}

//删除用户
func DeleteUser(id int64) error {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	if user.UserName == "admin" {
		//fmt.Errorf("不能删除admin用户!")
		return errors.New("不能删除Admin用户!")
	}
	_, err = o.Delete(&user)
	if err != nil {
		return err
	}
	return nil
}

//根据手机号码获取用户信息
func GetUserInfoByAccount(userName string) (*User, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_user")
	if userName != "" {
		qs = qs.Filter("UserName", userName)
	}
	user := new(User)
	//注意：如果不存在记录，也会报错（err不会为空）
	qs.One(user)
	return user, nil
}

//根据用户数据总个数
func GetRecordNum(search string) int64 {

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if search != "" {
		qs = qs.Filter("Name", search)
	}
	var us []User
	num, err := qs.All(&us)
	if err == nil {
		return num
	} else {
		return 0
	}
}

func SearchUsers(pageSize, pageNo int, search string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_user")

	if search != "" {
		cond := orm.NewCondition()
		cond = cond.And("UserName__icontains", search).Or("NickName__icontains", search)
		qs = qs.SetCond(cond)
	}

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var users []User
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).All(&users)

		if err == nil {
			resultMap["total"] = cnt
			for i, v := range users {
				v.CreateTimeFmt = common.FormatTime(v.CreateTime)
				v.LastLoginTimeFmt = common.FormatTime(v.LastLoginTime)
				users[i] = v
			}
			resultMap["data"] = users
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = nil

	return resultMap
}

func SearchDataList(pagesize, pageno int, search string) (users []User) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_user")
	if search != "" {
		qs = qs.Filter("UserName", search)
	}
	var us []User
	cnt, err := qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Println("count", cnt)
	}
	return us
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
