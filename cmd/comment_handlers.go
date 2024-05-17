package main

import (
	"errors"
	"net/http"

	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/models"
	"github.com/deniskhan22bd/Golang/ProjectGolang/pkg/validator"
)

func (app *application) GetCommentsByBookID(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	comments, err := app.models.Comments.GetByBookID(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"comments": comments}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) GetComment(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	comment, err := app.models.Comments.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"comment": comment}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) CreateComment(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Content string `json:"content"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	comment := &models.Comment{
		Content: input.Content,
		BookID:  id,
	}

	err = app.models.Comments.Insert(comment)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"comment": comment}, nil)
}

func (app *application) DeleteComment(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	comment, err := app.models.Comments.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.models.Comments.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "success", "deleted_comment": comment}, nil)
}

func (app *application) UpdateComment(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	comment, err := app.models.Comments.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Content string `json:"content"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comment.Content = input.Content

	v := validator.New()
	models.ValidateComment(v, comment)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Comments.Update(comment)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"comment": comment}, nil)
}
