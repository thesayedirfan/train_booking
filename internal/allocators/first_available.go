package allocators

import (
	

	"github.com/thesayedirfan/train-booking/internal/entity"
	"github.com/thesayedirfan/train-booking/internal/errors"
)

type DefaultAllocator struct{}

func NewDefaultAllocator() *DefaultAllocator {
	return &DefaultAllocator{}
}


func (a *DefaultAllocator) Allocate(seats []entity.Seat, user entity.User) (int64, error) {
    for i := range seats {
        if seats[i].User == nil {
            seats[i] = entity.Seat{
                Number: int64(i + 1),
                User:   &user,
            }
            return seats[i].Number, nil
        }
    }
    return 0, errors.ErrSeatNotAvailable
}