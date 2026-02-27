package repository

import (
	"context"
	"database/sql"
	"errors"
	"sispa-backend/internal/domain"
)

type sqlTransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) domain.TransactionRepository {
	return &sqlTransactionRepository{db: db}
}

func (r *sqlTransactionRepository) Store(c context.Context, invoice *domain.Invoice, items []domain.InvoiceItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Re-throw panic after cleanup
		} else if err != nil {
			tx.Rollback() // Rollback if any error occurred
		}
	}()

	queryHeader := `
		INSERT INTO invoices (nota_number, sub_total, discount_amount, grand_total, status, date) 
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := tx.ExecContext(c, queryHeader,
		invoice.NotaNumber,
		invoice.Subtotal,
		invoice.DiscountAmount,
		invoice.GrandTotal,
		invoice.Status,
		invoice.Date, // Assuming date is a string "YYYY-MM-DD"
	)

	if err != nil {
		return 0, err // This triggers the defer Rollback
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	invoiceID := int(id)

	// INSERT THE ITEMS
	queryItem := `
		INSERT INTO invoices_items (invoice_id, therapist_id, service_id, customer_id, price) 
		VALUES (?, ?, ?, ?, ?)
	`

	// OPTIMIZATION BIAR GAK NGULANG MANGGIL EXEC ORI DI DALAM LOOP
	stmt, err := tx.Prepare(queryItem)
	if err != nil {
		return 0, err
	}
	// CLOSE CONNECTION SETELAH FUNCTION SELESAI
	defer stmt.Close()

	for _, item := range items {
		_, err = stmt.ExecContext(c, invoiceID, item.TherapistID, item.ServiceID, item.CustomerID, item.Price)
		if err != nil {
			return 0, err // If one item fails, the WHOLE invoice is cancelled
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return invoiceID, nil
}

func (r *sqlTransactionRepository) GetAllHeader(c context.Context) ([]domain.Invoice, error) {
	query := `
		SELECT id, nota_number, sub_total, discount_amount, grand_total, status, payment_method, DATE_FORMAT(date, '%Y-%m-%d') as date, created_at
		FROM invoices
	`

	transactions := []domain.Invoice{}

	rows, err := r.db.QueryContext(c, query)
	if err != nil {
		return transactions, err
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Invoice

		err := rows.Scan(&t.ID, &t.NotaNumber, &t.Subtotal, &t.DiscountAmount, &t.GrandTotal, &t.Status, &t.PaymentMethod, &t.Date, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *sqlTransactionRepository) GetHeaderByID(c context.Context, id int) (*domain.Invoice, error) {
	query := `
		SELECT id, nota_number, sub_total, discount_amount, grand_total, status, payment_method, DATE_FORMAT(date, '%Y-%m-%d') as date, created_at FROM invoices
		WHERE id = ?
	`

	var h domain.Invoice

	err := r.db.QueryRowContext(c, query, id).Scan(&h.ID, &h.NotaNumber, &h.Subtotal, &h.DiscountAmount, &h.GrandTotal, &h.Status, &h.PaymentMethod, &h.Date, &h.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &h, nil

}

func (r *sqlTransactionRepository) GetDetailByID(c context.Context, id int) ([]domain.InvoiceItem, error) {
	query := `
		SELECT id, invoice_id, therapist_id, customer_id, service_id, price FROM invoices_items
		WHERE invoice_id = ?
	`
	details := []domain.InvoiceItem{}

	rows, err := r.db.QueryContext(c, query, id)

	if err != nil {
		return details, err
	}

	defer rows.Close()

	for rows.Next() {
		var d domain.InvoiceItem
		err := rows.Scan(&d.ID, &d.InvoiceID, &d.TherapistID, &d.CustomerID, &d.ServiceID, &d.Price)
		if err != nil {
			return nil, err
		}
		details = append(details, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return details, nil

}

func (r *sqlTransactionRepository) UpdateStatus(c context.Context, id int, status string, PaymentMethod string) error {
	query := `
		UPDATE invoices
		SET status = ?, payment_method = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(c, query, status, PaymentMethod, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		// 💡 AHA! The query ran, but no invoice had that ID!
		return errors.New("invoice not found or status already matches")
	}
	return nil
}

func (r *sqlTransactionRepository) UpdateTransaction(c context.Context, id int, invoice *domain.Invoice, items []domain.InvoiceItem) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Re-throw panic after cleanup
		} else if err != nil {
			tx.Rollback() // Rollback if any error occurred
		}
	}()

	queryHeader := `
		UPDATE invoices 
		SET nota_number = ?, sub_total = ?, discount_amount = ?, grand_total = ?, status = ?, date = ? 
		WHERE id = ?
	`

	_, err = tx.ExecContext(c, queryHeader,
		invoice.NotaNumber,
		invoice.Subtotal,
		invoice.DiscountAmount,
		invoice.GrandTotal,
		invoice.Status,
		invoice.Date, // Assuming date is a string "YYYY-MM-DD"
		id,
	)

	if err != nil {
		return err // This triggers the defer Rollback
	}

	invoiceID := int(id)

	_, err = tx.ExecContext(c, "DELETE FROM invoices_items WHERE invoice_id = ?", invoiceID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// INSERT THE ITEMS
	queryItem := `
		INSERT INTO invoices_items (invoice_id, therapist_id, service_id, customer_id, price) 
		VALUES (?, ?, ?, ?, ?)
	`

	// OPTIMIZATION BIAR GAK NGULANG MANGGIL EXEC ORI DI DALAM LOOP
	stmt, err := tx.Prepare(queryItem)
	if err != nil {
		return err
	}
	// CLOSE CONNECTION SETELAH FUNCTION SELESAI
	defer stmt.Close()

	for _, item := range items {
		_, err = stmt.ExecContext(c, invoiceID, item.TherapistID, item.ServiceID, item.CustomerID, item.Price)
		if err != nil {
			return err // If one item fails, the WHOLE invoice is cancelled
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
