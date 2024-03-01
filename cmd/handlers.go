package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/models"
	"github.com/gorilla/mux"
)

func (app *application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is running! \nWeb app about books by Khan Denis")
}

func (app *application) GetBooks(w http.ResponseWriter, r *http.Request) {

}

func (app *application) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vars_id := vars["id"]

	id, err := strconv.Atoi(vars_id)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	book, err := app.models.Books.Get(id)

	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, book)
}

func (app *application) CreateBook(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title         string `json:"title"`
		Author        string `json:"author"`
		PublishedYear int    `json:"publishedYear"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		app.models.Books.ErrorLog.Println(err)
		return
	}

	book := &models.Book{
		Title:         input.Title,
		Author:        input.Author,
		PublishedYear: input.PublishedYear,
	}

	err = app.models.Books.Insert(book)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, book)
}

func (app *application) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	err = app.models.Books.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	book, err := app.models.Books.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Title         *string `json:"title"`
		Author        *string `json:"author"`
		PublishedYear *int    `json:"publishedYear"`
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		app.models.Books.ErrorLog.Println(err)
		return
	}

	if input.Title != nil {
		book.Title = *input.Title
	}

	if input.Author != nil {
		book.Author = *input.Author
	}

	if input.PublishedYear != nil {
		book.PublishedYear = *input.PublishedYear
	}

	err = app.models.Books.Update(book)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, book)

}
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
