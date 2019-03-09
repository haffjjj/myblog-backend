package post

import (
	"context"

	"github.com/haffjjj/myblog-backend/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type mongoPostRepository struct {
	mgoClient *mongo.Client
}

//NewMongoPostRespository ...
func NewMongoPostRespository(c *mongo.Client) Repository {
	return &mongoPostRepository{c}
}

//GetById ...
func (m *mongoPostRepository) GetByID(i string) (*models.Post, error) {
	var collection = m.mgoClient.Database("myblog").Collection("posts")

	var post models.Post

	IDHex, err := primitive.ObjectIDFromHex(i)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), bson.D{{"_id", IDHex}}).Decode(&post)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (m *mongoPostRepository) GetGroups(p models.Pagination) ([]*models.PostsGroup, error) {

	var collection = m.mgoClient.Database("myblog").Collection("posts")
	var postsGroups []*models.PostsGroup

	//aggregate to get data
	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
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
			{"$skip", p.Start},
		},
		bson.D{
			{"$limit", p.Limit},
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

	if postsGroups == nil {
		return []*models.PostsGroup{}, nil
	}

	return postsGroups, nil
}

func (m *mongoPostRepository) GetGroupsByTag(t string, p models.Pagination) ([]*models.PostsGroup, error) {

	var collection = m.mgoClient.Database("myblog").Collection("posts")
	var postsGroups []*models.PostsGroup

	//aggregate to get data
	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"tag", bson.D{
					{"$regex", t},
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
			{"$skip", p.Start},
		},
		bson.D{
			{"$limit", p.Limit},
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

	if postsGroups == nil {
		return []*models.PostsGroup{}, nil
	}

	return postsGroups, nil
}
