package http

import (
	"log"
	"net/http"
	"sispa-backend/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	usecase domain.CustomerUsecase
}

func NewCustomerHandler(u domain.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{usecase: u}
}

// TODO: Refactor status codes to net/http
// ROUTES ATAU PATH BUAT API
func (c *CustomerHandler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/customers")
	{
		group.POST("", c.Register)
		group.GET("", c.GetAll)
		group.GET("/:id", c.GetByID)
		group.PUT("/:id", c.Update)
		group.DELETE("/:id", c.Delete)
	}
}

// DAFTAR CUSTOMER BARU
// TODO: Add return data from usecase
func (c *CustomerHandler) Register(ctx *gin.Context) {
	var input domain.Customer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Printf("[BIND ERROR] Failed to bind the request to JSON: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	timeoutContext := ctx.Request.Context()

	err := c.usecase.RegisterNewCustomer(timeoutContext, &input)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Register New Customer: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": input})
}

// DAPETIN SEMUA DATA CUSTOMER
func (c *CustomerHandler) GetAll(ctx *gin.Context) {
	customers, err := c.usecase.GetAll(ctx)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Get All Customer: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customers == nil {
		customers = []domain.Customer{}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": customers})
}

// DAPETIN DATA CUSTOMER SESUAI NOMOR HP, NGAMBIL DARI PARAM
func (c *CustomerHandler) GetByPhone(ctx *gin.Context) {
	// KODE BUAT NGAMBIL PARAM
	phoneString := ctx.Param("phone")

	customers, err := c.usecase.GetByPhone(ctx, phoneString)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Get Customer detail by phone: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customers == nil {
		log.Printf("[DB ERROR] Failed to Get Customer detail by phone: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "There is no customers with that phone number existed"})
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

// DAPETIN DATA CUSTOMER SESUAI ID, NGAMBIL DARI PARAM
func (c *CustomerHandler) GetByID(ctx *gin.Context) {
	// CODE BUAT NGAMBIL PARAM
	IDstring := ctx.Param("id")
	ID, err := strconv.Atoi(IDstring)

	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	customers, err := c.usecase.GetByID(ctx, ID)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Get Customer detail by ID: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customers == nil {
		log.Printf("[DB ERROR] Failed to Get Customer detail by ID: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "There is no customers with that phone number existed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": customers})
}

// UPDATE DATA CUSTOMER SESUAI NOMOR ID, NGAMBIL DARI PARAM
func (c *CustomerHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input domain.Customer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Printf("[BIND ERROR] Failed to bind the request to JSON: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	input.ID = id

	customer, err := c.usecase.UpdateCustomer(ctx, &input)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Update Customer: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": customer})

}

// DAPETIN DATA CUSTOMER SESUAI NOMOR ID, NGAMBIL DARI PARAM
func (c *CustomerHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = c.usecase.DeleteCustomer(ctx, id)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Delete Customer: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully"})
}
