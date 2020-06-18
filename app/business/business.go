package business

import (
	"Go_Practice/app/repository"
	"encoding/json"
	"fmt"
	"log"

	"git.cchntek.com/rd3-pkg/go-curl"
)

type Request struct {
	Redis repository.IRedis // 定義繼承 interface 接口
}

// IBusiness interface 接口
type IBusiness interface {
	Add(a, b int64) int64
	Api(a, b int) int
	GetRedis(key string) (value string, err error)
}

func NewBusiness() IBusiness {
	// 初始化
	return &Request{
		Redis: repository.RedisIns(),
	}
}

// Add 加法
func (r *Request) Add(a, b int64) int64 {
	fmt.Println("I'm in add business")
	return a + b
}

type Resp struct {
	Res int
}

// Api 呼叫 api
func (r *Request) Api(a, b int) int {
	fmt.Println("I'm in api business")

	url := fmt.Sprintf("http://localhost:7897/service?a=%v&b=%v", a, b)

	fmt.Println(url)

	// 創建一個 Request 實例
	req := curl.NewRequest()

	// API 請求
	resp, err := req.
		SetURL(url).
		Get()

	// 錯誤處理
	if err != nil {
		log.Fatal("Error Msg", err)
	}

	// 取回傳內容
	res := Resp{}

	if err := json.Unmarshal(resp.Body, &res); err != nil {
		fmt.Println(err)
	}

	return res.Res
}

// GetRedis 取 redis 值
func (r *Request) GetRedis(key string) (value string, err error) {

	value, err = r.Redis.Get(key)
	return
}
