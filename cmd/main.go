package main

import (
	"database/sql"
	"flag"
	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/jsonlog"
	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/models"
	"os"

	_ "github.com/lib/pq"
)

type config struct {
	port int
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
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
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

	// Call app.serve() to start the server.
	err = app.serve()
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
