package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (apiCfg *apiConfig) handlerGetByStatus(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Status string `json:"status"`
	}

	var sTodos []Todo

	decoder := json.NewDecoder(r.Body)
	parms := parameters{}

	err := decoder.Decode(&parms)

	if err != nil {
		log.Fatal(err)
	}

	todos, erro := apiCfg.DB.GetTodosByStatus(r.Context(), parms.Status)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, todo := range todos {
		sTodos = append(sTodos, databaseTodoToTodo(todo))
	}

	res, er := json.Marshal(sTodos)

	if er != nil {
		log.Fatal(er)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
