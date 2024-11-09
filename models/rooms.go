package models

import "time"

type Room struct {
	ID        int       `db:"id"`
	RoomName  string    `db:"room_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
