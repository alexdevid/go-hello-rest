package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, _ := strconv.Atoi(vars["todoId"])
	todos := Todos{
		Todo{Id: 1, Name: "Write presentation"},
		Todo{Id: 2, Name: "Host meetup"},
	}

	var todo Todo

	for _, value := range todos {
		if value.Id == todoId {
			todo = value
		}
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}
