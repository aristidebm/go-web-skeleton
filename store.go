package main

import "time"

type TodoRepository struct {
	todos []Todo
}

func (r *TodoRepository) ListTodos() ([]Todo, error) {
	return r.todos[:], nil
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: []Todo{
			{
				Title:   "Learn Go",
				Status:  PENDING,
				Created: time.Now().UTC(),
			},
			{
				Title:   "Learn Rust",
				Status:  COMPLETED,
				Created: time.Now().UTC(),
			},
			{
				Title:   "Learn System Design",
				Status:  CREATED,
				Created: time.Now().UTC(),
			},
			{
				Title:   "Learn Tennis",
				Status:  CREATED,
				Created: time.Now().UTC(),
			},
			{
				Title:   "Learn Piano",
				Status:  CREATED,
				Created: time.Now().UTC(),
			},
		},
	}
}
