package handler

import (
	"context"


	"github.com/thesayedirfan/train-booking/internal/entity"
	"github.com/thesayedirfan/train-booking/internal/service"
	pb "github.com/thesayedirfan/train-booking/proto"
)

type TicketHandler struct {
	pb.UnimplementedTicketServiceServer
	service *service.TicketService
}

func NewTicketHandler(service *service.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

func (h *TicketHandler) PurchaseTicket(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	user := entity.User{
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Email:     req.User.Email,
	}

	train := entity.Train{
		Name: req.Train.Name,
		From: req.Train.From,
		To:   req.Train.To,
	}

	ticket, err := h.service.PurchaseTicket(user, train, req.Price,req.Section)
	if err != nil {
		return nil, err
	}

	return &pb.PurchaseResponse{
		Ticket: &pb.Ticket{
			ID: ticket.ID,
			User: &pb.User{
				FirstName: ticket.User.FirstName,
				LastName:  ticket.User.LastName,
				Email:     ticket.User.Email,
			},
			Train: &pb.Train{
				Name: ticket.Train.Name,
				From: ticket.Train.From,
				To:   ticket.Train.To,
			},
			Price:      ticket.Price,
			Section:    ticket.Section,
			SeatNumber: ticket.SeatNumber,
		},
	}, nil
}

func (h *TicketHandler) GetReceipt(ctx context.Context, req *pb.ReceiptRequest) (*pb.ReceiptResponse, error) {
	ticket, err := h.service.GetTicketDetails(req.TicketId)
	if err != nil {
		return nil, err
	}

	return &pb.ReceiptResponse{
		TicketId: ticket.ID,
		User: &pb.User{
			FirstName: ticket.User.FirstName,
			LastName:  ticket.User.LastName,
			Email:     ticket.User.Email,
		},
		Train: &pb.Train{
			Name: ticket.Train.Name,
			From: ticket.Train.From,
			To:   ticket.Train.To,
		},
		Price:      ticket.Price,
		SeatNumber: ticket.SeatNumber,
	}, nil
}

func (h *TicketHandler) ViewSectionUsers(ctx context.Context, req *pb.SectionRequest) (*pb.SectionUsersResponse, error) {
	sectionSeats, err := h.service.GetSectionUsers(req.Section)

	if err != nil {
		return nil, err
	}

	var protoSeats []*pb.Seat
	for _, seat := range sectionSeats {
		if seat.User == nil {
			continue	
		}
		 protoSeats = append(protoSeats, &pb.Seat{
			Number: seat.Number,
			User: &pb.User{
				FirstName: seat.User.FirstName,
				LastName:  seat.User.LastName,
				Email:     seat.User.Email,
			},
		})
	}

	return &pb.SectionUsersResponse{
		Users: protoSeats,
	}, nil
}


func (h *TicketHandler) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	err := h.service.RemoveUser(req.TicketId)
	return &pb.RemoveUserResponse{Success: err == nil}, err
}

func (h *TicketHandler) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.ModifySeatResponse, error) {
	seatNumber, err := h.service.ModifyTicketSeat(req.TicketId, req.NewSection,req.SeatNumber)
	if err != nil {
		return nil, err
	}

	return &pb.ModifySeatResponse{
		Success:       true,
		NewSeatNumber: seatNumber,
	}, nil
}
