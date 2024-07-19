package notifications

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

func (s *Store) GetNotifications() ([]types.Notification, error) {
    query := `
    SELECT 
        n.id, 
        n.received_user_id, 
        n.sent_user_id, 
        n.received_project_id, 
        n.sent_project_id, 
        n.synergy_type,
        n.type, 
        n.title, 
        n.message, 
        n.status, 
        n.created_at
    FROM 
        notifications n
    ORDER BY 
        n.id ASC`

    rows, err := s.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notifications []types.Notification
    for rows.Next() {
        var notification types.Notification
        if err := rows.Scan(
            &notification.ID,
            &notification.ReceivedUserID,
            &notification.SentUserID,
            &notification.ReceivedProjectID,
            &notification.SentProjectID,
            &notification.SynergyType,
            &notification.Type,
            &notification.Title,
            &notification.Message,
            &notification.Status,
            &notification.CreatedAt,
        ); err != nil {
            return nil, err
        }
        notifications = append(notifications, notification)
    }
    return notifications, nil
}

func (s *Store) GetNotificationByID(id int) (*types.Notification, error) {
	var notification types.Notification
	err := s.db.QueryRow("SELECT id, received_user_id, sent_user_id, received_project_id, sent_project_id, synergy_type, type, title, message, status, created_at FROM notifications WHERE id=$1", id).Scan(
		&notification.ID, &notification.ReceivedUserID, &notification.SentUserID, &notification.ReceivedProjectID, &notification.SentProjectID, &notification.SynergyType, &notification.Type, &notification.Title, &notification.Message, &notification.Status, &notification.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (s *Store) CreateNotification(notification *types.Notification) error {
	_, err := s.db.Exec("INSERT INTO notifications (received_user_id, sent_user_id, received_project_id, sent_project_id, synergy_type, type, title, message, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		notification.ReceivedUserID, notification.SentUserID, notification.ReceivedProjectID, notification.SentProjectID, notification.SynergyType, notification.Type, notification.Title, notification.Message, notification.Status, notification.CreatedAt)
	return err
}

func (s *Store) UpdateNotification(id int, notification *types.Notification) error {
	result, err := s.db.Exec("UPDATE notifications SET received_user_id=$1, sent_user_id=$2, received_project_id=$3, sent_project_id=$4, synergy_type=$5, type=$6, title=$7, message=$8, status=$9, created_at=$10 WHERE id=$11",
		notification.ReceivedUserID, notification.SentUserID, notification.ReceivedProjectID, notification.SentProjectID, notification.SynergyType, notification.Type, notification.Title, notification.Message, notification.Status, notification.CreatedAt, id)
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

func (s *Store) DeleteNotification(id int) error {
	result, err := s.db.Exec("DELETE FROM notifications WHERE id=$1", id)
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

func (s *Store) GetNotificationByTitle(title string) (*types.Notification, error) {
    query := `
    SELECT 
        id, received_user_id, sent_user_id, received_project_id, sent_project_id, synergy_type, type, title, message, status, created_at
    FROM 
        notifications
    WHERE 
        title=$1`
    
    var notification types.Notification
    err := s.db.QueryRow(query, title).Scan(
        &notification.ID,
        &notification.ReceivedUserID,
        &notification.SentUserID,
        &notification.ReceivedProjectID,
        &notification.SentProjectID,
        &notification.SynergyType,
        &notification.Type,
        &notification.Title,
        &notification.Message,
        &notification.Status,
        &notification.CreatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &notification, nil
}

func (s *Store) GetNotificationsByUser(userId int) ([]types.DetailedNotification, error) {
    query := `
    SELECT 
        n.id, n.received_user_id, n.sent_user_id, n.received_project_id, n.sent_project_id, n.synergy_type, n.type, n.title, n.message, n.status, n.created_at,
        p.name, p.description, p.status, p.user_id, p.subcategory_id, p.category_id, p.created_at, p.updated_at, p.photo,
        sc.name AS subcategory_name,
        c.name AS category_name
    FROM 
        notifications n
    LEFT JOIN projects p ON n.received_project_id = p.id
    LEFT JOIN subcategories sc ON p.subcategory_id = sc.id
    LEFT JOIN categories c ON p.category_id = c.id
    WHERE 
        n.received_user_id = $1
    ORDER BY 
        n.id ASC`

    rows, err := s.db.Query(query, userId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notifications []types.DetailedNotification
    for rows.Next() {
        var notification types.DetailedNotification
        if err := rows.Scan(
            &notification.ID,
            &notification.ReceivedUserID,
            &notification.SentUserID,
            &notification.ReceivedProjectID,
            &notification.SentProjectID,
            &notification.SynergyType,
            &notification.Type,
            &notification.Title,
            &notification.Message,
            &notification.Status,
            &notification.CreatedAt,
            &notification.Project.Name,
            &notification.Project.Description,
            &notification.Project.Status,
            &notification.Project.UserID,
            &notification.Project.SubcategoryID,
            &notification.Project.CategoryID,
            &notification.Project.CreatedAt,
            &notification.Project.UpdatedAt,
            &notification.Project.Photo,
            &notification.SubcategoryName,
            &notification.CategoryName,
        ); err != nil {
            return nil, err
        }
        notifications = append(notifications, notification)
    }
    return notifications, nil
}

