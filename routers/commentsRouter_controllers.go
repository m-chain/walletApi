package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/addContact`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/deleteContact`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"],
		beego.ControllerComments{
			Method: "QueryAll",
			Router: `/getContactAll/:address`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/getContactInfo/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiContactController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/updateContact`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/addQuestion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/deleteQuestion/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"],
		beego.ControllerComments{
			Method: "QueryAll",
			Router: `/getQuestionAll`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/getQuestionInfo/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiQuestionController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/updateQuestion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiSignController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/addSignData`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiSignController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/getSignStatus/:qrCode`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiSignController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/updateSignData`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiVersionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiVersionController"],
		beego.ControllerComments{
			Method: "GetVersionLogs",
			Router: `/GetVersionLogs`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:ApiVersionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:ApiVersionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getVersionInfo/:platType`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:UserController"] = append(beego.GlobalControllerRouter["walletApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "AddUser",
			Router: `/User/AddUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:UserController"] = append(beego.GlobalControllerRouter["walletApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/User/DeleteUser/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:UserController"] = append(beego.GlobalControllerRouter["walletApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "InitAdd",
			Router: `/User/InitAdd/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:UserController"] = append(beego.GlobalControllerRouter["walletApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "InitList",
			Router: `/User/InitList/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:UserController"] = append(beego.GlobalControllerRouter["walletApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/User/List/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:UserController"] = append(beego.GlobalControllerRouter["walletApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Qrcode",
			Router: `/User/Qrcode`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:VersionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:VersionController"],
		beego.ControllerComments{
			Method: "AddVersion",
			Router: `/Version/AddVersion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:VersionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:VersionController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/Version/DeleteVersion/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:VersionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:VersionController"],
		beego.ControllerComments{
			Method: "InitAdd",
			Router: `/Version/InitAdd/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:VersionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:VersionController"],
		beego.ControllerComments{
			Method: "InitList",
			Router: `/Version/InitList/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["walletApi/controllers:VersionController"] = append(beego.GlobalControllerRouter["walletApi/controllers:VersionController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/Version/List/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
