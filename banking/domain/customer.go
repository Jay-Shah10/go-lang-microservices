package domain

import "github.com/Jay-Shah10/go-lang-microservices/banking/err"

// defining the domain of the hexagonal arch.
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirty string
	Status      string
}

// Defining the interface to be used with the domain. Adding the repository.
type CustomerRepository interface {
	FindAll() ([]Customer, *err.AppErrors)
	ById(string) (*Customer, *err.AppErrors) // this is a pointer because we want to send nil if the customer does not exist.
}
