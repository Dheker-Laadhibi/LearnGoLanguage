package main

import (
	f "fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPage struct {
	PageTitle string
	Todos     []Todo
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	f.Fprintf(w, "Welcome to IndexPage")
}
func todosHandler(w http.ResponseWriter, r *http.Request) {
	data := TodoPage{
		PageTitle: "All Todos",
		Todos: []Todo{
			{Title: "Learn GO", Done: false},
			{Title: "Learn Python", Done: true},
			{Title: "Learn Java", Done: true},
		},
	}

	t, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		log.Fatal("Error encountered while parsing the template : ", err)
	}
	t.Execute(w, data)
}


func main() {
	// todos route and todosHandler is the controller
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/todos", todosHandler)
	http.ListenAndServe(":8080", nil)

}
