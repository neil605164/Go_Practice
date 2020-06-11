package business

import "fmt"

type Request struct{}

// BInter interface 接口
type BInter interface {
	Add(a, b int8) int8
	Api() string
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

// Api 呼叫 api
func (r *Request) Api() string {
	fmt.Println("I'm in business")
	return ""
}
