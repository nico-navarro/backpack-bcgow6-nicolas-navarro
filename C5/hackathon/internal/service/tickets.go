package service

import "errors"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                             int
	Name, Email, Destination, Date string
	Price                          int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.Tickets = append(b.Tickets, t)
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (searchedTicket Ticket, err error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			searchedTicket = ticket
			return
		}
	}
	err = errors.New("no se encontró el id especificado")
	return
}

func (b *bookings) Update(id int, t Ticket) (searchedTicket Ticket, err error) {
	for index, ticket := range b.Tickets {
		if ticket.Id == id {
			b.Tickets[index] = t
			searchedTicket = t
			return
		}
	}
	err = errors.New("no se encontró el id especificado")
	return
}

// delete https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func (b *bookings) Delete(id int) (int, error) {
	for index, ticket := range b.Tickets {
		if ticket.Id == id {
			b.Tickets[index] = b.Tickets[len(b.Tickets)-1]
			b.Tickets = b.Tickets[:len(b.Tickets)-1]
			return id, nil
		}
	}
	err := errors.New("no se encontró el id especificado")
	return 0, err
}
