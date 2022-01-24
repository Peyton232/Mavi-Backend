package database

import (
	"context"
	"time"

	"github.com/Peyton232/Mavi-Backend/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

// ----------------------------------------------DELETION-----------------------------------------------
// Deletion
func (db *DB) RemoveUser(userid string) *model.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := db.FindUserByID(userid)
	if user == nil {
		return nil
	}

	collection.FindOneAndDelete(ctx, bson.M{"userid": userid})
	return user
}

// ----------------------------------------------QUERIES-----------------------------------------------
func (db *DB) FindUserByID(ID string) *model.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	collection.FindOne(ctx, bson.M{"userid": ID}).Decode(&user)
	return &user
}
