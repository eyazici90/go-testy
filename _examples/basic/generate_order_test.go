package basic

import (
	"time"
)

type generators struct {
	orderGenerator *orderGenerator
}

type orderGenerator struct {
	ordr Order
}

func a() *generators {
	return &generators{}
}
func an() *generators { return a() }

func (o *generators) order() *orderGenerator {
	return &orderGenerator{}
}

func (o *orderGenerator) createdAt(t time.Time) *orderGenerator {
	o.ordr.created = t
	return o
}

func (o *orderGenerator) as(s Status) *orderGenerator {
	o.ordr.status = s
	return o
}

func (o *orderGenerator) was() *orderGenerator { return o }

func (o *orderGenerator) generate() *Order { return &o.ordr }
