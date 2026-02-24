package usecase

import (
	"errors"
	"sispa-backend/internal/domain"
)

type ServiceUsecase struct {
	repo domain.ServiceRepository
}

func NewServiceUsecase(r domain.ServiceRepository) *ServiceUsecase {
	return &ServiceUsecase{repo: r}
}

// TODO: tambahin ctx := c.Request.Context()
// Locgic buat daftar layanan baru
// kalo nama + harga kosong maka akan return error
func (u *ServiceUsecase) RegisterNewService(service *domain.Service) error {
	if service.ServiceName == "" {
		return errors.New("service name cannot be empty")
	}

	if service.PriceMale == 0 && service.PriceFemale == 0 && service.PriceTourist == 0 {
		return errors.New("cannot register the service: at least one price must be filled")
	}

	err := u.repo.Create(service)
	if err != nil {
		return err
	}

	return nil
}

// buat ambil semua data layanan
// klo kosong return slices []
func (u *ServiceUsecase) GetAll() ([]domain.Service, error) {

	services, err := u.repo.GetAll()

	if err != nil {
		return nil, err
	}

	if services == nil {
		services = []domain.Service{}
		return services, nil
	}

	return services, nil
}

// buat ambil data layanan sesuai ID
// klo gada return nil ajah
func (u *ServiceUsecase) GetByID(id int) (*domain.Service, error) {

	services, err := u.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return services, nil
}

// update data layanan sesuai ID
func (u *ServiceUsecase) Update(service *domain.Service) (*domain.Service, error) {

	services, err := u.repo.Update(service)

	if err != nil {
		return nil, err
	}

	return services, nil
}

// hapus data layanan sesuai ID
func (u *ServiceUsecase) Delete(id int) error {

	err := u.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
