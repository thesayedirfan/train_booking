package main

import (
	
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/thesayedirfan/train-booking/handler"

	"github.com/thesayedirfan/train-booking/internal/allocators"
	"github.com/thesayedirfan/train-booking/internal/entity"
	"github.com/thesayedirfan/train-booking/internal/repository"
	"github.com/thesayedirfan/train-booking/internal/service"
	pb "github.com/thesayedirfan/train-booking/proto"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	train := &entity.Train{
		Sections: make(map[string][]entity.Seat),
	}

	allocator := allocators.NewDefaultAllocator()
	repo := repository.NewInMemoryRepository(train, []string{"A", "B"}, 10, allocator)

	
	 ticketService := service.NewTicketService(repo)


	ticketHandler := handler.NewTicketHandler(ticketService)

	 pb.RegisterTicketServiceServer(grpcServer, ticketHandler)

	
	 log.Println("Starting gRPC server on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
