package handler

import (
	"encoding/json"
	"net/http"
	"user_management/service"
)

type NationalityHandler struct {
	service service.NationalityService
}

func NewNationalityHandler(service service.NationalityService) *NationalityHandler {
	return &NationalityHandler{service}
}

func (h *NationalityHandler) GetAllNationalities(w http.ResponseWriter, r *http.Request) {
	nationalities, err := h.service.GetAllNationalities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to JSON response
	json.NewEncoder(w).Encode(nationalities)
}

type CustomerHandler struct {
	service service.CustomersService
}

func NewCustomerHandler(service service.CustomersService) *CustomerHandler {
	return &CustomerHandler{service}
}

func (h *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to JSON response
	json.NewEncoder(w).Encode(customers)
}

type FamilyHandler struct {
	service service.FamilyService
}

func NewFamilyHandler(service service.FamilyService) *FamilyHandler {
	return &FamilyHandler{service}
}

func (h *FamilyHandler) GetAllFamilies(w http.ResponseWriter, r *http.Request) {
	families, err := h.service.GetAllFamilies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to JSON response
	json.NewEncoder(w).Encode(families)
}
