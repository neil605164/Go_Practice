package main

import (
	"fmt"
	"log"

	"HTTP_GO/conf"
)

func main() {
	config, err := util.GetChannelConfig()
	if err != nil {
		log.Printf("get emayReminder channel config faild! err: #%v", err)
	}
	fmt.Println(config)
}
