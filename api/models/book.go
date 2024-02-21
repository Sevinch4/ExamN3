package models

import "time"

type Book struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	AuthorName string    `json:"author_name"`
	PageNumber int       `json:"page_number"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateBook struct {
	Name       string `json:"name"`
	AuthorName string `json:"author_name"`
	PageNumber int    `json:"page_number"`
}

type Update struct {
	ID         string `json:"-"`
	Name       string `json:"name"`
	AuthorName string `json:"author_name"`
}

type UpdatePage struct {
	ID         string `json:"-"`
	PageNumber int    `json:"page_number"`
}
type BookGetRequest struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
}

type BookResponse struct {
	Books []Book `json:"books"`
	Count int    `json:"count"`
}
