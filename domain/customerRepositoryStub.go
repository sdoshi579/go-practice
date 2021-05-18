package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) FindById(id string) (*Customer, error) {
	return &s.customers[0], nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Sharukh", "Mumbai", "400086", "1980-01-01", "1"},
		{"2", "Amitabh", "Mumbai", "400086", "1980-01-01", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}