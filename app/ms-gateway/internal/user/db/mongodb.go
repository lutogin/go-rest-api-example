package db

import (
	"context"
	"errors"
	"ms-gateway/internal/user"
	userDto "ms-gateway/internal/user/dto"
	"ms-gateway/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, payload userDto.CreateUserDto) (id string, err error) {
	result, err := d.collection.InsertOne(ctx, payload)
	if err != nil {
		return "", err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		errMsg := "error during getti ng oid"
		d.logger.Errorln(errMsg)
		d.logger.Traceln(payload)
		return "", errors.New(errMsg)
	}
	return oid.Hex(), nil
}

func (d *db) GetById(ctx context.Context, payload userDto.GetUserByIdDto) (user user.UserEntity, err error) {
	oid, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if err = result.Err(); err != nil {
		d.logger.Errorf("Error during looking for user by id: %s /n", payload.ID)
		d.logger.Traceln(payload.ID)
		return user, err
	}
	if err = result.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (d *db) GetByFilter(ctx context.Context, payload userDto.GetUsersDto) (user []user.UserEntity, err error) {
	filter := bson.M{payload}
	cursor, err := d.collection.Find(ctx, filter)
	if err != nil {
		d.logger.Errorf("Error during looking for user by filter: %s /n", payload)
		d.logger.Traceln(payload)
		return user, err
	}
	if err = cursor.Decode(&user); err != nil {
		return u, err
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, payload userDto.UpdateUserDto) (user user.UserEntity, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, payload userDto.DeleteUserDto) (err error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(database mongo.Database, collection string, logger *logging.Logger) user.Repository {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
