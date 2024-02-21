package service

import (
	"exam3/api/models"
	"exam3/pkg/logger"
	"exam3/storage"
	"golang.org/x/net/context"
)

type bookService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewBookService(storage storage.IStorage, log logger.ILogger) bookService {
	return bookService{
		storage: storage,
		log:     log,
	}
}
func (b bookService) Create(ctx context.Context, createBook models.CreateBook) (models.Book, error) {
	b.log.Info("Book create service layer", logger.Any("createBook", createBook))

	bookID, err := b.storage.Book().Create(ctx, createBook)
	if err != nil {
		b.log.Error("error in service layer while creating book", logger.Error(err))
		return models.Book{}, err
	}

	book, err := b.storage.Book().GetByID(ctx, bookID)
	if err != nil {
		b.log.Error("error in service layer while getting book", logger.Error(err))
		return models.Book{}, err
	}

	return book, nil
}
func (b bookService) Get(ctx context.Context, bookID string) (models.Book, error) {
	b.log.Info("Book get service layer", logger.Any("bookID", bookID))

	book, err := b.storage.Book().GetByID(ctx, bookID)
	if err != nil {
		b.log.Error("error in service layer while getting book", logger.Error(err))
		return models.Book{}, err
	}

	return book, nil
}

func (b bookService) GetList(ctx context.Context, request models.BookGetRequest) (models.BookResponse, error) {
	b.log.Info("Book get list service layer", logger.Any("request", request))

	books, err := b.storage.Book().GetList(ctx, request)
	if err != nil {
		b.log.Error("error in service layer while getting book list", logger.Error(err))
		return models.BookResponse{}, err
	}

	return books, nil
}

func (b bookService) Update(ctx context.Context, updateBook models.Update) (models.Book, error) {
	b.log.Info("Book update service layer", logger.Any("updateBook", updateBook))

	bookID, err := b.storage.Book().Update(ctx, updateBook)
	if err != nil {
		b.log.Error("error in service layer while updating book", logger.Error(err))
		return models.Book{}, err
	}

	updatedBook, err := b.storage.Book().GetByID(ctx, bookID)
	if err != nil {
		b.log.Error("error in service layer while getting book", logger.Error(err))
		return models.Book{}, err
	}

	return updatedBook, nil
}

func (b bookService) Delete(ctx context.Context, bookID string) error {
	b.log.Info("Book delete service layer", logger.Any("bookID", bookID))

	err := b.storage.Book().Delete(ctx, bookID)
	return err
}

func (b bookService) UpdatePage(ctx context.Context, page models.UpdatePage) (models.Book, error) {
	b.log.Info("Book update page service layer", logger.Any("updatePage", page))

	id, err := b.storage.Book().UpdatePage(ctx, page)
	if err != nil {
		b.log.Error("error in service layer while updating page", logger.Error(err))
		return models.Book{}, err
	}
	updatedPage, err := b.storage.Book().GetByID(ctx, id)
	if err != nil {
		b.log.Error("error in service layer while getting book", logger.Error(err))
		return models.Book{}, err
	}
	return updatedPage, nil
}
