package main

import (
	"fmt"
	"net/http"

	"github.com/qazpalm/gig-agg/internal/store/sqlite"
	"github.com/qazpalm/gig-agg/internal/routes"
	"github.com/qazpalm/gig-agg/internal/middleware"
	"github.com/qazpalm/gig-agg/internal/session"
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

	// Create session store
	sessionStore := session.NewSessionStore()

	_ = userStore
	_ = gigStore
	_ = venueStore
	_ = artistStore
	_ = genreStore

	// Create a new mux router
	mux := http.NewServeMux()

	// Register middleware
	sessionMiddleware := middleware.NewSessionMiddleware(sessionStore, userStore)

	// Register grouped routes
	routes.RegisterHomeRoutes(mux, sessionMiddleware)
	//routes.RegisterAdminRoutes(mux)
	//routes.RegisterAPIRoutes(mux)
	//routes.RegisterAuthRoutes(mux)

	// Start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	fmt.Println("Server stopped")
}