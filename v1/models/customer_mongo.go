package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	CustomerID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PersonalData PersonalData       `json:"personal_data,omitempty" bson:"personal_data,omitempty"`
	Bill         []Bill             `json:"bills,omitempty" bson:"bills,omitempty"`
}

type PersonalData struct {
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Address   string `json:"address,omitempty" bson:"address,omitempty"`
	Phone     string `json:"phone,omitempty" bson:"phone,omitempty"`
	BirthDate string `json:"age,omitempty" bson:"age,omitempty"`
	Gender    string `json:"gender,omitempty" bson:"gender,omitempty"`
	Picture   string `json:"picture_path,omitempty" bson:"picture_path,omitempty"`
}
