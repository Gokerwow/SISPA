package domain

import (
	"context"
	"time"
)

type Invoice struct {
	ID             int        `json:"id"`
	NotaNumber     string     `json:"nota_number"`     // e.g. "INV-20260218-001"
	Subtotal       int        `json:"subtotal"`        // Sum of all items prices
	DiscountAmount int        `json:"discount_amount"` // The shop-wide discount
	GrandTotal     int        `json:"grand_total"`     // Subtotal - Discount
	Status         string     `json:"status"`          // 'pending', 'paid'
	Date           *string    `json:"date"`            // "YYYY-MM-DD"
	CreatedAt      *time.Time `json:"created_at,omitempty"`
}

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
	Subtotal   int    `json:"subtotal"`    // Sum of all items prices
	GrandTotal int    `json:"grand_total"` // Subtotal - Discount

	// 💡 This is the magic part: A list inside the struct!
	Items []TransactionItemRequest `json:"items" binding:"dive"`
}

type TransactionItemRequest struct {
	CustomerID  int `json:"customer_id"`
	TherapistID int `json:"therapist_id"`
	ServiceID   int `json:"service_id"`
	// We don't ask frontend for Price. Backend fetches it for security.
}

type TransactionRepository interface {
	Store(invoice *Invoice, items []InvoiceItem) error
	GetAllHeader() ([]Invoice, error)
	GetHeaderByID(id int) (*Invoice, error)
	GetDetailByID(id int) ([]InvoiceItem, error)
	UpdateStatus(id int, status string) (*Invoice, []InvoiceItem, error)
}

type TransactionsUsecase interface {
	CreateTransaction(ctx context.Context, req *TransactionRequest) error
	GetAllInvoices(ctx context.Context) ([]TransactionRequest, error)
}
