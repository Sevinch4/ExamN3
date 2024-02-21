package postgres

import (
	"context"
	"exam3/api/models"
	"exam3/pkg/logger"
	"exam3/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewBookRepo(db *pgxpool.Pool, log logger.ILogger) storage.IBook {
	return &bookRepo{
		db:  db,
		log: log,
	}
}

func (b *bookRepo) Create(ctx context.Context, createBook models.CreateBook) (string, error) {
	query := `INSERT INTO books(id, name, author_name, page_number) 
                          values($1, $2, $3, $4)`

	id := uuid.New()
	if _, err := b.db.Exec(ctx, query, id, createBook.Name, createBook.AuthorName, createBook.PageNumber); err != nil {
		b.log.Error("error is while inserting data", logger.Error(err))
		return "", err
	}

	return id.String(), nil
}

func (b *bookRepo) GetByID(ctx context.Context, id string) (models.Book, error) {
	book := models.Book{}
	query := `SELECT id, name, author_name, page_number, created_at, updated_at FROM books WHERE id = $1 and deleted_at = 0`

	if err := b.db.QueryRow(ctx, query, id).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorName,
		&book.PageNumber,
		&book.CreatedAt,
		&book.UpdatedAt,
	); err != nil {
		fmt.Println("error is while selecting", err.Error())
		return models.Book{}, err
	}

	return book, nil
}

func (b *bookRepo) GetList(ctx context.Context, request models.BookGetRequest) (models.BookResponse, error) {
	var (
		query, countQuery string
		filter            string
		count             = 0
		books             []models.Book
		page              = request.Page
		offset            = (page - 1) * request.Limit
	)

	if request.BookName != "" {
		filter += fmt.Sprintf(` and name ilike '%%%s%%' `, request.BookName)
	}

	if request.AuthorName != "" {
		filter += fmt.Sprintf(` and author_name ilike '%%%s%%'`, request.AuthorName)
	}

	countQuery = `SELECT count(1) FROM books where deleted_at = 0 ` + filter

	if err := b.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		b.log.Error("error is while scanning count", logger.Error(err))
		return models.BookResponse{}, err
	}

	query = `SELECT id, name, author_name, page_number, created_at, updated_at FROM books where deleted_at = 0 ` + filter + `  ORDER BY created_at desc  LIMIT $1 OFFSET $2`

	rows, err := b.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		b.log.Error("error is while selecting all ", logger.Error(err))
		return models.BookResponse{}, err
	}

	for rows.Next() {
		book := models.Book{}
		if err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.AuthorName,
			&book.PageNumber,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			b.log.Error("error is while scanning all", logger.Error(err))
			return models.BookResponse{}, err
		}
		books = append(books, book)
	}

	return models.BookResponse{
		Books: books,
		Count: count,
	}, err
}

func (b *bookRepo) Update(ctx context.Context, update models.Update) (string, error) {
	query := `UPDATE books set name = $1, author_name = $2, updated_at = now() WHERE id = $3`

	rowsAffected, err := b.db.Exec(ctx, query, &update.Name, &update.AuthorName, &update.ID)
	if err != nil {
		b.log.Error("error is while updating")
		return "", err
	}

	if r := rowsAffected.RowsAffected(); r == 0 {
		b.log.Error("error is while rows affected", logger.Error(err))
		return "", err
	}

	return update.ID, err
}

func (b *bookRepo) Delete(ctx context.Context, id string) error {
	query := `UPDATE books SET deleted_at = extract(epoch from current_timestamp) WHERE id = $1`

	rowsAffected, err := b.db.Exec(ctx, query, id)
	if err != nil {
		b.log.Error("error is while deleting", logger.Error(err))
		return err
	}

	if r := rowsAffected.RowsAffected(); r == 0 {
		b.log.Error("error is while rows affected")
		return err
	}

	return err
}

func (b *bookRepo) UpdatePage(ctx context.Context, page models.UpdatePage) (string, error) {
	query := `UPDATE books SET page_number = $2, updated_at = now() WHERE id = $1`
	rowsAffected, err := b.db.Exec(ctx, query, &page.ID, &page.PageNumber)
	if err != nil {
		b.log.Error("error is while updating page", logger.Error(err))
		return "", err
	}

	if r := rowsAffected.RowsAffected(); r == 0 {
		b.log.Error("error is while rows affected")
		return "", err
	}

	return page.ID, nil
}
