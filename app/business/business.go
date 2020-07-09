package business

import (
	"Go_Practice/app/global/structs"
	"Go_Practice/app/repository"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"git.cchntek.com/rd3-pkg/go-curl"
)

type Request struct {
	Redis repository.IRedis // 定義繼承 interface 接口
	DB    repository.IDB
}

// IBusiness interface 接口
type IBusiness interface {
	Add(a, b int64) int64
	Api(a, b int) int
	GetRedis(key string) (value string, err error)
	StoreDBInfo(req structs.RawData) (err error)
	GetDBUserInfo() (resp []structs.DBResp, err error)
}

func NewBusiness() IBusiness {
	// 初始化
	return &Request{
		Redis: repository.RedisIns(),
		DB:    repository.DBIns(),
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

// StoreDBInfo 存值進入 DB
func (r *Request) StoreDBInfo(req structs.RawData) (err error) {

	// 其他邏輯處理
	if req.Age == 27 {
		return errors.New("The age is wrong")
	}

	// 將值存入 DB
	if err = r.DB.SetUserInfo(req); err != nil {
		return
	}
	fmt.Println(req)

	return
}

// GetDBUserInfo 取 DB 資料
func (r *Request) GetDBUserInfo() (resp []structs.DBResp, err error) {

	// 取 DB
	dbData, err := r.DB.GetUserInfo()
	if err != nil {
		return
	}

	fmt.Println(dbData)

	// 整理回傳資料 + 初始化
	resp = []structs.DBResp{}
	for k := range dbData {
		tmp := structs.DBResp{}

		tmp.ID = dbData[k].ID
		tmp.Name = dbData[k].Name
		tmp.Phone = dbData[k].Phone
		tmp.Age = dbData[k].Age

		resp = append(resp, tmp)
	}
	return
}
