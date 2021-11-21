package main

import (
	"encoding/json"
	"fmt"
	"github.com/iyear/biligo"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	GetVideoInfoUrl = "https://api.bilibili.com/x/space/arc/search"
	GetCommentUrl   = "https://api.bilibili.com/x/v2/reply"
)

// GetBvidListByMid 通过mid获取bvid列表
func GetBvidListByMid(mid int64) []string {
	vlist := getVlistFromMid(mid)
	return getBvidListFromVlist(vlist)
}

func getVlistFromMid(mid int64) []interface{} {
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

func getVideoInfoByBvid(bvid string) *biligo.VideoInfo {
	var videoInfo *biligo.VideoInfo
	client := biligo.NewCommClient(&biligo.CommSetting{})
	videoInfo, err := client.VideoGetInfo(biligo.BV2AV(bvid))
	if err != nil {
		fmt.Println(err)
	}
	return videoInfo
}

// GetDanmuListByBvidList 通过bvid列表获取相对应的弹幕
func GetDanmuListByBvidList(bvidList []string) map[string][]string {
	videoNums := len(bvidList)
	bvidToDanmu := map[string][]string{}
	danmuMutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(videoNums)
	for i := 0; i < videoNums; i++ {
		go func(bvid string) {
			defer wg.Done()
			videoInfo := getVideoInfoByBvid(bvid)
			danmuList, err := GetDanmuByCid(videoInfo.CID)
			if err != nil {
				fmt.Println(err)
				return
			}
			danmuMutex.Lock()
			defer danmuMutex.Unlock()
			bvidToDanmu[bvid] = danmuList
		}(bvidList[i])
	}
	wg.Wait()
	return bvidToDanmu
}

// GetCommentListByBvidList 通过bvid列表获取相对应的评论
func GetCommentListByBvidList(bvidList []string) map[string][]string {
	videoNums := len(bvidList)
	bvidToComments := map[string][]string{}
	commentsMutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(videoNums)
	for i := 0; i < videoNums; i++ {
		go func(bvid string) {
			defer wg.Done()
			comments := getCommentsByBvid(bvid)
			commentsMutex.Lock()
			defer commentsMutex.Unlock()
			bvidToComments[bvid] = comments
		}(bvidList[i])
	}
	wg.Wait()
	return bvidToComments
}

// GetCommentsByBvid 通过bvid获取评论列表
func getCommentsByBvid(bvid string) []string {
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
