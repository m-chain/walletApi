package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "DeleteQuestion",
            Router: `/Question/DeleteQuestion/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/Question/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/Question/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:QuestionController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/Question/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"],
        beego.ControllerComments{
            Method: "AddMessage",
            Router: `/Message/AddMessage`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"],
        beego.ControllerComments{
            Method: "DeleteMessage",
            Router: `/Message/DeleteMessage/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/Message/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/Message/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:SysMsgController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/Message/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/User/AddUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "CCQrcode",
            Router: `/User/CCQrcode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/User/DeleteUser/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/User/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "InitCCQrcode",
            Router: `/User/InitCCQrcode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/User/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/User/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:UserController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:UserController"],
        beego.ControllerComments{
            Method: "Qrcode",
            Router: `/User/Qrcode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:VersionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:VersionController"],
        beego.ControllerComments{
            Method: "AddVersion",
            Router: `/Version/AddVersion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:VersionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:VersionController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/Version/DeleteVersion/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:VersionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:VersionController"],
        beego.ControllerComments{
            Method: "InitAdd",
            Router: `/Version/InitAdd/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:VersionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:VersionController"],
        beego.ControllerComments{
            Method: "InitList",
            Router: `/Version/InitList/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/controller:VersionController"] = append(beego.GlobalControllerRouter["walletApi/src/controller:VersionController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/Version/List/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
