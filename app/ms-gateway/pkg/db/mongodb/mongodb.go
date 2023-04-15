package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnectOpt struct {
	ctx      context.Context
	host     string
	port     string
	user     string // optional field
	password string // optional field
	database string
	auth     string
}

func NewClient(opt MongoConnectOpt) (db *mongo.Database, err error) {
	var mongoConnect string

	if opt.user != "" && opt.password != "" {
		mongoConnect = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			opt.user,
			opt.password,
			opt.host,
			opt.port,
		)
	} else {
		mongoConnect = fmt.Sprintf("mongodb://%s:%s",
			opt.host,
			opt.port,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnect))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect mongoDB: %v", err))
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to ping mongoDB: %v", err))
	}

	return client.Database(opt.database), nil
}
