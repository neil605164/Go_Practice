package business

import (
	"Go_Practice/app/global/structs"
	"Go_Practice/app/repository"
	"Go_Practice/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestRequest_StoreDBInfo(t *testing.T) {
	// 開啟 gin 測試模式
	gin.SetMode(gin.TestMode)

	// new mock
	mDB := new(mocks.IDB)
	mRedis := new(mocks.IRedis)

	type fields struct {
		Redis repository.IRedis
		DB    repository.IDB
	}
	type args struct {
		req structs.RawData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "check db no error",
			fields: fields{
				Redis: mRedis,
				DB:    mDB,
			},
			args: args{
				req: structs.RawData{
					Name:  "Neil",
					Phone: "09XX-XXX-OOO",
					Age:   26,
				},
			},
			wantErr: false,
		},
		{
			name: "check db has error",
			fields: fields{
				Redis: mRedis,
				DB:    mDB,
			},
			args: args{
				req: structs.RawData{
					Name:  "Neil",
					Phone: "09XX-XXX-SSS",
					Age:   27,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {

		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// 平行處理
			t.Parallel()

			// 需要跳脫的 mockery
			mDB.On("SetUserInfo", mock.Anything).Return(nil)

			r := &Request{
				Redis: tt.fields.Redis,
				DB:    tt.fields.DB,
			}
			if err := r.StoreDBInfo(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Request.StoreDBInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !mDB.AssertExpectations(t) {
				t.Fail()
			}
		})

	}
}
