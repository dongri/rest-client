package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Ping ...
type Ping struct {
	Method      string
	ContentType string
	Name        string
}

func main() {
	defaultPort := "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on port %s.\n", port)
	http.ListenAndServe("localhost:"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]
	var name string
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
		if contentType == "application/json" {
			length, err := strconv.Atoi(r.Header.Get("Content-Length"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			body := make([]byte, length)
			length, err = r.Body.Read(body)
			if err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			type Params struct {
				Name string `json:"name"`
			}
			p := new(Params)
			if err = json.Unmarshal(body[:length], &p); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			name = p.Name
		} else if contentType == "application/x-www-form-urlencoded" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "parse form error", http.StatusInternalServerError)
				return
			}
			name = r.Form.Get("name")
		} else {
			http.Error(w, "content-type error", http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodGet || r.Method == http.MethodDelete {
		name = r.URL.Query().Get("name")
	}
	ping := Ping{
		Method:      r.Method,
		ContentType: contentType,
		Name:        name,
	}
	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
