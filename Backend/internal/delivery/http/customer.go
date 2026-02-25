package http

import (
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
func (c *CustomerHandler) Register(ctx *gin.Context) {
	var input domain.Customer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	timeoutContext := ctx.Request.Context()

	err := c.usecase.RegisterNewCustomer(timeoutContext, &input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": input})
}

// DAPETIN SEMUA DATA CUSTOMER
func (c *CustomerHandler) GetAll(ctx *gin.Context) {
	customers, err := c.usecase.GetAll(ctx)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if customers == nil {
		customers = []domain.Customer{}
	}

	ctx.JSON(201, customers)
}

// DAPETIN DATA CUSTOMER SESUAI NOMOR HP, NGAMBIL DARI PARAM
func (c *CustomerHandler) GetByPhone(ctx *gin.Context) {
	// KODE BUAT NGAMBIL PARAM
	phoneString := ctx.Param("phone")

	customers, err := c.usecase.GetByPhone(ctx, phoneString)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if customers == nil {
		ctx.JSON(400, gin.H{"Message": "There is no customers with that phone number existed"})
		return
	}

	ctx.JSON(201, customers)
}

// DAPETIN DATA CUSTOMER SESUAI ID, NGAMBIL DARI PARAM
func (c *CustomerHandler) GetByID(ctx *gin.Context) {
	// CODE BUAT NGAMBIL PARAM
	IDstring := ctx.Param("id")
	ID, err := strconv.Atoi(IDstring)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	customers, err := c.usecase.GetByID(ctx, ID)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if customers == nil {
		ctx.JSON(400, gin.H{"Message": "There is no customers with that phone number existed"})
		return
	}

	ctx.JSON(201, customers)
}

// UPDATE DATA CUSTOMER SESUAI NOMOR ID, NGAMBIL DARI PARAM
func (c *CustomerHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var input domain.Customer

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}

	input.ID = id

	customer, err := c.usecase.UpdateCustomer(ctx, &input)

	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, customer)

}

// DAPETIN DATA CUSTOMER SESUAI NOMOR ID, NGAMBIL DARI PARAM
func (c *CustomerHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	err = c.usecase.DeleteCustomer(ctx, id)

	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"Message": "Deleted Successfully"})

}
