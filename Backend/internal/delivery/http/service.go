package http

import (
	"sispa-backend/internal/domain"
	"sispa-backend/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	usecase *usecase.ServiceUsecase
}

func NewServiceHandler(u *usecase.ServiceUsecase) *ServiceHandler {
	return &ServiceHandler{usecase: u}
}

func (s *ServiceHandler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/services")
	{
		group.POST("", s.Create)
		group.GET("", s.GetAll)
		group.GET("/:id", s.GetByID)
		group.PUT("/:id", s.Update)
		group.DELETE("/:id", s.Delete)
	}
}

func (s *ServiceHandler) Create(ctx *gin.Context) {
	var input domain.Service

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	err := s.usecase.RegisterNewService(&input)

	if err != nil {
		ctx.JSON(409, gin.H{"error": err.Error()})
	}

	ctx.JSON(201, gin.H{"Message": "Service Created Successfully"})

}

func (s *ServiceHandler) GetAll(ctx *gin.Context) {

	services, err := s.usecase.GetAll()

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if services == nil {
		ctx.JSON(400, gin.H{"error": "Theres no service data existed"})
		return
	}

	ctx.JSON(201, services)

}

func (s *ServiceHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	service, err := s.usecase.GetByID(idINT)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if service == nil {
		ctx.JSON(404, gin.H{"error": "Theres no service data existed"})
		return
	}

	ctx.JSON(201, service)

}

func (s *ServiceHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var input domain.Service

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	input.ID = idINT

	service, err := s.usecase.Update(&input)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, service)

}

func (s *ServiceHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	err = s.usecase.Delete(idINT)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"Message": "Service Successfully deleted"})

}
