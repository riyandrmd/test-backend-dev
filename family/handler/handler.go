package handler

import (
	"net/http"
	"test-backend-dev/family"
	"test-backend-dev/helper"

	"github.com/gorilla/mux"
)

func NewHandler(use family.UsecaseFamily, r *mux.Router) {
	eng := handlerFamily{use}

	r.HandleFunc("/family", eng.GetFamily).Methods("GET")
	r.HandleFunc("/family", eng.PostFamily).Methods("POST")
	r.HandleFunc("/family/{id}", eng.UpdateFamily).Methods("PUT")
	r.HandleFunc("/family/{id}", eng.DeleteFamily).Methods("DELETE")
}

type handlerFamily struct {
	usecase family.UsecaseFamily
}

func (h handlerFamily) GetFamily(w http.ResponseWriter, r *http.Request) {
	res, code, err := h.usecase.GetFamily(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": res})
}

func (h handlerFamily) PostFamily(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.PostFamily(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": "Succes Post Data"})
}

func (h handlerFamily) UpdateFamily(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.UpdateFamily(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": "Succes Update Data"})
}

func (h handlerFamily) DeleteFamily(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.DeleteFamily(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succses": "Succses Delete Data"})
}
