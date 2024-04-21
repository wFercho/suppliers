package api

import (
	"net/http"
	"suppliers/db"
)

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type CreateSupplierRequest struct {
	Name         string   `json:"name"`
	Emails       []string `json:"emails"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

type UpdateSupplierRequest struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Emails       *[]string `json:"emails"`
	PhoneNumbers *[]string `json:"phoneNumbers"`
}

type CreateSupplieRequest struct {
	Name string `json:"name"`
}

type UpdateSupplieRequest struct {
	Name string `json:"name"`
}

type CreateSupplierSupplieRequest struct {
	Name       string `json:"name"`
	SupplierID int    `json:"supplierId"`
	SupplieID  int    `json:"supplieId"`
}

type APIServer struct {
	listenAddr string
	store      db.Storage
}
