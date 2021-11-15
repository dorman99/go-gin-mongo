package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/dorman99/go-gin-mongo/dto"
	"github.com/dorman99/go-gin-mongo/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository interface {
	Create(createPostDto dto.CreatePostDto) (entity.Post, error)
}

type postRepository struct {
	collectionName string
	collection     *mongo.Collection
}

func NewPostReposistory(db *mongo.Database, collectionName string) PostRepository {
	return &postRepository{
		collection:     db.Collection(collectionName),
		collectionName: collectionName,
	}
}

func (p *postRepository) Create(createPostDto dto.CreatePostDto) (entity.Post, error) {
	var post = entity.Post{
		Id:         primitive.NewObjectID(),
		Title:      createPostDto.Title,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	doc, err := p.collection.InsertOne(context.Background(), post)
	if err != nil {
		fmt.Println(err)
		return post, err
	}
	fmt.Println(doc)
	return post, nil
}
