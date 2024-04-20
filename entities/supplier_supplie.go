package entities

import "time"

type SupplierSupplie struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	SupplierID int       `json:"supplierId"`
	SupplieID  int       `json:"supplieId"`
	CreatedAt  time.Time `json:"createdAt"`
}

func NewSupplierSupplie(name string, supplierID, supplieID int) *SupplierSupplie {
	date := time.Now().UTC()
	return &SupplierSupplie{
		Name:       name,
		SupplierID: supplierID,
		SupplieID:  supplieID,
		CreatedAt:  date,
	}
}
