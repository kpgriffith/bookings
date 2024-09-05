package dbrepo

import (
	"context"
	"time"

	"github.com/kpgriffith/bookings/internal/models"
)

func (p *postgresDbRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the db
func (p *postgresDbRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `insert into public.reservations 
	(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
	values
	($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	var reservationId int
	err := p.DB.QueryRowContext(ctx, stmt,
		res.FirstName, res.LastName, res.Email, res.Phone, res.StartDate,
		res.EndDate, res.RoomId, time.Now(), time.Now()).Scan(&reservationId)
	if err != nil {
		return 0, err
	}

	return reservationId, nil
}

// InsertRoomRestrictions inserts an entry for room restrictions
func (p *postgresDbRepo) InsertRoomRestriction(roomRestriction models.RoomRestriction) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `insert into public.room_restrictions 
	(room_id, reservation_id, restriction_id, start_date, end_date, created_at, updated_at)
	values
	($1,$2,$3,$4,$5,$6,$7) returning id`

	var roomRestrictionId int
	err := p.DB.QueryRowContext(ctx, stmt,
		roomRestriction.RoomId, roomRestriction.ReservationId, roomRestriction.RestrictionId, roomRestriction.StartDate,
		roomRestriction.EndDate, time.Now(), time.Now()).Scan(&roomRestrictionId)
	if err != nil {
		return 0, err
	}

	return roomRestrictionId, nil
}

// SearchAvailabiltyByDatesByRoomId return true if the room is available, false if it's not for the dates given
func (p *postgresDbRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `select count(id) 
	from public.room_restrictions
	where  $1 < end_date 
	and $2 > start_date
	and room_id = $3`

	var numRows int
	err := p.DB.QueryRowContext(ctx, stmt, start, end, roomId).Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// SearchAvailabilityForAllRoomsByDates return a slice of models.Room if there are rooms available, nil if not.
func (p *postgresDbRepo) SearchAvailabilityForAllRoomsByDates(start, end time.Time) ([]*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `select r.id, r.room_name 
	from public.rooms r 
	where r.id not in (
		select rr.room_id from public.room_restrictions rr 
		where $1 < rr.end_date and $2 > rr.start_date
	)`

	rows, err := p.DB.QueryContext(ctx, stmt, start, end)
	if err != nil {
		return nil, err
	}

	var rooms []*models.Room
	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
