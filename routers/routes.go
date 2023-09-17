package routers

import (
	handlerC "test-backend-dev/customer/handler"
	repoCst "test-backend-dev/customer/repository"
	ucCst "test-backend-dev/customer/usecase"
	handlerF "test-backend-dev/family/handler"
	repoFm "test-backend-dev/family/repository"
	ucFamily "test-backend-dev/family/usecase"
	handlerN "test-backend-dev/nation/handler"
	repoNation "test-backend-dev/nation/repository"
	ucNation "test-backend-dev/nation/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Routers struct {
	R  *mux.Router
	Db *gorm.DB
}

func (r Routers) NewRouters() {
	eng := r.R.PathPrefix("/api").Subrouter()

	repo1 := repoNation.NewRepo(r.Db)
	UC1 := ucNation.NewUsecase(repo1)
	handlerN.NewHandler(UC1, eng)

	repo2 := repoCst.NewRepo(r.Db)
	UC2 := ucCst.NewUsecase(repo2)
	handlerC.NewHandler(UC2, eng)

	repo3 := repoFm.NewRepo(r.Db)
	UC3 := ucFamily.NewUsecase(repo3)
	handlerF.NewHandler(UC3, eng)
}
