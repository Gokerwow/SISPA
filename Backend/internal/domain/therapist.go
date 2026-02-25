package domain

import (
	"context"
	"time"
)

type Therapist struct {
	ID             int        `json:"id"`
	Name           string     `json:"name" binding:"required"`
	Phone          string     `json:"phone" binding:"required"`
	Address        string     `json:"address" binding:"required"`
	CommissionRate int        `json:"commission_rate,omitempty"` // Renamed for clarity
	JoinedDate     *string    `json:"joined_date"`               // Using string for easy Vue integration
	ExitDate       *string    `json:"exit_date,omitempty"`
	Status         string     `json:"status"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

type TherapistRepository interface {
	Create(c context.Context, therapist *Therapist) error
	GetByID(c context.Context, id int) (*Therapist, error)
	Update(c context.Context, service *Therapist) (*Therapist, error)
	GetAll(c context.Context) ([]Therapist, error)
	Delete(c context.Context, id int) error
}

type TherapistUsecase interface {
	RegisterNewTherapist(c context.Context, therapist *Therapist) error
	GetAllTherapist(c context.Context) ([]Therapist, error)
	GetByID(c context.Context, id int) (*Therapist, error)
	Update(c context.Context, service *Therapist) (*Therapist, error)
	Delete(c context.Context, id int) error
}
