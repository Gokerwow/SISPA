package http

import (
	"net/http"
	"sispa-backend/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	usecase domain.TransactionsUsecase
}

func NewTransactionHandler(u domain.TransactionsUsecase) *TransactionHandler {
	return &TransactionHandler{usecase: u}
}

func (h *TransactionHandler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/transactions")
	{
		group.GET("", h.GetAllHeader)
		group.POST("", h.CreateTransaction)
		group.GET("/:id", h.GetAllInvoiceDetails)
		group.PUT("/:id", h.GetAllInvoiceDetails)
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input domain.TransactionRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	invoice, err := h.usecase.CreateTransaction(c, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (h *TransactionHandler) GetAllHeader(c *gin.Context) {
	headers, err := h.usecase.GetAllInvoiceHeaders(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, headers)
}

func (h *TransactionHandler) GetAllInvoiceDetails(c *gin.Context) {
	IDString := c.Param("id")
	IDConv, err := strconv.Atoi(IDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	invoice, err := h.usecase.GetAllInvoiceDetails(c, IDConv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (h *TransactionHandler) UpdateInvoiceStatus(c *gin.Context) {
	IDString := c.Param("id")
	IDConv, err := strconv.Atoi(IDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req domain.UpdateStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecase.UpdateInvoiceStatus(c, IDConv, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
