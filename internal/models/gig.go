package models

import "time"

type Gig struct {
	ID          int				`json:"id"`
	Name        string			`json:"name"`
	Description string			`json:"description"`
	Artists	 	[]Artist 		`json:"artists"`
	VenueID     int 			`json:"venue_id"`
	GenreIDs   	[]int 			`json:"genre_ids"`
	DateTime    time.Time 		`json:"date_time"`
	TicketURL   string 			`json:"ticket_url"`

	CreatedAt   time.Time 		`json:"created_at"`
}