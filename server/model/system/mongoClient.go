package system

import "go.mongodb.org/mongo-driver/mongo"

type MongoClient struct {
	Database *mongo.Database
}

func (m *MongoClient) Comment() *mongo.Collection {
	return m.Database.Collection("comment")
}
func (m *MongoClient) Reply() *mongo.Collection {
	return m.Database.Collection("reply")
}
