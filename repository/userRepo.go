package repository

import (
	"context"
	"golang_noti/auth"
	"golang_noti/configs"
	"golang_noti/constants"
	"golang_noti/dto"
	"golang_noti/model"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	GetAllUser() ([]model.Users, error)
	CreateUser(body dto.CreateUser) constants.MessageResult
	UpdateUserById(id primitive.ObjectID, body dto.UpdateUser) constants.MessageResult
	DeleteUserById(id primitive.ObjectID) constants.MessageResult
}

type DB struct {
	collection *mongo.Collection
}

func UserRepo() UserRepository {
	return &DB{
		collection: configs.GetCollention(configs.ConfigDB(), "users"),
	}
}

func (u *DB) GetAllUser() ([]model.Users, error) {
	users := []model.Users{}
	filter := bson.M{}
	projection := bson.M{"password": 1}
	cursor, err := u.collection.Find(context.Background(), filter, options.Find().SetProjection(projection))

	for cursor.Next(context.Background()) {
		user := model.Users{}
		if err := cursor.Decode(&user); err == nil {
			users = append(users, user)
		}
	}

	defer cursor.Close(context.Background())

	return users, err
}

func (u *DB) CreateUser(body dto.CreateUser) constants.MessageResult {
	salt := auth.GenerateRandomSalt()
	netPassword := auth.HashedPassword(body.Password, salt)
	user := model.Users{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.Phone,
		Salt:      salt,
		Password:  netPassword,
		Role:      body.Role,
		Status:    body.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	filter := bson.M{"email": body.Email}
	var decode model.Users
	findByEmail := u.collection.FindOne(context.Background(), filter).Decode(&decode)
	if findByEmail == nil {
		return constants.MessageResult{
			Message: "Email Duplicate",
			Status:  false,
		}
	}

	_, err := u.collection.InsertOne(context.Background(), user)

	if err != nil {
		return constants.MessageResult{
			Message: "error",
			Status:  false,
		}
	}

	return constants.MessageResult{
		Status: true,
	}
}

func (u *DB) UpdateUserById(id primitive.ObjectID, body dto.UpdateUser) constants.MessageResult {

	filter := bson.M{"_id": id}
	var decode model.Users
	findById := u.collection.FindOne(context.Background(), filter).Decode(&decode)

	if findById != nil {
		return constants.MessageResult{
			Message: "User not found",
			Status:  false,
		}
	}

	updateUser := model.Users{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.Phone,
		UpdatedAt: time.Now(),
	}

	update := bson.M{"$set": updateUser}

	_, err := u.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return constants.MessageResult{
			Message: "Update User fail",
			Status:  false,
		}
	}

	return constants.MessageResult{
		Status: true,
	}
}
func (u *DB) DeleteUserById(id primitive.ObjectID) constants.MessageResult {
	filter := bson.M{"_id": id}
	var result model.Users
	findById := u.collection.FindOne(context.Background(), filter).Decode(&result)

	if findById != nil {
		return constants.MessageResult{
			Message: "User not found",
			Status:  false,
		}
	}

	deleteResult, _ := u.collection.DeleteOne(context.Background(), filter)

	if deleteResult.DeletedCount == 0 {
		return constants.MessageResult{
			Message: "Deleting user fail",
			Status:  false,
		}
	}

	return constants.MessageResult{
		Status: true,
	}

}
