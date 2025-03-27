package sqlite

import (
	"database/sql"
	
	"github.com/qazpalm/gig-agg/internal/models"
	"github.com/qazpalm/gig-agg/internal/store"
)

type sqliteGigStore struct {
	db *sql.DB
}

func NewGigStore(db *sql.DB) store.GigStore {
	return &sqliteGigStore{db}
}

func (s *sqliteGigStore) CreateGig(gig *models.Gig) (int64, error) {
	query := `INSERT INTO gigs (name, description, venue_id date_time, ticket_url) VALUES (?, ?, ?, ?, ?)`
	result, err := s.db.Exec(query, gig.Name, gig.Description, gig.VenueID, gig.DateTime, gig.TicketURL)
	if err != nil {
		return 0, err
	}

	gigId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Insert artists into gig_artists table
	for _, artist := range gig.Artists {
		query = `INSERT INTO gig_artists (gig_id, artist_id) VALUES (?, ?)`
		_, err = s.db.Exec(query, gigId, artist.ID)
		if err != nil {
			return 0, err
		}
	}

	// Insert genres into gig_genres table
	for _, genreID := range gig.GenreIDs {
		query = `INSERT INTO gig_genres (gig_id, genre_id) VALUES (?, ?)`
		_, err = s.db.Exec(query, gigId, genreID)
		if err != nil {
			return 0, err
		}
	}

	return gigId, nil
}

func (s *sqliteGigStore) GetGig(id int) (*models.Gig, error) {
	query := `SELECT id, name, description, venue_id, date_time, ticket_url, created_at FROM gigs WHERE id = ?`
	row := s.db.QueryRow(query, id)

	gig := &models.Gig{}
	err := row.Scan(&gig.ID, &gig.Name, &gig.Description, &gig.VenueID, &gig.DateTime, &gig.TicketURL, &gig.CreatedAt)
	if err != nil {
		return nil, err
	}

	// Get artists
	query = `SELECT artist_id FROM gig_artists WHERE gig_id = ?`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var artistID int
		err = rows.Scan(&artistID)
		if err != nil {
			return nil, err
		}

		artist, err := s.GetArtist(artistID)
		if err != nil {
			return nil, err
		}

		gig.Artists = append(gig.Artists, *artist)
	}

	// Get genres
	query = `SELECT genre_id FROM gig_genres WHERE gig_id = ?`
	rows, err = s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var genreID int
		err = rows.Scan(&genreID)
		if err != nil {
			return nil, err
		}

		genre, err := s.GetGenre(genreID)
		if err != nil {
			return nil, err
		}

		gig.GenreIDs = append(gig.GenreIDs, genre.ID)
	}

	return gig, nil
}

func (s *sqliteGigStore) GetGigs(count int, offset int) ([]*models.Gig, error) {
	query := `SELECT id, name, description, venue_id, date_time, ticket_url, created_at FROM gigs LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gigs := []*models.Gig{}
	for rows.Next() {
		gig := &models.Gig{}
		err = rows.Scan(&gig.ID, &gig.Name, &gig.Description, &gig.VenueID, &gig.DateTime, &gig.TicketURL, &gig.CreatedAt)
		if err != nil {
			return nil, err
		}

		// Get artists
		query = `SELECT artist_id FROM gig_artists WHERE gig_id = ?`
		artistRows, err := s.db.Query(query, gig.ID)
		if err != nil {
			return nil, err
		}

		for artistRows.Next() {
			var artistID int
			err = artistRows.Scan(&artistID)
			if err != nil {
				return nil, err
			}

			artist, err := s.GetArtist(artistID)
			if err != nil {
				return nil, err
			}

			gig.Artists = append(gig.Artists, *artist)
		}

		// Get genres
		query = `SELECT genre_id FROM gig_genres WHERE gig_id = ?`
		genreRows, err := s.db.Query(query, gig.ID)
		if err != nil {
			return nil, err
		}

		for genreRows.Next() {
			var genreID int
			err = genreRows.Scan(&genreID)
			if err != nil {
				return nil, err
			}

			genre, err := s.GetGenre(genreID)
			if err != nil {
				return nil, err
			}

			gig.GenreIDs = append(gig.GenreIDs, genre.ID)
		}

		gigs = append(gigs, gig)
	}

	return gigs, nil
}

func (s *sqliteGigStore) UpdateGig(gig *models.Gig) error {
	query := `UPDATE gigs SET name = ?, description = ?, venue_id = ?, date_time = ?, ticket_url = ? WHERE id = ?`
	_, err := s.db.Exec(query, gig.Name, gig.Description, gig.VenueID, gig.DateTime, gig.TicketURL, gig.ID)
	if err != nil {
		return err
	}

	// Delete existing artists
	query = `DELETE FROM gig_artists WHERE gig_id = ?`
	_, err = s.db.Exec(query, gig.ID)
	if err != nil {
		return err
	}

	// Insert artists into gig_artists table
	for _, artist := range gig.Artists {
		query = `INSERT INTO gig_artists (gig_id, artist_id) VALUES (?, ?)`
		_, err = s.db.Exec(query, gig.ID, artist.ID)
		if err != nil {
			return err
		}
	}

	// Delete existing genres
	query = `DELETE FROM gig_genres WHERE gig_id = ?`
	_, err = s.db.Exec(query, gig.ID)
	if err != nil {
		return err
	}

	// Insert genres into gig_genres table
	for _, genreID := range gig.GenreIDs {
		query = `INSERT INTO gig_genres (gig_id, genre_id) VALUES (?, ?)`
		_, err = s.db.Exec(query, gig.ID, genreID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *sqliteGigStore) DeleteGig(id int) error {
	query := `DELETE FROM gigs WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}