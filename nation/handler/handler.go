package handler

import (
	"net/http"
	"test-backend-dev/helper"
	"test-backend-dev/nation"

	"github.com/gorilla/mux"
)

func NewHandler(use nation.UsecaseNation, r *mux.Router) {
	eng := handlerNation{use}

	r.HandleFunc("/nation", eng.GetNation).Methods("GET")
	r.HandleFunc("/nation", eng.PostNation).Methods("POST")
	r.HandleFunc("/nation/{id}", eng.UpdateNation).Methods("PUT")
	r.HandleFunc("/nation/{id}", eng.DeleteNation).Methods("DELETE")
}

type handlerNation struct {
	usecase nation.UsecaseNation
}

func (h handlerNation) GetNation(w http.ResponseWriter, r *http.Request) {
	res, code, err := h.usecase.GetNation(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": res})
}

func (h handlerNation) PostNation(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.PostNation(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": "Succes Post Data"})
}

func (h handlerNation) UpdateNation(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.UpdateNation(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": "Succes Update Data"})
}

func (h handlerNation) DeleteNation(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.DeleteNation(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succses": "Succses Delete Data"})
}
