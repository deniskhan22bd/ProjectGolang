package models

import (
	"database/sql"
	"log"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}