package repo

import (
	"context"
	"fmt"
	"time"
	"todolist/internal/models"
	"todolist/internal/usecase"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	*mongo.Client
}

func NewRepo(db *mongo.Client) usecase.CRUD {
	return &Repo{db}
}

func (r *Repo) Create(ctx context.Context, t models.TodoList) error {
	collection := r.Client.Database("test").Collection("tasks")

	t.ID = primitive.NewObjectID()

	InsertRes, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	fmt.Println("Inserted a single document: ", InsertRes.InsertedID.(primitive.ObjectID))

	return nil
}

func (r *Repo) Read(ctx context.Context, t models.TodoList) (models.TodoList, error) {
	var res models.TodoList
	collection := r.Client.Database("test").Collection("tasks")

	filter := bson.M{"ID": t.ID}

	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return models.TodoList{}, err
	}

	fmt.Println("got a sing document: ", res)
	return models.TodoList{}, nil
}

func (r *Repo) UpdateByID(ctx context.Context, t models.TodoList) error {
	filter := bson.M{"ID": t.ID}
	collection := r.Client.Database("test").Collection("tasks")

	update := bson.M{"$set": bson.M{"Title": t.Title, "ActiveAt": t.ActiveAt}}

	res, err := collection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	fmt.Println("Updated a single document: ", res)

	return nil
}

func (r *Repo) MarkStatusByID(ctx context.Context, t models.TodoList) error {
	filter := bson.M{"ID": t.ID}
	collection := r.Client.Database("test").Collection("tasks")

	update := bson.M{"$set": bson.M{"StatusDone": true}}

	res, err := collection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		return err
	}

	fmt.Println("Updated a single document: ", res)

	return nil
}

func (r *Repo) List(ctx context.Context) ([]models.TodoList, error) {
	var results []models.TodoList
	filter := bson.M{"ActiveAt": bson.M{"$lte": time.Now()}}
	collection := r.Client.Database("test").Collection("tasks")

	options := options.Find()

	res, err := collection.Find(context.TODO(), filter, options)

	if err != nil {
		return []models.TodoList{}, err
	}

	for res.Next(context.TODO()) {
		var doc models.TodoList
		if err := res.Decode(&doc); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}

	if err := res.Err(); err != nil {
		return []models.TodoList{}, err
	}

	return results, nil
}

func (r *Repo) Delete(ctx context.Context, t models.TodoList) error {
	filter := bson.M{"ID": t.ID}
	collection := r.Client.Database("test").Collection("tasks")

	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Println("Deleted a single document: ", res)
	return nil
}
