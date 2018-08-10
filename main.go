package main

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

func main() {
	someText := "Neil_Hsieh"
	hash, err := hashTextTo32Bytes(someText)
	if err != nil {
		log.Println("Error Hash")
	}
	fmt.Printf("%s\n %s", hash, err)
}

func hashTextTo32Bytes(hashThis string) (hashed string, err error) {

	// 檢查是否有帶入字串
	if len(hashThis) == 0 {
		return "", errors.New("No input supplied")
	}

	hasher := sha256.New()
	hasher.Write([]byte(hashThis))

	fmt.Println("hasher: ", hasher)

	stringToSHA256 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return stringToSHA256[:32], nil
}
