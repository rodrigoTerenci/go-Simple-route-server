package main

import (
	"net/http"
	"text/template"
)

//Task is ...
type Task struct {
	Name string
	Done bool
}

//People is..
type People struct {
	Name     string
	LastName string
	Age      int
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/tasks", TasksHandler)
	http.ListenAndServe(":8887", nil)
}

//TasksHandler is ...
func TasksHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("tasks.html"))

	tasks := []Task{
		{Name: "Tarefa 1", Done: false},
		{Name: "Tarefa 2", Done: true},
	}
	tmpl.Execute(w, tasks)

}

//HelloHandler is ...
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("hello.html"))

	people := []People{
		{Name: "Ze", LastName: "Carvalho Terenci", Age: 35},
		{Name: "Rodrigo", LastName: "Delicia", Age: 38},
	}
	tmpl.Execute(w, people)

}
