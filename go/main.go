package main

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/gorilla/mux"
)

type Task struct {
	DomainID string `json:"domainid"`
	Date     string `json:"date"`
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	TypeID   int    `json:"typeid,omitempty"`
}

var tasks [] Task

func GetTaskEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, t := range tasks{
		if t.ID == params["id"]{
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

func CreateTaskEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var task Task
	_ = json.NewDecoder(req.Body).Decode(&task)
	task.DomainID = params["domainid"]
	task.Date     = params["date"]
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for i, t := range tasks {
		if t.ID == params["id"]{
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

func main() {
	router := mux.NewRouter()
	tasks = append(tasks, Task{ID:"1", Name:"DuohongTask", TypeID:2})

	router.HandleFunc("/{domainid}/{date}/task/{id}", GetTaskEndpoint   ).Methods("GET")
	router.HandleFunc("/{domainid}/{date}/task",      CreateTaskEndpoint).Methods("POST")
	router.HandleFunc("/{domainid}/{date}/task/{id}", DeleteTaskEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
