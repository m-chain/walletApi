package model

import (
	"errors"
	"fmt"
	"time"
	"walletApi/src/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//应用版本
type AppVersion struct {
	Id            int64     `orm:"auto" from:"id" description:"主键ID"`
	AppName       string    `orm:"size(80)" valid:"Required" form:"appName" description:"应用名称"`
	Version       string    `orm:"size(20)" valid:"Required" form:"version" description:"版本"`
	PlatType      int       `orm:"size(10)" valid:"Required" form:"platType" description:"平台类型 1、android 2、IOS"`
	UpgradeType   int       `orm:"size(10)" valid:"Required" form:"upgradeType" description:"升级类型　1 、强制升级 2、可忽略"`
	IsCurrent     int       `orm:"size(10)" valid:"Required" form:"isCurrent" description:"是否是当前版本 1、表示是 2、表示不是"`
	AppAddr       string    `orm:"size(100)" form:"appAddr" description:"下载地址"`
	AppDesc       string    `orm:"size(100)" form:"appDesc" description:"应用描述"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	CreateTimeFmt string    `orm:"-" description:"创建时间"`
}

type Version struct {
	Id            int64     `json:"id" description:"记录ID"`
	AppName       string    `json:"appName" description:"app名称"`
	Version       string    `json:"version" description:"版本号"`
	PlatType      int       `json:"platType" description:"平台类型 1、andorid 2、Ios"`
	UpgradeType   int       `json:"upgradeType" description:"升级类型　1 、强制升级 2、可忽略"`
	IsCurrent     int       `json:"isCurrent" description:"是否是当前版本 1、表示是 2、表示不是"`
	AppAddr       string    `json:"appAddr" description:"下载地址"`
	Size          string    `orm:"-" json:"size" description:"文件大小"`
	AppDesc       string    `json:"appDesc" description:"app描述"`
	CreateTime    time.Time `json:"createTime" description:"创建时间"`
	CreateTimeFmt string    `json:"createTimeFmt"`
}

func (u *AppVersion) TableName() string {
	return "t_app_version"
}

//添加版本
func AddVersion(version *AppVersion) error {
	o := orm.NewOrm()
	ver := new(AppVersion)
	//未找到记录也会报错
	o.Raw("SELECT * FROM t_app_version WHERE app_name = ? AND plat_type=? AND version=?", version.AppName, version.PlatType, version.Version).QueryRow(ver)
	if ver != nil && ver.AppName != "" {
		return errors.New("版本信息已存在，不能重复添加!")
	}
	beginTx := false
	if version.IsCurrent == 1 {
		o.Begin()
		_, err := o.Raw("UPDATE t_app_version SET is_current = 2 WHERE app_name = ? AND plat_type=?", version.AppName, version.PlatType).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
		beginTx = true
	}
	_, err := o.Insert(version)
	if beginTx {
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}
	return err
}

//更新版本信息
func UpdateVersion(version *AppVersion) error {
	o := orm.NewOrm()
	ver := new(AppVersion)
	o.Raw("SELECT * FROM t_app_version WHERE id<>? AND app_name = ? AND plat_type=? AND version=?", version.Id, version.AppName, version.PlatType, version.Version).QueryRow(ver)
	if ver != nil && ver.AppName != "" {
		return errors.New("修改失败,版本号已存在!")
	}
	beginTx := false
	if version.IsCurrent == 1 {
		o.Begin()
		_, err := o.Raw("UPDATE t_app_version SET is_current = 2 WHERE id<>? AND app_name = ? AND plat_type=?", version.Id, version.AppName, version.PlatType).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
		beginTx = true
	}
	_, err := o.Update(version, "AppName", "Version", "PlatType", "UpgradeType", "IsCurrent", "AppAddr", "AppDesc")
	if beginTx {
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}
	return err
}

//删除版本
func DeleteVersion(id int64) error {
	o := orm.NewOrm()
	version := AppVersion{Id: id}
	err := o.Read(&version)
	if err != nil {
		return err
	}
	if version.IsCurrent == 1 {
		return fmt.Errorf("不能删除正在使用的版本!")
	}
	_, err = o.Delete(&version)
	if err != nil {
		return err
	}
	return nil
}

//根据应用名称、平台类型查询当前使用的版本
func GetCurrentVersion(appName string, platType int64) (*AppVersion, error) {
	o := orm.NewOrm()
	version := new(AppVersion)
	err := o.Raw("SELECT * FROM t_app_version WHERE is_current=1 AND app_name = ? AND plat_type=?", appName, platType).QueryRow(version)
	if err != nil {
		return nil, err
	}
	return version, nil
}

//根据id查询版本详细信息
func (v *AppVersion) GetVersionInfo(id string) (*AppVersion, error) {
	o := orm.NewOrm()
	version := new(AppVersion)
	err := o.Raw("SELECT * FROM t_app_version WHERE id = ?", id).QueryRow(version)
	if err != nil {
		return nil, err
	}
	return version, nil
}

func (v *AppVersion) List(pageSize, pageNo int, search string) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_app_version")

	cond := orm.NewCondition()
	if search != "" {
		if common.IsValidVersion(search) {
			cond = cond.And("Version__exact", search)
		} else {
			cond = cond.And("AppName__icontains", search).Or("AppDesc__icontains", search)
		}
		qs = qs.SetCond(cond)
	}

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var versions []AppVersion
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-Id").All(&versions)

		if err == nil {
			resultMap["total"] = cnt
			for i, v := range versions {
				v.CreateTimeFmt = common.FormatTime(v.CreateTime)
				versions[i] = v
			}
			resultMap["data"] = versions
			return resultMap
		}
	}

	resultMap["total"] = 0
	resultMap["data"] = nil

	return resultMap
}

func ListByPlatType(pageSize, pageNo, platType int) map[string]interface{} {
	o := orm.NewOrm()
	qs := o.QueryTable("t_app_version")

	cond := orm.NewCondition()
	cond = cond.And("PlatType__exact", platType)
	qs = qs.SetCond(cond)

	resultMap := make(map[string]interface{})
	cnt, err := qs.Count()

	if err == nil && cnt > 0 {
		var versions []AppVersion
		var arr []Version
		_, err := qs.Limit(pageSize, (pageNo-1)*pageSize).OrderBy("-CreateTime").All(&versions)

		if err == nil && len(versions) > 0 {
			resultMap["total"] = cnt
			for _, ob := range versions {
				ob.CreateTimeFmt = common.FormatTime(ob.CreateTime)
				var version Version
				version.Id = ob.Id
				version.AppName = ob.AppName
				version.Version = ob.Version
				version.AppAddr = ob.AppAddr
				version.AppDesc = ob.AppDesc
				version.CreateTime = ob.CreateTime
				version.UpgradeType = ob.UpgradeType
				version.CreateTimeFmt = ob.CreateTimeFmt
				version.IsCurrent = ob.IsCurrent
				version.PlatType = ob.PlatType

				arr = append(arr, version)
			}
			resultMap["data"] = arr
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
	orm.RegisterModel(new(AppVersion))
}
