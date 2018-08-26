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

	// Slice 練習二
	fmt.Println("")
	fmt.Println("********第二種練習********")
	bar := []string{"A", "B"}
	fmt.Println("bar:", bar)

	// 將bar的第1個位子的值，丟給temp
	temp := bar[:1]
	fmt.Println("temp:", temp)
	fmt.Println("================")

	// 改變temp值，bar會一同更改
	s1 := append(temp, "C")
	fmt.Println("bar:", bar)
	fmt.Println("s1:", s1)
	fmt.Println("================")

	s2 := append(temp, "D")
	fmt.Println("bar:", bar)
	fmt.Println("s2:", s2)
	fmt.Println("================")

	// 無法正確更改bar的內容，因為bar的長度被定義為2
	// 顧雖然s3有變動，但是不會影響bar的內容
	s3 := append(temp, "E", "F")
	fmt.Println("bar:", bar)
	fmt.Println("s3:", s3)
	fmt.Println("================")
}
