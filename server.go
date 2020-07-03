package main

import (
    "fmt"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{book_id}/pages/{page_id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["book_id"]
		pageID := vars["page_id"]

		fmt.Fprintf(w, "You are requested the book: %s on page %s\n", bookID, pageID)
	})

	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", r)
}