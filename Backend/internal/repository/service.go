package repository

import (
	"context"
	"database/sql"
	"sispa-backend/internal/domain"
)

type sqlServiceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) domain.ServiceRepository {
	return &sqlServiceRepository{db: db}
}

// BUAT LAYANAN BARU
func (r *sqlServiceRepository) Create(c context.Context, service *domain.Service) error {
	query := `
		INSERT INTO services (name, price_male, price_female, price_tourist, is_active)
		VALUES (?,?,?,?,?)
	`
	_, err := r.db.ExecContext(c, query, service.ServiceName, service.PriceMale, service.PriceFemale, service.PriceTourist, service.IsActive)

	return err
}

// 2. AMBIL SEMUA DATA LAYANAN
func (r *sqlServiceRepository) GetAll(c context.Context) ([]domain.Service, error) {
	query := `
		SELECT id, name, price_male, price_female, price_tourist, is_active, created_at FROM services
	`

	rows, err := r.db.QueryContext(c, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	services := []domain.Service{}

	for rows.Next() {
		var c domain.Service
		err := rows.Scan(&c.ID, &c.ServiceName, &c.PriceMale, &c.PriceFemale, &c.PriceTourist, &c.IsActive, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}

// 3. DAPETIN DATA SESUAI ID
func (r *sqlServiceRepository) GetByID(c context.Context, id int) (*domain.Service, error) {
	query := `
		SELECT id, name, price_male, price_female, price_tourist, is_active, created_at FROM services WHERE id=?
	`

	var s domain.Service

	err := r.db.QueryRowContext(c, query, id).Scan(&s.ID, &s.ServiceName, &s.PriceMale, &s.PriceFemale, &s.PriceTourist, &s.IsActive, &s.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &s, nil
}

// 4. UPDATE DATA SESUAI ID
func (r *sqlServiceRepository) Update(c context.Context, service *domain.Service) (*domain.Service, error) {
	query := `
        UPDATE services 
        SET name = ?, price_male = ?, price_female = ?, price_tourist = ?, is_active = ?
        WHERE id = ?
    `

	_, err := r.db.ExecContext(c, query,
		service.ServiceName,
		service.PriceMale,
		service.PriceFemale,
		service.PriceTourist,
		service.IsActive,
		service.ID,
	)

	if err != nil {
		return nil, err
	}

	return service, nil
}

// 5. DELETE DATA LAYANAN SESUAI ID
func (r *sqlServiceRepository) Delete(c context.Context, id int) error {
	query := `
		DELETE FROM services WHERE id=?
	`

	_, err := r.db.ExecContext(c, query, id)

	if err != nil {
		return err
	}

	return nil
}
