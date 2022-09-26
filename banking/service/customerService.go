package service

import (
	"github.com/Jay-Shah10/go-lang-microservices/banking/domain"
)

// This is our primary port. Since we know that all ports are interfaces, this is an interface.
// Goal is to connect our primary port to the secondary port, which is the customerRepositoryStub.
type CustomerService interface {
	getAllCustomer() ([]domain.Customer, error)
}

// servie implementation for our port. This is the Business Logic.
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// receiver function.
func (s DefaultCustomerService) getAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// helper function for DefaultCustomerService.
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}