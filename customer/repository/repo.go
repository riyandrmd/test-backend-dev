package repository

import (
	"net/http"
	"strconv"
	"test-backend-dev/entity"

	"gorm.io/gorm"
)

type repoHandler struct {
	db *gorm.DB
}

func NewRepo(Db *gorm.DB) repoHandler {
	return repoHandler{Db}
}

func (repo repoHandler) GetCustomer() (data []entity.Customer, code int, err error) {
	if err = repo.db.Preload("Keluarga").Find(&data).Error; err != nil {
		code = 404
		return
	}
	code = 200
	return
}

func (repo repoHandler) PostCustomer(data entity.Customer) (int, error) {
	var nation entity.Nation

	if err := repo.db.Where("negara = ?", data.Kewarganegaraan).First(&nation).Error; err != nil {
		return http.StatusBadRequest, err
	}

	data.Nationality_id = nation.Nationality_id

	if err := repo.db.Create(&data).Error; err != nil {
		return http.StatusBadRequest, err
	}

	return 200, nil
}

func (repo repoHandler) UpdateCustomer(key string, data entity.Customer) (int, error) {
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		return 500, err
	}

	if err := repo.db.First(&entity.Customer{}, "cst_id = ?", int(keyInt)).Error; err != nil {
		return 404, err
	}

	if err := repo.db.Where("cst_id = ?", uint(keyInt)).Updates(&data).Error; err != nil {
		return 400, err
	}

	return 200, nil
}

func (repo repoHandler) DeleteCustomer(key string) (int, error) {
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		return 500, err
	}

	if err := repo.db.First(&entity.Customer{}, "cst_id = ?", int(keyInt)).Error; err != nil {
		return 404, err
	}

	if err := repo.db.Where("cst_id = ?", int(keyInt)).Delete(entity.Customer{}).Error; err != nil {
		return 500, err
	}

	return 200, nil
}
