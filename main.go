package main

import (
	"fmt"
	"sync"
)

func do(i int, wg *sync.WaitGroup) {
	fmt.Printf("start job: %d\n", i)
	fmt.Printf("end job: %d\n", i)
	fmt.Println("==================")
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go do(1, &wg)
	go do(2, &wg)
	go do(3, &wg)

	wg.Wait()
	fmt.Println("Done")
}
