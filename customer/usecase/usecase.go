package usecase

import (
	"encoding/json"
	"net/http"
	"test-backend-dev/customer"
	"test-backend-dev/entity"

	"github.com/gorilla/mux"
)

func NewUsecase(repo customer.RepoCustomer) usecaseHandle {
	return usecaseHandle{repo}
}

type usecaseHandle struct {
	repo customer.RepoCustomer
}

func (use usecaseHandle) GetCustomer(w http.ResponseWriter, r *http.Request) ([]entity.Customer, int, error) {
	res, code, err := use.repo.GetCustomer()

	if err != nil {
		return res, code, err
	}

	return res, code, err
}

func (use usecaseHandle) PostCustomer(w http.ResponseWriter, r *http.Request) (int, error) {
	var data entity.Customer

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return http.StatusInternalServerError, err
	}

	code, err := use.repo.PostCustomer(data)

	if err != nil {
		return code, err
	}

	return code, err
}

func (use usecaseHandle) UpdateCustomer(w http.ResponseWriter, r *http.Request) (int, error) {
	param := mux.Vars(r)
	var data entity.Customer

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return http.StatusInternalServerError, err
	}

	code, err := use.repo.UpdateCustomer(param["id"], data)

	if err != nil {
		return code, err
	}

	return code, err
}

func (use usecaseHandle) DeleteCustomer(w http.ResponseWriter, r *http.Request) (int, error) {
	param := mux.Vars(r)

	code, err := use.repo.DeleteCustomer(param["id"])

	if err != nil {
		return code, err
	}

	return code, err
}
