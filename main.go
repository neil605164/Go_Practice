package main

import (
	"fmt"
	"time"
)

// var i interface{}

func main() {
	// i = 1
	// convrt(i)

	// i = "word"
	// convrt(i)

	// i = float64(45.55)
	// convrt(i)

	// i = float32(45.55)
	// convrt(i)

	// myselect()
	// myselect2()
	// myselect3()
	myselect4()
}

// switch 用法
func convrt(i interface{}) {
	switch t := i.(type) {
	case string:
		println("i is string:", t)
	case int:
		println("i is int:", t)
	case float64:
		println("i is float64:", t)
	default:
		println("type not found")
	}
}

// Select 用法
// 只能用於接收 channel 否則會出錯
// default 會直接執行，所以沒有 default 的 select 就會 blocking
// 完全沒有接受 value 的 channel 會造成 panic
func myselect() {
	ch := make(chan int, 1)

	ch <- 1
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	}
}

func myselect2() {
	ch := make(chan int, 1)
	timeout := make(chan bool, 1)

	// go func() {
	// 	time.Sleep(time.Second * 1) // timeout
	// 	timeout <- true
	// }()

	go func() {
		time.Sleep(time.Second * 3) // timeout 02
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println("random 01")
	// timeout 超時機制
	case <-timeout:
		fmt.Println("timeout")
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 02")
	}
}

func myselect3() {
	ch := make(chan int, 1)

	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("channel val is", <-ch)
		fmt.Println("channel val is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}

func myselect4() {
	i := 0
	ch := make(chan string)
	defer func() {
		close(ch)
	}()

	go func() {
	Loop:
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {
			case m := <-ch:
				println(m)
				break Loop
			default:
			}
		}
	}()

	time.Sleep(4 * time.Second)
	ch <- "stop"
}
