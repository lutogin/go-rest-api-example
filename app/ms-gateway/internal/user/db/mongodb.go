package db

import (
	"context"
	"errors"
	"log"
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
	oid, err := primitive.ObjectIDFromHex(payload.Id)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if err = result.Err(); err != nil {
		d.logger.Errorf("Error during looking for user by id: %s /n", payload.Id)
		d.logger.Traceln(payload.Id)
		return user, err
	}
	if err = result.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (d *db) GetByFilter(ctx context.Context, payload userDto.GetUsersDto) (user []user.UserEntity, err error) {
	// Marshal the anonymous struct into BSON bytes
	bsonBytes, err := bson.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshaling userDto.GetUsersDto struct:", err)
	}

	// Unmarshal the BSON bytes into a bson.M
	var filter bson.M
	err = bson.Unmarshal(bsonBytes, &filter)
	if err != nil {
		log.Fatal("Error unmarshaling BSON bytes:", err)
	}

	cursor, err := d.collection.Find(ctx, filter)
	if err != nil {
		d.logger.Errorf("Error during looking for user by filter: %s /n", payload)
		d.logger.Traceln(payload)
		return user, err
	}
	if err = cursor.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (d *db) Update(ctx context.Context, payload userDto.UpdateUserDto) (err error) {
	oid, err := primitive.ObjectIDFromHex(payload.Id)
	if err != nil {
		d.logger.Errorf("Failed to convert ObjectIDFromHex: %v", err)
		return err
	}

	filter := bson.M{"_id": oid}
	payloadBytes, err := bson.Marshal(payload)
	if err != nil {
		d.logger.Errorf("Failed to bson marshal: %v", err)
		return err
	}

	var updatePayload bson.M
	err = bson.Unmarshal(payloadBytes, &updatePayload)
	if err != nil {
		d.logger.Errorf("Failed to Unmarshal: %v", err)
		return err
	}

	delete(updatePayload, "Id") // drop ID from payload to update

	update := bson.M{"$set": updatePayload}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		d.logger.Errorf("Failed to update user: %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		d.logger.Tracef("not found at update by ID: %s", payload.Id)
		return errors.New("not found")
	}
	return nil
}

func (d *db) Delete(ctx context.Context, payload userDto.DeleteUserDto) (err error) {
	uid, err := primitive.ObjectIDFromHex(payload.Id)
	if err != nil {
		d.logger.Errorf("Failed to ObjectIDFromHex for ID: %s", payload.Id)
		return err
	}
	filter := bson.M{"_id": uid}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		d.logger.Tracef("User ID: %s", payload.Id)
		d.logger.Errorf("Failed to delete user: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		d.logger.Errorf("Not found user with ID: %s", payload.Id)
	}

	return nil
}

func NewRepository(database *mongo.Database, collection string, logger *logging.Logger) user.Repository {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
