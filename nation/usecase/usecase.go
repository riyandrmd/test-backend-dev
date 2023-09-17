package usecase

import (
	"encoding/json"
	"net/http"
	"test-backend-dev/entity"
	"test-backend-dev/nation"

	"github.com/gorilla/mux"
)

func NewUsecase(repo nation.RepoNation) usecaseHandle {
	return usecaseHandle{repo}
}

type usecaseHandle struct {
	repo nation.RepoNation
}

func (use usecaseHandle) GetNation(w http.ResponseWriter, r *http.Request) ([]entity.Nation, int, error) {
	res, code, err := use.repo.GetNation()

	if err != nil {
		return res, code, err
	}

	return res, code, err
}

func (use usecaseHandle) PostNation(w http.ResponseWriter, r *http.Request) (int, error) {
	var data entity.Nation

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return http.StatusInternalServerError, err
	}

	code, err := use.repo.PostNation(data)

	if err != nil {
		return code, err
	}

	return code, err
}

func (use usecaseHandle) UpdateNation(w http.ResponseWriter, r *http.Request) (int, error) {
	param := mux.Vars(r)
	var data entity.Nation

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return http.StatusInternalServerError, err
	}

	code, err := use.repo.UpdateNation(param["id"], data)

	if err != nil {
		return code, err
	}

	return code, err
}

func (use usecaseHandle) DeleteNation(w http.ResponseWriter, r *http.Request) (int, error) {
	param := mux.Vars(r)

	code, err := use.repo.DeleteNation(param["id"])

	if err != nil {
		return code, err
	}

	return code, err
}
