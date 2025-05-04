package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Author      string             `json:"author,omitempty" bson:"author,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	ISBN        string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
}
