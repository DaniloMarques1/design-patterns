package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	client *mongo.Client
}

func NewMongoDatabase() (*MongoDatabase, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017/persons"))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return &MongoDatabase{client: client}, nil
}

func (m *MongoDatabase) Save(person *Person) error {
	collection := m.client.Database("persons").Collection("person")
	if _, err := collection.InsertOne(context.Background(), person, options.InsertOne()); err != nil {
		return err
	}

	return nil
}

func (m *MongoDatabase) FindByName(name string) (*Person, error) {
	collection := m.client.Database("persons").Collection("person")
	person := &Person{}
	if err := collection.FindOne(context.Background(), bson.M{"name": name}, options.FindOne()).Decode(person); err != nil {
		return nil, err
	}
	return person, nil
}
