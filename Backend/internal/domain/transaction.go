package domain

import (
	"context"
	"time"
)

// For invoice to db
type Invoice struct {
	ID             int        `json:"id"`
	NotaNumber     string     `json:"nota_number"`     // e.g. "INV-20260218-001"
	Subtotal       int        `json:"sub_total"`       // Sum of all items prices
	DiscountAmount int        `json:"discount_amount"` // The shop-wide discount
	GrandTotal     int        `json:"grand_total"`     // Subtotal - Discount
	Status         string     `json:"status"`          // 'pending', 'paid'
	Date           *string    `json:"date"`            // "YYYY-MM-DD"
	CreatedAt      *time.Time `json:"created_at,omitempty"`
}

// For invoice items to db
type InvoiceItem struct {
	ID          int `json:"id"`
	InvoiceID   int `json:"invoice_id"`   // The link to the parent
	TherapistID int `json:"therapist_id"` // Who did the work
	ServiceID   int `json:"service_id"`   // What service
	CustomerID  int `json:"customer_id"`  // Who received it
	Price       int `json:"price"`        // Snapshot of price at that moment
}

// This matches the JSON from Vue
type TransactionRequest struct {
	NotaNumber string `json:"nota_number"` // Optional, or generate in backend
	Date       string `json:"date"`        // "2026-02-18"
	Status     string `json:"status"`      // "pending"
	Discount   int    `json:"discount"`    // Global discount for invoice
	Subtotal   int    `json:"sub_total"`   // Sum of all items prices
	GrandTotal int    `json:"grand_total"` // Subtotal - Discount

	// 💡 This is the magic part: A list inside the struct!
	Items []TransactionItemRequest `json:"items" binding:"dive"`
}

// For getting invoice items to TransactionRequest
type TransactionItemRequest struct {
	CustomerID  int `json:"customer_id"`
	TherapistID int `json:"therapist_id"`
	ServiceID   int `json:"service_id"`
	// We don't ask frontend for Price. Backend fetches it for security.
}

// to return all invoice details
type InvoiceResponse struct {
	Invoice               // Anonymous embedding (promotes all Invoice fields)
	Items   []InvoiceItem `json:"items"`
}

// struct for status update from frontend
type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type TransactionRepository interface {
	Store(c context.Context, invoice *Invoice, items []InvoiceItem) (int, error)
	GetAllHeader(c context.Context) ([]Invoice, error)
	GetHeaderByID(c context.Context, id int) (*Invoice, error)
	GetDetailByID(c context.Context, id int) ([]InvoiceItem, error)
	UpdateStatus(c context.Context, id int, status string) error
}

type TransactionsUsecase interface {
	CreateTransaction(c context.Context, req *TransactionRequest) (*InvoiceResponse, error)
	GetAllInvoiceHeaders(c context.Context) ([]Invoice, error)
	GetAllInvoiceDetails(c context.Context, id int) (*InvoiceResponse, error)
	UpdateInvoiceStatus(c context.Context, id int, status string) (*InvoiceResponse, error)
}
