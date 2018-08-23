package main

import (
	"errors"
	"fmt"
)

func checkoutUserIsExist(username string) (bool, error) {
	if username == "bar" {
		return true, errors.New(fmt.Sprintf("The %s user is exist", username))
	}

	if username == "foo" {
		return true, fmt.Errorf("The %s user is exist", username)
	}
	return false, nil
}

func main() {
	fmt.Println("Hello World")

	if _, err := checkoutUserIsExist("bar"); err != nil {
		fmt.Println(err)
	}

	if _, err := checkoutUserIsExist("foo"); err != nil {
		fmt.Println(err)
	}
}
