package api

import (
	"log"
	"net/http"

	"suppliers/db"
)

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string, store db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {

	sMux := http.NewServeMux()

	sMux.HandleFunc("GET /supplier/{id}", makeHTTPHandleFunc(s.handleGetSupplierByID))
	sMux.HandleFunc("POST /supplier", makeHTTPHandleFunc(s.handleCreateSupplier))
	sMux.HandleFunc("PATCH /supplier", makeHTTPHandleFunc(s.handleUpdateSupplier))
	sMux.HandleFunc("DELETE /supplier/{id}", makeHTTPHandleFunc(s.handleDeleteSupplier))

	sMux.HandleFunc("GET /supplier", makeHTTPHandleFunc(s.handleFilterSuppliers))

	log.Println("Suppliers API running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, sMux)
}
