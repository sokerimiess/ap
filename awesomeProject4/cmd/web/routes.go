package main

import (
	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet)) // Moved down
	mux.Get("/students", http.HandlerFunc(app.students))
	mux.Get("/staff", http.HandlerFunc(app.staff))
	mux.Get("/applicants", http.HandlerFunc(app.applicants))
	mux.Get("/researches", http.HandlerFunc(app.researches))
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}
