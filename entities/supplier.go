package entities

import "time"

type Supplier struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Emails       []string  `json:"emails"`
	PhoneNumbers []string  `json:"phoneNumbers"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewSupplier(name string, emails, phoneNumbers []string) *Supplier {
	date := time.Now().UTC()
	return &Supplier{
		// ID:           rand.Intn(300),
		Name:         name,
		Emails:       emails,
		PhoneNumbers: phoneNumbers,
		CreatedAt:    date,
		UpdatedAt:    date,
	}
}
