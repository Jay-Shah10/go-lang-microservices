package domain

import "github.com/Jay-Shah10/go-lang-microservices/banking/err"

// defining the domain of the hexagonal arch.
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// Defining the interface to be used with the domain. Adding the repository.
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *err.AppErrors) // options: status=1 == active, status=0 == inactive, and status=''
	ById(id string) (*Customer, *err.AppErrors)         // this is a pointer because we want to send nil if the customer does not exist.

}
