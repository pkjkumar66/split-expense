package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Email           string    `json:"email" db:"email"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	Name            string    `json:"name" db:"name"`
	DefaultCurrency string    `json:"default_currency" db:"default_currency"`
	AvatarURL       *string   `json:"avatar_url" db:"avatar_url"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type Group struct {
	ID          uuid.UUID `json:"id" db:"id"`
	OwnerID     uuid.UUID `json:"owner_id" db:"owner_id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	AvatarURL   *string   `json:"avatar_url" db:"avatar_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Members     []User    `json:"members,omitempty"`
}

type Membership struct {
	ID       uuid.UUID `json:"id" db:"id"`
	GroupID  uuid.UUID `json:"group_id" db:"group_id"`
	UserID   uuid.UUID `json:"user_id" db:"user_id"`
	Role     string    `json:"role" db:"role"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}

type Expense struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	GroupID     *uuid.UUID     `json:"group_id" db:"group_id"`
	CreatedBy   uuid.UUID      `json:"created_by" db:"created_by"`
	Amount      float64        `json:"amount" db:"amount"`
	Currency    string         `json:"currency" db:"currency"`
	Description string         `json:"description" db:"description"`
	Date        time.Time      `json:"date" db:"date"`
	Category    *string        `json:"category" db:"category"`
	ReceiptURL  *string        `json:"receipt_url" db:"receipt_url"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
	Shares      []ExpenseShare `json:"shares,omitempty"`
	Creator     *User          `json:"creator,omitempty"`
}

type ExpenseShare struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ExpenseID   uuid.UUID `json:"expense_id" db:"expense_id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	ShareAmount float64   `json:"share_amount" db:"share_amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	User        *User     `json:"user,omitempty"`
}

type Settlement struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	GroupID     *uuid.UUID `json:"group_id" db:"group_id"`
	Payer       uuid.UUID  `json:"payer" db:"payer"`
	Payee       uuid.UUID  `json:"payee" db:"payee"`
	Amount      float64    `json:"amount" db:"amount"`
	Currency    string     `json:"currency" db:"currency"`
	Method      string     `json:"method" db:"method"`
	Description *string    `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	PayerUser   *User      `json:"payer_user,omitempty"`
	PayeeUser   *User      `json:"payee_user,omitempty"`
}

type Balance struct {
	UserID   uuid.UUID `json:"user_id"`
	GroupID  uuid.UUID `json:"group_id"`
	Currency string    `json:"currency"`
	Amount   float64   `json:"amount"`
	User     *User     `json:"user,omitempty"`
}

// Request/Response DTOs
type SignUpRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	DefaultCurrency string `json:"default_currency"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateGroupRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreateExpenseRequest struct {
	GroupID     *uuid.UUID                  `json:"group_id"`
	Amount      float64                     `json:"amount" binding:"required,gt=0"`
	Currency    string                      `json:"currency" binding:"required"`
	Description string                      `json:"description" binding:"required"`
	Date        time.Time                   `json:"date" binding:"required"`
	Category    string                      `json:"category"`
	Shares      []CreateExpenseShareRequest `json:"shares" binding:"required,min=1"`
}

type CreateExpenseShareRequest struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	ShareAmount float64   `json:"share_amount" binding:"required,gte=0"`
}

type CreateSettlementRequest struct {
	GroupID     *uuid.UUID `json:"group_id"`
	Payee       uuid.UUID  `json:"payee" binding:"required"`
	Amount      float64    `json:"amount" binding:"required,gt=0"`
	Currency    string     `json:"currency" binding:"required"`
	Method      string     `json:"method"`
	Description string     `json:"description"`
}
