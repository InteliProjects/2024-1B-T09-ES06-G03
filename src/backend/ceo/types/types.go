package types

import (
	"fmt"
	"time"
)

type Rating struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	Level     string    `json:"level"` // Enum: '1', '2', '3', '4'
	UserID    int       `json:"user_id"`
	ProjectID int       `json:"project_id"`
}

type RatingPayload struct {
	Date      CustomDate `json:"date" validate:"required"`
	Level     string     `json:"level" validate:"required,oneof=1 2 3 4"` // Ensures the level is one of the enum values
	UserID    int        `json:"user_id" validate:"required"`
	ProjectID int        `json:"project_id" validate:"required"`
}

type Notification struct {
	ID                int       `json:"id"`
	ReceivedUserID    int       `json:"received_user_id"`    // ID do usuário que recebe a notificação
	SentUserID        int       `json:"sent_user_id"`        // ID do usuário que enviou a notificação
	ReceivedProjectID int       `json:"received_project_id"` // ID do projeto relacionado ao usuário que recebeu a notificação
	SentProjectID     int       `json:"sent_project_id"`     // ID do projeto relacionado ao usuário que enviou a notificação
	SynergyType       string    `json:"synergy_type"`        // Tipo de sinergia (Aprendizagem, Integração, Unificação)
	Type              string    `json:"type"`                // Tipo de notificação (Solicitação, Atualização, etc.)
	Title             string    `json:"title"`
	Message           string    `json:"message"`
	Status            bool      `json:"status"` // Status da leitura da notificação
	CreatedAt         time.Time `json:"created_at"`
}
type Project struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	UserID        int       `json:"user_id"`
	SubcategoryID int       `json:"subcategory_id"`
	CategoryID    int       `json:"category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Photo         string    `json:"photo"`
    Local         string    `json:"local"` 
}

type DetailedNotification struct {
	Notification
	Project         Project `json:"project"`
	SubcategoryName string  `json:"subcategory_name"`
	CategoryName    string  `json:"category_name"`
}

type NotificationPayload struct {
	ReceivedUserID   int        `json:"received_user_id" validate:"required"`
	SentUserID       int        `json:"sent_user_id" validate:"required"`
	ReceivedProjectID int      `json:"received_project_id"`
	SentProjectID    int        `json:"sent_project_id"`
	SynergyType      string    `json:"synergy_type" validate:"required"`
	Type             string     `json:"type" validate:"required,oneof='Solicitação' 'Atualização' 'Novo Projeto' 'Outro'"`
	Title            string     `json:"title" validate:"required"`
	Message          string     `json:"message" validate:"required"`
	Status           bool       `json:"status"`
	CreatedAt        CustomDate `json:"created_at"`
}

type NotificationStore interface {
	GetNotifications() ([]Notification, error)
	GetNotificationByID(id int) (*Notification, error)
	CreateNotification(notification *Notification) error
	UpdateNotification(id int, notification *Notification) error
	DeleteNotification(id int) error
	GetNotificationByTitle(title string) (*Notification, error)
	GetNotificationsByUser(userId int) ([]DetailedNotification, error)
}

type RatingStore interface {
	GetRatings() ([]Rating, error)
	GetRatingByID(id int) (*Rating, error)
	CreateRating(rating *Rating) error
	UpdateRating(id int, rating *Rating) error
	DeleteRating(id int) error
	GetRatingsByUserID(userID int) ([]Rating, error)
}

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	// Trimming quotes
	str := string(b)
	str = str[1 : len(str)-1]

	// Parsing the date
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	cd.Time = t
	return nil
}
