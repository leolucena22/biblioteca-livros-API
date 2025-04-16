package model

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required,min=3"`
	ISBN      string    `json:"isbn" validate:"required,isbn"`
	Year      int       `json:"year" validate:"gte=1900"`
	AuthorID  uint      `json:"author_id"` // Para relacionamento
	CreatedAt time.Time `json:"created_at"`
}
