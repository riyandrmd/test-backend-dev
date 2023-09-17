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

func (repo repoHandler) GetFamily() (data []entity.Family_list, code int, err error) {
	if err = repo.db.Find(&data).Error; err != nil {
		code = 404
		return
	}
	code = 200
	return
}

func (repo repoHandler) PostFamily(data entity.Family_list) (int, error) {
	if err := repo.db.Create(&data).Error; err != nil {
		return 400, err
	}

	return 200, nil
}

func (repo repoHandler) UpdateFamily(key string, data entity.Family_list) (int, error) {
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		return 500, err
	}

	if err := repo.db.First(&entity.Family_list{}, "fl_id = ?", int(keyInt)).Error; err != nil {
		return 404, err
	}

	if err := repo.db.Where("fl_id = ?", uint(keyInt)).Updates(&data).Error; err != nil {
		return 400, err
	}

	return 200, nil
}

func (repo repoHandler) DeleteFamily(key string) (int, error) {
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		return 500, err
	}

	if err := repo.db.First(&entity.Family_list{}, "fl_id = ?", int(keyInt)).Error; err != nil {
		return 404, err
	}

	if err := repo.db.Where("fl_id = ?", int(keyInt)).Delete(entity.Family_list{}).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return 200, nil
}
