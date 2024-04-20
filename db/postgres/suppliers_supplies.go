package postgres

import (
	e "suppliers/entities"
)

func (s *PostgresStore) CreateSupplierSupplie(ss *e.SupplierSupplie) error {

	query := `INSERT INTO suppliers_supplies (name, supplier_id, supplie_id, created_at) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id`

	var id int
	err := s.db.QueryRow(query, ss.Name, ss.SupplierID, ss.SupplieID, ss.CreatedAt).Scan(&id)

	if err != nil {
		return err
	}

	ss.ID = id
	return nil
}
