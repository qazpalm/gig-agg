package main

import (
	"fmt"
	"github.com/qazpalm/gig-agg/pkg/apiclient"
)

func main() {
	// Initialize the API client with the localhost url and API key
	client := apiclient.NewClient("http://127.0.0.1:8080/api", "api_key_1")

	err := client.CreateArtist(apiclient.Artist{
		Name:        "New Artist",
		Description: "This is a new artist",
		SpotifyID:   "spotify_id",
		GenreIDs:    []int{1},
	})

	if err != nil {
		fmt.Printf("Error creating artist: %v\n", err)
		return
	}

	// Fetch the created artist
	artist, err := client.GetArtist(1)
	if err != nil {
		fmt.Printf("Error fetching artist: %v\n", err)
		return
	}

	fmt.Printf("Created artist: %+v\n", artist)
}