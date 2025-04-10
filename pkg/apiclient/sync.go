package apiclient

import (
	"time"
)

type DatabaseSync struct {
	Artists []Artist,
	Gigs    []Gig,
	Venues  []Venue,
	Genres  []Genre,
	Users   []User,

	LastUpdate time.Time,
}

func NewDatabaseSync() *DatabaseSync {
	return &DatabaseSync{
		Artists: make([]Artist, 0),
		Gigs:    make([]Gig, 0),
		Venues:  make([]Venue, 0),
		Genres:  make([]Genre, 0),
		Users:   make([]User, 0),
	}
}