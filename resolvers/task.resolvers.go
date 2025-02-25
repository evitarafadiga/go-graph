package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"

	"github.com/evitarafadiga/go-graph/generated"
	"github.com/evitarafadiga/go-graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var taskCollectionName = "task"

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, task model.TaskInput) (*model.Task, error) {
	collection := r.DB.Collection(taskCollectionName)

	newId := primitive.NewObjectID()

	newTask := model.Task{
		ID:      newId.Hex(),
		Name:    task.Name,
		Deleted: task.Deleted,
	}

	_, err := collection.InsertOne(ctx, newTask)
	if err != nil {
		return nil, err
	}

	return &newTask, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, task *model.TaskInput) (*model.Task, error) {
	collection := r.DB.Collection(taskCollectionName)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"name": task.Name,
			"deleted": task.Deleted,
		},
	}

	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updatedTask model.Task
	err = result.Decode(&updatedTask)
	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	collection := r.DB.Collection(taskCollectionName)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(ctx, filter)
	return true, err
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: Task - task"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	collection := r.DB.Collection(taskCollectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*model.Task
	for cursor.Next(ctx) {
		var task model.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
*/
