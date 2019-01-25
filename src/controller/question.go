package controller

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
	"walletApi/src/model"

	"github.com/astaxie/beego"
)

type QuestionController struct {
	beego.Controller
}

// @router /Question/InitList/ [get]
func (c *QuestionController) InitList() {
	c.TplName = "questionlist.html"
}

// @router /Question/InitAdd/ [get]
func (c *QuestionController) InitAdd() {
	idStr := c.GetString("id")
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		question, _ := model.GetQuestionInfo(id)
		if question != nil {
			c.Data["question"] = question
			if question.ImgName != "" {
				arr := strings.Split(question.ImgName, ",")
				fileMap := make(map[string]string, 0)
				for i, v := range arr {
					//一种是 ascii 字符，另一种为本地编码 (如:utf8) 的字符
					//fmt.Println("Hello, 世界", len("世界"), utf8.RuneCountInString("世界"))
					len := utf8.RuneCountInString(question.BaseUrl)

					r := question.BaseUrl[len-1 : len]
					if r == "/" {
						fileMap[fmt.Sprintf("附件%d", i+1)] = question.BaseUrl + v
					} else {
						fileMap[fmt.Sprintf("附件%d", i+1)] = question.BaseUrl + "/" + v
					}

				}
				c.Data["fileMap"] = fileMap
			}
		}
	}
	c.TplName = "questioninfo.html"
}

// @router /Question/List/ [post]
func (c *QuestionController) List() {

	keyWord := c.GetString("keyWord")

	pageNo, _ := c.GetInt("current")

	rowCount, _ := c.GetInt("rowCount")

	if pageNo == 0 {
		pageNo = 1
	}
	resultMap := model.SearchQuestions(rowCount, pageNo, keyWord)
	c.Data["json"] = map[string]interface{}{"rows": resultMap["data"], "rowCount": rowCount, "current": pageNo, "total": resultMap["total"]}

	c.ServeJSON()
}

// @router /Question/DeleteQuestion/:id [get]
func (c *QuestionController) DeleteQuestion() {
	result := new(model.Result)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	err := model.DeleteQuestion(id)

	if err != nil {
		result.Status = 1
		result.Msg = fmt.Sprintf("删除失败,%s", err.Error())
	} else {
		result.Status = 0
	}

	c.Data["json"] = result
	c.ServeJSON()
}
