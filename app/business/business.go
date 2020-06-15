package business

import (
	"encoding/json"
	"fmt"
	"log"

	"git.cchntek.com/rd3-pkg/go-curl"
)

type Request struct{}

// BInter interface 接口
type BInter interface {
	Add(a, b int8) int8
	Api(a, b int8) int
}

func NewBusiness() BInter {
	// 初始化
	return &Request{}
}

// Add 加法
func (r *Request) Add(a, b int8) int8 {
	fmt.Println("I'm in business")
	return a + b
}

type Resp struct {
	Res int
}

// Api 呼叫 api
func (r *Request) Api(a, b int8) int {
	fmt.Println("I'm in business")

	url := fmt.Sprintf("http://localhost:7897/service?a=%d&b=%d", a, b)

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
