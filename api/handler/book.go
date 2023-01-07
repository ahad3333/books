package handler

import (
	"add/models"
	"add/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreateBook(c *gin.Context) {

	var book models.CreateBook

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.InsertBook(h.db, book)
	if err != nil {
		log.Println("error whiling create book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := storage.GetByIdBook(h.db, models.BookPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) GetByIDBook(c *gin.Context) {

	id := c.Param("id")

	res, err := storage.GetByIdBook(h.db, models.BookPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

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

	res, err := storage.GetListBook(h.db, models.GetListBookRequest{
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

func (h *Handler) UpdateBook(c *gin.Context){
	var book models.UpdateBook

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Println("error whiling updete marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id := c.Param("id")

	err = storage.UpdateBook(h.db, models.UpdateBook{Id: id, Name: book.Name, Price: book.Price,Description: book.Description})
	if err != nil {
		log.Println("error whiling updete  book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	

	c.JSON(http.StatusCreated, "updete  book")
}

func (h *Handler) DeleteBook(c *gin.Context){
	id := c.Param("id")

	err := storage.DeleteBook(h.db, models.BookPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated,"delete book")
}
