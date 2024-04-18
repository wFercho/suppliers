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

type APIServer struct {
	listenAddr string
	store      db.Storage
}
