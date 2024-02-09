package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FromMongo struct {
	Server                string
	DbName                string
	PoliticiansCollection string
	VotesCollection       string
}

func (m *FromMongo) AllPoliticians() (Politicians, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Server))
	if err != nil {
		return nil, err
	}

	collection := client.Database(m.DbName).Collection(m.PoliticiansCollection)
	var result []Politician
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result, nil
}

func (m *FromMongo) AllVotes() (Votes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Server))
	if err != nil {
		return nil, err
	}

	collection := client.Database(m.DbName).Collection(m.VotesCollection)
	var result []Vote
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result, nil
}

func (m *FromMongo) PoliticianFromID(ID int) (Politician, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Server))
	if err != nil {
		return Politician{}, err
	}

	collection := client.Database(m.DbName).Collection(m.PoliticiansCollection)
	var result Politician
	err = collection.FindOne(ctx, bson.D{{Key: "id", Value: ID}}).Decode(&result)
	if err != nil {
		return Politician{}, err
	}

	return result, nil
}
