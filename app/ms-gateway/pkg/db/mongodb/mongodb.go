package mng

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnectOpt struct {
	Ctx      context.Context
	Host     string
	Port     string
	User     string // optional field
	Password string // optional field
	Database string
	Auth     string
}

func NewClient(opt MongoConnectOpt) (db *mongo.Database, err error) {
	var mongoConnect string

	if opt.User != "" && opt.Password != "" {
		mongoConnect = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			opt.User,
			opt.Password,
			opt.Host,
			opt.Port,
		)
	} else {
		mongoConnect = fmt.Sprintf("mongodb://%s:%s",
			opt.Host,
			opt.Port,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnect))

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect mongoDB: %v", err))
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to ping mongoDB: %v", err))
	}

	return client.Database(opt.Database), nil
}
