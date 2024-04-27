package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Update the routes() method to return a http.Handler instead of a *httprouter.Router.
func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	//Books handlers
	r.HandleFunc("/books", app.requirePermission("books:read", app.GetBooks)).Methods("GET")
	r.HandleFunc("/books", app.requirePermission("books:write", app.CreateBook)).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", app.requirePermission("books:read", app.GetBook)).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", app.requirePermission("books:write", app.UpdateBook)).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", app.requirePermission("books:delete", app.DeleteBook)).Methods("DELETE")
	r.HandleFunc("/books/{id:[0-9]+}/subscribe", app.requirePermission("books:read", app.SubcribeAtBook)).Methods("POST")

	//Users handlers
	r.HandleFunc("/users", app.registerUserHandler).Methods("POST")
	r.HandleFunc("/users/activate", app.activateUserHandler).Methods("PUT")
	r.HandleFunc("/users/login", app.createAuthenticationTokenHandler).Methods("POST")
	r.HandleFunc("/users/favorite", app.requirePermission("books:read", app.GetFavoriteBooks)).Methods("GET")
	http.Handle("/", r)

	// Wrap the router with the panic recovery middleware.
	return app.recoverPanic(app.authenticate((r)))

}
