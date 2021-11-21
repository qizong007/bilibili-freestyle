package main

import (
	"compress/flate"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	DanmuXMLUrl = "https://api.bilibili.com/x/v1/dm/list.so"
)

// GetDanmuByCid 通过用户cid获取弹幕
func GetDanmuByCid(cid int64) ([]string, error) {
	url := fmt.Sprintf("%s?oid=%d", DanmuXMLUrl, cid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(flate.NewReader(resp.Body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	reg := regexp.MustCompile(`<d.*?>(.*?)</d>`)
	list := reg.FindAllStringSubmatch(string(data), -1)
	res := make([]string, len(list))
	for i := range list {
		res[i] = list[i][1]
	}
	return res, nil
}
