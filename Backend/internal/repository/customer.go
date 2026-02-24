package repository

import (
	"context"
	"database/sql"
	"sispa-backend/internal/domain"
)

type sqlCustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) domain.CustomerRepository {
	return &sqlCustomerRepository{db: db}
}

// 1. Buat Data Customer Baru
func (r *sqlCustomerRepository) Create(c context.Context, customer *domain.Customer) error {
	query := `
		INSERT INTO customers (name, phone, birth_date, address, customer_type, note,  created_at)
		VALUES (?,?,?,?,?,?,?)
	`
	_, err := r.db.ExecContext(c, query, customer.Name, customer.Phone, customer.BirthDate, customer.Address, customer.CustomerType, customer.Note, customer.CreatedAt)

	return err
}

// 2. Ambil semua data customer
func (r *sqlCustomerRepository) GetAll(c context.Context) ([]domain.Customer, error) {
	query := `
		SELECT id, name, phone, birth_date, address, gender, created_at FROM customers
	`

	rows, err := r.db.QueryContext(c, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := []domain.Customer{}

	for rows.Next() {
		var c domain.Customer
		err := rows.Scan(&c.ID, &c.Name, &c.Phone, &c.BirthDate, &c.Address, &c.CustomerType, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

// todo : set this as a query
// 3. Ambil data customer sesuai nomor hp
func (r *sqlCustomerRepository) GetByPhone(c context.Context, phone string) (*domain.Customer, error) {
	query := `
		SELECT id, name, phone, address, gender, note, created_at FROM customers WHERE phone=?
	`

	var customer domain.Customer

	err := r.db.QueryRowContext(c, query, phone).Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Address, &customer.CustomerType, &customer.Note, &customer.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &customer, nil
}

// 3. Ambil data customer sesuai nomor hp
func (r *sqlCustomerRepository) GetByID(c context.Context, id int) (*domain.Customer, error) {
	query := `
		SELECT id, name, phone, birth_date, address, gender, note, created_at FROM customers WHERE id=?
	`

	var customer domain.Customer

	err := r.db.QueryRowContext(c, query, id).Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.BirthDate, &customer.Address, &customer.CustomerType, &customer.Note, &customer.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &customer, nil
}

// 4. Update data customer sesuai ID
func (r *sqlCustomerRepository) Update(c context.Context, customer *domain.Customer) (*domain.Customer, error) {
	query := `
        UPDATE customers 
        SET name = ?, phone = ?, birth_date = ?, address = ?, gender = ?, note = ?
        WHERE id = ?
    `

	_, err := r.db.ExecContext(c, query,
		customer.Name,
		customer.Phone,
		customer.BirthDate,
		customer.Address,
		customer.CustomerType,
		customer.Note,
		customer.ID,
	)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

// 5. Hapus data customer sesuai ID
func (r *sqlCustomerRepository) Delete(c context.Context, id int) error {
	query := `
		DELETE FROM customers WHERE id=?
	`

	_, err := r.db.ExecContext(c, query, id)

	if err != nil {
		return err
	}

	return nil
}
