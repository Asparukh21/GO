package data

import (
	"time"
)

type Antique struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty"`
	Color     string    `json:"color"`
	Version   int32     `json:"version"`
}
