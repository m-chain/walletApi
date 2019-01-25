package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	_ "walletApi/routers"
	"walletApi/src/common"
	"walletApi/src/model"
	"walletApi/src/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/beego/i18n"
)

//初始化
func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")

	//让beego也采用+8时区的时间
	//Beego的ORM插入Mysql后，时区不一致的解决方案
	//1、orm.RegisterDataBase("default", "mysql", "root:LPET6Plus@tcp(127.0.0.1:18283)/lpet6plusdb?charset=utf8&loc=Local")
	//2、orm.RegisterDataBase("default", "mysql", "db_test:dbtestqwe321@tcp(127.0.0.1:3306)/db_test?charset=utf8&loc=Asia%2FShanghai")
	//orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	//用户名:密码@tcp(url地址)/数据库
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8&loc=Local"
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)

	fmt.Printf("数据库连接成功！%s\n", conn)

}

func main() {
	//简单的设置 Debug 为 true 打印查询的语句,可能存在性能问题，不建议使用在产品模式
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.BConfig.WebConfig.Session.SessionName = "walletApisessionID"

	var FilterUser = func(ctx *context.Context) {
		noValidateMap := map[string]string{
			"/Home/Login":     "GET/POST",
			"/Home/InitReg":   "GET/POST",
			"/Home/InitPass":  "GET/POST",
			"/Home/Reg":       "POST",
			"/Home/UpdatePwd": "POST",
		}
		fmt.Println("****", ctx.Request.RequestURI, ctx.Request.Method)
		requireValidate := true
		for k, v := range noValidateMap {
			result := strings.Contains(v, ctx.Request.Method)
			if ctx.Request.RequestURI == k && result {
				requireValidate = false
				break
			}
		}
		if requireValidate {
			u, ok := ctx.Input.Session(common.USER_INFO).(*model.User)
			if !ok {
				ctx.Redirect(302, "/")
			}
			//必须是管理员才能操作
			if strings.Contains(ctx.Request.RequestURI, "/User/DeleteUser") || strings.Contains(ctx.Request.RequestURI, "/User/AddUser") {
				if u.UserName != "admin" {
					result := new(model.Result)
					result.Status = 1
					result.Msg = "无权限进行此操作!"
					jsonBytes, _ := json.Marshal(result)
					ctx.ResponseWriter.Write(jsonBytes)
					return
				}
			}
		}

	}
	beego.InsertFilter("/Home/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/Category/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/Goods/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/Order/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/User/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/File/*", beego.BeforeRouter, FilterUser)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	i18n.SetMessage("en-US", "conf/locale_zh-TW.ini")
	beego.AddFuncMap("i18n", i18n.Tr)

	runTask()

	beego.Run()
}

func runTask() {
	isRunning := false //是否正在运行
	isPause := false

	var blockNum int64
	blockNum = model.QueryMaxBlockNum()
	blockNum++

	//1、新建任务
	tk := toolbox.NewTask("myTask", "0/1 * * * * *", func() error {
		if !isRunning {
			isRunning = true
			err := service.SyncBlockInfo(blockNum)
			if err == nil {
				blockNum++
			}
			isRunning = false
		}
		return nil
	})

	tk2 := toolbox.NewTask("myTask2", "0/1 * * * * *", func() error {
		if !isPause {
			service.SyncTransferStatus(func(err error) {
				if err != nil {
					fmt.Println(err.Error())
					isPause = true
					//休眠1秒中
					time.Sleep(time.Duration(1) * time.Second)
					isPause = false
				}
			})
		}
		return nil
	})
	//2、运行任务
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	err = tk2.Run()
	if err != nil {
		fmt.Println(err)
	}
	//３、对任务进行管理
	toolbox.AddTask("myTask", tk)
	toolbox.AddTask("myTask2", tk2)
	toolbox.StartTask()
	//4、信息任务
	//time.Sleep(6 * time.Second)
	//toolbox.StopTask()
}

/*
var UrlManager = func(ctx *context.Context) {
        //数据库读取全部的url mapping数据
        urlMapping := model.GetUrlMapping()
        for baseurl,rule:=range urlMapping {
            if baseurl == ctx.Request.RequestURI {
                    ctx.Input.RunController = rule.controller
                    ctx.Input.RunMethod = rule.method
                    break
            }
        }
    }

    beego.InsertFilter("/*",beego.BeforeRouter,UrlManager)
*/
