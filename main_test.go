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
	r.GET("/ping", MyHandler)

	ctx.Request = httptest.NewRequest(http.MethodGet, "/ping", nil)

	// 開始執行
	r.ServeHTTP(w, ctx.Request)

	t.Log(ctx.Request.URL)
	t.Log(string(w.Body.Bytes()))

}

// TestHttpPost 測試 handlet 連線
func TestHttpPost(t *testing.T) {
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
		r.POST("/pong", MyHandler02)

		form := url.Values{}
		t.Run(tt.name, func(t *testing.T) {

			// 平行處理
			t.Parallel()

			// 設定form data
			for wk, wv := range tt.want {
				form.Add(wk, fmt.Sprint(wv))
			}

			ctx.Request = httptest.NewRequest(http.MethodPost, "/pong", strings.NewReader(form.Encode()))

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
