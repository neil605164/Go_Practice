package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {

	// 用完後關閉連線
	Conn()

	scan()
}

func Conn() *redis.Client {
	// 建立 Redis 客戶端
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 伺服器的位址
		Password: "",               // 伺服器的密碼，如果有的話
		DB:       0,                // 使用的資料庫
	})

	fmt.Println("連線建立成功")

	return client
}

func scan() {
	now := time.Now()
	var keys []string
	iter := client.Scan(0, "example:hello:*", 5000).Iterator()
	for iter.Next() {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	client.Del(keys...)
	fmt.Printf("程式執行時間: %s\n", time.Since(now))
}
