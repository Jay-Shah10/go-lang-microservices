package service

import (
	"github.com/Jay-Shah10/go-lang-microservices/banking/domain"
	"github.com/Jay-Shah10/go-lang-microservices/banking/err"
)

// This is our primary port. Since we know that all ports are interfaces, this is an interface.
// Goal is to connect our primary port to the secondary port, which is the customerRepositoryStub.
type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *err.AppErrors)
	GetCustomer(id string) (*domain.Customer, *err.AppErrors)
}

// service implementation for our port. This is the Business Logic.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// receiver function.
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *err.AppErrors) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

// receiver function.
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *err.AppErrors) {
	return s.repo.ById(id)
}

// helper function for DefaultCustomerService.
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

// This whole thing now connects our primary port to the secondary port.
// From this to the customerRepoStub. Service port is the primary port.
