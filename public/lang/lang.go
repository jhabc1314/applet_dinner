package lang

import "github.com/jhabc1314/applet_dinner/configs"

var messages map[string] string
type Lang interface {
	Content()
}

func Msg(key string) string {
	if msg, ok := messages[key]; ok {
		return msg
	}
	return "抱歉，系统异常(sorry,sys error)"
}

func init() {
	if configs.Lang == "cn" {
		messages = CnMessages()
	} else if configs.Lang == "en" {
		messages = EnMessages()
	}
}


