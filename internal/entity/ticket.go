package entity

type Ticket struct {
	ID string
	Train Train
	User User
	Price float64
	Section string
	SeatNumber int64
}