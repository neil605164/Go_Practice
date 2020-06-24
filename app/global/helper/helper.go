package helper

import "encoding/json"

// StructToMap struct型態 轉 map型態 (For DB 使用)
func StructToMap(myStruct interface{}) (myMap map[string]interface{}, err error) {

	// 轉形成map，才可以處理空字串,0,false值
	myMap = make(map[string]interface{})

	// 資料轉型
	byteData, err := json.Marshal(myStruct)
	if err != nil {
	}

	if err := json.Unmarshal(byteData, &myMap); err != nil {
	}

	return
}
