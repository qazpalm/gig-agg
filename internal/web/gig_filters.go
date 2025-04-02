package web

import (
	"time"
)

type GigFilters struct {
	VenueID   *int64
	ArtistID  *int64
	GenreID   *int64
	FromDate  *time.Time
	ToDate    *time.Time
	Query	  string
}