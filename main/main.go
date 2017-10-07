package main

import (
	"fmt"
)

func main() {
	FriendName := "Emily"
	switch FriendName {
	case "Amy", "Emily":
		fmt.Println("Hi beautiful girl")
	case "Neil", "Jhon":
		fmt.Println("Hi hamesome boy")
	}
}
