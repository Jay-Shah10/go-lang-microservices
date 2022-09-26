package domain

// This is the mock adapter. In this case this will be more of a struct. Will not be mock. stub adapter.
type CustomerRepositoryStub struct {
	customers []Customer
}

// Reciver function for the customerRepositoryStub.
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// helper function to make new customer repo stub.
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Jay", City: "Duluth", Zipcode: "12345", DateOfBirty: "1993-10-24"},
	}
	return CustomerRepositoryStub{customers}
}
