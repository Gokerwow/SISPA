package http

import (
	"log"
	"net/http"
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

func (s *ServiceHandler) Create(c *gin.Context) {
	var input domain.Service

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[BIND ERROR] Failed to bind the request to JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err := s.usecase.RegisterNewService(c, &input)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Register new service: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "Service Created Successfully"})

}

func (s *ServiceHandler) GetAll(c *gin.Context) {

	services, err := s.usecase.GetAll(c)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Get All services: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if services == nil {
		log.Printf("[DB ERROR] Failed to Get All therapist: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Theres no service data existed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": services})

}

func (s *ServiceHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	service, err := s.usecase.GetByID(c, idINT)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Get service detail: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if service == nil {
		log.Printf("[DB ERROR] Failed to Get service detail: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Theres no service data existed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": service})

}

func (s *ServiceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input domain.Service

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[BIND ERROR] Failed to bind the request to JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	input.ID = idINT

	service, err := s.usecase.Update(c, &input)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Update Service: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": service})

}

func (s *ServiceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	idINT, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = s.usecase.Delete(c, idINT)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Delete Service: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Service Successfully deleted"})

}
