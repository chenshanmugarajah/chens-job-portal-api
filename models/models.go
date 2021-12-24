package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID         primitive.ObjectID `json: "_id,omitempty" bson: "_id,omitempty"`
	Company    string             `json: "company,omitempty"`
	Title      string             `json: "title,omitempty"`
	Experience string             `json: "experience,omitempty"`
	Salary     string             `json: "salary,omitempty"`
	Link       string             `json: "link,omitempty"`
}
