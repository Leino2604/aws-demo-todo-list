package dynamodb

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/yourusername/todo-app/backend/internal/models"
)

type TodoRepository struct {
	db *dynamodb.DynamoDB
}

func NewTodoRepository() *TodoRepository {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)

	return &TodoRepository{db: db}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *models.Todo) error {
	av, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return fmt.Errorf("failed to marshal todo: %v", err)
	}

	_, err = r.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("Todos"),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to put item in DynamoDB: %v", err)
	}

	return nil
}

func (r *TodoRepository) GetTodos(ctx context.Context) ([]models.Todo, error) {
	result, err := r.db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("Todos"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %v", err)
	}

	var todos []models.Todo
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &todos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal items: %v", err)
	}

	return todos, nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, todo *models.Todo) error {
	av, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return fmt.Errorf("failed to marshal todo: %v", err)
	}

	_, err = r.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("Todos"),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to update item in DynamoDB: %v", err)
	}

	return nil
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id string) error {
	_, err := r.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("Todos"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to delete item from DynamoDB: %v", err)
	}

	return nil
}