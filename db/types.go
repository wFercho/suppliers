package db

import (
	e "suppliers/entities"
)

type Storage interface {
	SupplierStorage
	//SuppliesStorage
}

type SupplierStorage interface {
	CreateSupplier(*e.Supplier) error
	DeleteSupplier(int) error
	UpdateSupplier(*e.Supplier) (*e.Supplier, error)
	GetSupplierByID(int) (*e.Supplier, error)
	// FilterSuppliersByName(string) (*Supplier, error)
	// GetSuppliersByInputID(int) (*[]Supplier, error)
	// GetSuppliersByInputName(string) (*[]Supplier, error)
	// GetSuppliersByProductID(int) (*[]Supplier, error)
	// GetSuppliersByProductName(string) (*[]Supplier, error)
}

type SuppliesStorage interface {
	CreateSupplie(*e.Supplie) error
	DeleleteSupplie(int) error
	UpdateSupplie(int) error
	GetSupplieByID(int) (*e.Supplie, error)
	FilterSuppliesByName(string) (*[]e.Supplie, error)
}
