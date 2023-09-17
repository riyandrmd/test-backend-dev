package customer

import (
	"net/http"
	"test-backend-dev/entity"
)

type UsecaseCustomer interface {
	GetCustomer(http.ResponseWriter, *http.Request) ([]entity.Customer, int, error)
	PostCustomer(http.ResponseWriter, *http.Request) (int, error)
	UpdateCustomer(http.ResponseWriter, *http.Request) (int, error)
	DeleteCustomer(http.ResponseWriter, *http.Request) (int, error)
}

type RepoCustomer interface {
	GetCustomer() ([]entity.Customer, int, error)
	PostCustomer(entity.Customer) (int, error)
	UpdateCustomer(string, entity.Customer) (int, error)
	DeleteCustomer(string) (int, error)
}
