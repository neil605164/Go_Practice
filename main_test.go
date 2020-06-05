package main

import (
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

	// 建立新 http test
	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)

	// 定義測試用路徑，並指向特定 handler
	r.POST("/pong", MyHandler)

	tests := []struct {
		name       string                 // 測試名稱(說明)
		want       map[string]interface{} // 輸入的參數
		haveErr    bool                   // 是否有 error
		bindStruct interface{}            // 被繫結的結構體
		errMsg     string                 // 如果有錯，錯誤資訊
	}{
		{
			name:    "Set Value ABC",
			want:    map[string]interface{}{"A": "ABC"},
			haveErr: false,
			bindStruct: &struct {
				A string `form:"a" json:"a"`
			}{},
		},
	}

	for k := range tests {

		form := url.Values{}
		// 設定form data
		for wk, wv := range tests[k].want {
			form.Add(wk, fmt.Sprint(wv))
		}

		ctx.Request = httptest.NewRequest(http.MethodPost, "/pong", strings.NewReader(form.Encode()))

		// 這個很關鍵
		ctx.Request.Header.Add("Content-Type", binding.MIMEPOSTForm)

		// 開始執行
		r.ServeHTTP(w, ctx.Request)

		// 結果處理
		t.Log(ctx.Request.URL)
		t.Log(string(w.Body.Bytes()))

		assert.Equal(t, "{a:ABC}", w.Body.Bytes(), "correct")
	}
}
