package synergy

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

func (s *Store) GetSynergies() ([]types.DetailedSynergy, error) {
	query := `
		SELECT 
			s.id, 
			s.source_project_id, 
			sp.name, sp.description, sp.status, sp.user_id, sp.category_id, sp.subcategory_id, sp.created_at, sp.updated_at, sp.photo, sp.local, 
			s.target_project_id, 
			tp.name, tp.description, tp.status, tp.user_id, tp.category_id, tp.subcategory_id, tp.created_at, tp.updated_at, tp.photo, tp.local, 
			s.status, s.type, s.description
		FROM synergies s
		LEFT JOIN projects sp ON s.source_project_id = sp.id
		LEFT JOIN projects tp ON s.target_project_id = tp.id
		ORDER BY s.id ASC`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var synergies []types.DetailedSynergy
	for rows.Next() {
		var synergy types.DetailedSynergy
		if err := rows.Scan(
			&synergy.ID,
			&synergy.SourceProject.ID, &synergy.SourceProject.Name, &synergy.SourceProject.Description, &synergy.SourceProject.Status, &synergy.SourceProject.UserID, &synergy.SourceProject.CategoryID, &synergy.SourceProject.SubcategoryID, &synergy.SourceProject.CreatedAt, &synergy.SourceProject.UpdatedAt, &synergy.SourceProject.Photo, &synergy.SourceProject.Local,
			&synergy.TargetProject.ID, &synergy.TargetProject.Name, &synergy.TargetProject.Description, &synergy.TargetProject.Status, &synergy.TargetProject.UserID, &synergy.TargetProject.CategoryID, &synergy.TargetProject.SubcategoryID, &synergy.TargetProject.CreatedAt, &synergy.TargetProject.UpdatedAt, &synergy.TargetProject.Photo, &synergy.TargetProject.Local,
			&synergy.Status, &synergy.Type, &synergy.Description,
		); err != nil {
			return nil, err
		}
		synergies = append(synergies, synergy)
	}
	return synergies, nil
}

func (s *Store) GetSynergyByID(id int) (*types.DetailedSynergy, error) {
	query := `
		SELECT 
			s.id, 
			s.source_project_id, 
			sp.name, sp.description, sp.status, sp.user_id, sp.category_id, sp.subcategory_id, sp.created_at, sp.updated_at, sp.photo, sp.local, 
			s.target_project_id, 
			tp.name, tp.description, tp.status, tp.user_id, tp.category_id, tp.subcategory_id, tp.created_at, tp.updated_at, tp.photo, tp.local, 
			s.status, s.type, s.description
		FROM synergies s
		LEFT JOIN projects sp ON s.source_project_id = sp.id
		LEFT JOIN projects tp ON s.target_project_id = tp.id
		WHERE s.id = $1`

	var synergy types.DetailedSynergy
	err := s.db.QueryRow(query, id).Scan(
		&synergy.ID,
		&synergy.SourceProject.ID, &synergy.SourceProject.Name, &synergy.SourceProject.Description, &synergy.SourceProject.Status, &synergy.SourceProject.UserID, &synergy.SourceProject.CategoryID, &synergy.SourceProject.SubcategoryID, &synergy.SourceProject.CreatedAt, &synergy.SourceProject.UpdatedAt, &synergy.SourceProject.Photo, &synergy.SourceProject.Local,
		&synergy.TargetProject.ID, &synergy.TargetProject.Name, &synergy.TargetProject.Description, &synergy.TargetProject.Status, &synergy.TargetProject.UserID, &synergy.TargetProject.CategoryID, &synergy.TargetProject.SubcategoryID, &synergy.TargetProject.CreatedAt, &synergy.TargetProject.UpdatedAt, &synergy.TargetProject.Photo, &synergy.TargetProject.Local,
		&synergy.Status, &synergy.Type, &synergy.Description,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &synergy, nil
}

func (s *Store) CreateSynergy(synergy *types.Synergy) error {
	_, err := s.db.Exec("INSERT INTO synergies (source_project_id, target_project_id, status, type, description) VALUES ($1, $2, $3, $4, $5)",
		synergy.SourceProjectID, synergy.TargetProjectID, synergy.Status, synergy.Type, synergy.Description)
	return err
}

func (s *Store) UpdateSynergy(id int, synergy *types.Synergy) error {
	result, err := s.db.Exec("UPDATE synergies SET source_project_id = $1, target_project_id = $2, status = $3, type = $4, description = $5 WHERE id = $6",
		synergy.SourceProjectID, synergy.TargetProjectID, synergy.Status, synergy.Type, synergy.Description, id)
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

func (s *Store) DeleteSynergy(id int) error {
	_, err := s.db.Exec("DELETE FROM synergies WHERE id = $1", id)
	return err
}

func (s *Store) GetSynergiesByDescription(description string) ([]types.Synergy, error) {
	rows, err := s.db.Query("SELECT id, source_project_id, target_project_id, status, type, description FROM synergies WHERE description = $1", description)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var synergies []types.Synergy
	for rows.Next() {
		var synergy types.Synergy
		if err := rows.Scan(&synergy.ID, &synergy.SourceProjectID, &synergy.TargetProjectID, &synergy.Status, &synergy.Type, &synergy.Description); err != nil {
			return nil, err
		}
		synergies = append(synergies, synergy)
	}
	return synergies, nil
}
