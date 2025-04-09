package models

type Venue struct {
	ID      int			`json:"id"`
	Name    string		`json:"name"`
	Address string		`json:"address"`
	City    string		`json:"city"`

	Longitude 	float64 `json:"longitude"`
	Latitude 	float64 `json:"latitude"`
}