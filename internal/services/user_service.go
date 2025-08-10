package services

import (
	"database/sql"
	"fmt"

	"splitexpense/internal/models"

	"github.com/google/uuid"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetByID(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, name, default_currency, avatar_url, created_at, updated_at
		FROM users WHERE id = $1
	`
	err := s.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Name, &user.DefaultCurrency,
		&user.AvatarURL, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (s *UserService) Update(id uuid.UUID, updates map[string]interface{}) (*models.User, error) {
	// Build dynamic update query
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	for field, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", field, argIndex))
		args = append(args, value)
		argIndex++
	}

	if len(setParts) == 0 {
		return s.GetByID(id)
	}

	query := fmt.Sprintf(`
		UPDATE users SET %s, updated_at = NOW()
		WHERE id = $%d
		RETURNING id, email, name, default_currency, avatar_url, created_at, updated_at
	`, setParts[0], argIndex)

	for i := 1; i < len(setParts); i++ {
		query = query[:len(query)-len(fmt.Sprintf("WHERE id = $%d", argIndex))] + ", " + setParts[i] + fmt.Sprintf(" WHERE id = $%d", argIndex)
	}

	args = append(args, id)

	user := &models.User{}
	err := s.db.QueryRow(query, args...).Scan(
		&user.ID, &user.Email, &user.Name, &user.DefaultCurrency,
		&user.AvatarURL, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}
