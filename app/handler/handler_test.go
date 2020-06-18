package handler

import (
	"Go_Practice/app/business"
	"Go_Practice/mocks"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_MyHandler01(t *testing.T) {
	// 開啟 gin 測試模式
	gin.SetMode(gin.TestMode)

	// new mock
	m := new(mocks.IBusiness)

	type fields struct {
		BInter business.IBusiness
	}
	type args struct {
		c gin.Params
	}

	type Resp struct {
		A int `json:"a"`
	}

	tests := []struct {
		name   string
		fields fields
		args   gin.Params
	}{
		{
			name: "test mockery",
			fields: fields{
				BInter: m,
			},
			args: gin.Params{
				gin.Param{
					Key:   "num1",
					Value: "1",
				},
				gin.Param{
					Key:   "num2",
					Value: "456",
				},
			},
		},
		{
			name: "test mockery",
			fields: fields{
				BInter: m,
			},
			args: gin.Params{
				gin.Param{
					Key:   "num1",
					Value: "33",
				},
				gin.Param{
					Key:   "num2",
					Value: "456",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 平行處理
			t.Parallel()

			// 初始化 gin context
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			// 將 param 丟給 ctx
			ctx.Params = tt.args

			// 取參數
			strnum1 := ctx.Param("num1")
			require.NotEmpty(t, strnum1, "strnum1 is empty")
			num1, err := strconv.Atoi(strnum1)
			require.NoError(t, err, "strnum1 string to int error")

			strnum2 := ctx.Param("num2")
			require.NotEmpty(t, strnum2, "strnum2 is empty")
			num2, err := strconv.Atoi(strnum2)
			require.NoError(t, err, "strnum2 string to int error")

			// 定義預計忽略的 mock func
			m.On("Api", num1, num2).Return(num1 + num2)

			// 定義測試用路徑，並指向特定 handler
			hand := NewHandler()
			hand.BInter = m

			// 設定 request 內容
			url := fmt.Sprintf("/hand01?num1=%d&num2=%d", num1, num2)
			ctx.Request = httptest.NewRequest(http.MethodGet, url, nil)

			hand.MyHandler01(ctx)

			// 處理結果
			expected := num1 + num2
			actual := Resp{}

			err = json.Unmarshal(w.Body.Bytes(), &actual)
			require.NoError(t, err, "json unmarshal actual response error")

			t.Log(actual.A)
			t.Log(expected)

			assert.Equal(t, expected, actual.A, "value is not equal")
		})
	}

	if !m.AssertExpectations(t) {
		t.Fail()
	}
}
