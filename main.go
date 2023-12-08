package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id   int64  `bson:"_id"`
	Name string `bson:"name"`
}

func main() {
	client := Conn()
	// 關閉連接
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("your_database").Collection("your_collection")

	// first
	// First(collection)

	// Insert One
	// InsertOne(collection)

	// Insert Many
	// InsertMany(collection)

	// Insert Many Ignore duplicate
	InsertManyIgnoreDuplicate(collection)

	// Insert Many If duplicate then Ignore
	// InsertManyIgnoreDuplicateIfAlreadyExist(collection)
}

func Conn() *mongo.Client {

	// 設定連接選項
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")

	// 連接到MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 確保連接是有效的
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func First(collection *mongo.Collection) {

	userA := User{
		Id:   time.Now().UnixMilli(),
		Name: "neil",
	}

	temp := User{
		Id:   time.Now().UnixMilli(),
		Name: "neil",
	}
	if err := collection.FindOne(context.Background(), userA).Decode(&temp); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Fatal("error is ===>", err)
		}
	}
}

func InsertOne(collection *mongo.Collection) {

	user := User{
		Id:   time.Now().UnixMilli(),
		Name: "neil",
	}

	// Insert One
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("error is ===>", err)
	}

	fmt.Println("result.InsertedID =====>", result.InsertedID)
}

func InsertMany(collection *mongo.Collection) {

	userList := make([]any, 0)
	userA := User{
		Id:   time.Now().UnixMilli(),
		Name: "neil",
	}

	userList = append(userList, userA)

	time.Sleep(1 * time.Millisecond)

	userB := User{
		Id:   time.Now().UnixMilli(),
		Name: "wen",
	}

	userList = append(userList, userB)

	// Insert One
	result, err := collection.InsertMany(context.Background(), userList)
	if err != nil {
		log.Println("error is ===>", err)
	}

	for _, id := range result.InsertedIDs {
		fmt.Println("result.InsertedID =====>", id)
	}
}

func InsertManyIgnoreDuplicate(collection *mongo.Collection) {

	userList := make([]any, 0)
	userA := User{
		Id:   time.Now().UnixMilli(),
		Name: "neil",
	}

	userList = append(userList, userA)

	// time.Sleep(1 * time.Millisecond)

	userB := User{
		Id:   time.Now().UnixMilli(),
		Name: "wen",
	}

	userList = append(userList, userB)

	time.Sleep(1 * time.Millisecond)

	userC := User{
		Id:   time.Now().UnixMilli(),
		Name: "simon",
	}

	userList = append(userList, userC)

	var opt []*options.InsertManyOptions
	opt = append(opt, options.InsertMany().SetOrdered(false))

	// Insert One
	result, err := collection.InsertMany(context.Background(), userList, opt...)
	if err != nil {
		log.Println("error is ===>", err)
	}

	for _, id := range result.InsertedIDs {
		fmt.Println("result.InsertedID =====>", id)
	}
}

func InsertManyIgnoreDuplicateIfAlreadyExist(collection *mongo.Collection) {

	userA := User{
		Id:   time.Now().UnixMilli(),
		Name: "neil",
	}

	// Insert One
	result, err := collection.InsertOne(context.Background(), userA)
	if err != nil {
		log.Println("error is ===>", err)
	}

	fmt.Println("result.InsertedID ===>", result.InsertedID)

	time.Sleep(10 * time.Second)

	userList := make([]any, 0)
	userB := User{
		Id:   userA.Id,
		Name: "simon",
	}

	userList = append(userList, userB)

	userC := User{
		Id:   time.Now().UnixMilli(),
		Name: "wen",
	}

	userList = append(userList, userC)

	opt := options.InsertMany().SetOrdered(false)

	// Insert One
	resultList, err := collection.InsertMany(context.Background(), userList, opt)
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		log.Println("error is ===>", err)
	}

	for _, id := range resultList.InsertedIDs {
		fmt.Println("result.InsertedID =====>", id)
	}
}
