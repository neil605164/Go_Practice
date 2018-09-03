// 啟動服務後，執行以下指令
// go run main.go
// echo "GET http://localhost:8000" | vegeta attack -rate=100 -connections=1 -duration=1s | tee result.bin | vegeta report

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 全域變數
var globalDB *mgo.Database
var account = "neil_hsieh"
var input []chan string
var output []chan result
var maxUser = 100
var maxThread = 10

type result struct {
	Account string
	Result  float64
}

type currency struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Amount  float64       `bson:"amount"`
	Account string        `bson:"account"`
	Code    string        `bson:"code"`
}

// Random get random value
// 隨機慘生一個亂數（整數）
func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min+1) + min
}

func pay(w http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		number := Random(1, maxUser)
		channelNumber := number % maxThread
		account := "user" + strconv.Itoa(number)
		input[channelNumber] <- account
		for {
			select {
			case data := <-output[channelNumber]:
				fmt.Printf("%v\n", data)
				wg.Done()
				return
			}
		}
	}(&wg)

	wg.Wait()
	io.WriteString(w, "ok")
}

func main() {
	// 宣告兩個channel
	// input 型態為 string
	// output 型態為 result （struct）
	input = make([]chan string, maxThread)
	output = make([]chan result, maxThread)

	// 建立連線
	session, err := mgo.Dial("localhost:8082")
	if err != nil {
		log.Fatalf("[Mongo DB] error message: %v", err)
	}

	// 連接至指定資料database & table
	globalDB = session.DB("prac_golan_mongo")

	for i := range input {
		input[i] = make(chan string)
		output[i] = make(chan result)
	}

	// create 100 user
	for i := 0; i < maxUser; i++ {
		account := "user" + strconv.Itoa(i+1)
		user := currency{Account: account, Amount: 1000.00, Code: "USD"}
		if err := globalDB.C("bank").Insert(&user); err != nil {
			panic("insert error")
		}
	}

	// 建立一個goroutine，並將input丟入
	// 若input有值，開始執行case內的任務
	for i := range input {

		go func(input chan string, i int) {
			for {
				select {
				case account := <-input:
					entry := currency{}

					// step 1: get current amount
					err := globalDB.C("bank").Find(bson.M{"account": account}).One(&entry)

					if err != nil {
						panic(err)
					}

					//step 3: subtract current balance and update back to database
					entry.Amount = entry.Amount + 50.000 // 預計取的6000的結果
					err = globalDB.C("bank").UpdateId(entry.ID, &entry)

					if err != nil {
						panic("update error")
					}

					output[i] <- result{
						Account: account,
						Result:  entry.Amount,
					}
				}
			}
		}(input[i], i)
	}

	log.Println("Listen server on 8000 port")
	http.HandleFunc("/", pay)
	http.ListenAndServe(":8000", nil)
}
