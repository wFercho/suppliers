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

	var fields []string
	var values []any
	query := fmt.Sprintf("UPDATE %s SET ", SUPPLIER_TABLE)

	if updateName {
		fields = append(fields, "name")
		values = append(values, sup.Name)
	}

	if updateEmail {
		fields = append(fields, "emails")
		values = append(values, pq.Array(sup.Emails))
	}

	if updatePhoneNumbers {
		fields = append(fields, "phone_numbers")
		values = append(values, pq.Array(sup.PhoneNumbers))
	}

	l := len(fields)

	if l == 0 {
		return nil, fmt.Errorf("no field specified")
	}

	fields = append(fields, "updated_at")
	values = append(values, time.Now())

	for i, field := range fields {
		query += fmt.Sprintf("%s = $%d", field, i+1)
		if i < l {
			query += ", "
		}
	}

	query += fmt.Sprintf(" WHERE id = $%d", l+2)
	values = append(values, sup.ID)
	query += " RETURNING *"

	rows := s.db.QueryRow(query, values...)
	err := rows.Scan(&sup.ID, &sup.Name, pq.Array(&sup.Emails), pq.Array(&sup.PhoneNumbers), &sup.CreatedAt, &sup.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return sup, nil

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
