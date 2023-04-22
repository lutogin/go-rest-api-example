package mng

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ms-users/pkg/logging"
	"time"
)

type MongoConnectOpt struct {
	Host      string
	Port      string
	User      string // optional field
	Password  string // optional field
	Database  string
	UriScheme string
}

func NewClient(opt MongoConnectOpt, logger *logging.Logger) (db *mongo.Database, err error) {
	var mongoConnect string

	suffix := "?retryWrites=true&w=majority"
	if opt.User != "" && opt.Password != "" {
		mongoConnect = fmt.Sprintf("%s://%s:%s@%s:%s/%s%s",
			opt.UriScheme,
			opt.User,
			opt.Password,
			opt.Host,
			opt.Port,
			opt.Database,
			suffix,
		)
	} else {
		mongoConnect = fmt.Sprintf("%s://%s:%s/%s%s",
			opt.UriScheme,
			opt.Host,
			opt.Port,
			opt.Database,
			suffix,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnect))

	//defer func() {
	//	logger.Infoln("Disconnected from mongo.")
	//	if app-errors = client.Disconnect(ctx); app-errors != nil {
	//		panic(app-errors)
	//	}
	//}()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect mongoDB: %v", err))
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to ping mongoDB: %v", err))
	}
	logger.Infoln("Connected to mongo successfully")

	return client.Database(opt.Database), nil
}
