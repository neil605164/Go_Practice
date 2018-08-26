package main

import (
	"fmt"

	_ "HTTP_GO/bar"
	_ "HTTP_GO/foo"
)

var global = conver()

func conver() int {
	return 100
}

// 執行main前，會先執行init，但是會在func宣告之後，所以global才會是0
func init() {
	global = 0
}

func main() {
	fmt.Println("global is", global)
}
