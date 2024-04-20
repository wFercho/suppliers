package postgres

import (
	"database/sql"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=suppliers password=Devpass1234$ sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	if err := s.createSupplierTable(); err != nil {
		return err
	}
	if err := s.createSupplieTable(); err != nil {
		return err
	}
	if err := s.createSuppliersSuppliesTable(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createSupplierTable() error {
	query := `CREATE TABLE IF NOT EXISTS supplier (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		emails text[],
		phone_numbers text[],
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
		);`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createSupplieTable() error {
	query := `CREATE TABLE IF NOT EXISTS supplie (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
		);`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createSuppliersSuppliesTable() error {
	query := `CREATE TABLE IF NOT EXISTS suppliers_supplies(
		id SERIAL PRIMARY KEY,
		supplier_id INT REFERENCES supplier(id),
		supplie_id INT REFERENCES supplie(id),
		name VARCHAR(60),
		created_at TIMESTAMP NOT NULL
		);`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) buildQuery(params map[string]string) (query string, args []interface{}) {

	return "", nil
}
