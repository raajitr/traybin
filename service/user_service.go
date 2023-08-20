package service

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Username string `json:"userId"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type UserService struct {
	collections *mongo.Collection //TODO mongodb something
}

func NewUserService(db mongo.Database) *UserService {
	return &UserService{
		collections: db.Collection("user"),
	}
}

func (s UserService) GetId() {
	// fetch user with id
}

func (s UserService) Create() {
	// create user
}

func (s UserService) Update() {
	// update user
}

func (s UserService) Delete() {
	// delete user
}

