package repository

import (
	"Go_Practice/app/global/structs"
	"Go_Practice/app/model"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

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

	repo IDB
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
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

func (u *TUser) AfterTest(_, _ string) {
	// make sure all expectation were met
	require.NoError(u.T(), u.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(TUser))
}

func (u *TUser) TestDB_GetUserInfo() {

	// 預設回傳資料
	expected := []model.User{
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
	u.mock.ExpectQuery(query).WillReturnRows(rows)

	// 戳 DB Function
	res, err := u.repo.GetUserInfo()
	require.NoError(u.T(), err)

	// 檢查是否與預期的 value 相同
	require.Equal(u.T(), expected, res, "Return alue not same")
}

func (u *TUser) TestDB_SetUserInfo() {

	// Inset Value
	param := structs.RawData{
		Name:  "Jay",
		Phone: "09ZZ-XXX-OOO",
		Age:   24,
	}

	// sql query rule
	query := "INSERT INTO `user`"

	// transaction start
	u.mock.ExpectBegin()

	u.mock.ExpectExec(query).
		WithArgs(param.Name, param.Phone, param.Age, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// transaction end
	u.mock.ExpectCommit()

	// 戳 DB Function
	err := u.repo.SetUserInfo(param)
	require.NoError(u.T(), err)
}

func (u *TUser) TestDB_UpdateUserInfo() {
	// Inset Value
	param := make(map[string]interface{})
	param["id"] = 1
	param["name"] = "Jay"
	param["phone"] = "09ZZ-XXX-OOO"
	param["age"] = 24

	// sql query rule
	query := "UPDATE `user` SET (.+) WHERE (.+)"

	// transaction start
	u.mock.ExpectBegin()

	u.mock.ExpectExec(query).
		WithArgs(param["age"], param["id"], param["name"], param["phone"], AnyTime{}, param["id"]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// transaction end
	u.mock.ExpectCommit()

	// 戳 DB Function
	err := u.repo.UpdateUserInfo(param)
	require.NoError(u.T(), err)
}

func (u *TUser) TestDB_DeleteUserInfo() {
	id := 1

	// transaction start
	u.mock.ExpectBegin()

	u.mock.ExpectExec("DELETE FROM `user`  WHERE `user`.`id` = ?").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// transaction end
	u.mock.ExpectCommit()

	// 戳 DB Function
	err := u.repo.DeleteUserInfo(id)
	require.NoError(u.T(), err)
}
