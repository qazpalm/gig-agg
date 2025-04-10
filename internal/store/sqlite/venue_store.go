package sqlite

import (
	"database/sql"
	
	"github.com/qazpalm/gig-agg/internal/models"
	"github.com/qazpalm/gig-agg/internal/store"
)

type sqliteVenueStore struct {
	db *sql.DB
}

func NewVenueStore(db *sql.DB) store.VenueStore {
	return &sqliteVenueStore{db}
}

func (s *sqliteVenueStore) CreateVenue(venue *models.Venue) (int64, error) {
	query := `INSERT INTO venues (name, address, city, longitude, latitude) VALUES (?, ?, ?, ?, ?)`
	result, err := s.db.Exec(query, venue.Name, venue.Address, venue.City, venue.Longitude, venue.Latitude)
	if err != nil {
		return 0, err
	}

	venueId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return venueId, nil
}

func (s *sqliteVenueStore) GetVenue(id int) (*models.Venue, error) {
	query := `SELECT id, name, address, city, longitude, latitude FROM venues WHERE id = ?`
	row := s.db.QueryRow(query, id)

	venue := &models.Venue{}
	err := row.Scan(&venue.ID, &venue.Name, &venue.Address, &venue.City, &venue.Longitude, &venue.Latitude)
	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (s *sqliteVenueStore) GetVenues(count int, offset int) ([]*models.Venue, error) {
	query := `SELECT id, name, address, city, longitude, latitude FROM venues LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var venues []*models.Venue
	for rows.Next() {
		venue := &models.Venue{}
		err := rows.Scan(&venue.ID, &venue.Name, &venue.Address, &venue.City, &venue.Longitude, &venue.Latitude)
		if err != nil {
			return nil, err
		}
		venues = append(venues, venue)
	}

	return venues, nil
}

func (s *sqliteVenueStore) UpdateVenue(venue *models.Venue) error {
	query := `UPDATE venues SET name = ?, address = ?, city = ?, longitude = ?, latitude = ? WHERE id = ?`
	_, err := s.db.Exec(query, venue.Name, venue.Address, venue.City, venue.Longitude, venue.Latitude, venue.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteVenueStore) DeleteVenue(id int) error {
	query := `DELETE FROM venues WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteVenueStore) GetAllVenues() ([]*models.Venue, error) {
	query := `SELECT id, name, address, city, longitude, latitude FROM venues`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var venues []*models.Venue
	for rows.Next() {
		venue := &models.Venue{}
		err := rows.Scan(&venue.ID, &venue.Name, &venue.Address, &venue.City, &venue.Longitude, &venue.Latitude)
		if err != nil {
			return nil, err
		}
		venues = append(venues, venue)
	}

	return venues, nil
}