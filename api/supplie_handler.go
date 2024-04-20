package api

import (
	"encoding/json"
	"net/http"

	e "suppliers/entities"
)

func (s *APIServer) handleCreateSupplie(w http.ResponseWriter, r *http.Request) error {
	createSupplieReq := new(CreateSupplieRequest)

	if err := json.NewDecoder(r.Body).Decode(createSupplieReq); err != nil {
		return err
	}
	defer r.Body.Close()

	supplie := e.NewSupplie(createSupplieReq.Name)
	if err := s.store.CreateSupplie(supplie); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusCreated, supplie)

}
