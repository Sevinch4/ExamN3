package storage

import (
	"context"
	"exam3/api/models"
)

type IStorage interface {
	Close()
	Book() IBook
}

type IBook interface {
	Create(context.Context, models.CreateBook) (string, error)
	GetByID(context.Context, string) (models.Book, error)
	GetList(context.Context, models.BookGetRequest) (models.BookResponse, error)
	Update(context.Context, models.Update) (string, error)
	Delete(context.Context, string) error
	UpdatePage(context.Context, models.UpdatePage) (string, error)
}
