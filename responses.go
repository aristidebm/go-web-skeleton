package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func TextResponse(w http.ResponseWriter, content string, status int) error {
	h := w.Header()
	h.Set(http.CanonicalHeaderKey("Content-Type"), "text/plain")

	var buf bytes.Buffer
	if _, err := fmt.Fprintln(&buf, content); err != nil {
		return err
	}
	return response(w, &buf, status)
}

func JsonResponse(w http.ResponseWriter, content any, status int) error {
	h := w.Header()
	h.Set(http.CanonicalHeaderKey("Content-Type"), "application/json")

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(content); err != nil {
		// TODO: This is seally remove that
		panic(err)
	}
	return response(w, &buf, status)
}

func TemplateResponse(w http.ResponseWriter, content *template.Template, name string, data any, status int) error {
	h := w.Header()
	h.Set(http.CanonicalHeaderKey("Content-Type"), "text/html")

	var buf bytes.Buffer
	templ := &Template{templates: content}
	templ.Render(&buf, name, data)

	return response(w, &buf, status)
}

func response(w http.ResponseWriter, content io.Reader, status int) error {
	w.WriteHeader(status)
	if _, err := io.Copy(w, content); err != nil {
		return err
	}
	return nil
}
