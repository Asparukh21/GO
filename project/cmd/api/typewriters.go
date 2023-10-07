package main

import (
	"fmt"
	"net/http"
	"time"

	"project.alexedwards.net/internal/data"
	"project.alexedwards.net/internal/validator"
)

func (app *application) createtypewriteHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "create a new typewriter")

	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Color   string       `json:"string"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	antique := &data.Antique{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Color:   input.Color,
	}

	v := validator.New()

	if data.ValidateMovie(v, antique); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
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
