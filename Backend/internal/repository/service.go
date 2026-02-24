package repository

import (
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
func (r *sqlServiceRepository) Create(service *domain.Service) error {
	query := `
		INSERT INTO services (name, price_male, price_female, price_tourist, is_active)
		VALUES (?,?,?,?,?)
	`
	_, err := r.db.Exec(query, service.ServiceName, service.PriceMale, service.PriceFemale, service.PriceTourist, service.IsActive)

	return err
}

// 2. AMBIL SEMUA DATA LAYANAN
func (r *sqlServiceRepository) GetAll() ([]domain.Service, error) {
	query := `
		SELECT id, name, price_male, price_female, price_tourist, is_active, created_at FROM services
	`

	rows, err := r.db.Query(query)

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
func (r *sqlServiceRepository) GetByID(id int) (*domain.Service, error) {
	query := `
		SELECT id, name, price_male, price_female, price_tourist, is_active, created_at FROM services WHERE id=?
	`

	var c domain.Service

	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.ServiceName, &c.PriceMale, &c.PriceFemale, &c.PriceTourist, &c.IsActive, &c.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &c, nil
}

// 4. UPDATE DATA SESUAI ID
func (r *sqlServiceRepository) Update(service *domain.Service) (*domain.Service, error) {
	query := `
        UPDATE services 
        SET name = ?, price_male = ?, price_female = ?, price_tourist = ?, is_active = ?
        WHERE id = ?
    `

	_, err := r.db.Exec(query,
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
func (r *sqlServiceRepository) Delete(id int) error {
	query := `
		DELETE FROM services WHERE id=?
	`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
