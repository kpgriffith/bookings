package dbrepo

import (
	"errors"
	"time"

	"github.com/kpgriffith/bookings/internal/models"
)

func (p *testDbRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the db
func (p *testDbRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestrictions inserts an entry for room restrictions
func (p *testDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) (int, error) {
	return 1, nil
}

// SearchAvailabiltyByDatesByRoomId return true if the room is available, false if it's not for the dates given
func (p *testDbRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRoomsByDates return a slice of models.Room if there are rooms available, nil if not.
func (p *testDbRepo) SearchAvailabilityForAllRoomsByDates(start, end time.Time) ([]*models.Room, error) {
	var rooms []*models.Room
	return rooms, nil
}

func (p *testDbRepo) GetRoomById(id int) (*models.Room, error) {
	var room models.Room
	if id > 2 {
		return &room, errors.New("some error")
	}
	return &room, nil
}
