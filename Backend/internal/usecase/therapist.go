package usecase

import (
	"errors"
	"sispa-backend/internal/domain"
)

type TherapistUsecase struct {
	repo domain.TherapistRepository
}

func NewTherapistUsecase(r domain.TherapistRepository) *TherapistUsecase {
	return &TherapistUsecase{repo: r}
}

// daftarin terapis baru
func (u *TherapistUsecase) RegisterNewTherapist(therapist *domain.Therapist) error {
	if therapist.Name == "" {
		return errors.New("Therapist name cannot be empty")
	}

	if therapist.JoinedDate == nil || *therapist.JoinedDate == "" {
		return errors.New("Therapist joined date cannot be empty")
	}

	err := u.repo.Create(therapist)

	if err != nil {
		return err
	}

	return nil
}

// ambil semua data terapis
func (u *TherapistUsecase) GetAll() ([]domain.Therapist, error) {

	therapists, err := u.repo.GetAll()

	if err != nil {
		return nil, err
	}

	if therapists == nil {
		therapists = []domain.Therapist{}
		return therapists, nil
	}

	return therapists, nil
}

// ambil data terapis sesuai ID
func (u *TherapistUsecase) GetByID(id int) (*domain.Therapist, error) {

	therapist, err := u.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return therapist, nil
}

// update data terapis sesuai ID
func (u *TherapistUsecase) Update(therapist *domain.Therapist) (*domain.Therapist, error) {

	therapist, err := u.repo.Update(therapist)

	if err != nil {
		return nil, err
	}

	return therapist, nil
}

// hapus data terapis sesuai ID
func (u *TherapistUsecase) Delete(id int) error {

	err := u.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
