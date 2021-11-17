package bd

import (
	"context"
	"time"

	"github.com/nelsomar/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* RegisterUser, Insert user into database */
func RegisterUser(u models.Users) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	u.Password, _ = CryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjectID.String(), true, nil
}
