package initialization

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/model/system"
)

func InitMongoDb() *system.MongoClient {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?%s",
		viper.GetString("mongodb.username"),
		viper.GetString("mongodb.password"),
		viper.GetString("mongodb.host"),
		viper.GetString("mongodb.port"),
		viper.GetString("mongodb.param"),
	)
	connect, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	mongoClient := &system.MongoClient{Database: connect.Database(viper.GetString("mongodb.dbname"))}
	return mongoClient
}
