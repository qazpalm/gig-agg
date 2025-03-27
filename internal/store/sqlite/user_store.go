package sqlite

import (
	"database/sql"
	
	"github.com/qazpalm/gig-agg/internal/models"
	"github.com/qazpalm/gig-agg/internal/store"
)

type sqliteUserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) store.UserStore {
	return &sqliteUserStore{db}
}

func (s *sqliteUserStore) CreateUser(user *models.User) (int64, error) {
	query := `INSERT INTO users (username, email, password_hash, is_admin, remembered_token) VALUES (?, ?, ?, ?, ?)`
	result, err := s.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.IsAdmin, user.RememberedToken)
	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *sqliteUserStore) GetUser(id int) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, last_login, created_at, is_admin, remembered_token FROM users WHERE id = ?`
	row := s.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.LastLogin, &user.CreatedAt, &user.IsAdmin, &user.RememberedToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *sqliteUserStore) GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, last_login, created_at, is_admin, remembered_token FROM users WHERE username = ?`
	row := s.db.QueryRow(query, username)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.LastLogin, &user.CreatedAt, &user.IsAdmin, &user.RememberedToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *sqliteUserStore) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, last_login, created_at, is_admin, remembered_token FROM users WHERE email = ?`
	row := s.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.LastLogin, &user.CreatedAt, &user.IsAdmin, &user.RememberedToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *sqliteUserStore) GetUserByRememberedToken(token string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, last_login, created_at, is_admin, remembered_token FROM users WHERE remembered_token = ?`
	row := s.db.QueryRow(query, token)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.LastLogin, &user.CreatedAt, &user.IsAdmin, &user.RememberedToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *sqliteUserStore) GetUsers(count int, offset int) ([]*models.User, error) {
	query := `SELECT id, username, email, password_hash, last_login, created_at, is_admin, remembered_token FROM users LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, count, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.LastLogin, &user.CreatedAt, &user.IsAdmin, &user.RememberedToken)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *sqliteUserStore) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = ?, email = ?, password_hash = ?, last_login = ?, is_admin = ?, remembered_token = ? WHERE id = ?`
	_, err := s.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.LastLogin, user.IsAdmin, user.RememberedToken, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteUserStore) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

