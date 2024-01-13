package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/niicode/go-todo-app/internal/db"
)

type parameters struct {
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
}

func (p *parameters) UnmarshalJSON(data []byte) error {
	type Alias parameters
	aux := &struct {
		Description string `json:"description"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Description != "" {
		p.Description = sql.NullString{String: aux.Description, Valid: true}
	} else {
		p.Description = sql.NullString{Valid: false}
	}
	return nil
}

func (apiCfg *apiConfig) handlerCreateTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	parms := parameters{}

	err := decoder.Decode(&parms)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	todo, eror := apiCfg.DB.CreateTodo(r.Context(), db.CreateTodoParams{
		Title:       parms.Title,
		Description: parms.Description,
	})

	if eror != nil {
		log.Fatal(eror)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(eror.Error()))
	}

	res, erro := json.Marshal(databaseTodoToTodo(todo))

	if erro != nil {
		log.Fatal(erro)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
