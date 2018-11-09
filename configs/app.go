
package configs

//基本配置
const (
	MODE = "develop" //当前开发环境 product 生产环境
	TIME_FORMAT  = "2006-01-02 15:04:05"
	INFO = iota
	WARNING
	ERROR
)

var (
	Lang = "cn" //当前环境语言
	Db = db {
		Type : "mysql",
		Database : "",
		UserName : "jh",
		Port : 3309,
		Pwd :"",
	}
	LogFile = map[int]string{INFO:"./logs/info.log", WARNING:"./logs/warning.log", ERROR:"./logs/error.log"}
)


/**
* 数据库配置
*/
type db struct {
	Type string
	Database string
	UserName string
	Port int
	Pwd string
}

type redis struct {
	Ip string
}

type LogFormat struct {
	Date string
	LogContent string
}