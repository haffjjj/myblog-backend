package user

import (
	"context"

	"github.com/haffjjj/myblog-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoTagRepository struct {
	mgoClient *mongo.Client
}

//NewMongoUserRespository ...
func NewMongoUserRespository(c *mongo.Client) Repository {
	return &mongoTagRepository{c}
}

func (m *mongoTagRepository) GetByUsername(u string) (*models.User, error) {
	collection := m.mgoClient.Database("myblog").Collection("users")

	var user models.User

	err := collection.FindOne(context.TODO(), bson.D{{"username", u}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
