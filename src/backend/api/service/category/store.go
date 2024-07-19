package category

import (
	"database/sql"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCategories() ([]types.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	categories := make([]types.Category, 0)
	for rows.Next() {
		category, err := scanRowsIntoCategory(rows)
		if err != nil {
			return nil, err
		}

		categories = append(categories, *category)
	}

	return categories, nil
}

func (s *Store) GetCategoryByID(id int) (*types.Category, error) {
	category := &types.Category{}
	err := s.db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(
		&category.ID,
		&category.Name,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Store) CreateCategory(category *types.Category) error {
	_, err := s.db.Exec("INSERT INTO categories (name) VALUES ($1)", category.Name)
	return err
}

func (s *Store) UpdateCategory(id int, category *types.Category) error {
	result, err := s.db.Exec("UPDATE categories SET name = $1 WHERE id = $2", category.Name, category.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *Store) DeleteCategory(id int) error {
	_, err := s.db.Exec("DELETE FROM categories WHERE id = $1", id)
	return err
}

func scanRowsIntoCategory(rows *sql.Rows) (*types.Category, error) {
	category := new(types.Category)

	err := rows.Scan(
		&category.ID,
		&category.Name,
	)
	if err != nil {
		return nil, err
	}

	return category, nil
}
