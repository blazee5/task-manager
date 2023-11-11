package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

type ShortUser struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty" redis:"id"`
	Name  string             `json:"name" bson:"name" redis:"name"`
	Email string             `json:"email" bson:"email" redis:"email"`
}
