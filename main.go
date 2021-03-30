package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

func main() {
	GJsonWithFor()
}

func JsoniterUnmarshal() {
	// your data
	str := `{"horizontal":"wsSyajSg.jpg","vertical":"kvAgu5ek.jpg"}`

	// your map
	tmp := map[string]string{}

	if err := jsoniter.Unmarshal([]byte(str), &tmp); err != nil {
		fmt.Println("====>", err.Error())
		return
	}

	// fmt.Println("---->", tmp)
}

func JsonUnmarshal() {
	// your data
	str := `{"horizontal":"wsSyajSg.jpg","vertical":"kvAgu5ek.jpg"}`

	// your map
	tmp := map[string]string{}

	if err := jsoniter.Unmarshal([]byte(str), &tmp); err != nil {
		fmt.Println("====>", err.Error())
		return
	}

	// fmt.Println("---->", tmp)
}

func GJsonWithInterface() {
	// your data
	str := `{"horizontal":"wsSyajSg.jpg","vertical":"kvAgu5ek.jpg"}`

	// your map
	// tmp := map[string]string{}

	// 驗證是否為 json 格式
	if !gjson.Valid(str) {
		fmt.Println("invalid json")
		return
	}

	_, ok := gjson.Parse(str).Value().(map[string]interface{})
	if !ok {
		// not a map
		fmt.Println("not a map")
		return
	}

	//
	// value := gjson.Get(str, "..#..*")
	// fmt.Println("---->", m)
}

func GJsonWithFor() {
	// your data
	str := `{"horizontal":"wsSyajSg.jpg","vertical":"kvAgu5ek.jpg"}`

	// your map
	tmp := map[string]string{}

	// 驗證是否為 json 格式
	if !gjson.Valid(str) {
		fmt.Println("invalid json")
		return
	}

	// 取 json 值
	value := gjson.Parse(str)

	for k, v := range value.Map() {
		tmp[k] = v.Str
	}

	// fmt.Println(tmp)

}
