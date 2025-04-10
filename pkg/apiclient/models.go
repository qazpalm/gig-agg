package apiclient

import (
	"time"
	"fmt"
)

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

func (c *Client) CreateArtist(artist Artist) error {
	// Make the API call to create the artist
	err := c.doRequest("POST", "artist", artist, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetArtist(id int) (Artist, error) {
	artist := Artist{}
	err := c.doRequest("GET", fmt.Sprintf("artist/%d", id), nil, &artist)
	if err != nil {
		return Artist{}, err
	}
	return artist, nil
}

func (c *Client) UpdateArtist(artist Artist) error {
	// Make the API call to update the artist
	err := c.doRequest("PUT", fmt.Sprintf("artist/%d", artist.ID), artist, &artist)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteArtist(id int) error {
	// Make the API call to delete the artist
	err := c.doRequest("DELETE", fmt.Sprintf("artist/%d", id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateGenre(genre Genre) error {
	// Make the API call to create the genre
	err := c.doRequest("POST", "genre", genre, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetGenre(id int) (Genre, error) {
	genre := Genre{}
	err := c.doRequest("GET", fmt.Sprintf("genre/%d", id), nil, &genre)
	if err != nil {
		return Genre{}, err
	}
	return genre, nil
}

func (c *Client) UpdateGenre(genre Genre) error {
	// Make the API call to update the genre
	err := c.doRequest("PUT", fmt.Sprintf("genre/%d", genre.ID), genre, &genre)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteGenre(id int) error {
	// Make the API call to delete the genre
	err := c.doRequest("DELETE", fmt.Sprintf("genre/%d", id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateGig(gig Gig) error {
	// Make the API call to create the gig
	err := c.doRequest("POST", "gig", gig, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetGig(id int) (Gig, error) {
	gig := Gig{}
	err := c.doRequest("GET", fmt.Sprintf("gig/%d", id), nil, &gig)
	if err != nil {
		return Gig{}, err
	}
	return gig, nil
}

func (c *Client) UpdateGig(gig Gig) error {
	// Make the API call to update the gig
	err := c.doRequest("PUT", fmt.Sprintf("gig/%d", gig.ID), gig, &gig)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteGig(id int) error {
	// Make the API call to delete the gig
	err := c.doRequest("DELETE", fmt.Sprintf("gig/%d", id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateVenue(venue Venue) error {
	// Make the API call to create the venue
	err := c.doRequest("POST", "venue", venue, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetVenue(id int) (Venue, error) {
	venue := Venue{}
	err := c.doRequest("GET", fmt.Sprintf("venue/%d", id), nil, &venue)
	if err != nil {
		return Venue{}, err
	}
	return venue, nil
}

func (c *Client) UpdateVenue(venue Venue) error {
	// Make the API call to update the venue
	err := c.doRequest("PUT", fmt.Sprintf("venue/%d", venue.ID), venue, &venue)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteVenue(id int) error {
	// Make the API call to delete the venue
	err := c.doRequest("DELETE", fmt.Sprintf("venue/%d", id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateUser(user User) error {
	// Make the API call to create the user
	err := c.doRequest("POST", "user", user, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetUser(id int) (User, error) {
	user := User{}
	err := c.doRequest("GET", fmt.Sprintf("user/%d", id), nil, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (c *Client) UpdateUser(user User) error {
	// Make the API call to update the user
	err := c.doRequest("PUT", fmt.Sprintf("user/%d", user.ID), user, &user)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteUser(id int) error {
	// Make the API call to delete the user
	err := c.doRequest("DELETE", fmt.Sprintf("user/%d", id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

