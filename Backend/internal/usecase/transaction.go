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

func (u *TransactionsUsecase) CreateTransaction(c context.Context, transaction *domain.TransactionRequest) error {
	if transaction.Date == "" {
		return errors.New("Transaction Date is required")
	}

	timestamp := time.Now().Format("20060102-150405")
	notaNumber := "INV-" + timestamp

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

		service, err := u.serviceRepo.GetByID(item.ServiceID)
		if err != nil {
			return err
		}

		customer, err := u.customerRepo.GetByID(c, item.CustomerID)
		if err != nil {
			return err
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

	err := u.repo.Store(&invoice, invoiceItems)
	if err != nil {
		return err
	}

	return nil
}

func (u *TransactionsUsecase) GetAllInvoices(ctx context.Context) ([]domain.TransactionRequest, error) {
	header, err := u.repo.GetAllHeader(ctx)
}
