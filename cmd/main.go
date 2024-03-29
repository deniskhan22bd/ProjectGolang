package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/models"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models models.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8080", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:123@localhost/golang_project?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: models.NewModels(db),
	}

	app.run()
	// app.models.Books.GetAll()
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) run() {
	r := mux.NewRouter()
	r.HandleFunc("/health-check", app.HealthCheck).Methods("GET")

	r.HandleFunc("/books", app.GetBooks).Methods("GET")

	r.HandleFunc("/books/{id:[0-9]+}", app.GetBook).Methods("GET")

	r.HandleFunc("/books", app.CreateBook).Methods("POST")

	r.HandleFunc("/books/{id:[0-9]+}", app.UpdateBook).Methods("PUT")

	r.HandleFunc("/books/{id:[0-9]+}", app.DeleteBook).Methods("DELETE")

	http.Handle("/", r)

	log.Printf("starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}
