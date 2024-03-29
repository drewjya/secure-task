// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package repository

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID      `json:"id"`
	UserID    uuid.UUID      `json:"userId"`
	UserName  string         `json:"userName"`
	Picture   *string `json:"picture"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

type Token struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"userId"`
	SessionToken   string    `json:"sessionToken"`
	CreatedAt      time.Time `json:"createdAt"`
	LastAccessedAt time.Time `json:"lastAccessedAt"`
	ExpiresAt      time.Time `json:"expiresAt"`
	IsValid        bool      `json:"isValid"`
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	LastAccessed time.Time `json:"lastAccessed"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
