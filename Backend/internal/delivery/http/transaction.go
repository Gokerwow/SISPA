package http

import (
	"log"
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
		group.PATCH("/:id", h.UpdateInvoiceStatus)
		group.PUT("/:id", h.UpdateTransaction)
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input domain.TransactionRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[BIND ERROR] Failed to bind the input to JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	invoice, err := h.usecase.CreateTransaction(c, &input)

	if err != nil {
		log.Printf("[DB ERROR] Failed to Create transaction: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

func (h *TransactionHandler) GetAllHeader(c *gin.Context) {
	headers, err := h.usecase.GetAllInvoiceHeaders(c)
	if err != nil {
		log.Printf("[DB ERROR] Failed to Get All invoices headers: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": headers})
}

func (h *TransactionHandler) GetAllInvoiceDetails(c *gin.Context) {
	IDString := c.Param("id")
	IDConv, err := strconv.Atoi(IDString)

	if err != nil {
		log.Printf("[DB ERROR] Failed to convert ID string to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	invoice, err := h.usecase.GetAllInvoiceDetails(c, IDConv)
	if err != nil {
		log.Printf("[DB ERROR] Failed to Get All invoices Details: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

func (h *TransactionHandler) UpdateInvoiceStatus(c *gin.Context) {
	IDString := c.Param("id")
	IDConv, err := strconv.Atoi(IDString)

	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req domain.UpdateStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[BIND ERROR] Failed to bind the request to JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecase.UpdateInvoiceStatus(c, IDConv, req.Status, req.PaymentMethod)
	if err != nil {
		log.Printf("[DB ERROR] Failed to Update invoice status: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	IDString := c.Param("id")
	IDConv, err := strconv.Atoi(IDString)
	if err != nil {
		log.Printf("[CONVERT ERROR] Failed to convert ID string to int: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req domain.TransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[BIND ERROR] Failed to bind the request to JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecase.UpdateTransaction(c, IDConv, &req)
	if err != nil {
		log.Printf("[DB ERROR] Failed to Update invoice: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
