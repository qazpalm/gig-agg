package main

import (
	"fmt"
	"github.com/qazpalm/gig-agg/internal/store/sqlite"
)

func main() {
	db, err := sqlite.NewDB("gig-agg.db")
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return
	}
	defer db.Close()
	sqlite.InitialiseSchema(db)
	if err != nil {
		fmt.Printf("Error initialising schema: %v\n", err)
		return
	}

	// Create stores
	userStore 	:= sqlite.NewUserStore(db)
	gigStore 	:= sqlite.NewGigStore(db)
	venueStore 	:= sqlite.NewVenueStore(db)
	artistStore := sqlite.NewArtistStore(db)
	genreStore 	:= sqlite.NewGenreStore(db)

	_ = userStore
	_ = gigStore
	_ = venueStore
	_ = artistStore
	_ = genreStore
}