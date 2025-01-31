package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"

	"github.com/evitarafadiga/go-graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	
	var userCollectionName = "user"

	collection := r.DB.Collection(userCollectionName)

	newId := primitive.NewObjectID()

	newUser := model.User{
		ID:       newId.Hex(),
		Username: user.Username,
		Password: user.Password,
	}

	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, user *model.UserInput) (*model.User, error) {
	collection := r.DB.Collection(userCollectionName)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": objectID}
	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"password": user.Password,
		},
	}

	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updatedUser model.User
	err = result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	collection := r.DB.Collection(userCollectionName)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return true, err
	}

	filter := bson.M{"id": objectID}

	_, err = collection.DeleteOne(ctx, filter)
	return true, err
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	collection := r.DB.Collection(userCollectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*model.User
	for cursor.Next(ctx) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	
*/
