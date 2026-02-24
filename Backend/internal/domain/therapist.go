package domain

import "time"

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
	Create(therapist *Therapist) error
	GetByID(id int) (*Therapist, error)
	Update(service *Therapist) (*Therapist, error)
	GetAll() ([]Therapist, error)
	Delete(id int) error
}
