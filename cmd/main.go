package main

import (
	"database/sql"
	"flag"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/jsonlog"
	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/models"

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
	logger *jsonlog.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8080", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:123@localhost/golang_project?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Init logger
	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: cfg,
		models: models.NewModels(db),
		logger: logger,
	}
	// Again, we use the PrintInfo() method to write a "starting server" message at the
	// INFO level. But this time we pass a map containing additional properties (the
	// operating environment and server address) as the final parameter.
	logger.PrintInfo("starting server", map[string]string{
		"addr": cfg.port,
		"env":  cfg.env,
	})

	err = app.run()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) run() error {
	r := mux.NewRouter()

	r.HandleFunc("/books", app.GetBooks).Methods("GET")
	r.HandleFunc("/books", app.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", app.GetBook).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", app.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", app.DeleteBook).Methods("DELETE")

	r.HandleFunc("/users", app.registerUserHandler).Methods("POST")
	r.HandleFunc("/users/activate", app.activateUserHandler).Methods("POST")
	r.HandleFunc("/users/login", app.createAuthenticationTokenHandler).Methods("POST")
	http.Handle("/", r)
	err := http.ListenAndServe(app.config.port, r)
	if err != nil {
		return err
	}
	return nil
}
