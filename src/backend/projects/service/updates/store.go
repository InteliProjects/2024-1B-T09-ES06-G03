package update

import (
	"database/sql"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUpdates() ([]types.Update, error) {
	rows, err := s.db.Query("SELECT id, title, description, date, created_at, synergy_id FROM updates ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updates []types.Update
	for rows.Next() {
		var update types.Update
		if err := rows.Scan(&update.ID, &update.Title, &update.Description, &update.Date, &update.CreatedAt, &update.SynergyID); err != nil {
			return nil, err
		}
		updates = append(updates, update)
	}
	return updates, nil
}

func (s *Store) GetUpdateByID(id int) (*types.Update, error) {
	var update types.Update
	err := s.db.QueryRow("SELECT id, title, description, date, created_at, synergy_id FROM updates WHERE id = $1", id).Scan(
		&update.ID, &update.Title, &update.Description, &update.Date, &update.CreatedAt, &update.SynergyID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &update, nil
}

func (s *Store) CreateUpdate(update *types.Update) error {
	_, err := s.db.Exec("INSERT INTO updates (title, description, date, synergy_id) VALUES ($1, $2, $3, $4)",
		update.Title, update.Description, update.Date, update.SynergyID)
	return err
}

func (s *Store) UpdateUpdate(id int, update *types.Update) error {
	result, err := s.db.Exec("UPDATE updates SET title = $1, description = $2, date = $3, synergy_id = $4 WHERE id = $5",
		update.Title, update.Description, update.Date, update.SynergyID, id)
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

func (s *Store) DeleteUpdate(id int) error {
	_, err := s.db.Exec("DELETE FROM updates WHERE id = $1", id)
	return err
}

func (s *Store) GetUpdateByTitle(title string) ([]types.Update, error) {
	rows, err := s.db.Query("SELECT id, title, description, date, created_at, synergy_id FROM updates WHERE title = $1", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updates []types.Update
	for rows.Next() {
		var update types.Update
		if err := rows.Scan(&update.ID, &update.Title, &update.Description, &update.Date, &update.CreatedAt, &update.SynergyID); err != nil {
			return nil, err
		}
		updates = append(updates, update)
	}
	return updates, nil
}

