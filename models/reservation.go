package models

import "time"

type Reservation struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Phone     string    `db:"phone"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	RoomID    int       `db:"room_id"`
	Room      Room      `belongs_to:"room" fk_id:"RoomID"` // Association with Room
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
