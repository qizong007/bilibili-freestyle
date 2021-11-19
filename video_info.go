package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyear/biligo"
	"io/ioutil"
	"net/http"
)

const (
	GetVideoInfoUrl = "https://api.bilibili.com/x/space/arc/search"
	GetCommentUrl   = "https://api.bilibili.com/x/v2/reply"
)

func GetBvidListByMid(mid int64) []string {
	vlist := GetVlistFromMid(mid)
	return getBvidListFromVlist(vlist)
}

func GetVlistFromMid(mid int64) []interface{} {
	// pn = page, ps = limit
	url := fmt.Sprintf("%s?mid=%d&pn=1&ps=50", GetVideoInfoUrl, mid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var ret map[string]interface{}
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	vlist := ret["data"].(map[string]interface{})["list"].(map[string]interface{})["vlist"]
	return vlist.([]interface{})
}

func getBvidListFromVlist(vlist []interface{}) []string {
	bvList := make([]string, len(vlist))
	for i, video := range vlist {
		bvList[i] = video.(map[string]interface{})["bvid"].(string)
	}
	return bvList
}

func GetVideoInfoByBvid(bvid string) *biligo.VideoInfo {
	var videoInfo *biligo.VideoInfo
	client := biligo.NewCommClient(&biligo.CommSetting{})
	videoInfo, err := client.VideoGetInfo(biligo.BV2AV(bvid))
	if err != nil {
		fmt.Println(err)
	}
	return videoInfo
}

func GetCommentsByAid(bvid string) []string {
	// pn = page
	url := fmt.Sprintf("%s?pn=1&type=1&oid=%d&sort=2", GetCommentUrl, biligo.BV2AV(bvid))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var ret map[string]interface{}
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	replies := ret["data"].(map[string]interface{})["replies"].([]interface{})
	comments := make([]string, len(replies))
	for i, reply := range replies {
		comments[i] = reply.(map[string]interface{})["content"].(map[string]interface{})["message"].(string)
	}
	return comments
}
