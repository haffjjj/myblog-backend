package tag

import (
	"context"

	"github.com/haffjjj/myblog-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoTagRepository struct {
	mgoClient *mongo.Client
}

//NewMongoTagRespository ...
func NewMongoTagRespository(c *mongo.Client) Repository {
	return &mongoTagRepository{c}
}

func (m *mongoTagRepository) Get() ([]*models.Tag, error) {
	var collection = m.mgoClient.Database("myblog").Collection("posts")
	var tags []*models.Tag

	//aggregate to get data
	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$unwind", "$tag"},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", "$tag"},
				{"tag", bson.D{
					{"$first", "$tag"},
				}},
			}},
		},
	})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem models.Tag
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	if tags == nil {
		return []*models.Tag{}, nil
	}

	return tags, nil
}
