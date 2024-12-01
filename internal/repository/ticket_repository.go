package repository

import (
	"sync"
	"github.com/thesayedirfan/train-booking/internal/entity"
	"github.com/thesayedirfan/train-booking/internal/errors"
	"github.com/thesayedirfan/train-booking/pkg/uuid"
)

type TicketRepository struct {
	Tickets   map[string]*entity.Ticket
	Trains    *entity.Train
	Allocator entity.Allocator
	mu        sync.Mutex
}

func NewInMemoryRepository(train *entity.Train, sections []string, capacity int, allocator entity.Allocator) *TicketRepository {

	if train.Sections == nil {
		train.Sections = make(map[string][]entity.Seat)
	}

	for _, section := range sections {
		if _, exists := train.Sections[section]; !exists {
			train.Sections[section] = make([]entity.Seat, capacity)
			for i := 0; i < capacity; i++ {
				train.Sections[section][i] = entity.Seat{
					Number: int64(i + 1),
					User:   nil,
				}
			}
		}
	}

	repo := &TicketRepository{
		Tickets:   make(map[string]*entity.Ticket),
		Trains:    train,
		Allocator: allocator,
	}
	return repo
}

func (r *TicketRepository) PurchaseTicket(ticket *entity.Ticket) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	section := r.Trains.Sections[ticket.Section]
	availableSeat, err := r.Allocator.Allocate(section, ticket.User)
	if err != nil {
		return "", err
	}

	ticket.ID = uuid.GenerateShortUUID()
	ticket.SeatNumber = availableSeat

	r.Tickets[ticket.ID] = ticket

	return ticket.ID, nil
}

func (r *TicketRepository) GetDetails(ticketID string) (*entity.Ticket, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ticket, exists := r.Tickets[ticketID]
	if !exists {
		return nil, errors.ErrTicketNotFound
	}

	return ticket, nil
}

func (r *TicketRepository) GetSectionUsers(section string) ([]entity.Seat, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	sectionSeats, exists := r.Trains.Sections[section]
	if !exists {
		return nil, errors.ErrSectionInvalid
	}

	return sectionSeats, nil
}

func (r *TicketRepository) RemoveUser(ticketID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	ticket, exists := r.Tickets[ticketID]
	if !exists {
		return errors.ErrTicketNotFound
	}

	for i := range r.Trains.Sections[ticket.Section] {
		seat := &r.Trains.Sections[ticket.Section][i]
		if seat.User != nil && ticket.ID == ticketID {
			seat.User = nil
			break
		}
	}

	delete(r.Tickets, ticketID)

	return nil
}

func (r *TicketRepository) ModifySeat(ticketID, section string, seatNumber int64) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()


	ticket, exists := r.Tickets[ticketID]
	if !exists {
		return 0, errors.ErrTicketNotFound
	}

	sectionSeats, exists := r.Trains.Sections[section]
	if !exists {
		return 0, errors.ErrSectionInvalid
	}

	if seatNumber < 1 || int(seatNumber) > len(sectionSeats) {
		return 0, errors.ErrSeatInvalid
	}

	if sectionSeats[seatNumber-1].User != nil {
		return 0, errors.ErrSeatOccupied
	}

	for i := range r.Trains.Sections[ticket.Section] {
		seat := &r.Trains.Sections[ticket.Section][i]
		if seat.User != nil && ticket.ID == ticketID {
			seat.User = nil
			break
		}
	}

    ticket.Section = section
    ticket.SeatNumber = seatNumber

    sectionSeats[seatNumber-1].User = &ticket.User
    sectionSeats[seatNumber-1].Number = seatNumber

    r.Tickets[ticketID] = ticket


	return seatNumber, nil
}
