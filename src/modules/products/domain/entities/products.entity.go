package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          *int64             `json:"id"`
	MongoID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Price       *float64           `json:"price" bson:"price,omitempty"`
}
