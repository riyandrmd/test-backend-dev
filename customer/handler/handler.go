package handler

import (
	"net/http"
	"test-backend-dev/customer"
	"test-backend-dev/helper"

	"github.com/gorilla/mux"
)

func NewHandler(use customer.UsecaseCustomer, r *mux.Router) {
	eng := handlerCustomer{use}

	r.HandleFunc("/customer", eng.GetCustomer).Methods("GET")
	r.HandleFunc("/customer", eng.PostCustomer).Methods("POST")
	r.HandleFunc("/customer/{id}", eng.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customer/{id}", eng.DeleteCustomer).Methods("DELETE")
}

type handlerCustomer struct {
	usecase customer.UsecaseCustomer
}

func (h handlerCustomer) GetCustomer(w http.ResponseWriter, r *http.Request) {
	res, code, err := h.usecase.GetCustomer(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": res})
}

func (h handlerCustomer) PostCustomer(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.PostCustomer(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": "Succes Post Data"})
}

func (h handlerCustomer) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.UpdateCustomer(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succes": "Succes Update Data"})
}

func (h handlerCustomer) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	code, err := h.usecase.DeleteCustomer(w, r)

	if err != nil {
		helper.Response(w, code, map[string]any{"Error": err.Error()})
		return
	}

	helper.Response(w, code, map[string]any{"Succses": "Succses Delete Data"})
}
