package models

import "time"

type Restriction struct {
	ID              int       `db:"id"`
	RestrictionName string    `db:"restriction_name"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
