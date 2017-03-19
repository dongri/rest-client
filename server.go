package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pressly/chi"
	"github.com/unrolled/render"
)

func main() {
	defaultPort := "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	renderEngine := render.New(render.Options{
		IndentJSON:    true,
		IsDevelopment: true,
	})

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		output := map[string]string{
			"name": "yourserver",
		}
		renderEngine.JSON(w, http.StatusOK, output)
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		output := []interface{}{
			map[string]string{
				"name":  "taro",
				"email": "taro@domain.com",
			},
			map[string]string{
				"name":  "jiro",
				"email": "jiro@domain.com",
			},
		}
		renderEngine.JSON(w, http.StatusOK, output)
	})

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.Form.Get("name")
		email := r.Form.Get("email")
		output := map[string]string{
			"name":  name,
			"email": email,
		}
		renderEngine.JSON(w, http.StatusCreated, output)
	})

	r.Put("/users", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.Form.Get("name")
		email := r.Form.Get("email")
		output := map[string]string{
			"name":  name,
			"email": email,
		}
		renderEngine.JSON(w, http.StatusOK, output)
	})

	r.Delete("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		renderEngine.JSON(w, http.StatusOK, nil)
	})

	fmt.Printf("Server listening on port %s.\n", port)
	http.ListenAndServe(":"+port, r)
}
