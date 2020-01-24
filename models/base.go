package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbAddress := os.Getenv("database_url")
	dbName := os.Getenv("database_name")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))
	log.Println(client)
	log.Println(err)
	log.Println(dbAddress)
	if err != nil {
		panic(err)
	}
	db = client.Database(dbName)
}

func GetDB() *mongo.Database {
	return db
}
