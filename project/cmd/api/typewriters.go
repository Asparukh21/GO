package main

import (
	"fmt"
	"net/http"
	"time"

	"project.alexedwards.net/internal/data"
)

func (app *application) createtypewriteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new typewriter")
}

func (app *application) showtypewriteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	antique := data.Antique{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Underwood Vintage TypeWriter",
		Year:      1966,
		Runtime:   102,
		Color:     "Black",
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"Antique TypeWriter": antique}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
