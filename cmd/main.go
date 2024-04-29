package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/jsonlog"
	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/peterbourgon/ff/v3"
)

type config struct {
	port       int
	env        string
	migrations string
	db         struct {
		dsn string
	}
}

type application struct {
	config config
	models models.Models
	logger *jsonlog.Logger
}

func main() {
	fs := flag.NewFlagSet("demo-app", flag.ContinueOnError)

	var (
		cfg        config
		migrations = fs.String("migrations", "file://pkg/migrations", "Path to migration files folder. If not provided, migrations do not applied")
		port       = fs.Int("port", 8080, "API server port")
		env        = fs.String("env", "development", "Environment (development|staging|production)")
		dbDsn      = fs.String("dsn", "postgres://postgres:123@localhost:5432/golang_project?sslmode=disable", "PostgreSQL DSN")
	)
	// Init logger
	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)
	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars()); err != nil {
		logger.PrintFatal(err, nil)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	cfg.port = *port
	cfg.env = *env
	cfg.db.dsn = *dbDsn
	cfg.migrations = *migrations

	logger.PrintInfo("starting application with configuration", map[string]string{
		"port":       fmt.Sprintf("%d", cfg.port),
		"env":        cfg.env,
		"db":         cfg.db.dsn,
		"migrations": cfg.migrations,
	})

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

	if cfg.migrations != "" {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return nil, err
		}
		m, err := migrate.NewWithDatabaseInstance(
			cfg.migrations,
			"postgres", driver)
		if err != nil {
			return nil, err
		}
		m.Up()
	}

	return db, nil
}
