package entities

import "time"

type Supplie struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewSupplie(name string) *Supplie {
	date := time.Now().UTC()
	return &Supplie{
		Name:      name,
		UpdatedAt: date,
		CreatedAt: date,
	}
}
