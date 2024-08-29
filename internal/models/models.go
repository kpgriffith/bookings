package models

import "time"

// User model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation model
type Reservation struct {
	ID        int
	UserId    int
	RoomId    int
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Room      Room
}

// RoomRestriction model
type RoomRestriction struct {
	ID            int
	ReservationId int
	RoomId        int
	RestrictionId int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Reservation   Reservation
	Room          Room
	Restriction   Restriction
}
