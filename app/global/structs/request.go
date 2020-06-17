package structs

type UrlQuery struct {
	Num1 int `json:"num1" form:"num1" binding:"required"`
	Num2 int `json:"num2" form:"num2"`
}
