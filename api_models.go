package main

import (
	"time"
)

type TodoStatus string

const (
	CREATED   TodoStatus = "created"
	PENDING              = "pending"
	COMPLETED            = "completed"
)

type Todo struct {
	Title   string     `json:"title"`
	Status  TodoStatus `json:"status"`
	Created time.Time  `json:"created"`
}

type LoginForm struct {
	Email      string
	Password   string
	RememberMe bool
}
