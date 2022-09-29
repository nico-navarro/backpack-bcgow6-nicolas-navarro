package file

import (
	"bufio"
	"hackathon/internal/service"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Path string
}

// type Ticket struct {
// 	Id                             int
// 	Name, Email, Destination, Date string
// 	Price                          int
// }

func (f *File) Read() (tickets []service.Ticket, err error) {
	file, _ := os.Open(f.Path) // NOTE: handle error
	fscanner := bufio.NewScanner(file)
	ticketIndex := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		lineSlice := strings.Split(line, ",")
		ticket := service.Ticket{}
		ticket.Id, _ = strconv.Atoi(lineSlice[0]) // NOTE: handle error
		ticket.Name = lineSlice[1]
		ticket.Email = lineSlice[2]
		ticket.Destination = lineSlice[3]
		ticket.Date = lineSlice[4]
		ticket.Price, _ = strconv.Atoi(lineSlice[5]) // NOTE: handle error
		tickets = append(tickets, ticket)
		ticketIndex++
	}
	return
}

func (f *File) Write(service.Ticket) error {
	return nil
}
