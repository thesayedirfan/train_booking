package allocators_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thesayedirfan/train-booking/internal/allocators"
	"github.com/thesayedirfan/train-booking/internal/entity"
	"github.com/thesayedirfan/train-booking/internal/errors"
)

func TestDefaultAllocator_Allocate(t *testing.T) {
	allocator := allocators.NewDefaultAllocator()
	seats := make([]entity.Seat, 5)

	for i := range seats {
		seats[i] = entity.Seat{
			Number: int64(i + 1),
			User:   nil,
		}
	}

	// Allocating first user
	user1 := entity.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	seatNumber, err := allocator.Allocate(seats, user1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), seatNumber)
	assert.NotNil(t, seats[0].User)
	assert.Equal(t, user1, *seats[0].User)

	// Allocating to remaining user
	for i := 1; i < len(seats); i++ {
		user := entity.User{
			FirstName: "User",
			LastName:  string(rune(i + 1)),
			Email:     "user" + string(rune(i+1)) + "@example.com",
		}
		_, err := allocator.Allocate(seats, user)
		assert.NoError(t, err)
		assert.NotNil(t, seats[i].User)
	}

	// Allocating when all users all occupied

	userExtra := entity.User{
		FirstName: "Extra",
		LastName:  "User",
		Email:     "extra.user@example.com",
	}
	seatNumber, err = allocator.Allocate(seats, userExtra)
	assert.ErrorIs(t, err, errors.ErrSeatNotAvailable)
	assert.Equal(t, int64(0), seatNumber)
}
