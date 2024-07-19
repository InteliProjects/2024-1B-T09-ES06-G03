package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Project struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Status        string     `json:"status"`
	UserID        int        `json:"user_id"`
	CategoryID    int        `json:"category_id"`
	SubcategoryID int        `json:"subcategory_id"`
	CreatedAt     CustomDate `json:"created_at"`
	UpdatedAt     CustomDate `json:"updated_at"`
	Photo         string     `json:"photo"`
	Local         string     `json:"local"`
}

type ProjectPayload struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	UserID        int    `json:"user_id"`
	SubcategoryID int    `json:"subcategory_id"`
	CategoryID    int    `json:"category_id"`
	Photo         string `json:"photo"`
	Local         string `json:"local"`
}

type ProjectDetails struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Status          string     `json:"status"`
	UserID          int        `json:"user_id"`
	CategoryID      int        `json:"category_id"`
	SubcategoryID   int        `json:"subcategory_id"`
	CreatedAt       CustomDate `json:"created_at"`
	UpdatedAt       CustomDate `json:"updated_at"`
	Photo           string     `json:"photo"`
	SynergyCount    int        `json:"synergy_count"`
	InterestedCount int        `json:"interested_count"`
	CeoName         string     `json:"ceo_name"`
	CategoryName    string     `json:"category_name"`
	SubcategoryName string     `json:"subcategory_name"`
	CeoPhoto        string     `json:"ceo_photo"`
	CompanyName     string     `json:"company_name"`
	Local           string     `json:"local"`
}

type Synergy struct {
	ID              int    `json:"id"`
	SourceProjectID int    `json:"source_project_id"`
	TargetProjectID int    `json:"target_project_id"`
	Status          string `json:"status"`
	Type            string `json:"type"`
	Description     string `json:"description"`
}

type SynergyPayload struct {
	SourceProjectID int    `json:"source_project_id"`
	TargetProjectID int    `json:"target_project_id"`
	Status          string `json:"status"`
	Type            string `json:"type"`
	Description     string `json:"description"`
}

type DetailedSynergy struct {
	ID            int            `json:"id"`
	SourceProject ProjectDetails `json:"source_project"`
	TargetProject ProjectDetails `json:"target_project"`
	Status        string         `json:"status"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
}

type InterestedUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SynergyDetails struct {
	ID            int            `json:"id"`
	SourceProject ProjectDetails `json:"source_project"`
	TargetProject ProjectDetails `json:"target_project"`
	Status        string         `json:"status"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
}

type InterestedAndSynergiesResponse struct {
	InterestedUsers []InterestedUser `json:"interested_users"`
	Synergies       []SynergyDetails `json:"synergies"`
}

type Update struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
	CreatedAt   CustomDate `json:"created_at"`
	SynergyID   int        `json:"synergy_id"`
}

type UpdatePayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	SynergyID   int    `json:"synergy_id"`
}

type CustomDate struct {
	time.Time
}

type ProjectStore interface {
	GetProjects() ([]ProjectDetails, error)
	GetProjectByID(id int) (*Project, error)
	CreateProject(project *Project) error
	UpdateProject(id int, project *Project) error
	DeleteProject(id int) error
	GetProjectsByName(name string) ([]Project, error)
	GetProjectsByCeoID(ceoId int) ([]Project, error)
	GetEvaluations() ([]Evaluation, error)
    GetInterestedAndSynergiesByProjectID(projectID int) (*InterestedAndSynergiesResponse, error)

}

type SynergyStore interface {
	GetSynergies() ([]DetailedSynergy, error)
	GetSynergyByID(id int) (*DetailedSynergy, error)
	CreateSynergy(synergy *Synergy) error
	UpdateSynergy(id int, synergy *Synergy) error
	DeleteSynergy(id int) error
	GetSynergiesByDescription(description string) ([]Synergy, error)
}

type UpdateStore interface {
	GetUpdates() ([]Update, error)
	GetUpdateByID(id int) (*Update, error)
	CreateUpdate(update *Update) error
	UpdateUpdate(id int, update *Update) error
	DeleteUpdate(id int) error
	GetUpdateByTitle(title string) ([]Update, error)
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", cd.Time.Format(time.RFC3339))), nil
}

func (cd CustomDate) Value() (driver.Value, error) {
	return cd.Time, nil
}

func (cd *CustomDate) Scan(value interface{}) error {
	if value == nil {
		*cd = CustomDate{Time: time.Time{}}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*cd = CustomDate{v}
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type CustomDate", value)
	}
}

// Estrutura Rating em types.go
type Rating struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	Level     int       `json:"level"`
	UserID    int       `json:"user_id"`
	ProjectID int       `json:"project_id"`
}

type Evaluation struct {
	IDProponente int `json:"id_proponente"`
	IDProjeto    int `json:"id_projeto"`
	Avaliacao    int `json:"avaliacao"`
}

type PredictRequest struct {
	UserID int          `json:"user_id"`
	Data   []Evaluation `json:"data"`
}

type PredictResponse struct {
	ID int `json:"id"`
}
