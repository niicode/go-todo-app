package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerGetTodo(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	parms := parameters{}

	err := decoder.Decode(&parms)

	if err != nil {
		log.Fatal(err)
	}

	todo, ero := apiCfg.DB.GetTodo(r.Context(), parms.ID)

	if ero != nil {
		log.Fatal(ero)
	}

	res, er := json.Marshal(databaseTodoToTodo(todo))

	if er != nil {
		log.Fatal(er)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
