package http

import (
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

func (h *TherapistHandler) Create(ctx *gin.Context) {
	var input domain.Therapist

	if err := ctx.ShouldBindJSON(&input); err != nil {
		// 💡 This will tell you EXACTLY which field is the traitor
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.RegisterNewTherapist(&input)

	if err != nil {
		ctx.JSON(409, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"Message": "Therapist Created Successfully"})

}

func (h *TherapistHandler) GetAll(ctx *gin.Context) {
	therapists, err := h.usecase.GetAll()

	if err != nil {
		ctx.JSON(409, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, therapists)

}
func (h *TherapistHandler) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	therapists, err := h.usecase.GetByID(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, therapists)

}

func (h *TherapistHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var input domain.Therapist

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	input.ID = idINT

	therapist, err := h.usecase.Update(&input)

	if err != nil {
		// This will show you if it's a "Duplicate Entry", "Data Truncation",
		// or "Table not found" error.
		ctx.JSON(409, gin.H{
			"error":      "Backend Error: " + err.Error(),
			"suggestion": "Check if you are sending all required fields",
		})
		return
	}

	ctx.JSON(201, therapist)

}

func (h *TherapistHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.usecase.Delete(idINT)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"Message": "Therapist Successfully deleted"})

}
