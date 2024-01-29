package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	title := vars["title"]
	numbers := vars["numbers"]
	fmt.Fprint(w, " you've requested the book ", title, "on page", numbers)
}
func main() {

	r := mux.NewRouter()
	//books/go-books/page/100
	r.HandleFunc("/books/{title}/page/{numbers}", Handler)
	http.ListenAndServe(":8080", r)

}
