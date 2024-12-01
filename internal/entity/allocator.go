package entity

type Allocator interface {
    Allocate(seats []Seat, user User) (int64, error)
}