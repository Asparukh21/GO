package data

import (
	"time"

	"project.alexedwards.net/internal/validator"
)

type Antique struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Color     string    `json:"color"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, antique *Antique) {
	v.Check(antique.Title != "", "title", "must be provided")
	v.Check(len(antique.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(antique.Year != 0, "year", "must be provided")
	v.Check(antique.Year >= 1888, "year", "must be greater than 1888")
	v.Check(antique.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(antique.Runtime != 0, "runtime", "must be provided")
	v.Check(antique.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(antique.Color != "", "color", "must be provided")
	v.Check(len(antique.Color) <= 500, "color", "must not be more than 500 bytes long")

}
