package model

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// redisPool 存放redis連線池的全域變數
var redisPool *redis.Pool

func init() {
	redisPool = &redis.Pool{
		MaxIdle:     100,               // int 最大可允許的閒置連線數
		MaxActive:   10000,             // int 最大建立的連線數，默認為0不限制(reids 預設最大連線量)
		IdleTimeout: 300 * time.Second, // 連線過期時間，默認為0表示不做過期限制
		Wait:        true,              // 當連線超出限制數量後，是否等待到空閒連線釋放
		Dial: func() (c redis.Conn, err error) {
			// 使用redis封裝的Dial進行tcp連接
			host := "127.0.0.1"
			port := "6379"
			pwd := "rd3Redis"

			// 組合連接資訊
			var connectionString = fmt.Sprintf("%s:%s", host, port)
			c, err = redis.Dial(
				"tcp",
				connectionString,
				redis.DialPassword(pwd),
				redis.DialConnectTimeout(5*time.Second), // 建立連線 time out 時間 5 秒
				redis.DialReadTimeout(5*time.Second),    // 讀取資料 time out 時間 5 秒
				redis.DialWriteTimeout(5*time.Second),   // 寫入資料 time out 時間 5 秒
			)

			return
		}, // 連接redis的函数
		TestOnBorrow: func(redis redis.Conn, t time.Time) (err error) {
			// 每5秒ping一次redis
			if time.Since(t) < (5 * time.Second) {
				return
			}

			_, err = redis.Do("PING")
			if err != nil {
				return err
			}

			return
		}, // 定期對 redis server 做 ping/pong 測試

	}

}

// RedisPoolConnect 回傳連線池的 Redis 連線
func RedisPoolConnect() *redis.Pool {
	return redisPool
}
