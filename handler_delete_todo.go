package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	type deleteParms struct {
		ID uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	parms := deleteParms{}

	err := decoder.Decode(&parms)

	if err != nil {
		log.Fatal(err)
	}

	erro := apiCfg.DB.DeleteTodo(r.Context(), parms.ID)

	if erro != nil {
		log.Fatal(erro)
	}

	res, er := json.Marshal("Deleted")

	if er != nil {
		log.Fatal(er)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
