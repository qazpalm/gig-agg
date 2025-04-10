package sqlite

import (
	"database/sql"
	
	"github.com/qazpalm/gig-agg/internal/models"
	"github.com/qazpalm/gig-agg/internal/store"
)

type sqliteGenreStore struct {
	db *sql.DB
}

func NewGenreStore(db *sql.DB) store.GenreStore {
	return &sqliteGenreStore{db}
}

func (s *sqliteGenreStore) CreateGenre(genre *models.Genre) (int64, error) {
	query := `INSERT INTO genres (name) VALUES (?)`
	result, err := s.db.Exec(query, genre.Name)
	if err != nil {
		return 0, err
	}

	genreId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return genreId, nil
}

func (s *sqliteGenreStore) GetGenre(id int) (*models.Genre, error) {
	query := `SELECT id, name FROM genres WHERE id = ?`
	row := s.db.QueryRow(query, id)

	genre := &models.Genre{}
	err := row.Scan(&genre.ID, &genre.Name)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (s *sqliteGenreStore) GetGenres(count int, offset int) ([]*models.Genre, error) {
	query := `SELECT id, name FROM genres LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*models.Genre
	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (s *sqliteGenreStore) UpdateGenre(genre *models.Genre) error {
	query := `UPDATE genres SET name = ? WHERE id = ?`
	_, err := s.db.Exec(query, genre.Name, genre.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteGenreStore) DeleteGenre(id int) error {
	query := `DELETE FROM genres WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteGenreStore) GetAllGenres() ([]*models.Genre, error) {
	query := `SELECT id, name FROM genres`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*models.Genre
	for rows.Next() {
		genre := &models.Genre{}
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}