package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	defaultPort := "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on port %s.\n", port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func parseRequest(r *http.Request) (interface{}, error) {
	var p interface{}
	if r.Method != http.MethodPost {
		return nil, errors.New("Unsupported method error")
	}
	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]
	if contentType == "application/json" {
		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			return nil, err
		}
		body := make([]byte, length)
		length, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err = json.Unmarshal(body[:length], &p); err != nil {
			return nil, err
		}
	} else if contentType == "application/x-www-form-urlencoded" {
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Unsupported Content-Type")
	}
	return &p, nil
}
