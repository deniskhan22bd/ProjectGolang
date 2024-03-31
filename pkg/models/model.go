package models

import (
	"database/sql"
	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/jsonlog"
	"os"
)

type Models struct {
	Books BookModel
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)
	return Models{
		Books: BookModel{
			DB:       db,
			Logger: logger,
		},
		Users: UserModel{
			DB:       db,
			Logger: logger,
		},
	}
}
