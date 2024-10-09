package db

import (
	"auth-go-app/models"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Init() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	return nil
}
func Close() {
	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
func CheckIfUserExists(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := Client.Database("userdb").Collection("users")
	filter := bson.M{"email": email}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Println("error")
	}
	return count > 0
}
func Save(u *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := Client.Database("userdb").Collection("users")
	u.CreatedAt = time.Now()
	

	_, dbError := collection.InsertOne(ctx, u)
	return dbError
}
