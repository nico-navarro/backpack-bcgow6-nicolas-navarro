package main

import (
	"fmt"
	"hackathon/internal/file"
	"hackathon/internal/service"
)

func main() {
	csvFile := file.File{Path: "tickets.csv"}
	tickets, _ := csvFile.Read()
	bookings := service.NewBookings(tickets)
	ticket, _ := bookings.Read(1) // NOTE: handle error
	fmt.Println(ticket)
}
