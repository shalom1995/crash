package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	//LOG_PATH = "./log_deal_backup.txt"
	//LOG_PATH = "../log_spider_backup.txt"
	//LOG_PATH = "../backups/log_main_spider.txt"
	LOG_PATH             = "../address/log_address_bsc_1611216890.log"
	LOG_PATH_address     = "../address/log_address_bsc_1611318096.log"
	LOG_PATH_address_sub = "../address-sub/log_address-sub_bsc_1611318063.log"
	LOG_PATH_collection   = "../collection/log_collection_bsc_1626012082.log"
	LOG_PATH_assets1155  = "../assets1155/log_assets1155_main_1612766448.log"
)

type GetErc struct {
	UserAddr     string `json:"userAddr"`
	ContractAddr string `json:"contractAddr"`
	ChainID      int    `json:"chainID"`
	Message      string `json:"message"`
}

type GetAllNFT struct {
	UserAddr string `json:"userAddr"`
	ChainID  int    `json:"chainID"`
	Message  string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

func main() {
	//go handle(LOG_PATH_address)
	//go handle(LOG_PATH_address_sub)
	handle(LOG_PATH_collection)
	//handle(LOG_PATH_assets1155)

	//c := make(chan int)
	//<-c
}

func handle(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("openFile failed error:", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	url := "开始分析swap交易是否是闪电贷"
	Contain(r, url, path+url+".log")
	//DisContain(r,"solidity.ABIDecodeString error",LOG_PATH+"_deal.log")
	//JsonMar(r)
}

func OpenLogFile() *bufio.Reader {
	file, err := os.OpenFile(LOG_PATH, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("openFile failed error:", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	return r
}

func Contain(reader *bufio.Reader, subStr string, dealedPath string) {
	requestLog, err := os.Create(dealedPath)
	if err != nil {
		log.Fatal(err)
	}
	defer requestLog.Close()
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}

		if strings.Contains(string(buf), subStr) {
			_, err := requestLog.Write(buf)
			if err != nil {
				log.Fatal("write file failed", err)
			}
		}
	}
}

func DisContain(reader *bufio.Reader, subStr string, dealedPath string) {
	requestLog, err := os.Create(dealedPath)
	if err != nil {
		log.Fatal(err)
	}
	defer requestLog.Close()
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}

		if !strings.Contains(string(buf), subStr) {
			_, err := requestLog.Write(buf)
			if err != nil {
				log.Fatal("write file failed", err)
			}
		}
	}
}

func JsonMar(reader *bufio.Reader, dealedPath string) {
	requestLog, err := os.Create(dealedPath)
	if err != nil {
		log.Fatal(err)
	}
	defer requestLog.Close()
	msg := Message{}
	i := 0
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}
		err = json.Unmarshal(buf, &msg)
		if err != nil {
			fmt.Println("json.Unmarshal err: ", err)
		}
		fmt.Println(msg.Message)
		i++
		if i == 1 {
			break
		}
	}
}
