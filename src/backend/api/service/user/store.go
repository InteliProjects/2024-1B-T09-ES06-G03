package user

import (
	"database/sql"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetUserByEmail returns nil user and nil error if no user is found, distinguishing it from other errors.
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := &types.User{}
	err := s.db.QueryRow("SELECT id, name, email, password, company, instagram, linkedin, photo, description FROM users WHERE email = $1", email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Company, &user.Instagram, &user.Linkedin, &user.Photo, &user.Description,
	)
	if err == sql.ErrNoRows {
		return nil, nil // No user found
	}
	if err != nil {
		return nil, err // Actual error
	}
	return user, nil
}

// GetUserByID retrieves a user by ID or returns nil if not found.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	user := &types.User{}
	err := s.db.QueryRow("SELECT id, name, email, password, company, instagram, linkedin, photo, description FROM users WHERE id = $1", id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Company, &user.Instagram, &user.Linkedin, &user.Photo, &user.Description,
	)
	if err == sql.ErrNoRows {
		return nil, nil // No user found
	}
	if err != nil {
		return nil, err // Actual error
	}
	return user, nil
}

// CreateUser inserts a new user into the database.
func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (name, email, password, company, instagram, linkedin, photo, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		user.Name, user.Email, user.Password, user.Company, user.Instagram, user.Linkedin, &user.Photo, &user.Description,
	)
	return err
}
