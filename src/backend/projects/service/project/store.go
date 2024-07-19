package project

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

func (s *Store) GetProjects() ([]types.ProjectDetails, error) {
	query := `
		SELECT 
			p.id, 
			p.name, 
			p.description, 
			p.status, 
			p.user_id, 
			p.category_id, 
			p.subcategory_id, 
			p.created_at, 
			p.updated_at, 
			p.photo, 
			p.local, -- Adicionado aqui
			(SELECT COUNT(*) FROM synergies WHERE source_project_id = p.id OR target_project_id = p.id) AS synergy_count, 
			(SELECT COUNT(*) FROM ratings WHERE project_id = p.id AND level = '4') AS interested_count,
			u.name AS ceo_name,
			c.name AS category_name,
			s.name AS subcategory_name,
			u.photo AS ceo_photo,
			u.company AS company_name
		FROM 
			projects p
			JOIN users u ON p.user_id = u.id
			JOIN categories c ON p.category_id = c.id
			JOIN subcategories s ON p.subcategory_id = s.id
		ORDER BY 
			p.id ASC`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []types.ProjectDetails
	for rows.Next() {
		var project types.ProjectDetails
		if err := rows.Scan(
			&project.ID, &project.Name, &project.Description, &project.Status, &project.UserID,
			&project.CategoryID, &project.SubcategoryID, &project.CreatedAt, &project.UpdatedAt,
			&project.Photo, &project.Local, &project.SynergyCount, &project.InterestedCount, &project.CeoName,
			&project.CategoryName, &project.SubcategoryName, &project.CeoPhoto, &project.CompanyName); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (s *Store) GetProjectByID(id int) (*types.Project, error) {
	var project types.Project
	err := s.db.QueryRow("SELECT id, name, description, status, user_id, category_id, subcategory_id, created_at, updated_at, photo, local FROM projects WHERE id = $1", id).Scan(
		&project.ID, &project.Name, &project.Description, &project.Status, &project.UserID, &project.CategoryID, &project.SubcategoryID, &project.CreatedAt, &project.UpdatedAt, &project.Photo, &project.Local,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *Store) CreateProject(project *types.Project) error {
	_, err := s.db.Exec("INSERT INTO projects (name, description, status, user_id, category_id, subcategory_id, created_at, updated_at, photo, local) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", project.Name, project.Description, project.Status, project.UserID, project.CategoryID, project.SubcategoryID, project.CreatedAt, project.UpdatedAt, project.Photo, project.Local)
	return err
}

func (s *Store) UpdateProject(id int, project *types.Project) error {
	result, err := s.db.Exec("UPDATE projects SET name = $1, description = $2, status = $3, user_id = $4, category_id = $5, subcategory_id = $6, updated_at = $7, photo = $8, local = $9 WHERE id = $10", project.Name, project.Description, project.Status, project.UserID, project.CategoryID, project.SubcategoryID, project.UpdatedAt, project.Photo, project.Local, id)
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

func (s *Store) DeleteProject(id int) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = $1", id)
	return err
}

func (s *Store) GetProjectsByName(name string) ([]types.Project, error) {
	rows, err := s.db.Query("SELECT id, name, description, status, user_id, category_id, subcategory_id, created_at, updated_at, photo FROM projects WHERE name = $1 ORDER BY id ASC", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []types.Project
	for rows.Next() {
		var project types.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.Status, &project.UserID, &project.CategoryID, &project.SubcategoryID, &project.CreatedAt, &project.UpdatedAt, &project.Photo); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (s *Store) GetProjectsByCeoID(ceoId int) ([]types.Project, error) {
	rows, err := s.db.Query("SELECT id, name, description, status, user_id, category_id, subcategory_id, photo FROM projects WHERE user_id = $1 ORDER BY id ASC", ceoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []types.Project
	for rows.Next() {
		var project types.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.Status, &project.UserID, &project.CategoryID, &project.SubcategoryID, &project.Photo); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (s *Store) GetEvaluations() ([]types.Evaluation, error) {
	rows, err := s.db.Query("SELECT user_id, project_id, level FROM ratings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var evaluations []types.Evaluation
	for rows.Next() {
		var eval types.Evaluation
		if err := rows.Scan(&eval.IDProponente, &eval.IDProjeto, &eval.Avaliacao); err != nil {
			return nil, err
		}
		evaluations = append(evaluations, eval)
	}

	return evaluations, nil
}

func (s *Store) GetInterestedAndSynergiesByProjectID(projectID int) (*types.InterestedAndSynergiesResponse, error) {
	interestedQuery := `
        SELECT u.id, u.name, u.email
        FROM ratings r
        JOIN users u ON r.user_id = u.id
        WHERE r.project_id = $1 AND r.level = '4'`

	synergiesQuery := `
        SELECT 
            s.id, 
            sp.id, sp.name, sp.description, sp.status, sp.user_id, sp.category_id, sp.subcategory_id, sp.created_at, sp.updated_at, sp.photo, sp.local,
            tp.id, tp.name, tp.description, tp.status, tp.user_id, tp.category_id, tp.subcategory_id, tp.created_at, tp.updated_at, tp.photo, tp.local,
            s.status, s.type, s.description
        FROM synergies s
        LEFT JOIN projects sp ON s.source_project_id = sp.id
        LEFT JOIN projects tp ON s.target_project_id = tp.id
        WHERE s.source_project_id = $1 OR s.target_project_id = $1`

	// Fetch interested users
	rows, err := s.db.Query(interestedQuery, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interestedUsers []types.InterestedUser
	for rows.Next() {
		var user types.InterestedUser
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		interestedUsers = append(interestedUsers, user)
	}

	// Fetch synergies
	rows, err = s.db.Query(synergiesQuery, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var synergies []types.SynergyDetails
	for rows.Next() {
		var synergy types.SynergyDetails
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

	return &types.InterestedAndSynergiesResponse{
		InterestedUsers: interestedUsers,
		Synergies:       synergies,
	}, nil
}
