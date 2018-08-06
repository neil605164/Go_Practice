package main

import "fmt"

// 定義class型態(struct)
// type person struct {
// 	Name string `json:"name"`
// 	Sex  string `json:"sex"`
// 	Age  int    `json:"age"`
// }
type person struct {
	name string
	sex  string
	age  int
}

func main() {
	// fmt.Println(person{"Neil", "Male", 25})
	// fmt.Println(&person{name: "Ann", sex: "Female", age: 40})
	// s := person{name: "Sean", sex: "Male", age: 25}
	// fmt.Println(s.name)

	// sp := &s
	// fmt.Println(sp)

	// s.age = 26
	// fmt.Println(s)

	// 提供用戶資料
	// data 存放person物件的資料
	data := person{
		name: "Neil",
		sex:  "male",
		age:  25,
	}

	// 顯示當前data 資訊
	fmt.Println(data) //{Neil male 25}

	// 呼叫person的myInfo func
	data.myInfo()

	data.retrunMyInfo()

	// 執行完returnMyInfo後，印出新的data資料
	fmt.Println(data) //{Neil male 40}
}

func (c *person) myInfo() {
	fmt.Println("My name is ", c.name)
	fmt.Println("I'm ", c.age, " yaers old")
	fmt.Println("I'm a", c.sex)
}

func (c *person) retrunMyInfo() (a *person) {
	c.age = 40
	c.name = "Amy"
	c.sex = "Female"
	return
}
