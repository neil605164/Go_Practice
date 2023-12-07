package main

import (
	"context"
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

	db := client.Database("your_database")
	collection1 := db.Collection("your_collection_1")
	collection2 := db.Collection("your_collection_2")

	// Transaction Start
	TransactionDemo(client, collection1, collection2)
}

func Conn() *mongo.Client {

	// 設定連接選項
	clientOptions := options.Client().ApplyURI("mongodb://root:12345678@mongo-im-dev.cluster-cr6bxddhtfoj.ap-southeast-1.docdb.amazonaws.com:27017/?replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false")

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

func TransactionDemo(client *mongo.Client, collection1, collection2 *mongo.Collection) {

	transaction := func(sessionContext mongo.SessionContext) (any, error) {
		if _, err := collection1.InsertOne(sessionContext, map[string]interface{}{
			"_id": time.Now().UnixMilli(),
			"key": "value",
		}); err != nil {
			return nil, err
		}

		if _, err := collection2.InsertOne(sessionContext, map[string]interface{}{
			"_id": 1234,
			"key": "value",
		}); err != nil {
			return nil, err
		}

		return nil, nil
	}

	session, err := client.StartSession()
	if err != nil {
		log.Fatal("StartSession error is ===>", err)
	}
	defer session.EndSession(context.Background())

	result, err := session.WithTransaction(context.Background(), transaction)
	if err != nil {
		log.Fatal("WithTransaction error is ===>", err)
	}

	// if result != nil {
	fmt.Println("Transaction completed successfully", result)
	// } else {
	// 	fmt.Println("Transaction aborted.", result)
	// }
}
