package main

import (
	"fmt"
	"github.com/qazpalm/gig-agg/internal/store"
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
	userStore := store.NewUserStore(db)
	gigStore := store.NewGigStore(db)
	venueStore := store.NewVenueStore(db)
	artistStore := store.NewArtistStore(db)
	genreStore := store.NewGenreStore(db)

	_ = userStore
	_ = gigStore
	_ = venueStore
	_ = artistStore
	_ = genreStore
}