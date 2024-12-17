package main

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Make(fun HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := fun(w, r); err != nil {
			// TODO: handle API errors
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
