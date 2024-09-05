package repository

import (
	"time"

	"github.com/kpgriffith/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) (int, error)
	SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error)
	SearchAvailabilityForAllRoomsByDates(start, end time.Time) ([]*models.Room, error)
}
