package repository

import (
	"Go_Practice/app/model"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TUser struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo     IDB
	expected []model.User
}

func (u *TUser) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	// 	建立 DB 連線
	db, u.mock, err = sqlmock.New()
	require.NoError(u.T(), err)

	u.DB, err = gorm.Open("mysql", db)
	require.NoError(u.T(), err)

	// 開啟 Debug 模式
	u.DB.LogMode(true)

	// 初始化 repo
	u.repo = testCreateRepository(u.DB)
}
func TestInit(t *testing.T) {
	suite.Run(t, new(TUser))
}

func (u *TUser) TestDB_GetUserInfo() {
	defer u.DB.Close()

	// 預設回傳資料
	u.expected = []model.User{
		{
			ID:    1,
			Name:  "Neil",
			Phone: "09XX-XXX-OOO",
			Age:   26,
		},
	}

	query := "SELECT (.+) FROM `user`"
	rows := sqlmock.NewRows([]string{"id", "name", "phone", "age"}).
		AddRow("1", "Neil", "09XX-XXX-OOO", "26")

	// 執行語法
	u.mock.ExpectQuery(query).
		WillReturnRows(rows)

	res, err := u.repo.GetUserInfo()

	require.NoError(u.T(), err)

	require.Equal(u.T(), u.expected, res, "Return alue not same")

	require.NoError(u.T(), u.mock.ExpectationsWereMet())
}
