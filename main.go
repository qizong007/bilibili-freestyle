package main

import "fmt"

const (
	BVID = "BV1SL411x7aT"
)

func main() {
	//var mid int64 = 11720660
	//list := GetBvidListByMid(mid)
	//for i := range list {
	//	fmt.Println(list[i])
	//}

	comments := GetCommentsByAid("BV1FT4y1d7pz")
	for i := range comments {
		fmt.Println(comments[i])
	}

	//videoInfo := GetVideoInfoByBvid(BVID)
	//cid := videoInfo.CID
	//fmt.Println("cid", cid)
	//list, err := GetDanmuByCid(cid)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for i := range list {
	//	fmt.Println(list[i])
	//}
}
