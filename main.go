package main

import "fmt"

func modify(foo []string) {
	foo[1] = "c"
	fmt.Println("Modify foo", foo)
}

func addValue(foo []string) []string {
	fmt.Println("Now foo is: ", foo)
	foo = append(foo, "c")
	fmt.Println("addValue foo", foo)
	return foo
}

func addValue2(foo *[]string) {
	fmt.Println("Now foo is: ", foo)
	*foo = append(*foo, "d")

}

func main() {
	// 宣告一個slice，並賦予一兩個值
	foo := []string{"a", "b"}

	// 印出目前slice結果
	fmt.Println("Before foo: ", foo)
	fmt.Println("================")

	// 因為slice是一個call by reference
	// 所以當呼叫modify func時，帶入foo
	// 此時修改"a","b"的值會一併被改變
	// 但foo長度以固定，顧增加值是不被與允許的
	modify(foo)
	fmt.Println("After foo: ", foo)
	fmt.Println("================")

	// 可用以下兩種func新增、刪減foo的長度 + 內容
	// 方法一： return 結果
	foo = addValue(foo)
	fmt.Println("After foo: ", foo)
	fmt.Println("================")

	// 方法一： 傳址
	addValue2(&foo)
	fmt.Println("After foo: ", foo)
	fmt.Println("================")
}
