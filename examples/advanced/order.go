package advanced

import (
	"errors"
	"time"
)

type Status int

const (
	Submitted Status = iota
	Paid
	Shipped
	Cancelled
)

type Order struct {
	created    time.Time
	status     Status
	customerID int
}

func (o *Order) Ship() { o.status = Shipped }

func (o *Order) Pay() Status {
	o.status = Paid
	return o.status
}

func (o *Order) Valid() (bool, error) {
	return false, errors.New("fake error")
}
