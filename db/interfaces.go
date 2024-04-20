package db

import (
	e "suppliers/entities"
)

type Storage interface {
	SupplierStorage
	SuppliesStorage
	SupplierSuppliesStorage
}

type SupplierStorage interface {
	CreateSupplier(*e.Supplier) error
	DeleteSupplier(int) error
	UpdateSupplier(*e.Supplier) (*e.Supplier, error)
	GetSupplierByID(int) (*e.Supplier, error)
	FilterSuppliersByName(string) (*[]e.Supplier, error)
	FilterSuppliersBySupplieID(int) (*[]e.Supplier, error)
	// FilterSuppliersBySupplieName(string) (*[]Supplier, error)
	// FilterSuppliersByProductID(int) (*[]Supplier, error)
	// FilterSuppliersByProductName(string) (*[]Supplier, error)
}

type SuppliesStorage interface {
	CreateSupplie(*e.Supplie) error
	// DeleleteSupplie(int) error
	// UpdateSupplie(int) error
	// GetSupplieByID(int) (*e.Supplie, error)
	// FilterSuppliesByName(string) (*[]e.Supplie, error)
}

type SupplierSuppliesStorage interface {
	CreateSupplierSupplie(*e.SupplierSupplie) error
}
