package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	ReadPage := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	}
	r.HandleFunc("/books/{title}/page/{page}", ReadPage)

	type Info struct {
		Title  string `json:"title"`
		Page   int    `json:"page"`
		Author string `json:"author"`
	}

	books := map[string]Info{}

	CreateBook := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var info Info
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			books[info.Title] = info
			log.Println("create book success")
		} else {
			log.Println("create book fail")
		}
	}

	ReadBook := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		w.Header().Set("Content-Type", "application/json")

		jsonResponse, jsonError := json.Marshal(books[title])
		if jsonError != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		}
	}

	UpdateBook := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		title := vars["title"]
		w.Header().Set("Content-Type", "application/json")

		var info Info
		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			books[title] = info

			resp := []byte(`{"status":"OK"}`)
			w.Write(resp)
		}

	}

	DeleteBook := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		w.Header().Set("Content-Type", "application/json")

		delete(books, title)
		resp := []byte(`{"status":"OK"}`)
		w.Write(resp)
	}

	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
