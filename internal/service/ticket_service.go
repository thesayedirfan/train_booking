package service

import (
	"github.com/thesayedirfan/train-booking/internal/entity"
	"github.com/thesayedirfan/train-booking/internal/repository"
)

type TicketService struct {
	repo *repository.TicketRepository
}

func NewTicketService(repo *repository.TicketRepository) *TicketService {
	return &TicketService{
		repo: repo,
	}
}

func (s *TicketService) PurchaseTicket(user entity.User, train entity.Train, price float64, section string) (*entity.Ticket, error) {
	ticket := &entity.Ticket{
		User:    user,
		Train:   train,
		Price:   price,
		Section: section,
	}

	ticketID, err := s.repo.PurchaseTicket(ticket)
	if err != nil {
		return nil, err
	}
	return s.repo.GetDetails(ticketID)
}

func (s *TicketService) GetTicketDetails(ticketID string) (*entity.Ticket, error) {
	return s.repo.GetDetails(ticketID)
}

func (s *TicketService) GetSectionUsers(section string) ([]entity.Seat, error) {
	return s.repo.GetSectionUsers(section)
}

func (s *TicketService) RemoveUser(ticketID string) error {
	return s.repo.RemoveUser(ticketID)
}

func (s *TicketService) ModifyTicketSeat(ticketID, section string, seatNumber int64) (int64, error) {
	return s.repo.ModifySeat(ticketID, section, seatNumber)
}
