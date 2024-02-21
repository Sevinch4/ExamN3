package handler

import (
	"context"
	"exam3/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// CreateBook godoc
// @Router       /book [POST]
// @Summary      Creates a new book
// @Description  create a new book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book body models.CreateBook false "basket"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBook(c *gin.Context) {
	book := models.CreateBook{}
	if err := c.ShouldBind(&book); err != nil {
		handleResponse(c, h.log, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	createdBook, err := h.service.Book().Create(ctx, book)
	if err != nil {
		handleResponse(c, h.log, "error is while creating book", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "success", http.StatusCreated, createdBook)
}

// GetBook godoc
// @Router       /book/{id} [GET]
// @Summary      Get book
// @Description  get book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "bookID"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book, err := h.service.Book().Get(ctx, id)
	if err != nil {
		handleResponse(c, h.log, "error is while getting book", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, book)
}

// GetBookList godoc
// @Router       /books [GET]
// @Summary      Get books list
// @Description  get books list
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        book-name query string false "book-name"
// @Param        author-name query string false "author-name"
// @Success      201  {object}  models.BookResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBookList(c *gin.Context) {
	var (
		page, limit          int
		bookName, authorName string
		err                  error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	bookName = c.Query("book-name")
	authorName = c.Query("author-name")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	books, err := h.service.Book().GetList(ctx, models.BookGetRequest{
		Page:       page,
		Limit:      limit,
		BookName:   bookName,
		AuthorName: authorName,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, books)
}

// UpdateBook godoc
// @Router       /book/{id} [PUT]
// @Summary      Update book
// @Description  update book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "bookID"
// @Param        book body models.Update false "basket"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	book := models.Update{}
	if err := c.ShouldBind(&book); err != nil {
		handleResponse(c, h.log, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	book.ID = id

	updatedBook, err := h.service.Book().Update(ctx, book)
	if err != nil {
		handleResponse(c, h.log, "error is while updating book", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "updated", http.StatusOK, updatedBook)
}

// DeleteBook godoc
// @Router       /book/{id} [DELETE]
// @Summary      Delete book
// @Description  delete book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "bookID"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := h.service.Book().Delete(ctx, id); err != nil {
		handleResponse(c, h.log, "error is while deleting book", http.StatusInternalServerError, err.Error())
	}

	handleResponse(c, h.log, "deleted", http.StatusOK, "deleted")
}

// UpdateBookPage godoc
// @Router       /book/{id} [PATCH]
// @Summary      Update book page
// @Description  update book page
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "bookID"
// @Param        page body models.UpdatePage false "page"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBookPage(c *gin.Context) {
	id := c.Param("id")
	page := models.UpdatePage{}

	if err := c.ShouldBind(&page); err != nil {
		handleResponse(c, h.log, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	page.ID = id

	updatedBook, err := h.service.Book().UpdatePage(ctx, page)
	if err != nil {
		handleResponse(c, h.log, "error is while updating page", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "updated", http.StatusOK, updatedBook)
}
