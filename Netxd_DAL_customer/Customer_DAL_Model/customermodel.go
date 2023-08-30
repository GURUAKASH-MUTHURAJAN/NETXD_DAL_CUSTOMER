package customermodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	Bank_id   string    `json:"bank_id" bson:"bank_id"`
	Balance   float64   `json:"balance" bson:"balance"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	IsActive  bool      `json:"is_active" bson:"is_active"`
}
type CustomerResponse struct {
	Customer_Id primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}
