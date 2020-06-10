package repository

import (
	"Demo/entity"
	"log"

	//"Demo/errors"
	"context"
	//"errors"
	"go.mongodb.org/mongo-driver/bson"
	//"log"
)

type Repository interface {
	FindByUsername(user *entity.User) (bool, error)
	FindByEmail(user *entity.User) (bool, error)
	//FindByEmailAndPassword(email, password string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	//DoesEmailExist(email string) (bool, error)
}

type repo struct {}

func NewRepository() Repository {
	return &repo{}
}

func (*repo) FindByUsername(user *entity.User) (bool, error) {
	var result *entity.User
	err := userCollection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}
	}
	return true, err
}

func (*repo) FindByEmail(user *entity.User) (bool, error) {
	var result entity.User
	err := userCollection.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}
	}
	return true, err
}

//func (*repo) FindByEmailAndPassword(email, password string) (*entity.User, error) {
//
//}

func (*repo) Create(user *entity.User) (*entity.User, error) {
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatalf("Could not add new User %v", err)
		//return nil, nil
	}
	return user, nil
}

//func (*repo) DoesEmailExist(email string) (bool, error) {
//
//}