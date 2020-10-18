package sample

import (
	"context"
	"errors"

	"github.com/hokauz/go-clean-api/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoRespository mongodb repo
type MongoRespository struct {
	Collection *mongo.Collection
	Context    context.Context
}

// NewMongoRespository create new repository
func NewMongoRespository(ctx context.Context, coll *mongo.Collection) Repository {
	return &MongoRespository{Collection: coll, Context: ctx}
}

// ReadOne -
func (r *MongoRespository) ReadOne(id string) (data *entity.Sample, err error) {
	// TODO add filter to injection
	objID, _ := primitive.ObjectIDFromHex(id)
	err = r.Collection.FindOne(r.Context, bson.M{"_id": objID}).Decode(&data)
	return
}

// Create -
func (r *MongoRespository) Create(data *entity.Sample) (id string, err error) {
	// TODO add filter to injection
	data.ID = primitive.NewObjectID()

	res, err := r.Collection.InsertOne(r.Context, data)
	if err != nil {
		err = handlerErr(err)
		return
	}

	id = res.InsertedID.(primitive.ObjectID).Hex()
	return
}

// Update -
func (r *MongoRespository) Update(id string, data *entity.Sample) (d *entity.Sample, err error) {
	// TODO add filter to injection
	objID, _ := primitive.ObjectIDFromHex(id)
	data.ID = objID
	filter := bson.M{"_id": objID}
	after := options.After

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{
		primitive.E{Key: "$set", Value: data},
		primitive.E{Key: "$currentDate", Value: bson.D{
			primitive.E{Key: "last_modified", Value: true},
		}},
	}

	err = r.Collection.FindOneAndUpdate(r.Context, filter, update, &opt).Decode(&d)
	if err != nil {
		err = handlerErr(err)
		return
	}

	return
}

// Delete -
func (r *MongoRespository) Delete(id string) (err error) {
	// TODO add filter to injection
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	res := r.Collection.FindOneAndDelete(r.Context, filter)
	if res.Err() != nil {
		err = handlerErr(res.Err())
		return
	}

	return
}

func handlerErr(err error) error {
	if err.Error() == "mongo: no documents in result" {
		return errors.New("Nenhum dado encontrado")
	}

	return err
}
