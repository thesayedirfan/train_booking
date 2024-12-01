# Cloudbees Assigment
## Project Title: Train Ticket Booking System
### Overview
This project is a Train Ticket Booking System implemented in Go (Golang) using gRPC and adhering to the Clean Architecture principles. 
It provides APIs for  purchase for a ticket, seat allocation, shows the details of the receipt for the user,view the users and seat they are allocated by the requested section,remove a user from the train and an API to modify a user's seat


### Installation
1. have latest version of go installed
    *  verify the installation of go

    ```bash
    go version
    ```
2. Clone the repository 
    ```bash
    git clone https://github.com/thesayedirfan/train_booking.git

    cd train_booking
    ```
3. Start The Project
    - start the server
    ```bash
    make start_server
    ```
    - start the client
    ```bash
    make start_client
    ```
4. Run the tests
    ```bash
    go test ./...
    ```

## Project Structure

```bash
├── Makefile
├── README.md
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── go.mod
├── go.sum
├── handler
│   └── grpc.go
├── internal
│   ├── allocators
│   │   ├── first_available.go
│   │   └── first_available_test.go
│   ├── entity
│   │   ├── allocator.go
│   │   ├── ticket.go
│   │   ├── train.go
│   │   └── user.go
│   ├── errors
│   │   └── errors.go
│   ├── repository
│   │   ├── ticket_repository.go
│   │   └── ticket_repository_test.go
│   └── service
│       └── ticket_service.go
├── pkg
│   └── uuid
│       └── generator.go
└── proto
    ├── ticketservice.pb.go
    ├── ticketservice.proto
    └── ticketservice_grpc.pb.go

```


## gRPC API Usage

### Available API Calls

| **Method**          | **Request Example**                                                                                               | **Description**                     |
|----------------------|-------------------------------------------------------------------------------------------------------------------|-------------------------------------|
| `PurchaseTicket`     | ```{user: {firstName: "John", lastName: "Doe", email: "john.doe@example.com"}, train: {from: "London", to: "Paris"}, price: 20.0, section: "A"}``` | Purchase a ticket for a specified train, user, and section. |
| `GetReceipt`         | `{ticketId: "12345"}`                                                                                            | Retrieve receipt details for a specific ticket by ID. |
| `RemoveUser`         | `{ticketId: "12345"}`                                                                                            | Remove a user and their ticket reservation by ticket ID. |
| `ModifySeat`         | `{ticketId: "12345", newSection: "B", seatNumber: 3}`                                                            | Modify the seat assignment for a specific ticket. |
| `ViewSectionUsers`   | `{section: "A"}`                                                                                                 | View all users assigned to a specific section. |

---