package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/thesayedirfan/train-booking/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	purchaseResp1,err  := client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		User: &pb.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		Train: &pb.Train{
			From:  "London",
			To:    "Paris",
		},
		Price: 20.0,
		Section: "A",
	})
	if err != nil {
		log.Fatalf("Error purchasing ticket: %v", err)
	}

	purchaseResp2, err := client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		User: &pb.User{
			FirstName: "Pan",
			LastName:  "Doe",
			Email:     "pan.doe@example.com",
		},
		Train: &pb.Train{
			From:  "London",
			To:    "Paris",
		},
		Price: 30.0,
		Section: "B",
	})
	if err != nil {
		log.Fatalf("Error purchasing ticket: %v", err)
	}

	purchaseResp3, err := client.PurchaseTicket(ctx, &pb.PurchaseRequest{
		User: &pb.User{
			FirstName: "San",
			LastName:  "Doe",
			Email:     "san.doe@example.com",
		},
		Train: &pb.Train{
			From:  "London",
			To:    "Paris",
		},
		Price: 40.0,
		Section: "A",
	})
	if err != nil {
		log.Fatalf("Error purchasing ticket: %v", err)
	}

	fmt.Println(purchaseResp3)


	ticketDetails1 , err := client.GetReceipt(ctx,&pb.ReceiptRequest{
		TicketId: purchaseResp1.Ticket.ID,
	})

	if err != nil {
		log.Fatalf("Error purchasing ticket: %v", err)
	}

	 fmt.Println(ticketDetails1)


	client.RemoveUser(ctx,&pb.RemoveUserRequest{
		TicketId: purchaseResp3.Ticket.ID,
	})


	client.ModifySeat(ctx,&pb.ModifySeatRequest{
		TicketId: purchaseResp2.Ticket.ID,
		NewSection: "B",
		SeatNumber: 3,
	})


	sectionA , err := client.ViewSectionUsers(ctx,&pb.SectionRequest{
		Section: "A",
	})

	if err != nil {
		log.Fatalf("Error purchasing ticket: %v", err)
	}

	fmt.Println("============================")

	fmt.Println(sectionA)

	fmt.Println("============================")

	sectionB , err := client.ViewSectionUsers(ctx,&pb.SectionRequest{
		Section: "B",
	})

	if err != nil {
		log.Fatalf("Error purchasing ticket: %v", err)
	}

	fmt.Println("============================")

	fmt.Println(sectionB)

	fmt.Println("============================")
	



}