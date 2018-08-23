package main

import (
	"errors"
	"fmt"
)

//定義一個ErrUserNameExist的class
type errUserNameExist struct {
	username string
}

// 自己定義error
func (e errUserNameExist) Error() string {
	return fmt.Sprintf("The %s user is exist", e.username)
}

func checkoutUserIsExist(username string) (bool, error) {
	if username == "bar" {
		return true, errors.New(fmt.Sprintf("user %v is exsit", username))
	}

	if username == "foo" {
		return true, errUserNameExist{username}
	}

	if username == "tee" {
		return true, &errUserNameExist{username}
	}

	return false, nil
}

func isErrUserNameExist(err error) bool {
	_, ok := err.(errUserNameExist)
	if !ok {
		_, ok = err.(*errUserNameExist)
	}
	return ok
}

func main() {
	fmt.Println("Hello World")

	if _, err := checkoutUserIsExist("bar"); err != nil {
		fmt.Printf("%T", err)
		fmt.Println("")
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("isErrUserNameExist is false")
		}
		fmt.Println(err)
	}
	fmt.Println("================")

	if _, err := checkoutUserIsExist("foo"); err != nil {
		fmt.Printf("%T", err)
		fmt.Println("")
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("isErrUserNameExist is false")
		}
		fmt.Println(err)
	}
	fmt.Println("================")

	if _, err := checkoutUserIsExist("tee"); err != nil {
		fmt.Printf("%T", err)
		fmt.Println("")
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("isErrUserNameExist is false")
		}
		fmt.Println(err)
	}

}
