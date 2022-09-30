package file

import (
	"bufio"
	"fmt"
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

func (f *File) Write(ticket service.Ticket) (err error) {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	line := fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", ticket.Id, ticket.Name, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	file.WriteString(line)
	return
}

func (f *File) WriteAll(tickets []service.Ticket) (err error) {
	_, err = os.Create(f.Path)
	for _, ticket := range tickets {
		f.Write(ticket)
	}
	return
}
