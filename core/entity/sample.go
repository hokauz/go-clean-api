package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Sample -
type Sample struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Text string             `json:"text" bson:"text"`
}
