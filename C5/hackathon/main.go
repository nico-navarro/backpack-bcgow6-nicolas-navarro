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

	//CREATE
	newTicket := service.Ticket{
		Id:          1001,
		Name:        "Nico Navarro",
		Email:       "nico.navarro@meli.cl",
		Destination: "Ecuador",
		Date:        "10:11",
		Price:       1000,
	}
	bookings.Create(newTicket)

	//READ
	ticket, _ := bookings.Read(1001) // NOTE: handle error
	fmt.Println(ticket)

	//UPDATE
	newData := service.Ticket{
		Id:          1001,
		Name:        "Nico Navarro",
		Email:       "nico.navarro@meli.cl",
		Destination: "Brasil",
		Date:        "10:11",
		Price:       1000,
	}
	bookings.Update(1001, newData)
	updatedTicket, _ := bookings.Read(1001) // NOTE: handle error
	fmt.Println(updatedTicket)

	//DELETE
	fmt.Println(bookings)
	deletedTicket, _ := bookings.Delete(998)
	fmt.Println(deletedTicket) // NOTE: handle error

}
