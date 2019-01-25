package api

import (
	"strconv"
	"walletApi/src/common"
	"walletApi/src/model"
)

// 有关应用版本的相关操作
type ApiVersionController struct {
	BaseController
}

// @Title Get
// @Description 获取某平台的版本信息
// @Param	platType		path 	string	true		"平台类型(1、Android;2、IOS)"
// @Success 200 {object} model.VersionResult
// @Failure 403 :platType不能为空
// @router /getVersionInfo/:platType [get]
func (o *ApiVersionController) Get() {
	platType := o.Ctx.Input.Param(":platType")
	var result model.VersionResult
	if platType != "" {
		_type, _ := strconv.ParseInt(platType, 10, 64)
		ob, err := model.GetCurrentVersion("wallet", _type)
		if err != nil {
			result.Status = false
			result.Msg = err.Error()
		} else {
			result.Status = true

			var version model.Version
			version.Id = ob.Id
			version.AppName = ob.AppName
			version.Version = ob.Version
			version.AppAddr = ob.AppAddr
			version.AppDesc = ob.AppDesc
			version.CreateTime = ob.CreateTime
			version.UpgradeType = ob.UpgradeType
			version.CreateTimeFmt = common.FormatFullTime(version.CreateTime)
			version.IsCurrent = ob.IsCurrent
			version.PlatType = ob.PlatType

			if version.AppAddr != "" {
				fileSize := common.GetFileSize(version.AppAddr)
				version.Size = common.FormatFileSize(fileSize)
			}

			result.Data = version
		}
	} else {
		result.Status = false
		result.Msg = o.Tr("device_type_empty")
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// @Title List
// @Description 获取版本日志信息
// @Param	platType		formData 	int	true		"平台类型(1、Android;2、IOS)"
// @Param	page		formData 	int	true		"第几页"
// @Param	pageSize		formData 	int	true		"每页多少条记录""
// @Success 200 {object} model.VersionListResult
// @router /GetVersionLogs [post]
func (o *ApiVersionController) GetVersionLogs() {
	platType, _ := o.GetInt("platType", 0)
	page, _ := o.GetInt("page", 1)
	pageSize, _ := o.GetInt("pageSize", 20)
	resultMap := model.ListByPlatType(pageSize, page, platType)
	var result model.VersionListResult
	if resultMap["data"] == nil {
		result.Status = false
		result.Msg = o.Tr("not_found_msg")
		o.Data["json"] = result
	} else {
		resultMap["status"] = true
		o.Data["json"] = resultMap
	}
	o.ServeJSON()
}
