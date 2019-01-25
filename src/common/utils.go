package common

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/beego/i18n"

	"github.com/astaxie/beego"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
	uuid "github.com/satori/go.uuid"
)

const USER_INFO = "userInfo"

const DATETIMEFORMAT = "2006-01-02 15:04"
const DATETIMEFULLFORMAT = "2006-01-02 15:04:05"
const DATEFORMAT = "2006-01-02"
const PLATSIGN = "platsign"

type SignTipItem struct {
	TipMsg     string `json:"tipMsg" description:"提示信息"`
	SuccessMsg string `json:"successMsg" description:"成功提示信息"`
	FailMsg    string `json:"failMsg" description:"失败提示信息"`
}

var tenToAny map[int64]string = map[int64]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

var SignTips map[int]SignTipItem = map[int]SignTipItem{
	1:  SignTipItem{"tip_sign_deploye_cc", "tip_sign_cc_success", ""},
	2:  SignTipItem{"tip_sign_common", "tip_sign_cc_init_success", ""},
	3:  SignTipItem{"tip_sign_common", "tip_sign_cc_upgrade_success", ""},
	4:  SignTipItem{"tip_sign_token", "tip_sign_token_success", ""},
	5:  SignTipItem{"tip_sign_common", "tip_sign_multi_sign", ""},
	6:  SignTipItem{"tip_sign_common", "tip_sign_success", ""},
	7:  SignTipItem{"tip_sign_common", "tip_sign_addmanager", ""},
	8:  SignTipItem{"tip_sign_common", "tip_sign_replacemanager", ""},
	9:  SignTipItem{"tip_sign_common", "tip_sign_deletemanager", ""},
	10: SignTipItem{"tip_sign_common", "tip_sign_manager_value", ""},
	11: SignTipItem{"tip_sign_common", "tip_sign_token_value", ""},
	12: SignTipItem{"tip_sign_common", "tip_sign_cc_value", ""},
	13: SignTipItem{"tip_sign_common", "tip_sign_returnconfig", ""},
	14: SignTipItem{"tip_sign_common", "tip_sign_cc_delete", ""},
	15: SignTipItem{"tip_sign_common", "tip_sign_master_value", ""},
	16: SignTipItem{"tip_sign_common", "tip_sign_token_icon", ""},
}

type Callback func(result interface{}, err error)

func GetMD5Str(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
	//return string(cipherStr[:])
}

func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	str := uuid.String()
	str = strings.Replace(str, "-", "", -1)
	return str
}

//get real language
func GetLangType(fullName string) string {
	if fullName == "en-US" {
		return "en"
	} else if fullName == "zh-TW" {
		return "tw"
	} else {
		return "zh"
	}
}

//http get request
func HttpGet(url, lang string, callback Callback) {
	lang = GetLangType(lang)

	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		callback(nil, err)
		return
	}
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-type", "application/json")
	request.Header.Add("language", lang)

	//处理返回结果
	resp, err := client.Do(request)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			callback(nil, err)
			return
		}
		callback(body, err)
	} else {
		callback(nil, errors.New("The request failed!"))
	}
}

//http get request
func SyncHttpGet(url, lang string) ([]byte, error) {
	lang = GetLangType(lang)

	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-type", "application/json")
	request.Header.Add("language", lang)

	//处理返回结果
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	} else {
		return nil, errors.New("The request failed!")
	}
}

//http post request
func HttpPost(url, lang string, params map[string]interface{}, callback Callback) {
	lang = GetLangType(lang)

	client := &http.Client{}

	paramBytes, _ := json.Marshal(params)
	reader := bytes.NewBuffer(paramBytes)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-type", "application/json")
	request.Header.Add("language", lang)

	resp, err := client.Do(request)
	if err != nil {
		callback(nil, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			callback(nil, err)
			return
		}
		callback(body, err)
	} else {
		callback(nil, errors.New("The request failed!"))
	}
}

//sync http post request
func SyncHttpPost(url, lang string, params map[string]interface{}) ([]byte, error) {
	lang = GetLangType(lang)

	client := &http.Client{}

	paramBytes, _ := json.Marshal(params)
	reader := bytes.NewBuffer(paramBytes)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-type", "application/json")
	request.Header.Add("language", lang)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	} else {
		return nil, errors.New("The request failed!")
	}
}

//http get request
func HttpGet_old(url string, callback Callback) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}
	if resp.StatusCode == 200 {
		callback(body, err)
	} else {
		callback(body, errors.New("The request failed!"))
	}
}

//http get request
func SyncHttpGet_old(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil, err
	}
	if resp.StatusCode == 200 {
		return body, err
	} else {
		return body, errors.New("The request failed!")
	}
}

//http post request
func HttpPost_old(url string, data url.Values, callback Callback) {
	//data := url.Values{"apikey": {blockNum}, "mobile": {mobile}, "tpl_value": {tpl_value}}
	resp, err := http.PostForm(url, data)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		callback(nil, err)
		return
	}

	if resp.StatusCode == 200 {
		callback(body, err)
	} else {
		callback(body, errors.New("The request failed!"))
	}
}

//sync http post request
func SyncHttpPost_old(url string, data url.Values) ([]byte, error) {
	//data := url.Values{"apikey": {blockNum}, "mobile": {mobile}, "tpl_value": {tpl_value}}
	resp, err := http.PostForm(url, data)
	if err != nil {
		// handle error
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil, err
	}
	if resp.StatusCode == 200 {
		return body, err
	} else {
		return body, errors.New("The request failed!")
	}
}

//发送远程推送消息
func SendMessage(url string, reqData []byte) ([]byte, error) {
	reader := bytes.NewReader(reqData)
	appkey := beego.AppConfig.String("jpushappkey")
	secret := beego.AppConfig.String("jpushsecret")
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	input := []byte(appkey + ":" + secret)
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Basic "+encodeString)
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return respBytes, nil
}

//假定字符串的每节数都在5位以下
func ToNum(a string) string {
	c := strings.Split(a, ".")
	ret := make([]string, len(c))
	r := []string{"", "0", "00", "000", "0000"}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	for i := 0; i < len(c); i++ {
		len := len(c[i])
		ret = append(ret, r[len]+c[i])
	}
	return strings.Join(ret, "")
}

//验证版本格式是否正确
func IsValidVersion(version string) bool {
	//不能省掉括号()
	r, err := regexp.Compile(`(^\d+\.\d+\.\d+$)|(^\d+$)|(^\d+\.\d+$)`)
	if err != nil {
		return false
	}
	return r.MatchString(version)
}

func CmpVersion(a, b string) int {
	_a := ToNum(a)
	_b := ToNum(b)
	//比较字符串的大小即可, 不能转换成数字进行比较
	if _a == _b {
		return 0
	} else if _a > _b {
		return 1
	} else {
		return -1
	}
}

//精确获取两个大整型数据相除的结果
func BigIntDiv(aV, b *big.Int) string {
	bigA := big.NewInt(0)
	ltZero := false
	if aV.Cmp(big.NewInt(0)) == -1 {
		bigA = big.NewInt(0).Abs(aV)
		ltZero = true
	} else {
		bigA = aV
	}
	ip := big.NewInt(1)
	r := ip.Div(bigA, b)

	ip = big.NewInt(1)
	c := ip.Mul(r, b)

	ip = big.NewInt(1)
	d := ip.Sub(bigA, c)
	e := d.Cmp(big.NewInt(0))
	if e > 0 {
		n := len(b.String()) - len(d.String()) - 1
		var buffer bytes.Buffer
		for i := 0; i < n; i++ {
			buffer.WriteString("0")
		}
		buffer.WriteString(d.String())
		if ltZero {
			return fmt.Sprintf("-%v.%s", r, buffer.String())
		}
		return fmt.Sprintf("%v.%s", r, buffer.String())
	}
	if ltZero {
		return fmt.Sprintf("-%v", r)
	}
	return fmt.Sprintf("%v", r)
}

//根据指定的位数生成数字，如3位，会生成对应的1000
func GetNumberWithDigits(n int) *big.Int {
	if n > 0 {
		var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
		buffer.WriteString("1")
		for i := 0; i < n; i++ {
			buffer.WriteString("0")
		}
		result := new(big.Int)
		result.SetString(buffer.String(), 10)
		return result
	}
	return big.NewInt(1)
}

//获取某个月份起始日期和结止日期
func StartAndEndDayOfMonth(yearMonth string) (startDay, endDay string) {
	y, _ := strconv.Atoi(yearMonth[0:4])
	m, _ := strconv.Atoi(yearMonth[5:7])
	m = m + 1
	if m > 12 {
		y++
		m = 1
	}
	mStr := ""
	if m < 10 {
		mStr = fmt.Sprintf("0%d", m)
	} else {
		mStr = fmt.Sprintf("%d", m)
	}
	toBeCharge := fmt.Sprintf("%d-%s-01 00:00:00", y, mStr)                 //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	loc, _ := time.LoadLocation("Local")                                    //重要：获取时区
	theTime, _ := time.ParseInLocation(DATETIMEFULLFORMAT, toBeCharge, loc) //使用模板在对应时区转化为time.time类型

	year, month, _ := theTime.Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0).Format(DATEFORMAT)
	end := thisMonth.AddDate(0, 0, -1).Format(DATEFORMAT)

	return start, end
}

//将20060102150405格式转换成2006-01-02 15:04:05
func FormatDate(dateStr string) string {
	timeLayout := "20060102150405"                               //转化所需模板
	loc, _ := time.LoadLocation("Local")                         //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, dateStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime.Format(DATETIMEFULLFORMAT)
}

//获取两月份之间的年月
func MonthsBetweenYM(startYM, endYM string) []string {
	startY, _ := strconv.Atoi(startYM[0:4])
	startM, _ := strconv.Atoi(startYM[5:7])

	endY, _ := strconv.Atoi(endYM[0:4])
	endM, _ := strconv.Atoi(endYM[5:7])

	var arr []string
	if startY == endY {
		for i := startM; i <= endM; i++ {
			yearMonth := ""
			if i < 10 {
				yearMonth = fmt.Sprintf("%d-0%d", startY, i)
			} else {
				yearMonth = fmt.Sprintf("%d-%d", startY, i)
			}
			arr = append(arr, yearMonth)
		}
	} else if startY < endY {
		i := startY
		for i <= endY {
			s := 1
			e := 12
			if i == startY {
				s = startM
			}
			if i == endY {
				e = endM
			}
			for j := s; j <= e; j++ {
				yearMonth := ""
				if j < 10 {
					yearMonth = fmt.Sprintf("%d%s%d", i, "-0", j)
				} else {
					yearMonth = fmt.Sprintf("%d-%d", i, j)
				}
				arr = append(arr, yearMonth)
				if j == e {
					i++
				}
			}
		}
	}

	return arr
}

// 10进制转16进制
func DecimalToHex(num *big.Int) string {
	return "0x" + decimalToAny(num, 16)
}

// 10进制转任意进制
func decimalToAny(num *big.Int, n int64) string {
	newNumStr := ""
	var remainder *big.Int
	var remainderString string
	for num.Cmp(big.NewInt(0)) != 0 {
		remainder = big.NewInt(1).Mod(num, big.NewInt(n))
		r76 := remainder.Cmp(big.NewInt(76))
		r9 := remainder.Cmp(big.NewInt(9))
		if r76 == -1 && r9 > 0 {
			remainderString = tenToAny[remainder.Int64()]
		} else {
			remainderString = remainder.String()
		}
		newNumStr = remainderString + newNumStr
		num = big.NewInt(1).Div(num, big.NewInt(n))
	}
	if newNumStr == "" {
		newNumStr = "0"
	}
	return newNumStr
}

//文件下载
func DownloadFile(url string, fb func(totalLen, downLen int64)) ([]byte, error) {
	var (
		fsize   int64
		buf     = make([]byte, 32*1024)
		written int64
	)
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	//创建一个http client
	client := new(http.Client)
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	//读取服务器返回的文件大小
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return nil, err
	}

	if resp.Body == nil {
		return nil, errors.New("body is null")
	}
	defer resp.Body.Close()
	//下面是 io.copyBuffer() 的简化版本
	for {
		//读取bytes
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			//写入bytes
			nw, ew := buffer.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
		//没有错误了快使用 callback
		fb(fsize, written)
	}
	return buffer.Bytes(), err
}

//获取文件大小
func GetFileSize(url string) int64 {
	var fsize int64

	//创建一个http client
	client := new(http.Client)
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return 0
	}
	//读取服务器返回的文件大小
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return 0
	}

	defer resp.Body.Close()

	return fsize
}

//文件下载-同步方式
func SyncDownloadFile(url string) ([]byte, error) {
	var (
		buf     = make([]byte, 32*1024)
		written int64
	)
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	//创建一个http client
	client := new(http.Client)
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.Body == nil {
		return nil, errors.New("body is null")
	}
	defer resp.Body.Close()

	//下面是 io.copyBuffer() 的简化版本
	for {
		//读取bytes
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			//写入bytes
			nw, ew := buffer.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return buffer.Bytes(), err
}

func FloatNumber(number string, decimalUnits int) string {
	digits := GetNumberWithDigits(decimalUnits)
	bigAmount := new(big.Int)
	bigAmount.SetString(number, 10)

	end := BigIntDiv(bigAmount, digits)

	return FormatNumber(end)
}

// 16进制转10进制
func HexToBigInt(num string) *big.Int {
	if num == "" || num == "0" {
		return big.NewInt(0)
	}
	r := strings.Index(num, "0x")
	if r == 0 {
		num = num[2:]
	}
	return anyToDecimal(num, 16)
}

//map根据value找key
func findkey(in string) int64 {
	var result int64 = -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func anyToDecimal(num string, n int64) *big.Int {
	newNum := big.NewInt(0)
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := big.NewInt(findkey(value))
		if tmp.Int64() != -1 {
			ip := big.NewInt(1)
			newNum = ip.Mul(tmp, BigIntPow(n, int64(nNum))).Add(ip, newNum)
			nNum = nNum - 1
		} else {
			break
		}
	}
	return newNum
}

/**
* n 如果为16
* m 如果为9
*结果为 16的9次方
 */
func BigIntPow(n, m int64) *big.Int {
	bigSum := big.NewInt(1)
	var i int64
	for i = 0; i < m; i++ {
		bigSum = big.NewInt(1).Mul(bigSum, big.NewInt(n))
	}
	return bigSum
}

//已知时间戳和系统当前时间戳比较
func TimestampCmp(timeStr string) int {

	aBig := new(big.Int)
	aBig.SetString(timeStr, 10)

	bBig := new(big.Int)
	bBig.SetInt64(time.Now().Unix())

	return aBig.Cmp(bBig)
}

//将时间戳转换成日期字符串
func TimestampToStr(seconds int64) string {
	//获取本地location
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	time.LoadLocation("Local")          //重要：获取时区

	//时间戳转日期
	return time.Unix(seconds, 0).Format(timeLayout)
}

//将秒转换成分钟、小时、天、月
func ConverTimestamp(seconds int64, lang string) string {
	if seconds < 60 { //秒
		return fmt.Sprintf("%d%s", seconds, i18n.Tr(lang, "second"))
	} else if seconds < 3600 { //分钟
		return formatMinutes(seconds, lang)
	} else if seconds < 3600*24 { //小时
		return formatHours(seconds, lang)
	} else { //天
		if seconds%(3600*24) == 0 {
			return fmt.Sprintf("%d%s", seconds/(3600*24), i18n.Tr(lang, "day"))
		} else {
			days := seconds / (3600 * 24)
			leaveSeconds := seconds - days*(3600*24)
			hoursStr := formatHours(leaveSeconds, lang)
			return fmt.Sprintf("%d%s%s", days, i18n.Tr(lang, "day"), hoursStr)
		}
	}
}

func formatHours(seconds int64, lang string) string {
	hour := i18n.Tr(lang, "hour")
	minute := i18n.Tr(lang, "minute")
	second := i18n.Tr(lang, "second")
	if seconds < 60 { //秒
		return fmt.Sprintf("0%s0%s%d%s", hour, minute, seconds, second)
	} else if seconds < 3600 {
		return fmt.Sprintf("0%s%s", hour, formatMinutes(seconds, lang))
	} else if seconds < 3600*24 { //小时
		if seconds%3600 == 0 {
			return fmt.Sprintf("%d%s", (seconds / 3600), hour)
		} else {
			hours := seconds / 3600
			leaveSeconds := seconds - hours*3600
			minutesStr := formatMinutes(leaveSeconds, lang)
			return fmt.Sprintf("%d%s%s", hours, hour, minutesStr)
		}
	}
	return ""
}

func formatMinutes(seconds int64, lang string) string {
	minute := i18n.Tr(lang, "minute")
	second := i18n.Tr(lang, "second")
	if seconds < 60 { //秒
		return fmt.Sprintf("0%s%d%s", minute, seconds, second)
	} else {
		if seconds%60 == 0 {
			return fmt.Sprintf("%d%s", (seconds / 60), minute)
		} else {
			minutes := seconds / 60
			leaveSeconds := seconds - minutes*60
			return fmt.Sprintf("%d%s%d%s", minutes, minute, leaveSeconds, second)
		}
	}
}

//签名验证
func Verify(pubKeyHexStr, originStr, signHexStr string) (bool, error) {
	if pubKeyHexStr == "" || originStr == "" || signHexStr == "" {
		return false, nil
	}
	// Decode hex-encoded serialized public key.
	pubKeyBytes, err := hex.DecodeString(pubKeyHexStr)
	if err != nil {
		return false, err
	}
	pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	if err != nil {
		return false, err
	}

	// Decode hex-encoded serialized signature.
	sigBytes, err := hex.DecodeString(signHexStr)
	if err != nil {
		return false, err
	}
	signature, err := btcec.ParseSignature(sigBytes, btcec.S256())
	if err != nil {
		return false, err
	}

	// Verify the signature for the message using the public key.
	originHash := chainhash.HashB([]byte(originStr))

	return signature.Verify(originHash, pubKey), nil
}

//get address by public key
func GetAddress(pubKeyHexStr string) string {
	if pubKeyHexStr == "" {
		return ""
	}
	// Decode hex-encoded serialized public key.
	pubKeyBytes, err := hex.DecodeString(pubKeyHexStr)
	if err != nil {
		return ""
	}
	address, err := btcutil.NewAddressPubKey(pubKeyBytes, &chaincfg.MainNetParams)
	if err != nil {
		return ""
	}

	return beego.AppConfig.String("addressPrefix") + address.EncodeAddress()
}

//格式化数字金额字符串
func FormatNumber(numStr string) string {
	idx := strings.Index(numStr, ".")
	if idx == -1 {
		return numStr
	}
	//123.45
	endPrefix := numStr[idx+1:]
	len := len(endPrefix)
	if len >= 4 {
		return numStr[0:idx] + "." + endPrefix[0:4]
	} else {
		tempLen := 4 - len
		var buffer bytes.Buffer
		for i := 0; i < tempLen; i++ {
			buffer.WriteString("0")
		}
		return numStr + buffer.String()
	}
}

//格式化文件大小
func FormatFileSize(fileSize int64) string {
	tempFloat, _ := strconv.ParseFloat(fmt.Sprintf("%d", fileSize), 64)
	if fileSize < 1024 {
		return fmt.Sprintf("%dB", fileSize)
	} else if fileSize < (1024 * 1024) {
		var temp = tempFloat / 1024
		return fmt.Sprintf("%.2fKB", temp)
	} else if fileSize < (1024 * 1024 * 1024) {
		var temp = tempFloat / (1024 * 1024)
		return fmt.Sprintf("%.2fMB", temp)
	} else {
		var temp = tempFloat / (1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fGB", temp)
	}
}

//将日期转换成字符串
func FormatTime(t time.Time) string {
	local, _ := time.LoadLocation("Local")
	return t.In(local).Format(DATETIMEFORMAT)
}

//将日期转换成字符串
func FormatFullTime(t time.Time) string {
	local, _ := time.LoadLocation("Local")
	return t.In(local).Format(DATETIMEFULLFORMAT)
}
