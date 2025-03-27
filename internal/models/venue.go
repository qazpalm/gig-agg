package models

type Venue struct {
	ID      int
	Name    string
	Address string
	City    string

	Longitude 	float64
	Latitude 	float64
}