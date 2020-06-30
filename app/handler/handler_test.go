package handler

import (
	"Go_Practice/app/business"
	"Go_Practice/app/global/structs"
	"Go_Practice/mocks"
	"bytes"
	"encoding/json"
	"errors"
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
			name: "test mockery part 2",
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

			// 設定 request 內容
			url := fmt.Sprintf("/hand01?num1=%d&num2=%d", num1, num2)
			ctx.Request = httptest.NewRequest(http.MethodGet, url, nil)

			// 定義測試用路徑，並指向特定 handler
			hand := NewHandler()
			hand.BInter = m

			hand.MyHandler01(ctx)

			// 處理結果
			expected := num1 + num2
			actual := Resp{}

			err = json.Unmarshal(w.Body.Bytes(), &actual)
			require.NoError(t, err, "json unmarshal actual response error")

			t.Log(actual.A)
			t.Log(expected)

			assert.Equal(t, expected, actual.A, "value is not equal")

			if !m.AssertExpectations(t) {
				t.Fail()
			}
		})
	}
}

func TestHandler_MyHandler02(t *testing.T) {
	// 開啟 gin 測試模式
	gin.SetMode(gin.TestMode)

	// new mock
	m := new(mocks.IBusiness)

	redisExpectValue := "1"

	type fields struct {
		BInter business.IBusiness
	}

	type resp struct {
		Res int `json:"res"`
	}

	tests := []struct {
		name   string
		fields fields
		args   gin.Params
		errMsg error
		harErr bool
	}{
		{
			name: "test get mock redis func",
			fields: fields{
				BInter: m,
			},
			args: gin.Params{
				gin.Param{
					Key:   "key",
					Value: "num1",
				},
			},
			errMsg: nil,
			harErr: false,
		},
		{
			name: "test get mock redis func",
			fields: fields{
				BInter: m,
			},
			args: gin.Params{
				gin.Param{
					Key:   "key",
					Value: "num2",
				},
			},
			errMsg: errors.New("redigo: nil returned"),
			harErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 平行處理
			t.Parallel()

			// 開始寫要跳脫的 mock
			m.On("GetRedis", tt.args.ByName("key")).Return(redisExpectValue, tt.errMsg)

			if tt.harErr {
				require.Error(t, tt.errMsg, "get redis value error")
			} else {
				// 初始化 gin context
				w := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(w)

				// 將參數丟入 ctx 內
				ctx.Params = tt.args

				// 設定 request 內容
				url := fmt.Sprintf("/hand02?key=%v", tt.args.ByName("key"))
				ctx.Request = httptest.NewRequest(http.MethodGet, url, nil)

				h := &Handler{
					BInter: tt.fields.BInter,
				}

				// 開始執行
				h.MyHandler02(ctx)

				// 處理結果
				num, err := strconv.Atoi(redisExpectValue)
				require.NoError(t, err, "redis value string to int error")

				actual := resp{}
				err = json.Unmarshal(w.Body.Bytes(), &actual)
				require.NoError(t, err, "json unmarshal actual response error")

				require.Equal(t, num+10, actual.Res, "value is not equal")

				if !m.AssertExpectations(t) {
					t.Fail()
				}
			}
		})
	}
}

func TestHandler_MyHandler03(t *testing.T) {
	// 開啟 gin 測試模式
	gin.SetMode(gin.TestMode)

	// new mock
	m := new(mocks.IBusiness)

	type fields struct {
		BInter business.IBusiness
	}

	type res struct {
		Res structs.RawData
	}

	tests := []struct {
		name   string
		fields fields
		args   structs.RawData
		errMsg error
		hasErr bool
	}{
		{
			name: "input correct value",
			fields: fields{
				BInter: m,
			},
			args: structs.RawData{
				Name:  "Neil",
				Phone: "09XX-XXX-OOO",
				Age:   26,
			},
			errMsg: nil,
			hasErr: false,
		},
		{
			name: "input correct value",
			fields: fields{
				BInter: m,
			},
			args: structs.RawData{
				Name:  "Neil",
				Phone: "09XX-XXX-OOO",
				Age:   26,
			},
			errMsg: errors.New("something error"),
			hasErr: true,
		},
	}
	for _, tt := range tests {

		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			// 平行處理
			t.Parallel()

			// 開始寫要跳脫的 mock
			m.On("StoreDBInfo", tt.args).Return(tt.errMsg)

			// 初始化 gin context
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			// 處理帶入的參數
			byteData, err := json.Marshal(tt.args)
			assert.NoError(t, err, "json marshal error")

			// 設定 request 內容
			ctx.Request = httptest.NewRequest(http.MethodPost, "/hand03", bytes.NewBuffer(byteData))

			h := &Handler{
				BInter: tt.fields.BInter,
			}
			h.MyHandler03(ctx)

			// 處理結果
			resp := res{}
			err = json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err, "json unmarshal error")

			expected := structs.RawData{}
			err = json.Unmarshal(byteData, &expected)
			assert.NoError(t, err, "json unmarshal error")

			if !tt.hasErr {
				assert.Equal(t, expected, resp.Res, "same response")
			} else {
				assert.Error(t, tt.errMsg, "something error")
			}

			if !m.AssertExpectations(t) {
				t.Fail()
			}
		})
	}
}
