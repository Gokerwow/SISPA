package usecase

import (
	"context"
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
func (u *TherapistUsecase) RegisterNewTherapist(c context.Context, therapist *domain.Therapist) error {
	if therapist.Name == "" {
		return errors.New("Therapist name cannot be empty")
	}

	if therapist.JoinedDate == nil || *therapist.JoinedDate == "" {
		return errors.New("Therapist joined date cannot be empty")
	}

	err := u.repo.Create(c, therapist)

	if err != nil {
		return err
	}

	return nil
}

// ambil semua data terapis
func (u *TherapistUsecase) GetAllTherapist(c context.Context) ([]domain.Therapist, error) {

	therapists, err := u.repo.GetAll(c)

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
func (u *TherapistUsecase) GetByID(c context.Context, id int) (*domain.Therapist, error) {

	therapist, err := u.repo.GetByID(c, id)

	if err != nil {
		return nil, err
	}

	return therapist, nil
}

// update data terapis sesuai ID
func (u *TherapistUsecase) Update(c context.Context, therapist *domain.Therapist) (*domain.Therapist, error) {

	therapist, err := u.repo.Update(c, therapist)

	if err != nil {
		return nil, err
	}

	return therapist, nil
}

// hapus data terapis sesuai ID
func (u *TherapistUsecase) Delete(c context.Context, id int) error {

	err := u.repo.Delete(c, id)

	if err != nil {
		return err
	}

	return nil
}
