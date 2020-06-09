package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

// TestHttpGet 測試 handlet 連線
func TestHttpGet(t *testing.T) {
	// 開啟 gin 測試模式
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)

	// 定義測試用路徑，並指向特定 handler
	r.GET("/hand01", MyHandler01)

	ctx.Request = httptest.NewRequest(http.MethodGet, "/hand01", nil)

	// 開始執行
	r.ServeHTTP(w, ctx.Request)

	t.Log(ctx.Request.URL)
	t.Log(string(w.Body.Bytes()))

}

// TestHttpPostWithFormData 透過 form data 測試 handler 連線
func TestHttpPostWithFormData(t *testing.T) {
	// 開啟 gin 測試模式
	gin.SetMode(gin.TestMode)

	// 測定測試內容
	tests := []struct {
		name string                 // 測試名稱(說明)
		want map[string]interface{} // 輸入的參數
	}{
		{
			name: "Set Value ABC",
			want: map[string]interface{}{"a": "ABC"},
		},
		{
			name: "Set Value 123",
			want: map[string]interface{}{"a": "123"},
		},
	}

	// 開始執行測試
	for _, tt := range tests {
		// 建立新 http test
		w := httptest.NewRecorder()
		ctx, r := gin.CreateTestContext(w)

		// 定義測試用路徑，並指向特定 handler
		r.POST("/hand02", MyHandler02)

		form := url.Values{}
		t.Run(tt.name, func(t *testing.T) {

			// 平行處理
			t.Parallel()

			// 設定form data
			for wk, wv := range tt.want {
				form.Add(wk, fmt.Sprint(wv))
			}

			ctx.Request = httptest.NewRequest(http.MethodPost, "/hand02", strings.NewReader(form.Encode()))

			// 這個很關鍵，設定 header 為 application/x-www-form-urlencoded
			ctx.Request.Header.Add("Content-Type", binding.MIMEPOSTForm)

			// 開始執行
			r.ServeHTTP(w, ctx.Request)

			// 處理回傳資料
			res := map[string]interface{}{}
			err := json.Unmarshal(w.Body.Bytes(), &res)
			assert.NoError(t, err, "json unmarshal error")

			for _, wv := range tt.want {
				if wv == "123" {
					assert.NotEqual(t, tt.want, res, "Not equal correct")
				} else {
					// 檢查是否相同
					assert.Equal(t, tt.want, res, "Equal correct") //true
				}
			}

		})
	}
}

// TestHttpPostWithRawData 透過 rawdata 測試 handler 連線
func TestHttpPostWithRawData(t *testing.T) {
	// 開啟 gin 測試模式
	gin.SetMode(gin.TestMode)

	// 定義參數規則
	type Fields struct {
		Name string
		Sex  string
	}

	// 設定參數內容
	tests := []struct {
		name string
		want Fields
	}{
		{
			name: "input neil and male",
			want: Fields{
				Name: "Neil",
				Sex:  "male",
			},
		},
		{
			name: "input linda and female",
			want: Fields{
				Name: "Linda",
				Sex:  "female",
			},
		},
	}

	// 開始執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 建立新 http test
			w := httptest.NewRecorder()
			ctx, r := gin.CreateTestContext(w)

			// 定義測試用路徑，並指向特定 handler
			r.POST("/hand03", MyHandler03)

			// 處理參數
			byteData, err := json.Marshal(tt.want)
			assert.NoError(t, err, "json unmarshal error")

			ctx.Request = httptest.NewRequest(http.MethodPost, "/hand03", strings.NewReader(string(byteData)))

			// 這個很關鍵，設定 header 為 application/json
			ctx.Request.Header.Add("Content-Type", binding.MIMEJSON)

			// 開始執行
			r.ServeHTTP(w, ctx.Request)

			// 處理回傳資料
			t.Log(string(w.Body.Bytes()))

			res := Fields{}
			err = json.Unmarshal(w.Body.Bytes(), &res)
			assert.NoError(t, err, "json unmarshal error")

			if res.Name == "Linda" {
				assert.NotEqual(t, tt.want, res, "It's Same")
			} else {
				assert.Equal(t, tt.want, res, "It's Not Same")
			}
		})
	}
}
