package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	return TextResponse(w, fmt.Sprintf("Hello %s !", name), http.StatusOK)
}

func ListTodos(w http.ResponseWriter, r *http.Request) error {
	repo := NewTodoRepository()
	data, err := repo.ListTodos()
	if err != nil {
		return err
	}
	return JsonResponse(w, data, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) error {
	data := NewTemplateData()
	data.Form = LoginForm{}
	templ, name, err := GetTemplate("ui/html/pages/login.html")
	if err != nil {
		return err
	}
	return TemplateResponse(w, templ, name, data, http.StatusOK)
}
