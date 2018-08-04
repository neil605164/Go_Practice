// 1. 啟動一個特定Port服務
// 2. 將訊息資料顯示在頁面上
// 3. 製作Router

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listen to 8080 Port")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
