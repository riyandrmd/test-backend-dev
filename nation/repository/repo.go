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

func (repo repoHandler) GetNation() (data []entity.Nation, code int, err error) {
	if err = repo.db.Preload("Customer").Find(&data).Error; err != nil {
		code = 404
		return
	}
	code = 200
	return
}

func (repo repoHandler) PostNation(data entity.Nation) (int, error) {
	if err := repo.db.Create(&data).Error; err != nil {
		return http.StatusBadRequest, err
	}

	return 200, nil
}

func (repo repoHandler) UpdateNation(key string, data entity.Nation) (int, error) {
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		return 500, err
	}

	if err := repo.db.First(&entity.Nation{}, "nationality_id = ?", int(keyInt)).Error; err != nil {
		return 404, err
	}

	if err := repo.db.Where("nationality_id = ?", uint(keyInt)).Updates(&data).Error; err != nil {
		return http.StatusBadRequest, err
	}

	return 200, nil
}

func (repo repoHandler) DeleteNation(key string) (int, error) {
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		return 500, err
	}

	if err := repo.db.First(&entity.Nation{}, "nationality_id = ?", int(keyInt)).Error; err != nil {
		return 404, err
	}

	if err := repo.db.Where("nationality_id = ?", int(keyInt)).Delete(entity.Nation{}).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return 200, nil
}
