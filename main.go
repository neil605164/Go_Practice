package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {

	// Connect DB
	session, err := dbCon()
	if err != nil {
		fmt.Println("===>", err)
		panic(err)
	}
	defer session.Close()

	// choose DataBase
	db, err := chooseDB("test", session)
	if err != nil {
		fmt.Println("####>", err)
		panic(err)
	}

	// Insert Data
	// if err = InsertData(db); err != nil {
	// 	fmt.Println("++++>", err)
	// 	panic(err)
	// }

	// Select Data
	res := Person{}
	res, err = SelectData(db, res)
	if err != nil {
		fmt.Println("---->", err)
		log.Fatal(err)
	}

	fmt.Println("Phone:", res.Phone)

	// Update Data
	row := Person{
		Name:  "Ale",
		Phone: "09SS-XXX-OOO",
	}
	if err = UpdateData(db, row); err != nil {
		fmt.Println("~~~~>", err)
		log.Fatal(err)
	}

	// Select Data
	res, err = SelectData(db, res)
	if err != nil {
		fmt.Println("$$$$>", err)
		log.Fatal(err)
	}

	fmt.Println("Phone:", res.Phone)

	// Delete Data
	if err = DeleteData(db, "Ale"); err != nil {
		fmt.Println("%%%%>", err)
		log.Fatal(err)
	}

	persons := []Person{}
	list, err := SelectAll(db, persons)
	if err != nil {
		fmt.Println("&&&&>", err)
		log.Fatal(err)
	}

	fmt.Println("Result:", list)
}

func dbCon() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}

// chooseDB 選擇 DB
func chooseDB(database string, session *mgo.Session) (*mgo.Database, error) {
	username := "root"
	password := "123456"

	// choose DB and DB Login
	db := session.DB(database)
	if err := db.Login(username, password); err != nil {
		return nil, err
	}

	return db, nil
}

func chooseTable(table string, db *mgo.Database) *mgo.Collection {
	return db.C(table)
}

// InsertData Insert Data
func InsertData(db *mgo.Database) error {
	// choose Table
	c := chooseTable("people", db)

	err := c.Insert(
		&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"},
	)

	return err
}

// SelectData Select Data
func SelectData(db *mgo.Database, res Person) (Person, error) {
	// choose Table
	c := chooseTable("people", db)

	err := c.Find(bson.M{"name": "Ale"}).One(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// UpdateData Update Data
func UpdateData(db *mgo.Database, row Person) error {

	// choose Table
	c := chooseTable("people", db)

	err := c.Update(bson.M{"name": "Ale"}, &row)
	if err != nil {
		return err
	}
	return nil
}

// DeleteData Delete Data
func DeleteData(db *mgo.Database, row string) error {
	// choose Table
	c := chooseTable("people", db)

	if err := c.Remove(bson.M{"name": row}); err != nil {
		return err
	}

	return nil
}

func SelectAll(db *mgo.Database, res []Person) ([]Person, error) {
	// choose Table
	c := chooseTable("people", db)

	if err := c.Find(bson.M{}).All(&res); err != nil {
		return res, err
	}
	return res, nil
}
