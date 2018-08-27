// 可以指定func名稱，決定指測試那一個func
//go test -v -run=TestNewWithAssert ./car/.

package car

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		// Fatal 會停在該地方，不再繼續執行
		t.Fatal("got errors:", err)
	}

	if c == nil {
		// Error 會在繼續向下執行，可以將上方Fatal更改為 Error在跑一次測試
		t.Error("Car should be nil")
	}

	if c == nil {
		t.Log("Car should be nil")
	}
}

func TestNewWithAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = New("foo", 100)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "foo", c.Name)
}

// 平行測試
func TestCar_SetName(t *testing.T) {
	type fields struct {
		Name  string
		Price float32
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "no input name",
			fields: fields{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "",
			},
			want: "foo",
		},
		{
			name: "input name",
			fields: fields{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "bar",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		// 產生local variable ，避免tt被覆蓋掉（必須）
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// 分別印出當下的tt.args參數
			log.Println(tt.args)
			fmt.Println("===========")

			c := &Car{
				Name:  tt.fields.Name,
				Price: tt.fields.Price,
			}
			if got := c.SetName(tt.args.name); got != tt.want {
				t.Errorf("Car.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
