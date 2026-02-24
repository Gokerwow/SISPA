package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID           int        `json:"id"`   // JSON biar si kode tau kalau namanya nanti di json bakalan nama, cuma nanti untuk structnya pakai yang kiri
	Name         string     `json:"name"` // omitempty, biar kalau jsonnya gada isinya maka gak negembaliin apapun dan ilang, kalau tanpa itu bakalan dikembaliin dengan zero value
	Phone        string     `json:"phone"`
	BirthDate    *string    `json:"birth_date"`
	Address      string     `json:"address"`
	CustomerType string     `json:"customer_type"`
	Note         string     `json:"note,omitempty"`
	TotalSpend   int        `json:"total_spend"`
	TotalVisit   int        `json:"total_visit"`
	LastVisit    *time.Time `json:"last_visit,omitempty"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
}

type CustomerRepository interface {
	Create(c context.Context, customer *Customer) error
	GetAll(c context.Context) ([]Customer, error)
	GetByPhone(c context.Context, phone string) (*Customer, error)
	GetByID(c context.Context, id int) (*Customer, error)
	Update(c context.Context, customer *Customer) (*Customer, error)
	Delete(c context.Context, id int) error
}

type CustomerUsecase interface {
	RegisterNewCustomer(ctx context.Context, customer *Customer) error
	GetAll(c context.Context) ([]Customer, error)
	GetByPhone(c context.Context, phone string) (*Customer, error)
	GetByID(c context.Context, id int) (*Customer, error)
	UpdateCustomer(c context.Context, customer *Customer) (*Customer, error)
	DeleteCustomer(c context.Context, id int) error
}
