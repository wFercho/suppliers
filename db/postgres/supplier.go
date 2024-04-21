package postgres

import (
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"

	e "suppliers/entities"
)

// type supplierPostgres struct {
// 	id            int
// 	name          string
// 	emails        pq.StringArray
// 	phone_numbers pq.StringArray
// 	created_at    time.Time
// 	updated_at    time.Time
// }

func (s *PostgresStore) CreateSupplier(sup *e.Supplier) error {
	query := fmt.Sprintf(`INSERT INTO %s (name, emails, phone_numbers, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`, SUPPLIER_TABLE)

	// fmt.Println(queryN)

	// query := `INSERT INTO supplier (name, emails, phone_numbers, created_at, updated_at)
	// VALUES ($1, $2, $3, $4, $5)
	// RETURNING id`

	id := 0
	err := s.db.QueryRow(
		query,
		sup.Name,
		pq.Array(sup.Emails),
		pq.Array(sup.PhoneNumbers),
		sup.CreatedAt,
		sup.UpdatedAt).Scan(&id)

	if err != nil {
		return err
	}
	sup.ID = id
	//defer s.db.Close()

	return nil
}

func (s *PostgresStore) DeleteSupplier(id int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", SUPPLIER_TABLE)

	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	//defer s.db.Close()

	return nil
}

func (s *PostgresStore) UpdateSupplier(sup *e.Supplier) (*e.Supplier, error) {

	updateEmail := sup.Emails != nil
	updatePhoneNumbers := sup.PhoneNumbers != nil
	updateName := len(sup.Name) > 0

	if updateEmail && updatePhoneNumbers && updateName {
		query := `UPDATE supplier SET name = $1, emails = $2, phone_numbers = $3, updated_at = $4 WHERE id = $5 RETURNING *`
		rows := s.db.QueryRow(query, sup.Name, pq.Array(sup.Emails), pq.Array(sup.PhoneNumbers), time.Now(), sup.ID)
		err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		return sup, nil
	}

	if updateEmail && updatePhoneNumbers && !updateName {
		query := `UPDATE supplier SET emails = $1, phone_numbers = $2, updated_at = $3 WHERE id = $4 RETURNING *`
		rows := s.db.QueryRow(query, pq.Array(sup.Emails), pq.Array(sup.PhoneNumbers), time.Now(), sup.ID)
		err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		return sup, nil
	}

	if updateEmail && !updatePhoneNumbers && updateName {
		query := `UPDATE supplier SET name = $1, emails = $2, updated_at = $3 WHERE id = $4 RETURNING *`
		rows := s.db.QueryRow(query, sup.Name, pq.Array(sup.Emails), time.Now(), sup.ID)
		err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		return sup, nil
	}

	if !updateEmail && updatePhoneNumbers && updateName {
		query := `UPDATE supplier SET name = $1, phone_numbers = $2, updated_at = $3 WHERE id = $4 RETURNING *`
		rows := s.db.QueryRow(query, sup.Name, pq.Array(sup.PhoneNumbers), time.Now(), sup.ID)
		err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		return sup, nil
	}

	if !updateEmail && !updatePhoneNumbers && updateName {
		query := `UPDATE supplier SET name = $1, updated_at = $2 WHERE id = $3 RETURNING *`
		rows := s.db.QueryRow(query, sup.Name, time.Now(), sup.ID)
		supp := &e.Supplier{}
		err := rows.Scan(&supp.ID, &supp.Name, pq.Array(&supp.Emails), pq.Array(&supp.PhoneNumbers), &supp.CreatedAt, &supp.UpdatedAt)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("%+v\n", supp)
		return supp, nil
	}

	if !updateEmail && updatePhoneNumbers && !updateName {
		query := `UPDATE supplier SET phone_numbers = $1, updated_at = $2 WHERE id = $3 RETURNING *`
		rows := s.db.QueryRow(query, pq.Array(sup.PhoneNumbers), time.Now(), sup.ID)
		err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		return sup, nil
	}

	if updateEmail && !updatePhoneNumbers && !updateName {
		query := `UPDATE supplier SET emails = $1 updated_at = $2 WHERE id = $3 RETURNING *`
		rows := s.db.QueryRow(query, pq.Array(sup.Emails), time.Now(), sup.ID)
		err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		return sup, nil
	}

	return nil, fmt.Errorf("no field specified")
}

func (s *PostgresStore) GetSupplierByID(id int) (*e.Supplier, error) {

	query := `SELECT * FROM supplier WHERE id = $1`

	rows := s.db.QueryRow(query, id)

	//sup := supplierPostgres{}
	sup := e.Supplier{}
	//rows.Scan(&sup.id, &sup.name, &sup.emails, &sup.phone_numbers, &sup.created_at, &sup.updated_at)
	err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
	//fmt.Printf("%+v\n", sup)
	//defer s.db.Close()

	if err != nil {
		return nil, err
	}

	return &sup, nil
}

func (s *PostgresStore) FilterSuppliersByName(name string) (*[]e.Supplier, error) {

	containsName := fmt.Sprint("%", strings.ToLower(name), "%")
	pQuery := fmt.Sprintf("SELECT * FROM supplier WHERE LOWER(name) LIKE '%s'", containsName)

	rows, err := s.db.Query(pQuery)

	if err != nil {
		return nil, err
	}

	suppliers := make([]e.Supplier, 0)
	for rows.Next() {
		sup := e.Supplier{}
		err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, sup)
	}

	return &suppliers, nil
}

func (s *PostgresStore) FilterSuppliersBySupplieID(supplieID int) (*[]e.Supplier, error) {

	subQuery := `SELECT supplier_id FROM suppliers_supplies WHERE supplie_id = $1`

	query := fmt.Sprintf("SELECT * FROM supplier JOIN (%s) ON id = supplier_id", subQuery)
	fmt.Println(query)

	return nil, nil
}
