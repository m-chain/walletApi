package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/addContact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/deleteContact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"],
        beego.ControllerComments{
            Method: "QueryAll",
            Router: `/getContactAll/:address`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"],
        beego.ControllerComments{
            Method: "Query",
            Router: `/getContactInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiContactController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/updateContact`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/addQuestion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/deleteQuestion/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"],
        beego.ControllerComments{
            Method: "QueryAll",
            Router: `/getQuestionAll`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"],
        beego.ControllerComments{
            Method: "Query",
            Router: `/getQuestionInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiQuestionController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/updateQuestion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "GetSignStatus",
            Router: `/getSignStatus/:qrCode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "QueryAllSignInfos",
            Router: `/queryAllSignInfos`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "QueryConfirmInfo",
            Router: `/queryConfirmInfo/:qrCode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "QuerySignInfo",
            Router: `/querySignInfo/:qrCode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiSignController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/updateSignData`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiVersionController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiVersionController"],
        beego.ControllerComments{
            Method: "GetVersionLogs",
            Router: `/GetVersionLogs`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:ApiVersionController"] = append(beego.GlobalControllerRouter["walletApi/src/api:ApiVersionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/getVersionInfo/:platType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryAllTokens",
            Router: `/queryAllTokens`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryBalance",
            Router: `/queryBalance/:address`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryChargeGas",
            Router: `/queryChargeGas`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryContractInfo",
            Router: `/queryContractInfo/:address/:version`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryContractList",
            Router: `/queryContractList/:address`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryJpushMsgInfo",
            Router: `/queryJpushMsgInfo/:msgId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryJpushMsgList",
            Router: `/queryJpushMsgList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryMasterTokenInfo",
            Router: `/queryMasterTokenInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryPublishCCRequireNum",
            Router: `/queryPublishCCRequireNum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryPublishTokenRequireNum",
            Router: `/queryPublishTokenRequireNum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryReturnGasConfig",
            Router: `/queryReturnGasConfig`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryReturnItemsOfWait",
            Router: `/queryReturnItemsOfWait`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryReturnNumOfWait",
            Router: `/queryReturnNumOfWait`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QuerySysMsgInfo",
            Router: `/querySysMsgInfo/:msgId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QuerySysMsgList",
            Router: `/querySysMsgList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryTokenInfo",
            Router: `/queryTokenInfo/:tokenId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryTransferCount",
            Router: `/queryTransferCount/:address`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryTransferDetails",
            Router: `/queryTransferDetails`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryTransferList",
            Router: `/queryTransferList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryTransferListByAddress",
            Router: `/queryTransferListByAddress`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryTransferListByToken",
            Router: `/queryTransferListByToken`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryTransferStatistics",
            Router: `/queryTransferStatistics`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "QueryVestingInfo",
            Router: `/queryVestingInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "Transfer",
            Router: `/sendRawTransaction`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "UpdatePubkey",
            Router: `/updateWalletPubkey`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["walletApi/src/api:WalletController"] = append(beego.GlobalControllerRouter["walletApi/src/api:WalletController"],
        beego.ControllerComments{
            Method: "WalletIsExist",
            Router: `/walletIsExist/:address`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
