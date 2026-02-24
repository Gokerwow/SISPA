package domain

import "time"

// HARUS UPPER KLLO MAU EXPORT
type Service struct {
	ID           int        `json:"id"`
	ServiceName  string     `json:"name"`
	PriceMale    int        `json:"price_male"`
	PriceFemale  int        `json:"price_female"`
	PriceTourist int        `json:"price_tourist"`
	IsActive     bool       `json:"is_active,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
}

type ServiceRepository interface {
	Create(service *Service) error
	GetByID(id int) (*Service, error)
	Update(service *Service) (*Service, error)
	GetAll() ([]Service, error)
	Delete(id int) error
}
