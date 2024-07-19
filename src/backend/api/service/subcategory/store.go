package subcategory

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

func (s *Store) GetSubcategories() ([]types.Subcategory, error) {
	rows, err := s.db.Query("SELECT * FROM subcategories ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subcategories := make([]types.Subcategory, 0)
	for rows.Next() {
		subcategory := &types.Subcategory{}
		err := rows.Scan(&subcategory.ID, &subcategory.Name, &subcategory.CategoryID)
		if err != nil {
			return nil, err
		}

		subcategories = append(subcategories, *subcategory)
	}

	return subcategories, nil
}

func (s *Store) GetSubcategoryByID(id int) (*types.Subcategory, error) {
	subcategory := &types.Subcategory{}
	err := s.db.QueryRow("SELECT * FROM subcategories WHERE id = $1", id).Scan(
		&subcategory.ID,
		&subcategory.Name,
		&subcategory.CategoryID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return subcategory, nil
}

func (s *Store) GetSubcategoriesByCategory(categoryID int) ([]types.Subcategory, error) {
	rows, err := s.db.Query("SELECT id, name, category_id FROM subcategories WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subcategories []types.Subcategory
	for rows.Next() {
		var subcategory types.Subcategory
		if err := rows.Scan(&subcategory.ID, &subcategory.Name, &subcategory.CategoryID); err != nil {
			return nil, err
		}
		subcategories = append(subcategories, subcategory)
	}

	return subcategories, nil
}

func (s *Store) CreateSubcategory(subcategory *types.Subcategory) error {
	_, err := s.db.Exec("INSERT INTO subcategories (name, category_id) VALUES ($1, $2)", subcategory.Name, subcategory.CategoryID)
	return err
}

func (s *Store) UpdateSubCategory(id int, subcategory *types.Subcategory) error {
	_, err := s.db.Exec("UPDATE subcategories SET name = $1 WHERE id = $2", subcategory.Name, subcategory.ID)
	return err
}

func (s *Store) DeleteSubcategory(id int) error {
	_, err := s.db.Exec("DELETE FROM subcategories WHERE id = $1", id)
	return err
}
