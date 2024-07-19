package rating

import (
	"database/sql"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetRatings() ([]types.Rating, error) {
	rows, err := s.db.Query("SELECT id, date, level, user_id, project_id FROM ratings ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []types.Rating
	for rows.Next() {
		var rating types.Rating
		if err := rows.Scan(&rating.ID, &rating.Date, &rating.Level, &rating.UserID, &rating.ProjectID); err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

func (s *Store) GetRatingByID(id int) (*types.Rating, error) {
	var rating types.Rating
	err := s.db.QueryRow("SELECT id, date,level, user_id, project_id FROM ratings WHERE id=$1", id).Scan(
		&rating.ID, &rating.Date, &rating.Level, &rating.UserID, &rating.ProjectID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (s *Store) CreateRating(rating *types.Rating) error {
	_, err := s.db.Exec("INSERT INTO ratings (date, level, user_id, project_id) VALUES ($1, $2, $3, $4)", rating.Date, rating.Level, rating.UserID, rating.ProjectID)
	return err
}

func (s *Store) UpdateRating(id int, rating *types.Rating) error {
	result, err := s.db.Exec("UPDATE ratings SET date=$1, level=$2, user_id=$3, project_id=$4 WHERE id=$5", rating.Date, rating.Level, rating.UserID, rating.ProjectID, id)
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

func (s *Store) DeleteRating(id int) error {
	result, err := s.db.Exec("DELETE FROM ratings WHERE id=$1", id)
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

func (s *Store) GetRatingsByUserID(userID int) ([]types.Rating, error) {
    // Query para buscar todas as avaliações que correspondam ao userID fornecido
    query := `SELECT id, date, level, user_id, project_id FROM ratings WHERE user_id = $1 ORDER BY id ASC`

    rows, err := s.db.Query(query, userID) // Executa a consulta com userID como parâmetro
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ratings []types.Rating
    for rows.Next() {
        var rating types.Rating
        if err := rows.Scan(&rating.ID, &rating.Date, &rating.Level, &rating.UserID, &rating.ProjectID); err != nil {
            return nil, err
        }
        ratings = append(ratings, rating)
    }

    return ratings, nil
}

