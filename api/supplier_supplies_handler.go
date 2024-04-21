package api

import (
	"encoding/json"
	"net/http"
	e "suppliers/entities"
)

func (s *APIServer) handleAddSupplierSupplie(w http.ResponseWriter, r *http.Request) error {

	createSSrequest := new(CreateSupplierSupplieRequest)

	if err := json.NewDecoder(r.Body).Decode(createSSrequest); err != nil {
		return err
	}

	defer r.Body.Close()

	ss := e.NewSupplierSupplie(createSSrequest.Name, createSSrequest.SupplierID, createSSrequest.SupplieID)
	if err := s.store.CreateSupplierSupplie(ss); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, ss)
}
