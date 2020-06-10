package repository

import (
	"Demo/entity"
	"context"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type PostRepository interface {
	FindAllPost(posts []entity.Post) ([]entity.Post, error)
	FindOnePost(post *entity.Post) (*entity.Post, error)
	Create(post *entity.Post) (*entity.Post, error)
}

type postRepo struct {}

func NewPostRepository() PostRepository {
	return &postRepo{}
}

func (*postRepo) FindAllPost(posts []entity.Post) ([]entity.Post, error) {
	cursor, err := postCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalf("Could not get all post %v", err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var post entity.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	return posts, nil
}

func (*postRepo) FindOnePost(post *entity.Post) (*entity.Post, error) {
	var result *entity.Post
	err := postCollection.FindOne(context.TODO(), bson.D{{"title", post.Title}}).Decode(&result)
	if err != nil {
		log.Fatalf("could not get the post %v", err)
	}
	return post, nil
}

func (*postRepo) Create(post *entity.Post) (*entity.Post, error) {
	_, err := postCollection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatalf("Failed adding a new post %v", err)
	}
	return post, nil
}
