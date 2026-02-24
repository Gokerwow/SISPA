package usecase

import (
	"context"
	"errors"
	"sispa-backend/internal/domain"
	"time"
)

type CustomerUsecase struct {
	repo domain.CustomerRepository
}

func NewCustomerUsecase(r domain.CustomerRepository) domain.CustomerUsecase {
	return &CustomerUsecase{repo: r}
}

// Logic buat daftarin customer baru
// klo nomor hp kosng gabis ditambah
// dicek lewat nomer hp, kalau ada berarti customer sudah terdaftar
func (u *CustomerUsecase) RegisterNewCustomer(ctx context.Context, customer *domain.Customer) error {
	if customer.Phone == "" {
		return errors.New("cannot register: phone number is mandatory for spa bookings")
	}
	if customer.Name == "" {
		return errors.New("cannot register: customer name is mandatory for spa bookings")
	}

	existed, _ := u.repo.GetByPhone(ctx, customer.Phone)
	if existed != nil {
		return errors.New("cannot register: customer already existed")
	}

	customer.CreatedAt = time.Now()
	error := u.repo.Create(ctx, customer)

	if error != nil {
		return error
	}

	return nil
}

// ambil semua data customer
func (u *CustomerUsecase) GetAll(ctx context.Context) ([]domain.Customer, error) {
	customers, err := u.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

// ambil data customer sesuai ID
func (u *CustomerUsecase) GetByID(ctx context.Context, id int) (*domain.Customer, error) {
	customers, err := u.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

// ambil data customer sesuai ID
func (u *CustomerUsecase) GetByPhone(ctx context.Context, phone string) (*domain.Customer, error) {
	customers, err := u.repo.GetByPhone(ctx, phone)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

// update data customer sesuai ID
func (u *CustomerUsecase) UpdateCustomer(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	customers, err := u.repo.Update(ctx, customer)

	if err != nil {
		return nil, err
	}

	return customers, nil
}

// hapus data customer sesuai ID
func (u *CustomerUsecase) DeleteCustomer(ctx context.Context, id int) error {
	err := u.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
