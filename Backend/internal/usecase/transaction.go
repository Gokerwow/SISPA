package usecase

import (
	"context"
	"errors"
	"sispa-backend/internal/domain"
	"time"
)

type TransactionsUsecase struct {
	repo         domain.TransactionRepository
	serviceRepo  domain.ServiceRepository
	customerRepo domain.CustomerRepository
}

func NewTransactionUsecase(tr domain.TransactionRepository, sr domain.ServiceRepository, cr domain.CustomerRepository) domain.TransactionsUsecase {
	return &TransactionsUsecase{repo: tr, serviceRepo: sr, customerRepo: cr}
}

func (u *TransactionsUsecase) CreateTransaction(c context.Context, transaction *domain.TransactionRequest) (*domain.InvoiceResponse, error) {
	if transaction.Date == "" {
		return nil, errors.New("Transaction Date is required")
	}

	timestamp := time.Now().Format("20060102-150405")

	var notaNumber string

	if transaction.NotaNumber == "" {
		notaNumber = "INV-" + timestamp
	} else {
		notaNumber = transaction.NotaNumber
	}

	invoice := domain.Invoice{
		NotaNumber:     notaNumber,
		Subtotal:       transaction.Subtotal,
		DiscountAmount: transaction.Discount,
		GrandTotal:     transaction.GrandTotal,
		Status:         transaction.Status,
		Date:           &transaction.Date,
	}

	invoiceItems := []domain.InvoiceItem{}

	for _, item := range transaction.Items {
		var invoiceItem domain.InvoiceItem

		service, err := u.serviceRepo.GetByID(c, item.ServiceID)
		if err != nil {
			return nil, err
		}

		customer, err := u.customerRepo.GetByID(c, item.CustomerID)
		if err != nil {
			return nil, err
		}

		var price int

		switch customer.CustomerType {
		case "male":
			price = service.PriceMale
		case "female":
			price = service.PriceFemale
		default:
			price = service.PriceTourist
		}

		invoiceItem = domain.InvoiceItem{
			ServiceID:   item.ServiceID,   // Comes from the Vue request
			TherapistID: item.TherapistID, // Comes from the Vue request
			CustomerID:  item.CustomerID,  // Comes from the Vue request
			Price:       price,            // 💡 Comes from the Database! (Security)
		}

		invoiceItems = append(invoiceItems, invoiceItem)
	}

	id, err := u.repo.Store(c, &invoice, invoiceItems)
	if err != nil {
		return nil, err
	}

	details, err := u.GetAllInvoiceDetails(c, id)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func (u *TransactionsUsecase) GetAllInvoiceDetails(c context.Context, id int) (*domain.InvoiceResponse, error) {
	header, err := u.repo.GetHeaderByID(c, id)
	if err != nil {
		return nil, err
	}

	items, err := u.repo.GetDetailByID(c, id)
	if err != nil {
		return nil, err
	}

	invoiceDetail := domain.InvoiceResponse{
		Invoice: *header, // We dereference the pointer to embed the actual data
		Items:   items,   // We pass the slice directly
	}

	return &invoiceDetail, nil

}

func (u *TransactionsUsecase) GetAllInvoiceHeaders(c context.Context) ([]domain.Invoice, error) {
	header, err := u.repo.GetAllHeader(c)
	if err != nil {
		return nil, err
	}

	return header, err
}

func (u *TransactionsUsecase) UpdateInvoiceStatus(c context.Context, id int, status string, PaymentMethod string) (*domain.InvoiceResponse, error) {
	err := u.repo.UpdateStatus(c, id, status, PaymentMethod)
	if err != nil {
		return nil, err
	}

	invoice, err := u.GetAllInvoiceDetails(c, id)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (u *TransactionsUsecase) UpdateTransaction(c context.Context, id int, req *domain.TransactionRequest) (*domain.InvoiceResponse, error) {
	if req.Date == "" {
		return nil, errors.New("Transaction Date is required")
	}

	invoice := domain.Invoice{
		NotaNumber:     req.NotaNumber,
		Subtotal:       req.Subtotal,
		DiscountAmount: req.Discount,
		GrandTotal:     req.GrandTotal,
		Status:         req.Status,
		Date:           &req.Date,
	}

	invoiceItems := []domain.InvoiceItem{}

	for _, item := range req.Items {
		var invoiceItem domain.InvoiceItem

		service, err := u.serviceRepo.GetByID(c, item.ServiceID)
		if err != nil {
			return nil, err
		}

		customer, err := u.customerRepo.GetByID(c, item.CustomerID)
		if err != nil {
			return nil, err
		}

		var price int

		switch customer.CustomerType {
		case "male":
			price = service.PriceMale
		case "female":
			price = service.PriceFemale
		default:
			price = service.PriceTourist
		}

		invoiceItem = domain.InvoiceItem{
			ServiceID:   item.ServiceID,   // Comes from the Vue request
			TherapistID: item.TherapistID, // Comes from the Vue request
			CustomerID:  item.CustomerID,  // Comes from the Vue request
			Price:       price,            // 💡 Comes from the Database! (Security)
		}

		invoiceItems = append(invoiceItems, invoiceItem)
	}

	err := u.repo.UpdateTransaction(c, id, &invoice, invoiceItems)
	if err != nil {
		return nil, err
	}

	details, err := u.GetAllInvoiceDetails(c, id)
	if err != nil {
		return nil, err
	}

	return details, nil
}
