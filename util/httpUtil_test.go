package util

import "testing"

func TestPostUrlRetry(t *testing.T) {
	url := "http://127.0.0.1:10000/new_address"
	param:=make(map[string]string)
	param["address"] = "hello shalom!"
	PostUrl(url,param,nil,map[string]string{"Connection": "close"})
}
