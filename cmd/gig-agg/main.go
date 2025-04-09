package main

import (
	"fmt"
	"net/http"

	"github.com/qazpalm/gig-agg/internal/store/sqlite"
	"github.com/qazpalm/gig-agg/internal/routes"
	"github.com/qazpalm/gig-agg/internal/middleware"
	"github.com/qazpalm/gig-agg/internal/session"
	"github.com/qazpalm/gig-agg/internal/apikeys"
	"github.com/qazpalm/gig-agg/internal/auth"
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

	// Create user auth manager
	userAuthManager := auth.NewUserAuthManager(userStore, sessionStore)

	// Create API key manager
	apiKeyManager := apikeys.NewAPIKeyManager()
	_ = apiKeyManager

	_ = userStore
	_ = gigStore

	// Create a new mux router
	mux := http.NewServeMux()

	// Serve static files
	staticFileServer := http.FileServer(http.Dir("./assets/static"))
    mux.Handle("/static/", http.StripPrefix("/static", staticFileServer))

	// Register middleware
	sessionMiddleware := middleware.NewSessionMiddleware(sessionStore, userStore)

	// Register grouped routes
	routes.RegisterHomeRoutes(mux, sessionMiddleware, userStore, sessionStore, userAuthManager)
	//routes.RegisterAdminRoutes(mux)
	routes.RegisterAPIRoutes(mux, artistStore, genreStore, venueStore, gigStore, apiKeyManager)
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