package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"log"
)

const (
	MID = 11720660 // 我的b站mid
)

func main() {
	// 1. 获取我的所有视频（bvid列表）
	bvidList := GetBvidListByMid(MID)
	videoNums := len(bvidList)
	fmt.Println("视频总数：", videoNums)
	// 2. 获取所有视频评论列表
	bvidToComments := GetCommentListByBvidList(bvidList)
	// 3. 获取所有视频弹幕列表
	bvidToDanmu := GetDanmuListByBvidList(bvidList)
	// 4. 分别分词，写入本地文件
	jieba := gojieba.NewJieba()
	defer jieba.Free()
	// 4.1 写入评论
	for _, comments := range bvidToComments {
		for i := range comments {
			if NeedSkip(comments[i]) {
				continue
			}
			words := jieba.Cut(comments[i], true)
			err := WriteStringsToFile(words, "./dictionary/comments.txt")
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
	// 4.2 写入弹幕
	for _, danmuList := range bvidToDanmu {
		for i := range danmuList {
			if NeedSkip(danmuList[i]) {
				continue
			}
			words := jieba.Cut(danmuList[i], true)
			err := WriteStringsToFile(words, "./dictionary/danmu.txt")
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
