package domain

import (
	"context"
	"time"
)

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
	Create(c context.Context, service *Service) error
	GetByID(c context.Context, id int) (*Service, error)
	Update(c context.Context, service *Service) (*Service, error)
	GetAll(c context.Context) ([]Service, error)
	Delete(c context.Context, id int) error
}

type ServiceUsecase interface {
	RegisterNewService(c context.Context, service *Service) error
	GetByID(c context.Context, id int) (*Service, error)
	Update(c context.Context, service *Service) (*Service, error)
	GetAll(c context.Context) ([]Service, error)
	Delete(c context.Context, id int) error
}
