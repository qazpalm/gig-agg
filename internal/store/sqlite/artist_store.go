package sqlite

import (
	"database/sql"
	
	"github.com/qazpalm/gig-agg/internal/models"
	"github.com/qazpalm/gig-agg/internal/store"
)

type sqliteArtistStore struct {
	db *sql.DB
}

func NewArtistStore(db *sql.DB) store.ArtistStore {
	return &sqliteArtistStore{db}
}

func (s *sqliteArtistStore) CreateArtist(artist *models.Artist) (int64, error) {
	artistQuery := `INSERT INTO artists (name, description, spotify_id) VALUES (?, ?, ?)`
	result, err := s.db.Exec(artistQuery, artist.Name, artist.Description, artist.SpotifyID)
	if err != nil {
		return 0, err
	}

	artistID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Insert genres into artist_genres table
	for _, genreID := range artist.GenreIDs {
		genreQuery := `INSERT INTO artist_genres (artist_id, genre_id) VALUES (?, ?)`
		_, err = s.db.Exec(genreQuery, artistID, genreID)
		if err != nil {
			return 0, err
		}
	}

	return artistID, nil
}

func (s *sqliteArtistStore) GetArtist(id int) (*models.Artist, error) {
	query := `SELECT id, name, description, spotify_id FROM artists WHERE id = ?`
	row := s.db.QueryRow(query, id)

	artist := &models.Artist{}
	err := row.Scan(&artist.ID, &artist.Name, &artist.Description, &artist.SpotifyID)
	if err != nil {
		return nil, err
	}

	// Get genres for artist
	genreQuery := `SELECT genre_id FROM artist_genres WHERE artist_id = ?`
	rows, err := s.db.Query(genreQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var genreID int
		err := rows.Scan(&genreID)
		if err != nil {
			return nil, err
		}

		artist.GenreIDs = append(artist.GenreIDs, genreID)
	}

	return artist, nil
}

func (s *sqliteArtistStore) GetArtists(count int, offset int) ([]*models.Artist, error) {
	query := `SELECT id, name, description, spotify_id FROM artists LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []*models.Artist
	for rows.Next() {
		artist := &models.Artist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.Description, &artist.SpotifyID)
		if err != nil {
			return nil, err
		}

		// Get genres for artist
		genreQuery := `SELECT genre_id FROM artist_genres WHERE artist_id = ?`
		genreRows, err := s.db.Query(genreQuery, artist.ID)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()

		for genreRows.Next() {
			var genreID int
			err := genreRows.Scan(&genreID)
			if err != nil {
				return nil, err
			}

			artist.GenreIDs = append(artist.GenreIDs, genreID)
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

func (s *sqliteArtistStore) UpdateArtist(artist *models.Artist) error {
	query := `UPDATE artists SET name = ?, description = ?, spotify_id = ? WHERE id = ?`
	_, err := s.db.Exec(query, artist.Name, artist.Description, artist.SpotifyID, artist.ID)
	if err != nil {
		return err
	}

	// Delete existing genres for artist
	deleteQuery := `DELETE FROM artist_genres WHERE artist_id = ?`
	_, err = s.db.Exec(deleteQuery, artist.ID)
	if err != nil {
		return err
	}

	// Insert genres into artist_genres table
	for _, genreID := range artist.GenreIDs {
		genreQuery := `INSERT INTO artist_genres (artist_id, genre_id) VALUES (?, ?)`
		_, err = s.db.Exec(genreQuery, artist.ID, genreID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *sqliteArtistStore) DeleteArtist(id int) error {
	query := `DELETE FROM artists WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteArtistStore) SearchArtists(query string, count int, offset int) ([]*models.Artist, error) {
	searchQuery := `SELECT id, name, description, spotify_id FROM artists WHERE name LIKE ? LIMIT ? OFFSET ?`
	rows, err := s.db.Query(searchQuery, "%"+query+"%", count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []*models.Artist
	for rows.Next() {
		artist := &models.Artist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.Description, &artist.SpotifyID)
		if err != nil {
			return nil, err
		}

		// Get genres for artist
		genreQuery := `SELECT genre_id FROM artist_genres WHERE artist_id = ?`
		genreRows, err := s.db.Query(genreQuery, artist.ID)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()

		for genreRows.Next() {
			var genreID int
			err := genreRows.Scan(&genreID)
			if err != nil {
				return nil, err
			}

			artist.GenreIDs = append(artist.GenreIDs, genreID)
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

func (s *sqliteArtistStore) GetArtistByGenres(genres []models.Genre, count int, offset int) ([]*models.Artist, error) {
	genreIDs := make([]int, len(genres))
	for i, genre := range genres {
		genreIDs[i] = genre.ID
	}

	query := `SELECT DISTINCT a.id, a.name, a.description, a.spotify_id
			  FROM artists a
			  JOIN artist_genres ag ON a.id = ag.artist_id
			  WHERE ag.genre_id IN (?)
			  LIMIT ? OFFSET ?`

	rows, err := s.db.Query(query, genreIDs, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []*models.Artist
	for rows.Next() {
		artist := &models.Artist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.Description, &artist.SpotifyID)
		if err != nil {
			return nil, err
		}

		// Get genres for artist
		genreQuery := `SELECT genre_id FROM artist_genres WHERE artist_id = ?`
		genreRows, err := s.db.Query(genreQuery, artist.ID)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()

		for genreRows.Next() {
			var genreID int
			err := genreRows.Scan(&genreID)
			if err != nil {
				return nil, err
			}

			artist.GenreIDs = append(artist.GenreIDs, genreID)
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

func (s *sqliteArtistStore) GetAllArtists() ([]*models.Artist, error) {
	query := `SELECT id, name, description, spotify_id FROM artists`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []*models.Artist
	for rows.Next() {
		artist := &models.Artist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.Description, &artist.SpotifyID)
		if err != nil {
			return nil, err
		}

		// Get genres for artist
		genreQuery := `SELECT genre_id FROM artist_genres WHERE artist_id = ?`
		genreRows, err := s.db.Query(genreQuery, artist.ID)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()

		for genreRows.Next() {
			var genreID int
			err := genreRows.Scan(&genreID)
			if err != nil {
				return nil, err
			}

			artist.GenreIDs = append(artist.GenreIDs, genreID)
		}

		artists = append(artists, artist)
	}

	return artists, nil
}
