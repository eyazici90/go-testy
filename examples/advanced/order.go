package advanced

import (
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
