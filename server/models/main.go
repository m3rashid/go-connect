package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDataLayer() (*mongo.Client, context.Context) {
	const mongodbUri = "mongodb://localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}

	return client, ctx
}

const databaseName = "jmiconnect"

const (
	AuthCollectionName            = "auth"
	AvatarsCollectionName         = "avatars"
	BookmarksCollectionName       = "bookmark"
	ChatsCollectionName           = "chats"
	ClassificationsCollectionName = "classifications"
	CommentsCollectionName        = "comments"
	FriendsCollectionName         = "friends"
	LikesCollectionName           = "likes"
	PostsCollectionName           = "posts"
	TopicsCollectionName          = "topics"
)

func GetCollection(db *mongo.Client, collectionName string) *mongo.Collection {
	return db.Database(databaseName).Collection(collectionName)
}
