package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	AuthorID      uuid.UUID `json:"author"`
	AllowComments bool      `json:"allow_comments"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (p Post) GetID() uuid.UUID {
	return p.ID
}

func (p Post) SetID(id uuid.UUID) {
	p.ID = id
}
