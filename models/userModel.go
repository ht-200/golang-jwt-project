package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `bson:"_id"`
	FirstName     *string            `json:"firstname" validate:"required, min=2 max=100"`
	LastName      *string            `json:"lastname" validate:"required, min=1 max=100"`
	Password      *string            `json:"password" validate:"required, min=8"`
	Email         *string            `json:"email" validate:"email, required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	User_Type     *string            `json:"userType" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_Token *string            `json:"refreshToken"`
	Created_At    time.Time          `json:"createdAt"`
	Updated_At    time.Time          `json:"updatedAt"`
	User_Id       *string            `json:"userId"`
}
