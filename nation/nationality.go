package nation

import (
	"net/http"
	"test-backend-dev/entity"
)

type UsecaseNation interface {
	GetNation(http.ResponseWriter, *http.Request) ([]entity.Nation, int, error)
	PostNation(http.ResponseWriter, *http.Request) (int, error)
	UpdateNation(http.ResponseWriter, *http.Request) (int, error)
	DeleteNation(http.ResponseWriter, *http.Request) (int, error)
}

type RepoNation interface {
	GetNation() ([]entity.Nation, int, error)
	PostNation(entity.Nation) (int, error)
	UpdateNation(string, entity.Nation) (int, error)
	DeleteNation(string) (int, error)
}
