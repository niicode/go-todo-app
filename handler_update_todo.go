package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/niicode/go-todo-app/internal/db"
)

func (apiCfg *apiConfig) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID     uuid.UUID `json:"id"`
		Status string    `json:"status"`
	}

	decoder := json.NewDecoder(r.Body)
	parms := parameters{}

	err := decoder.Decode(&parms)

	if err != nil {
		log.Fatal(err)
	}

	updateTodo, erro := apiCfg.DB.UpdateTodo(r.Context(), db.UpdateTodoParams{
		ID:     parms.ID,
		Status: parms.Status,
	})

	if erro != nil {
		log.Fatal(erro)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
	}

	res, er := json.Marshal(databaseTodoToTodo(updateTodo))

	if er != nil {
		log.Fatal(er)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
