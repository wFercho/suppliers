package postgres

import (
	e "suppliers/entities"
)

func (s *PostgresStore) CreateSupplie(supl *e.Supplie) error {

	query := `INSERT INTO supplie (name, updated_at, created_at) 
	VALUES ($1, $2, $3) 
	RETURNING id`

	var id int

	err := s.db.QueryRow(query, supl.Name, supl.UpdatedAt, supl.CreatedAt).Scan(&id)

	if err != nil {
		return err
	}

	supl.ID = id
	return nil
}
