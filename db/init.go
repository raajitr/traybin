package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	DBName  string `yaml:"dbname"`
	ConnStr string `yaml:"connstr"`
}

type Database struct {
	db             *mongo.Database
	client         *mongo.Client
	userCollection *mongo.Collection
}

func NewDatabase(ctx context.Context) (*Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("uri"))
	if err != nil {
		return nil, fmt.Errorf("mongo client creation: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("mongo client connection: %w", err)
	}
	db := client.Database("traybin")

	return &Database{
		db:             db,
		client:         client,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.client.Ping(ctx, readpref.Primary())
}

func (d *Database) Close(ctx context.Context) error {
	return d.client.Disconnect(ctx)
}