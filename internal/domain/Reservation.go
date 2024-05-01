package domain

import "time"

type Reservation struct {
	id              int64
	reservedCopyId  int64
	reservingUserId int64
	reservationDate time.Time
	status          string
}
