package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/yanyiwu/gojieba"
	"log"
	"time"
)

const (
	MID              = 11720660                    // 我的b站mid
	CommentsFilePath = "./dictionary/comments.txt" // 评论文件路径
	DanmuFilePath    = "./dictionary/danmu.txt"    // 弹幕文件路径
)

func main() {
	// 查看文件是否存在
	if !CheckFileIsExist(CommentsFilePath) || !CheckFileIsExist(DanmuFilePath) {
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
				words := jieba.CutForSearch(comments[i], true)
				err := WriteStringsToFile(words, CommentsFilePath)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
		// 4.2 写入弹幕
		for _, danmuList := range bvidToDanmu {
			for i := range danmuList {
				words := jieba.Cut(danmuList[i], true)
				err := WriteStringsToFile(words, DanmuFilePath)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
	// 5. 读出文件
	t1 := time.Now()
	color.Blue.Println("正在导入原材料...")
	comments := ReadFileToListByLine(CommentsFilePath)
	danmu := ReadFileToListByLine(DanmuFilePath)
	words := append([]string{}, comments...)
	words = append([]string{}, danmu...)
	color.Blue.Println("✅ 完成导入，共计词语：", len(words), "条, 耗时：", time.Since(t1))
	t2 := time.Now()
	words = FilterWords(words)
	color.Blue.Println("✅ 完成过滤，共计词语：", len(words), "条, 耗时：", time.Since(t2))
	StartGame(words)
}
