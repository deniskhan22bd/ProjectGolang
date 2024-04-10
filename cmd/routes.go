package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Update the routes() method to return a http.Handler instead of a *httprouter.Router.
func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/books", app.requirePermission("books:read", app.GetBooks)).Methods("GET")
	r.HandleFunc("/books", app.requirePermission("books:write", app.CreateBook)).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", app.requirePermission("books:read", app.GetBook)).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", app.requirePermission("books:write", app.UpdateBook)).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", app.requirePermission("books:delete", app.DeleteBook)).Methods("DELETE")

	r.HandleFunc("/users", app.registerUserHandler).Methods("POST")
	r.HandleFunc("/users/activate", app.activateUserHandler).Methods("POST")
	r.HandleFunc("/users/login", app.createAuthenticationTokenHandler).Methods("POST")
	http.Handle("/", r)

	// Wrap the router with the panic recovery middleware.
	return app.recoverPanic(app.authenticate((r)))

}
