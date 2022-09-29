package main

import (
	"fmt"
	"hackathon/internal/file"
	"hackathon/internal/service"
)

func main() {
	//var tickets []service.Ticket
	csvFile := file.File{Path: "tickets.csv"}
	tickets, _ := csvFile.Read()
	fmt.Println(tickets[144])
	fmt.Println(tickets[644])
	fmt.Println(tickets[914])
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
