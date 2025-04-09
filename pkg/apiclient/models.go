package apiclient

type Artist struct {
	ID   		int		`json:"id"`
	Name 		string	`json:"name"`
	Description string	`json:"description"`
	SpotifyID 	string	`json:"spotify_id"`
	GenreIDs 	[]int	`json:"genre_ids"`
}

type Genre struct {
	ID   int		`json:"id"`
	Name string		`json:"name"`
}

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

type Venue struct {
	ID      int			`json:"id"`
	Name    string		`json:"name"`
	Address string		`json:"address"`
	City    string		`json:"city"`

	Longitude 	float64 `json:"longitude"`
	Latitude 	float64 `json:"latitude"`
}

type User struct {
	ID       		int  		`json:"id"`
	Username 		string 		`json:"username"`
	Email    		string	 	`json:"email"`
	PasswordHash 	string 		`json:"password_hash"`
	LastLogin 		time.Time 	`json:"last_login"`
	CreatedAt 		time.Time 	`json:"created_at"`

	RememberedToken string 		`json:"remembered_token"`

	IsAdmin 		bool 		`json:"is_admin"`
}

