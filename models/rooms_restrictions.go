package models

import "time"

type RoomRestriction struct {
	ID            int         `db:"id"`
	StartDate     time.Time   `db:"start_date"`
	EndDate       time.Time   `db:"end_date"`
	RoomID        int         `db:"room_id"`
	RestrictionID int         `db:"restriction_id"`
	ReservationID int         `db:"reservation_id"`
	Restriction   Restriction `belongs_to:"restriction" fk_id:"RestrictionID"` // Association with Room
	Reservation   Reservation `belongs_to:"reservation" fk_id:"ReservationID"` // Association with Room
	CreatedAt     time.Time   `db:"created_at"`
	UpdatedAt     time.Time   `db:"updated_at"`
}
