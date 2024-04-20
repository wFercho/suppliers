package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	e "suppliers/entities"
)

func (s *APIServer) handleCreateSupplier(w http.ResponseWriter, r *http.Request) error {
	createSupplierReq := new(CreateSupplierRequest)
	//createSupplierReq := CreateSupplierRequest{}
	//if err := json.NewDecoder(r.Body).Decode(&createSupplierReq); err != nil {
	if err := json.NewDecoder(r.Body).Decode(createSupplierReq); err != nil {
		return err
	}

	defer r.Body.Close()

	supplier := e.NewSupplier(createSupplierReq.Name, createSupplierReq.Emails, createSupplierReq.PhoneNumbers)

	if err := s.store.CreateSupplier(supplier); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusCreated, supplier)
}

func (s *APIServer) handleDeleteSupplier(w http.ResponseWriter, r *http.Request) error {
	_id := r.PathValue("id")

	if _id == "" {
		return fmt.Errorf("provide a validate id")
	}

	id, err := strconv.Atoi(_id)

	if err != nil {
		return err
	}
	if err := s.store.DeleteSupplier(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, id)
}

func (s *APIServer) handleGetSupplierByID(w http.ResponseWriter, r *http.Request) error {
	_id := r.PathValue("id")

	//fmt.Printf("id provided %s\n", _id)
	// if _id == "" {
	// 	return fmt.Errorf("provide a validate id")
	// }

	id, err := strconv.Atoi(_id)

	if err != nil {
		return fmt.Errorf("provide a validate id")
	}

	sup, err := s.store.GetSupplierByID(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, sup)
}

func (s *APIServer) handleUpdateSupplier(w http.ResponseWriter, r *http.Request) error {
	updateSupplierReq := new(e.Supplier)
	//createSupplierReq := CreateSupplierRequest{}
	//if err := json.NewDecoder(r.Body).Decode(&createSupplierReq); err != nil {
	if err := json.NewDecoder(r.Body).Decode(updateSupplierReq); err != nil {
		return err
	}
	defer r.Body.Close()

	//supplier := e.NewSupplier(createSupplierReq.Name, createSupplierReq.Emails, createSupplierReq.PhoneNumbers)

	sup, err := s.store.UpdateSupplier(updateSupplierReq)

	if err != nil {
		return err

	}
	return WriteJSON(w, http.StatusOK, sup)

}

func (s *APIServer) handleFilterSuppliers(w http.ResponseWriter, r *http.Request) error {
	qName := r.URL.Query().Get("name")
	qSupplieID := r.URL.Query().Get("supplie_id")
	//qSupplieName := r.URL.Query().Get("supplie_name")
	// qProductID := r.URL.Query().Get("product_id")
	// qProductName := r.URL.Query().Get("product_name")

	/*
		supplier name
		supplie id
		supplie name
		product id
		product name
	*/
	if qSupplieID != "" {
		//suppliers, err := s.store.FilterSuppliersBySupplieID()
	}

	if qName != "" {
		suppliers, err := s.store.FilterSuppliersByName(qName)

		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, suppliers)
	}

	return nil
}
