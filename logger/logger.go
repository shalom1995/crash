package logger

import (
	"crashPri/util"
	"fmt"
	"github.com/rs/zerolog"
	"log"
	"os"
	"time"
)

var Logger zerolog.Logger

func init() {
	dir, err := util.GetAppPath()
	if err != nil {
		log.Fatal("GetAppPath failed, error: ", err)
	}

	logPath := dir + fmt.Sprintf("/log_%d.log",  time.Now().Unix())
	file := &os.File{}
	if util.FileNotExist(logPath) {
		file, err = os.Create(logPath)
		if err != nil {
			log.Fatal("create log file failed, error: ", err)
		}
	}else {
		file, err = os.OpenFile(logPath,os.O_RDWR,0666)
		if err != nil {
			log.Fatal("open log file failed, error: ", err)
		}
	}

	Logger = zerolog.New(file).With().Timestamp().Logger()
	Logger.Level(zerolog.InfoLevel)
}
