package errors

import (
	"errors"
)

// Train Errors
var (
	ErrTrainWithIDAlreadyExits = errors.New("train with id already exists")
	ErrTrainWithIDDoesNotExits = errors.New("train with id does not exits")
)
// Seat Errors
var (
	ErrSeatNotAvailable = errors.New("no seats available")
	ErrSeatInvalid   = errors.New("invalid seat number")
    ErrSeatOccupied  = errors.New("seat is already occupied")
)

// Ticket Error
var (
	ErrTicketNotFound = errors.New("ticket not found")
)
// Section Error
var (
	ErrSectionInvalid = errors.New("section is invalid")
)