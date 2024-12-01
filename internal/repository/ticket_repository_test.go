package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thesayedirfan/train-booking/internal/allocators"
	"github.com/thesayedirfan/train-booking/internal/entity"
	"github.com/thesayedirfan/train-booking/internal/repository"
)

func setup() (*repository.TicketRepository, *entity.Train) {
	train := &entity.Train{
		Name:     "Express",
		From:     "Station A",
		To:       "Station B",
		Sections: make(map[string][]entity.Seat),
	}

	sections := []string{"A", "B"}
	capacity := 10
	allocator := allocators.NewDefaultAllocator()

	repo := repository.NewInMemoryRepository(train, sections, capacity, allocator)
	return repo, train
}

func TestPurchaseTicket(t *testing.T) {
	repo, _ := setup()

	ticket := &entity.Ticket{
		Train:   *repo.Trains,
		User:    entity.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		Price:   50.0,
		Section: "A",
	}

	ticketID, err := repo.PurchaseTicket(ticket)
	assert.NoError(t, err)
	assert.NotEmpty(t, ticketID)

	ticketFromRepo, err := repo.GetDetails(ticketID)
	assert.NoError(t, err)
	assert.Equal(t, ticket.User.Email, ticketFromRepo.User.Email)
}

func TestGetSectionUsers(t *testing.T) {
	repo, _ := setup()

	ticket := &entity.Ticket{
		Train:   *repo.Trains,
		User:    entity.User{FirstName: "Alice", LastName: "Smith", Email: "alice.smith@example.com"},
		Price:   100.0,
		Section: "B",
	}

	_, err := repo.PurchaseTicket(ticket)
	assert.NoError(t, err)

	sectionSeats, err := repo.GetSectionUsers("B")
	assert.NoError(t, err)

	var users []*entity.User
	for _, seat := range sectionSeats {
		if seat.User != nil {
			users = append(users, seat.User)
		}
	}

	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Alice", users[0].FirstName)
}

func TestRemoveUser(t *testing.T) {
	repo, _ := setup()

	ticket := &entity.Ticket{
		Train:   *repo.Trains,
		User:    entity.User{FirstName: "Bob", LastName: "Johnson", Email: "bob.johnson@example.com"},
		Price:   30.0,
		Section: "A",
	}

	ticketID, err := repo.PurchaseTicket(ticket)
	assert.NoError(t, err)

	err = repo.RemoveUser(ticketID)
	assert.NoError(t, err)

	sectionSeats, _ := repo.GetSectionUsers("A")
	for _, seat := range sectionSeats {
		if seat.User != nil {
			assert.NotEqual(t, ticket.User.Email, seat.User.Email)
		}
	}
}

func TestModifySeat(t *testing.T) {
	repo, _ := setup()

	ticket := &entity.Ticket{
		Train:   *repo.Trains,
		User:    entity.User{FirstName: "Charlie", LastName: "Brown", Email: "charlie.brown@example.com"},
		Price:   40.0,
		Section: "A",
	}

	ticketID, err := repo.PurchaseTicket(ticket)
	assert.NoError(t, err)

	newSeatNumber := int64(5)
	newSection := "B"
	seatNumber, err := repo.ModifySeat(ticketID, newSection, newSeatNumber)
	assert.NoError(t, err)
	assert.Equal(t, newSeatNumber, seatNumber)

	ticketFromRepo, _ := repo.GetDetails(ticketID)
	assert.Equal(t, newSection, ticketFromRepo.Section)
	assert.Equal(t, newSeatNumber, ticketFromRepo.SeatNumber)
}
