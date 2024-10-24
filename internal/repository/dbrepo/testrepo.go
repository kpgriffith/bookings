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
	// if room id is 2 fail, otherwise pass
	if res.RoomId == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestrictions inserts an entry for room restrictions
func (p *testDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) (int, error) {
	if roomRestriction.RoomId == 1000 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// SearchAvailabiltyByDatesByRoomId return true if the room is available, false if it's not for the dates given
func (p *testDbRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error) {
	if roomId == 1000 {
		return false, errors.New("some error")
	}
	return true, nil
}

// SearchAvailabilityForAllRoomsByDates return a slice of models.Room if there are rooms available, nil if not.
func (p *testDbRepo) SearchAvailabilityForAllRoomsByDates(start, end time.Time) ([]*models.Room, error) {
	var rooms []*models.Room

	successStart, _ := time.Parse("2006-01-02", "2050-01-01")
	if start.Equal(successStart) {
		rooms = append(rooms, &models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		})
	}

	errorStart, _ := time.Parse("2006-01-02", "2100-01-01")
	if start.Equal(errorStart) {
		return rooms, errors.New("failed to query db")
	}

	return rooms, nil
}

func (p *testDbRepo) GetRoomById(id int) (*models.Room, error) {
	var room models.Room
	if id > 2 {
		return &room, errors.New("some error")
	}
	return &room, nil
}
