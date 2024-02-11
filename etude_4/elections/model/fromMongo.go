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

	if err = cur.All(context.Background(), &result); err != nil {
		log.Fatal(err)
		return nil, err
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

	if err = cur.All(context.Background(), &result); err != nil {
		log.Fatal(err)
		return nil, err
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

func (m *FromMongo) Winner() (Politician, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Server))
	if err != nil {
		return Politician{}, err
	}
	collection := client.Database(m.DbName).Collection(m.VotesCollection)

	resp := []bson.M{}
	groupStage := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: "$politician_id"},
		{Key: "votes", Value: bson.D{{Key: "$sum", Value: 1}}},
	}}}

	sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: "votes", Value: -1}}}}

	limitStage := bson.D{{Key: "$limit", Value: 1}}

	opts := options.Aggregate().SetMaxTime(2 * time.Second)
	cursor, err := collection.Aggregate(
		context.TODO(),
		mongo.Pipeline{groupStage, sortStage, limitStage},
		opts)
	if err != nil {
		return Politician{}, err
	}

	if err = cursor.All(context.TODO(), &resp); err != nil {
		return Politician{}, err
	}

	id := int(resp[0]["_id"].(int32)) // Convert id to int
	p, err := m.PoliticianFromID(id)
	if err != nil {
		return Politician{}, err
	}

	return p, nil
}
