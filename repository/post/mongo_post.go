package post

import (
	"context"

	"github.com/haffjjj/myblog-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoPostRepository struct {
	mgoClient *mongo.Client
}

//NewMongoPostRespository ...
func NewMongoPostRespository(c *mongo.Client) Repository {
	return &mongoPostRepository{c}
}

func (m *mongoPostRepository) GetGroups() ([]*models.PostsGroup, error) {

	var collection = m.mgoClient.Database("myblog").Collection("posts")
	var postsGroups []*models.PostsGroup

	//aggregate to get data
	cur, err := collection.Aggregate(context.TODO(), []bson.D{
		bson.D{
			{"$matc", bson.D{
				{"tag", bson.D{
					{"$regex", ""},
				}},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"count", bson.D{{"$sum", 1}}},
				{"data", bson.D{{"$push", "$$ROOT"}}},
			}},
		},
		bson.D{
			{"$unwind", "$data"},
		},
		bson.D{
			{"$replaceRoot", bson.D{
				{"newRoot", bson.D{
					{"$mergeObjects", bson.A{"$data", "$$ROOT"}},
				}},
			}},
		},
		bson.D{
			{"$group", bson.D{
				{"_id", nil},
				{"count", bson.D{
					{"$first", "$count"},
				}},
				{"data", bson.D{
					{"$push", "$data"},
				}},
			}},
		},
	})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var elem models.PostsGroup
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		postsGroups = append(postsGroups, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	return postsGroups, nil
}
