package family

import (
	"net/http"
	"test-backend-dev/entity"
)

type UsecaseFamily interface {
	GetFamily(http.ResponseWriter, *http.Request) ([]entity.Family_list, int, error)
	PostFamily(http.ResponseWriter, *http.Request) (int, error)
	UpdateFamily(http.ResponseWriter, *http.Request) (int, error)
	DeleteFamily(http.ResponseWriter, *http.Request) (int, error)
}

type RepoFamily interface {
	GetFamily() ([]entity.Family_list, int, error)
	PostFamily(entity.Family_list) (int, error)
	UpdateFamily(string, entity.Family_list) (int, error)
	DeleteFamily(string) (int, error)
}
