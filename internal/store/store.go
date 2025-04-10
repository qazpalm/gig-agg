package store

import (
	"github.com/qazpalm/gig-agg/internal/models"
)

type UserStore interface {
	CreateUser(user *models.User) (int64, error)
	GetUser(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByRememberedToken(token string) (*models.User, error)
	GetUsers(count int, offset int) ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error

	GetAllUsers() ([]*models.User, error)
}

type GigStore interface {
	CreateGig(gig *models.Gig) (int64, error)
	GetGig(id int) (*models.Gig, error)
	GetGigs(count int, offset int) ([]*models.Gig, error)
	UpdateGig(gig *models.Gig) error
	DeleteGig(id int) error

	GetAllGigs() ([]*models.Gig, error)
}

type VenueStore interface {
	CreateVenue(venue *models.Venue) (int64, error)
	GetVenue(id int) (*models.Venue, error)
	GetVenues(count int, offset int) ([]*models.Venue, error)
	UpdateVenue(venue *models.Venue) error
	DeleteVenue(id int) error

	GetAllVenues() ([]*models.Venue, error)
}

type ArtistStore interface {
	CreateArtist(artist *models.Artist) (int64, error)
	GetArtist(id int) (*models.Artist, error)
	GetArtists(count int, offset int) ([]*models.Artist, error)
	UpdateArtist(artist *models.Artist) error
	DeleteArtist(id int) error

	GetAllArtists() ([]*models.Artist, error)
}

type GenreStore interface {
	CreateGenre(genre *models.Genre) (int64, error)
	GetGenre(id int) (*models.Genre, error)
	GetGenres(count int, offset int) ([]*models.Genre, error)
	UpdateGenre(genre *models.Genre) error
	DeleteGenre(id int) error

	GetAllGenres() ([]*models.Genre, error)
}
