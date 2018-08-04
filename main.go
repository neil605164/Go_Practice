// 1. 啟動一個特定Port服務
// 2. 將訊息資料顯示在頁面上
// 3. 製作Router

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// router setting
	http.HandleFunc("/", homeHandle)
	http.HandleFunc("/api", apiHandle)

	// server port listen
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// 首頁api
func homeHandle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Write([]byte(path))
}

// 測試api
func apiHandle(w http.ResponseWriter, r *http.Request) {
	// 定義元資料
	data := map[string]string{
		"status": "OK",
		"name":   "Neil",
	}

	// 將『任何型態』資料，轉換為json格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("JSON錯誤" + err.Error()))

		return
	}

	// 定義 header 型態
	w.Header().Set("content-type", "application/json")

	// 回傳至web頁面
	w.Write(jsonData)
}
