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

	// choose Table
	c := chooseTable("people", db)

	// Insert Data
	// if err = InsertData(c, db); err != nil {
	// 	fmt.Println("++++>", err)
	// 	panic(err)
	// }

	// Select Data
	res := Person{}
	res, err = SelectData(c, db, res)
	if err != nil {
		fmt.Println("---->", err)
		log.Fatal(err)
	}

	fmt.Println("Phone:", res.Phone)
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

// Insert Data
func InsertData(c *mgo.Collection, db *mgo.Database) error {
	err := c.Insert(
		&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"},
	)

	return err
}

// Select Data
func SelectData(c *mgo.Collection, db *mgo.Database, res Person) (Person, error) {
	err := c.Find(bson.M{"name": "Ale"}).One(&res)
	if err != nil {
		fmt.Println("---->", err)
		log.Fatal(err)
	}

	return res, nil
}
