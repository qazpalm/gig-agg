package models

type Artist struct {
	ID   		int		`json:"id"`
	Name 		string	`json:"name"`
	Description string	`json:"description"`
	SpotifyID 	string	`json:"spotify_id"`
	GenreIDs 	[]int	`json:"genre_ids"`
}