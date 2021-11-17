package bd

import (
	"context"
	"time"

	"github.com/nelsomar/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* CheckUserExists, this function check if the user exists into database */
func CheckUserExists(email string) (models.Users, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.Users

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
