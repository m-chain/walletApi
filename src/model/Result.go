package model

type Result struct {
	Status int64
	Msg    string
	Data   interface{}
}

type VersionResult struct {
	Status bool    `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string  `json:"msg" description:"status为false时的错误信息"`
	Data   Version `json:"data" description:"版本信息"`
}

type VersionListResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Total  int       `json:"total"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   []Version `json:"data" description:"版本日志列表"`
}

type JpushMsgListResult struct {
	Status bool           `json:"status" description:"true 表示成功；false 表示失败"`
	Total  int            `json:"total"`
	Msg    string         `json:"msg" description:"status为false时的错误信息"`
	Data   []JpushMessage `json:"data" description:"交易消息列表"`
}

type SysMsgListResult struct {
	Status bool         `json:"status" description:"true 表示成功；false 表示失败"`
	Total  int          `json:"total"`
	Msg    string       `json:"msg" description:"status为false时的错误信息"`
	Data   []SysMessage `json:"data" description:"系统消息列表"`
}

type SysMsgResult struct {
	Status bool        `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string      `json:"msg" description:"status为false时的错误信息"`
	Data   *SysMessage `json:"data" description:"系统消息"`
}

type TokenInfoResult struct {
	Status bool   `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string `json:"msg" description:"status为false时的错误信息"`
	Data   *Token `json:"data" description:"token信息"`
}

type TokenListResult struct {
	Status bool    `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string  `json:"msg" description:"status为false时的错误信息"`
	Data   []Token `json:"data" description:"token列表"`
}

type VestingListResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   []Vesting `json:"data" description:"股权信息列表"`
}

type ReturnListResult struct {
	Status bool          `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string        `json:"msg" description:"status为false时的错误信息"`
	Data   []VestingItem `json:"data" description:"返还记录列表"`
}

type TransResult struct {
	Status bool         `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string       `json:"msg" description:"status为false时的错误信息"`
	Data   *Transaction `json:"data" description:"交易信息"`
}

type StatisticsInfo struct {
	OutAmount int `json:"outAmount" description:"转出"`
	InAmount  int `json:"inAmount" description:"转入"`
}

type TransStatisticsResult struct {
	Status bool           `json:"status" description:"true 表示成功；false 表示失败"`
	Data   StatisticsInfo `json:"data" description:"统计信息"`
	Msg    string         `json:"msg" description:"status为false时的错误信息"`
}

type TransListResult struct {
	Status bool          `json:"status" description:"true 表示成功；false 表示失败"`
	Total  int           `json:"total" description:"交易总数"`
	Data   []Transaction `json:"data" description:"交易列表"`
	Msg    string        `json:"msg" description:"status为false时的错误信息"`
}

type JpushMsgResult struct {
	Status bool          `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string        `json:"msg" description:"status为false时的错误信息"`
	Data   *JpushMessage `json:"data" description:"交易消息"`
}

type JsonResult struct {
	Status bool   `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string `json:"msg" description:"status为false时的错误信息"`
	Data   string `json:"data" description:"结果"`
}

type ConfirmItem struct {
	ConfirmCount int `json:"confirmCount" description:"已签名的次数"`
	Threshold    int `json:"threshold" description:"需签名的次数"`
}

type ConfirmResult struct {
	Status bool        `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string      `json:"msg" description:"status为false时的错误信息"`
	Data   ConfirmItem `json:"data" description:"签名情况"`
}

//合约调用返回结果
type CCInvokeResult struct {
	Status          bool   `json:"status" description:"true 表示成功；false 表示失败"`
	Msg             string `json:"msg" description:"status为false时的错误信息"`
	TxId            string `json:"txId" description:"交易ID"`
	TokenId         string `json:"tokenId" description:"token标识"`
	ContractAddress string `json:"contractAddress" description:"合约地址"`
}

type SignInfoResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   *SignData `json:"data" description:"签名数据"`
}

type SignListResult struct {
	Status bool       `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string     `json:"msg" description:"status为false时的错误信息"`
	Total  int        `json:"total"`
	Data   []SignData `json:"data" description:"签名记录"`
}

type QuestionResult struct {
	Status bool     `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string   `json:"msg" description:"status为false时的错误信息"`
	Data   Question `json:"data" description:"意见反馈信息"`
}

type QuestionListResult struct {
	Status bool       `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string     `json:"msg" description:"status为false时的错误信息"`
	Data   []Question `json:"data" description:"意见反馈列表"`
}

type ContactResult struct {
	Status bool    `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string  `json:"msg" description:"status为false时的错误信息"`
	Data   Contact `json:"data" description:"联系人信息"`
}

type ContactListResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   []Contact `json:"data" description:"联系人列表"`
}
type ReturnGasConfigResponse struct {
	Status bool            `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string          `json:"msg" description:"status为false时的错误信息"`
	Data   ReturnGasConfig `json:"data" description:"数据信息"`
}

type ReturnGasConfig struct {
	InitReleaseRatio string `json:"initReleaseRatio" description:"立即返还比例（按百分比）"`
	Interval         string `json:"interval" description:"每次释放间隔时间以秒为单位"`
	ReleaseRatio     string `json:"releaseRatio" description:"释放比例（按百分比）"`
}

type ContractResult struct {
	Status bool      `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string    `json:"msg" description:"status为false时的错误信息"`
	Data   *Contract `json:"data" description:"合约信息"`
}
type ContractListResult struct {
	Status bool       `json:"status" description:"true 表示成功；false 表示失败"`
	Msg    string     `json:"msg" description:"status为false时的错误信息"`
	Data   []Contract `json:"data" description:"合约列表"`
}
