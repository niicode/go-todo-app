package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (apiCfg *apiConfig) handlerGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := apiCfg.DB.GetTodos(r.Context())

	var getTodos []Todo

	if err != nil {
		log.Fatal(err)
	}

	for _, todo := range todos {
		getTodos = append(getTodos, databaseTodoToTodo(todo))
	}

	res, erro := json.Marshal(getTodos)

	if erro != nil {
		log.Fatal(erro)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
