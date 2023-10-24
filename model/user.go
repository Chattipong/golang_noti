package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersInfo struct {
	Users []Users `json:"users"`
}
type Users struct {
	UesrID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Phone     string             `json:"phone" bson:"phone"`
	Salt      []byte             `json:"salt ,omitempty" bson:"salt,omitempty" `
	Password  string             `json:"password ,omitempty" bson:"password,omitempty" `
	Role      int                `json:"role" bson:"role"`
	Status    bool               `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
