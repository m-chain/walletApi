package model

import (
	"strings"
	"time"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//意见反馈
type Question struct {
	Id            int64     `orm:"auto" json:"id" description:"记录ID"`
	Title         string    `orm:"size(100)" json:"title" description:"标题"`
	Address       string    `orm:"size(80)" json:"address" description:"当前钱包地址"`
	Question      string    `orm:"size(1000)" json:"question" description:"问题描述"`
	Email         string    `orm:"size(40)" json:"email" description:"邮箱"`
	BaseUrl       string    `orm:"size(100)" json:"baseUrl" description:"图片根url"`
	ImgName       string    `orm:"size(2000)" json:"imgName" description:"图片名称，文件名加后缀名"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" json:"createTime" description:"创建日期"`
	UpdateTime    time.Time `orm:"auto_now;type(datetime)" json:"updateTime" description:"修改日期"`
	CreateTimeFmt string    `orm:"-" json:"createTimeFmt" description:"格式化日期输出"`
	UpdateTimeFmt string    `orm:"-" json:"updateTimeFmt" description:"格式化日期输出"`
}

func (q *Question) TableName() string {
	return "t_question"
}

//添加
func AddQuestion(quest *Question) error {
	o := orm.NewOrm()
	_, err := o.Insert(quest)
	return err
}

//修改
func UpdateQuestion(quest *Question) error {
	o := orm.NewOrm()
	//Update 默认更新所有的字段，可以更新指定的字段：
	var err error
	if quest.BaseUrl == "" || quest.ImgName == "" {
		_, err = o.Update(quest, "Title", "Question", "UpdateTime")
	} else {
		_, err = o.Update(quest, "Title", "Question", "BaseUrl", "ImgName", "UpdateTime")
	}

	return err
}

//删除
func DeleteQuestion(id int64) error {
	o := orm.NewOrm()
	quest := Question{Id: id}
	err := o.Read(&quest)
	if err != nil {
		return err
	}
	_, err = o.Delete(&quest)
	if err != nil {
		return err
	}
	return nil
}

//查询详情
func GetQuestionInfo(id int64) (*Question, error) {
	o := orm.NewOrm()
	quest := new(Question)
	err := o.Raw("SELECT * FROM t_question WHERE id = ?", id).QueryRow(quest)
	if err != nil {
		return nil, err
	}
	quest.CreateTimeFmt = common.FormatTime(quest.CreateTime)
	quest.UpdateTimeFmt = common.FormatTime(quest.UpdateTime)
	return quest, nil
}

//查询所有
func GetQuestionAll(address string) ([]Question, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("t_question")

	var quests []Question
	arr := strings.Split(address, ",")
	cond := orm.NewCondition()
	for _, v := range arr {
		cond = cond.Or("Address__exact", v)
	}
	_, err := qs.SetCond(cond).OrderBy("-UpdateTime").All(&quests)
	if err != nil {
		return nil, err
	}
	for i := range quests {
		quests[i].CreateTimeFmt = common.FormatTime(quests[i].CreateTime)
		quests[i].UpdateTimeFmt = common.FormatTime(quests[i].UpdateTime)
	}
	return quests, nil
}

//搜索相关问题
func SearchQuestions(pageSize, pageNo int, search string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_question")

	if search != "" {
		cond := orm.NewCondition()
		cond = cond.And("Question__icontains", search).Or("Title__icontains", search).Or("Email__icontains", search)
		qs = qs.SetCond(cond)
	}

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var questions []Question
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).All(&questions)

		if err == nil {
			resultMap["total"] = cnt
			for i, v := range questions {
				v.CreateTimeFmt = common.FormatTime(v.CreateTime)

				questions[i] = v
			}
			resultMap["data"] = questions
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
	orm.RegisterModel(new(Question))
}
