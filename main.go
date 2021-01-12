package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	outChan := make(chan string, 100)
	errChan := make(chan error, 100)
	finishChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func(val int, wg *sync.WaitGroup, out chan string, err chan error) {
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Microsecond)
			out <- fmt.Sprintf("finished job id: %d", val)

			if val == 15 {
				err <- errors.New("fail job in 15")
			}

			if val == 19 {
				time.Sleep(1 * time.Second)
			}

			wg.Done()
		}(i, &wg, outChan, errChan)
	}

	go func() {
		wg.Wait()
		// 關閉
		close(finishChan)
	}()

	// 自定義名稱
Loop:
	for {
		select {
		// 接收輸出內容
		case out := <-outChan:
			log.Println(out)
		// 接收錯誤訊息，或行為
		case err := <-errChan:
			log.Println(err)

			// 可以直接跳離服務，或者寫 log，db，alert 處理等等
			// break Loop
		// 結束後的行為
		case <-finishChan:
			break Loop
		// timeout 後行為
		case <-time.After(100 * time.Millisecond):
			log.Println("timeout")

			// 可以直接跳離服務，或者寫 log，db，alert 處理等等
			break Loop
		}
	}
}
