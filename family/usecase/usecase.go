package usecase

import (
	"encoding/json"
	"net/http"
	"test-backend-dev/entity"
	"test-backend-dev/family"

	"github.com/gorilla/mux"
)

func NewUsecase(repo family.RepoFamily) usecaseHandle {
	return usecaseHandle{repo}
}

type usecaseHandle struct {
	repo family.RepoFamily
}

func (use usecaseHandle) GetFamily(w http.ResponseWriter, r *http.Request) ([]entity.Family_list, int, error) {
	res, code, err := use.repo.GetFamily()

	if err != nil {
		return res, code, err
	}

	return res, code, err
}

func (use usecaseHandle) PostFamily(w http.ResponseWriter, r *http.Request) (int, error) {
	var data entity.Family_list

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return http.StatusInternalServerError, err
	}

	code, err := use.repo.PostFamily(data)

	if err != nil {
		return code, err
	}

	return code, err
}

func (use usecaseHandle) UpdateFamily(w http.ResponseWriter, r *http.Request) (int, error) {
	param := mux.Vars(r)
	var data entity.Family_list

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return http.StatusInternalServerError, err
	}

	code, err := use.repo.UpdateFamily(param["id"], data)

	if err != nil {
		return code, err
	}

	return code, err
}

func (use usecaseHandle) DeleteFamily(w http.ResponseWriter, r *http.Request) (int, error) {
	param := mux.Vars(r)

	code, err := use.repo.DeleteFamily(param["id"])

	if err != nil {
		return code, err
	}

	return code, err
}
