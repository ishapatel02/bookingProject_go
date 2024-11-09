package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID          uuid.UUID `db:"id"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"-"`
	AccessLevel int       `db:"access_level" json:"access_level"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
