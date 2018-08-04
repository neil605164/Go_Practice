package main

import (
	"log"
	"time"
)

func main() {
	// 一片土地
	theMine := []string{"rock", "ore", "rock", "ore", "rock"}
	// 宣告搜尋礦坑的通道
	foundOerChannel := make(chan string)
	// 宣告找到礦石的通道
	minedOerChannel := make(chan string)

	// 整理出一堆礦坑
	// foundOre := finder(theMine)
	// fmt.Println("找到礦物 ---->", foundOre)
	go finder(theMine, foundOerChannel)

	// 整理出一堆礦物
	// minerOre := miner(foundOre)
	// fmt.Println("挖好礦物 ---->", minerOre)
	go miner(foundOerChannel, minedOerChannel)

	// 提煉成黃金
	// gold := smelter(minerOre)
	// fmt.Println("提煉好礦物 ---->", gold)
	smelter(minedOerChannel)
}

// 找礦物者
func finder(mines []string, foundOerChannel chan string) {
	for i, mine := range mines {
		if mine == "ore" {
			log.Println("開始找礦坑")
			// oreMines = append(oreMines, mine)
			foundOerChannel <- mine
			log.Println("找到第", i, "礦坑")
		}
	}

	// 當沒有礦區要尋找礦場時，主動關閉通道
	close(foundOerChannel)
	return
}

// 挖礦物者
func miner(foundOerChannel, minedOerChannel chan string) {
	i := 0
	for {
		mine := <-foundOerChannel
		// 判斷是否還有礦區要挖礦
		if mine != "ore" {

			// 當沒有礦要挖掘時，主動關閉通道
			close(minedOerChannel)
			break
		}
		log.Println("開始挖掘第", i+1, "顆礦石")
		time.Sleep(time.Second)

		minedOerChannel <- mine
		log.Println("挖好第", i+1, "顆礦石")
		i++
	}
	// for _, mine := range mines {
	// 	log.Println("開始挖掘一個礦石")
	// 	time.Sleep(time.Second)
	// 	oreMines = append(oreMines, mine)
	// 	log.Println("挖好一顆礦石")
	// }
	// return
}

// 煉金人員
func smelter(minedOerChannel chan string) {
	i := 0
	for {
		mine := <-minedOerChannel
		// 判斷是否還有礦要提煉
		if mine != "ore" {
			break
		}
		log.Println("開始提煉第", i+1, "顆礦物")
		time.Sleep(time.Second)
		log.Println("提煉完成第", i+1, "顆礦物")
		i++
	}
	// for i, mine := range mines {
	// 	log.Println("開始提煉第", i+1, "顆礦物")
	// 	time.Sleep(time.Second)
	// 	gold = append(gold, mine)
	// 	log.Println("提煉完成第", i+1, "顆礦物")
	// }
	// return
}
