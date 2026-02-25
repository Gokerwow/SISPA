package http

import (
	"net/http"
	"sispa-backend/internal/domain"
	"sispa-backend/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TherapistHandler struct {
	usecase *usecase.TherapistUsecase
}

func NewTherapistHandler(u *usecase.TherapistUsecase) *TherapistHandler {
	return &TherapistHandler{usecase: u}
}

func (h *TherapistHandler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/therapists")
	{
		group.POST("", h.Create)
		group.GET("", h.GetAll)
		group.GET("/:id", h.GetByID)
		group.PUT("/:id", h.Update)
		group.DELETE("/:id", h.Delete)
	}
}

func (h *TherapistHandler) Create(c *gin.Context) {
	var input domain.Therapist

	if err := c.ShouldBindJSON(&input); err != nil {
		// 💡 This will tell you EXACTLY which field is the traitor
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.RegisterNewTherapist(c, &input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "Therapist Created Successfully"})

}

func (h *TherapistHandler) GetAll(c *gin.Context) {
	therapists, err := h.usecase.GetAllTherapist(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, therapists)

}

func (h *TherapistHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	therapists, err := h.usecase.GetByID(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, therapists)

}

func (h *TherapistHandler) Update(c *gin.Context) {
	id := c.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input domain.Therapist

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	input.ID = idINT

	therapist, err := h.usecase.Update(c, &input)

	if err != nil {
		// This will show you if it's a "Duplicate Entry", "Data Truncation",
		// or "Table not found" error.
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      "Backend Error: " + err.Error(),
			"suggestion": "Check if you are sending all required fields",
		})
		return
	}

	c.JSON(http.StatusOK, therapist)

}

func (h *TherapistHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.usecase.Delete(c, idINT)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Therapist Successfully deleted"})

}
