package main

import (
	"fmt"
	"github.com/gookit/color"
	"math/rand"
	"time"
)

var mq chan bool

func StartGame(words []string) {
	color.Green.Println("-----------------------------")
	color.Green.Println("｜     🎤 Freestyle         ｜")
	color.Green.Println("-----------------------------")
	color.Green.Println(" 🕶 你准备好了吗！！！！！！！！！")
	color.Green.Println("（ ⌨️ 输入任意字符开始/结束）")
	ans := ""
	fmt.Scanln(&ans)
	if ans != "" {
		num := len(words)
		mq = make(chan bool)
		// 用于暂停游戏
		go func() {
			// 开局后三秒才能暂停
			time.Sleep(time.Second * 3)
			in := ""
			fmt.Scanln(&in)
			mq <- true
		}()
		for {
			select {
			case <-mq:
				color.Red.Println("游戏结束！😈")
				return
			default:
				r := rand.New(rand.NewSource(time.Now().Unix()))
				color.Cyan.Println(words[r.Intn(num)])
				time.Sleep(time.Second * 5)
			}
		}
	}
}
