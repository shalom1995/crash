package logger

import (
	"crashPri/util"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T)  {
	dir, _ := util.GetAppPath()
	logPath := dir + `\log\log.txt`

	_, err := os.Create(logPath)
	if err != nil {
		log.Fatal("create log file failed, error: ", err)
	}
}
