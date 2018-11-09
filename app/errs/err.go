package errs

import (
	"encoding/json"
	"fmt"
	"github.com/jhabc1314/applet_dinner/configs"
	"os"
	"time"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		//Write(configs.ERROR, err)
	}
}

func Write(logLevel int, err error) {
	logFormat := configs.LogFormat{Date:time.Now().Format(configs.TIME_FORMAT), LogContent:err.Error()}
	f, err := os.Open(configs.LogFile[logLevel])
	HandleError(err)
	defer f.Close()
	b, err := json.Marshal(logFormat)
	f.Write(b)
}