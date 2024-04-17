package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tolazhewa/todo-backend-go/models"
)

func GetTodos() ([]models.Todo, error) {
	db := GetDB()
	queryStatement := /* sql */ `
	SELECT 
			id, 
			user_id, 
			title, 
			completed, 
			create_datetime, 
			update_datetime
	FROM todos`
	rows, err := db.Query(queryStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]models.Todo, 0)
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Completed, &todo.Created, &todo.Updated)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	return todos, nil
}

func GetTodo(todoId string) (models.Todo, error) {
	queryStatement := /* sql */ `
	SELECT 
			id, 
			user_id, 
			title, 
			completed, 
			create_datetime, 
			update_datetime
	FROM 
		todos 
	WHERE 
		id = $1`

	db := GetDB()
	row := db.QueryRow(queryStatement, todoId)
	var todo models.Todo
	err := row.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Completed, &todo.Created, &todo.Updated)
	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to get todo: %w", err)
	}
	return todo, nil
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	todo.Id = uuid.New().String()
	todo.Created = time.Now().Format(time.RFC3339)
	todo.Updated = time.Now().Format(time.RFC3339)

	queryStatement := /* sql */ `INSERT INTO "todos" (
			id, 
			user_id, 
			title, 
			completed, 
			create_datetime, 
			update_datetime
		) 
		VALUES 
			($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			user_id,
			title,
			completed,
			create_datetime,
			update_datetime
		`
	var newTodo models.Todo

	db := GetDB()
	err := db.QueryRow(
		queryStatement,
		todo.Id,
		todo.UserId,
		todo.Title,
		todo.Completed,
		todo.Created,
		todo.Updated,
	).Scan(
		&newTodo.Id,
		&newTodo.UserId,
		&newTodo.Title,
		&newTodo.Completed,
		&newTodo.Created,
		&newTodo.Updated,
	)

	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to create and scan new todo: %w", err)
	}

	return newTodo, nil
}

func UpdateTodo(todo models.Todo) (models.Todo, error) {
	todo.Updated = time.Now().Format(time.RFC3339)

	queryStatement := /* sql */ `
	UPDATE todos 
	SET 
		title=$1, 
		completed=$2, 
		update_datetime=$3
	WHERE 
		id=$4
	RETURNING
		
		`
	db := GetDB()
	err := db.QueryRow(
		queryStatement,
		todo.Title,
		todo.Completed,
		todo.Updated,
		todo.Id,
	).Scan(
		&todo.Id,
		&todo.UserId,
		&todo.Title,
		&todo.Completed,
		&todo.Created,
		&todo.Updated,
	)

	if err != nil {
		return models.Todo{}, fmt.Errorf("failed to update and scan new todo: %w", err)
	}
	return todo, nil
}

func DeleteTodo(todoId string) error {
	queryStatement := /* sql */ `DELETE FROM todos WHERE id=$1`
	db := GetDB()
	_, err := db.Exec(queryStatement, todoId)
	if err != nil {
		return fmt.Errorf("Failed to delete to do: %w", err)
	}
	return nil
}
