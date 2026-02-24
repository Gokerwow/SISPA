package server

import (
	"database/sql"
	"sispa-backend/internal/delivery/http"
	"sispa-backend/internal/repository"
	"sispa-backend/internal/usecase"
)

type AppHandlers struct {
	CustomerHandler  *http.CustomerHandler
	ServiceHandler   *http.ServiceHandler
	TherapistHandler *http.TherapistHandler
}

func InitHandlers(db *sql.DB) *AppHandlers {
	customerRepo := repository.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	customerHandler := http.NewCustomerHandler(customerUsecase)

	serviceRepo := repository.NewServiceRepository(db)
	serviceUsecase := usecase.NewServiceUsecase(serviceRepo)
	serviceHandler := http.NewServiceHandler(serviceUsecase)

	therapistRepo := repository.NewTherapistRepository(db)
	therapistUsecase := usecase.NewTherapistUsecase(therapistRepo)
	therapistHandler := http.NewTherapistHandler(therapistUsecase)

	return &AppHandlers{CustomerHandler: customerHandler, ServiceHandler: serviceHandler, TherapistHandler: therapistHandler}
}
