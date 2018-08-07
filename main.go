package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text1.txt")
	if err != nil {
		fmt.Println("File is not exit")
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Content is null")
		return
	}

	bs := make([]byte, stat.Size())
	fmt.Println(bs)
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
