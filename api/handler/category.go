package handler

import (
	"add/models"
	"add/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @ID create_category
// @Router /category [POST]
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "CreateCategoryRequestBody"
// @Success 201 {object} models.Category "GetCategoryBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateCategory(c *gin.Context) {

	var book models.CreateCategory

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.InsertCategory(h.db, book)
	if err != nil {
		log.Println("error whiling create Category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := storage.GetByIdCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id Category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}
// GetByIDCategory godoc
// @ID get_by_id_category
// @Router /category/{id} [GET]
// @Summary Get By ID Category
// @Description Get By ID Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Category "GetCategoryBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdCategory(c *gin.Context) {

	id := c.Param("id")

	res, err := storage.GetByIdCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}
// GetListCategory godoc
// @ID get_list_category
// @Router /category [GET]
// @Summary Get List Category
// @Description Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListCategoryResponse "GetCategoryListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListCategory(c *gin.Context) {
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

	res, err := storage.GetListCategory(h.db, models.GetListCategoryRequest{
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
// UpdateCategory godoc
// @ID update_category
// @Router /category/{id} [PUT]
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param category body models.UpdateCategorySwag true "UpdateCategoryRequestBody"
// @Success 202 {object} models.Category "UpdateCategoryBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateCategory(c *gin.Context){
	var category models.UpdateCategory

	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println("error whiling updete marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id := c.Param("id")

	err = storage.UpdateCategory(h.db, models.UpdateCategory{Id: id, Name: category.Name})
	if err != nil {
		log.Println("error whiling updete  Category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	

	c.JSON(http.StatusCreated, "updete  Category")
}

// DeleteCategory godoc
// @ID delete_category
// @Router /category/{id} [DELETE]
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteCategoryBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteCategory(c *gin.Context){
	id := c.Param("id")

	err := storage.DeleteCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated,"delete Category")
}
