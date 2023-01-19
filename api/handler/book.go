package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"add/models"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @ID CreateBook
// @Router /book [POST]
// @Summary CreateBook
// @Description CreateBook
// @Tags Book
// @Accept json
// @Produce json
// @Param book body models.CreateBook true "CreateBookRequestBody"
// @Success 201 {object} models.Book "GetBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateBook(c *gin.Context) {

	var book models.CreateBook

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Book().Insert(context.Background(), &book)
	if err != nil {
		log.Println("error whiling create book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.Book().GetByID(context.Background(), &models.BookPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIDBook godoc
// @ID Get_By_IDBook
// @Router /book/{id} [GET]
// @Summary GetByID Book
// @Description GetByID Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Book "GetByIDBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDBook(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Book().GetByID(context.Background(), &models.BookPrimeryKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListBook godoc
// @ID BookPrimeryKey
// @Router /book [GET]
// @Summary Get List Book
// @Description Get List Book
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListBookResponse "GetBookListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListBook(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res, err := h.storage.Book().GetList(context.Background(),&models.GetListBookRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateBook godoc
// @ID UpdateBook
// @Router /book/{id} [PUT]
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param book body models.UpdateBookSwag true "UpdateBookRequestBody"
// @Success 202 {object} models.Book "UpdateBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateBook(c *gin.Context) {

	var (
		book models.UpdateBook
	)


	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	book.Id = c.Param("id")
	 err = h.storage.Book().Update(context.Background(),&models.UpdateBook{
		Id: book.Id,
		Name: book.Name,
		Price: book.Price,
		Description: book.Description,
	})
	if err != nil {
		log.Printf("error whiling update 2: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	// if rowsAffected == 0 {
	// 	log.Printf("error whiling update rows affected: %v", err)
	// 	c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
	// 	return
	// }

	resp, err := h.storage.Book().GetByID(context.Background(),&models.BookPrimeryKey{Id: book.Id})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// DeleteBook godoc
// @ID DeleteBook
// @Router /book/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Book().Delete(context.Background(),&models.BookPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  book:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "delete book")
}
