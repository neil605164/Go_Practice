package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// 一片土地
	theMine := []string{"rock", "ore", "rock", "ore", "rock"}

	// 整理出一堆礦坑
	foundOre := finder(theMine)
	fmt.Println("找到礦物 ---->", foundOre)

	// 整理出一堆礦物
	minerOre := miner(foundOre)
	fmt.Println("挖好礦物 ---->", minerOre)

	// 提煉成黃金
	gold := smelter(minerOre)
	fmt.Println("提煉好礦物 ---->", gold)
}

// 找礦物者
func finder(mines []string) (oreMines []string) {
	for i, mine := range mines {
		if mine == "ore" {
			log.Println("開始找礦坑")
			oreMines = append(oreMines, mine)
			log.Println("找到第", i, "礦坑")
		}
	}
	return
}

// 挖礦物者
func miner(mines []string) (oreMines []string) {
	for _, mine := range mines {
		log.Println("開始挖掘一個礦石")
		time.Sleep(time.Second)
		oreMines = append(oreMines, mine)
		log.Println("挖好一顆礦石")
	}
	return
}

// 煉金人員
func smelter(mines []string) (gold []string) {
	for i, mine := range mines {
		log.Println("開始提煉第", i+1, "顆礦物")
		time.Sleep(time.Second)
		gold = append(gold, mine)
		log.Println("提煉完成第", i+1, "顆礦物")
	}
	return
}
