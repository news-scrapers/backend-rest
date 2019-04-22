package models

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var db *mongo.Database
func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbAddress := os.Getenv("database_url")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(dbAddress))
	db = client.Database("testing")
}

func GetDB() *mongo.Database {
	return db
}
