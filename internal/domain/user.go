package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) GetID() uuid.UUID {
	return u.ID
}

func (u User) SetID(id uuid.UUID) {
	u.ID = id
}
