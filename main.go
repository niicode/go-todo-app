package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/niicode/go-todo-app/internal/db"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	db_url := os.Getenv("db_Url")

	if db_url == "" {
		log.Fatal("Db url is not defined")
	}

	if port == "" {
		log.Fatal("The port must be defined")
	}

	//connect to db
	conn, err := sql.Open("postgres", db_url)

	if err != nil {
		log.Fatal(err)
	}

	qry := db.New(conn)

	apiCfg := &apiConfig{
		DB: qry,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1 := chi.NewRouter()
	v1.Post("/todo", apiCfg.handlerCreateTodo)
	v1.Put("/update_todo", apiCfg.handleUpdateTodo)
	v1.Delete("/del_todo", apiCfg.handleDeleteTodo)
	v1.Get("/get_todo", apiCfg.handlerGetTodo)
	v1.Get("/todos", apiCfg.handlerGetTodos)
	v1.Get("/get_by_status", apiCfg.handlerGetByStatus)

	router.Mount("/v1", v1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server is running on Port %v", port)

	erro := srv.ListenAndServe()

	if erro != nil {
		log.Fatal("Failed to run server", erro)
	}
}

