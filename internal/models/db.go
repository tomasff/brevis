package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Datastore interface {
	GetShortById(id string) (*Short, error)
	RemoveShort(id string) error
	CreateShort(url string) (*Short, error)
}

type DB struct {
	*mongo.Database
}

func (db *DB) GetShortById(id string) (*Short, error) {
	return nil, nil
}

func (db *DB) RemoveShort(id string) error {
	return nil
}

func (db *DB) CreateShort(url string) (*Short, error) {
	return nil, nil
}

func NewDB(dbName string, uri string) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5 * time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	database := client.Database(dbName)

	return &DB{database}, nil
}