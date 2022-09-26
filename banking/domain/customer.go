package domain

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
	FindAll() ([]Customer, error)
}
